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
	rt.router.GET("/photos/:photoID", rt.wrap(rt.getPhoto, "getPhoto", true))
	rt.router.DELETE("/photos/:photoID", rt.wrap(rt.deletePhoto, "deletePhoto", true))
	rt.router.GET("/photos/:photoID/likes/", rt.wrap(rt.getLikes, "getLikes", true))
	rt.router.GET("/photos/:photoID/comments/", rt.wrap(rt.getComments, "getComments", true))
	rt.router.POST("/photos/:photoID/comments/", rt.wrap(rt.commentPhoto, "commentPhoto", true))
	rt.router.DELETE("/photos/:photoID/comments/:commentID", rt.wrap(rt.uncommentPhoto, "uncommentPhoto", true))
	rt.router.GET("/users/", rt.wrap(rt.searchUser, "searchUser", true))
	rt.router.GET("/users/:userID", rt.wrap(rt.getUserProfile, "getUserProfile", true))
	rt.router.GET("/users/:userID/following/", rt.wrap(rt.getFollowing, "getFollowing", true))
	rt.router.GET("/users/:userID/followers/", rt.wrap(rt.getFollowers, "getFollowers", true))
	rt.router.PUT("/liked-photos/:photoID", rt.wrap(rt.likePhoto, "likePhoto", true))
	rt.router.DELETE("/liked-photos/:photoID", rt.wrap(rt.unlikePhoto, "unlikePhoto", true))
	rt.router.PUT("/following/:userID", rt.wrap(rt.followUser, "followUser", true))
	rt.router.DELETE("/following/:userID", rt.wrap(rt.unfollowUser, "unfollowUser", true))
	rt.router.PUT("/bans/:userID", rt.wrap(rt.banUser, "banUser", true))
	rt.router.DELETE("/bans/:userID", rt.wrap(rt.unbanUser, "unbanUser", true))
	rt.router.GET("/stream", rt.wrap(rt.getMyStream, "getMyStream", true))

	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
