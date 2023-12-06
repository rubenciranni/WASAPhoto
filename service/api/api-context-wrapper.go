package api

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/gofrs/uuid"
	"github.com/julienschmidt/httprouter"
	"github.com/rubenciranni/WASAPhoto/service/api/reqcontext"
	"github.com/rubenciranni/WASAPhoto/service/model/response"
	"github.com/sirupsen/logrus"
)

// httpRouterHandler is the signature for functions that accepts a reqcontext.RequestContext in addition to those
// required by the httprouter package.
type httpRouterHandler func(http.ResponseWriter, *http.Request, httprouter.Params, reqcontext.RequestContext)

// wrap parses the request and adds a reqcontext.RequestContext instance related to the request.
func (rt *_router) wrap(fn httpRouterHandler, requiresAuth bool) func(http.ResponseWriter, *http.Request, httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		reqUUID, err := uuid.NewV4()
		if err != nil {
			rt.baseLogger.WithError(err).Error("can't generate a request UUID")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		var ctx = reqcontext.RequestContext{
			ReqUUID: reqUUID,
		}

		// Create a request-specific logger
		ctx.Logger = rt.baseLogger.WithFields(logrus.Fields{
			"reqid":     ctx.ReqUUID.String(),
			"remote-ip": r.RemoteAddr,
		})

		if requiresAuth {
			ctx.Logger.Debug("checking authorization header")
			respondUnauthorized := func() {
				w.WriteHeader(http.StatusUnauthorized)
				response := response.Problem{
					Title:  "Unauthorized",
					Status: 401,
					Detail: "Request authorization header missing or invalid."}
				json.NewEncoder(w).Encode(response)
				rt.baseLogger.Debug("sending response")
				return
			}
			// Check authorization header format
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				ctx.Logger.Debugf("request authorization header missing")
				respondUnauthorized()
				return
			}
			if !strings.HasPrefix(authHeader, "Bearer ") {
				ctx.Logger.Debugf("request authorization header has wrong format")
				respondUnauthorized()
				return
			}
			// Check authorization token format
			authToken := strings.TrimPrefix(authHeader, "Bearer ")
			if !isValidUUID(authToken) {
				ctx.Logger.Debugf("request authorization token is not a valid UUID")
				respondUnauthorized()
				return
			}
			// Check authorization token existance
			if user, err := rt.db.GetUser(authToken); errors.Is(err, sql.ErrNoRows) {
				ctx.Logger.Debugf("request authorization token not found")
				respondUnauthorized()
				return
			} else if err != nil {
				ctx.Logger.WithError(err).Error("can't get the userId for username")
				w.WriteHeader(http.StatusInternalServerError)
				response := response.Problem{
					Title:  "Internal Server Error",
					Status: 501,
					Detail: "Cannot authorize user"}
				json.NewEncoder(w).Encode(response)
				rt.baseLogger.Debug("sending response")
				return
				// Adding User to context
			} else {
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
