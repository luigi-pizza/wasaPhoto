package api

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/luigi-pizza/wasaPhoto/service/api/reqcontext"
)

func (rt *_router) delete_comment(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// check parameter
	requestedComment, err := strconv.ParseUint(ps.ByName("commentId"), 10, 64)
	if err != nil {
		ctx.Logger.Error("delete-comment: unable to parse parameter - 'commentId'")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// check if requestedComment is a real comment_id
	isCID, authorId, photoId, err := rt.db.IsCommentId(requestedComment)
	if err != nil {
		ctx.Logger.WithError(err).Error("delete-comment: error in DB - 'IsCommentId(requestedCommentId)'")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if !isCID {
		ctx.Logger.WithField("requestedCommentId", requestedComment).Error("delete-comment: resource not found  - 'requestedCommentId'")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// check requestingUser is the comment's author
	if authorId != ctx.Uid {
		ctx.Logger.Error("delete-comment: forbidden action - 'UnauthorisedDeletionOfResource'")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// Authorised -> get result
	err = rt.db.Delete_comment(requestedComment, photoId)

	// check for DB errors
	if err != nil {
		ctx.Logger.WithError(err).Error("delete-comment: error in DB - 'Delete_comment(requestedComment)'")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// send result
	ctx.Logger.Debug("delete-comment: 204")
	w.WriteHeader(http.StatusNoContent)
}
