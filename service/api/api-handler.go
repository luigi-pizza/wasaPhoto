package api

import (
	"net/http"
)

// Ogni funzione in cui sei loggato deve checcare il bearer authorization
// userid non si deve passare perché è nell'header
// controlla se l'user è bannato o robe strane

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// Register routes

	rt.router.POST("/login", rt.wrap(rt.post_login, false))

	rt.router.PUT("/settings/username", rt.wrap(rt.put_username, true))

	rt.router.PUT("/banned_users/:userId", rt.wrap(rt.put_ban, true))
	rt.router.PUT("/followed_users/:userId", rt.wrap(rt.put_follow, true))
	rt.router.PUT("/photos/:postId/likes/self", rt.wrap(rt.put_like, true))
	rt.router.POST("/photos/", rt.wrap(rt.post_photo, true))
	rt.router.POST("/photos/:postId/comments/", rt.wrap(rt.post_comment, true))
	rt.router.GET("/users/:userId", rt.wrap(rt.get_userComplete, true))

	rt.router.DELETE("/banned_users/:userId", rt.wrap(rt.delete_ban, true))
	rt.router.DELETE("/followed_users/:userId", rt.wrap(rt.delete_follow, true))
	rt.router.DELETE("/photos/:postId/likes/self", rt.wrap(rt.delete_like, true))
	rt.router.DELETE("/photos/:postId/comments/:commentId", rt.wrap(rt.delete_comment, true))
	rt.router.DELETE("/photos/:postId", rt.wrap(rt.delete_photo, true))

	rt.router.GET("/users/", rt.wrap(rt.get_userList, true))

	rt.router.GET("/photos/:postId/comments/", rt.wrap(rt.get_commentList, true))

	rt.router.GET("/users/:userId/photos/", rt.wrap(rt.get_postList, true))
	rt.router.GET("/stream/", rt.wrap(rt.get_stream, true))

	rt.router.GET("/photos/:postId", rt.wrap(rt.get_photo, true))

	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
