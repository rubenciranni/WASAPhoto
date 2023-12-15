package api

import (
	"database/sql"
	"errors"
	"net/http"
	"strings"

	"github.com/gofrs/uuid"
	"github.com/julienschmidt/httprouter"
	"github.com/rubenciranni/WASAPhoto/service/api/reqcontext"
	"github.com/sirupsen/logrus"
)

// httpRouterHandler is the signature for functions that accepts a reqcontext.RequestContext in addition to those
// required by the httprouter package.
type httpRouterHandler func(http.ResponseWriter, *http.Request, httprouter.Params, reqcontext.RequestContext)

// wrap parses the request and adds a reqcontext.RequestContext instance related to the request.
func (rt *_router) wrap(fn httpRouterHandler, operationId string, requiresAuth bool) func(http.ResponseWriter, *http.Request, httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		reqUUID, err := uuid.NewV4()
		if err != nil {
			rt.baseLogger.WithError(err).Error("error generating a request UUID")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		var ctx = reqcontext.RequestContext{
			ReqUUID:        reqUUID,
			ReqOperationId: operationId,
		}

		// Create a request-specific logger
		ctx.Logger = rt.baseLogger.WithFields(logrus.Fields{
			"reqid":     ctx.ReqUUID.String(),
			"req-op-id": ctx.ReqOperationId,
			"remote-ip": r.RemoteAddr,
		})

		if requiresAuth {
			ctx.Logger.Debug("checking authorization header")
			// Check authorization header format
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				ctx.Logger.Error("request authorization header missing")
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			if !strings.HasPrefix(authHeader, "Bearer ") {
				ctx.Logger.Error("request authorization header has wrong format")
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			// Check authorization token format
			authToken := strings.TrimPrefix(authHeader, "Bearer ")
			if !isValidUUID(authToken) {
				ctx.Logger.Error("request authorization token is not a valid UUID")
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			// Check authorization token existance
			if user, err := rt.db.GetUser(authToken); errors.Is(err, sql.ErrNoRows) {
				ctx.Logger.Error("request authorization token not found")
				w.WriteHeader(http.StatusUnauthorized)
				return
			} else if err != nil {
				ctx.Logger.WithError(err).Error("error getting the userID by username from database")
				w.WriteHeader(http.StatusInternalServerError)
				return
			} else {
				// Authorization successful
				ctx.User = user
			}
		}

		// Call the next handler in chain (usually, the handler function for the path)
		fn(w, r, ps, ctx)
	}
}

func isValidUUID(u string) bool {
	_, err := uuid.FromString(u)
	return err == nil
}
