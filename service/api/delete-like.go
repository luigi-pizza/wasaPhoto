package api

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/luigi-pizza/wasaPhoto/service/api/reqcontext"
)

func (rt *_router) delete_like(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// check params
	requestedPhoto, err := strconv.ParseUint(ps.ByName("postId"), 10, 64)

	// check correctness of parameter
	if err != nil {
		ctx.Logger.Error("delete-like: unable to parse parameter - 'postId'")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// check if requestedPhoto is a real photo_id
	isPID, requestedUser, err := rt.db.IsPhotoId(requestedPhoto)
	if err != nil {
		ctx.Logger.WithError(err).Error("delete-like: error in DB - 'IsPhotoId(requestedPhotoId)'")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if !isPID {
		ctx.Logger.WithField("requestedPhotoId", requestedPhoto).Error("delete-like: resource not found  - 'requestedPhotoId'")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// check if requestedUser has banned requestingUser
	isBanned, err := rt.db.IsBanned(requestedUser, ctx.Uid)
	if err != nil {
		ctx.Logger.WithError(err).Error("delete-like: error in DB - 'IsBanned(requestedUser, ctx.Uid)'")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if isBanned {
		ctx.Logger.Error("delete-like: forbidden action - 'ErrBannedRequestingUserError'")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// Authorised -> get result
	err = rt.db.Delete_like(ctx.Uid, requestedPhoto)

	// check for DB errors
	if err != nil {
		ctx.Logger.WithError(err).Error("delete-like: error in DB - 'Delete_like(ctx.Uid, requestedPhoto)'")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// send result
	ctx.Logger.Debug("delete-like: 204")
	w.WriteHeader(http.StatusNoContent)
}
