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
	rt.router.DELETE("/photos/:photoId/comments/:commentId", rt.wrap(rt.uncommentPhoto, "uncommentPhoto", true))
	rt.router.GET("/users/", rt.wrap(rt.searchUser, "searchUser", true))
	rt.router.GET("/users/:userId", rt.wrap(rt.getUserProfile, "getUserProfile", true))
	rt.router.GET("/users/:userId/following/", rt.wrap(rt.getFollowing, "getFollowing", true))
	rt.router.GET("/users/:userId/followers/", rt.wrap(rt.getFollowers, "getFollowers", true))
	rt.router.PUT("/liked-photos/:photoId", rt.wrap(rt.likePhoto, "likePhoto", true))
	rt.router.DELETE("/liked-photos/:photoId", rt.wrap(rt.unlikePhoto, "unlikePhoto", true))
	rt.router.PUT("/following/:userId", rt.wrap(rt.followUser, "followUser", true))
	rt.router.DELETE("/following/:userId", rt.wrap(rt.unfollowUser, "unfollowUser", true))

	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
