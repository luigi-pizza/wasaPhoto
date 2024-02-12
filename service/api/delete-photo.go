package api

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/luigi-pizza/wasaPhoto/service/api/reqcontext"
	"github.com/luigi-pizza/wasaPhoto/service/filesystem"
)

func (rt *_router) delete_photo (w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// check params
	requestedPhoto, err := strconv.ParseUint(ps.ByName("postId"), 10, 64)
	if err != nil {
		ctx.Logger.Error("delete-photo: unable to parse parameter - 'postId'")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// check if requestedPhoto is a real photo_id
	isPID, authorId, err := rt.db.IsPhotoId(requestedPhoto)
	if err != nil {
		ctx.Logger.WithError(err).Error("delete-photo: error in DB - 'IsPhotoId(requestedPhotoId)'")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if !isPID {
		ctx.Logger.WithField("requestedPhotoId", requestedPhoto).Error("delete-photo: resource not found  - 'requestedPhotoId'")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// check requestingUser is the photo's author
	if authorId != ctx.Uid {
		ctx.Logger.Error("delete-photo: forbidden action - 'UnauthorisedDeletionOfResource'")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Authorised
	err = rt.db.Delete_photo(requestedPhoto)

	// check for DB errors
	if err != nil {
		ctx.Logger.WithError(err).Error("delete-photo: error in DB - 'Delete_photo(requestedPhoto)'")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// delete file
	err = filesystem.DeletePhoto(requestedPhoto)
	if err != nil {
		ctx.Logger.WithError(err).Error("delete-photo: error in filesystem - 'filesystem.DeletePhoto(requestedPhoto)'")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// send result
	ctx.Logger.Debug("delete-photo: 204")
	w.WriteHeader(http.StatusNoContent)
}