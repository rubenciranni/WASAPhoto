package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// Register routes
	rt.router.POST("/session", rt.wrap(rt.doLogin, "doLogin", false))
	rt.router.PUT("/settings/username", rt.wrap(rt.setMyUserName, "setMyUserName", true))
	rt.router.POST("/photos/", rt.wrap(rt.uploadPhoto, "uploadPhoto", true))
	rt.router.GET("/photos/", rt.wrap(rt.getPhotos, "getPhotos", true))
	rt.router.GET("/photos/:photoId", rt.wrap(rt.getPhoto, "getPhoto", true))
	rt.router.DELETE("/photos/:photoId", rt.wrap(rt.deletePhoto, "deletePhoto", true))
	rt.router.GET("/photos/:photoId/likes/", rt.wrap(rt.getLikes, "getLikes", true))
	rt.router.GET("/photos/:photoId/comments/", rt.wrap(rt.getComments, "getComments", true))
	rt.router.POST("/photos/:photoId/comments/", rt.wrap(rt.commentPhoto, "commentPhoto", true))

	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
