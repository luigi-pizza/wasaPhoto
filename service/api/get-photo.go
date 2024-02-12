package api

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/luigi-pizza/wasaPhoto/service/api/reqcontext"
	"github.com/luigi-pizza/wasaPhoto/service/filesystem"
)

func (rt *_router) get_photo(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// check params
	requestedPhoto, err := strconv.ParseUint(ps.ByName("postId"), 10, 64)
	if err != nil {
		ctx.Logger.Error("get-photo: unable to parse parameter - 'postId'")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// check if requestedPhoto is a real photoId
	isuid, requestedUser, err := rt.db.IsPhotoId(requestedPhoto)
	if err != nil {
		ctx.Logger.WithError(err).Error("get-photo: error in DB - 'IsPhotoId(requestedPhotoId)'")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if !isuid {
		ctx.Logger.WithField("requestedPhotoId", requestedPhoto).Error("get-photo: resource not found  - 'requestedPhotoId'")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// check if the author of the photo has banned requestingUser
	isBanned, err := rt.db.IsBanned(requestedUser, ctx.Uid)
	if err != nil {
		ctx.Logger.WithError(err).Error("get-photo: error in DB - 'IsBanned(requestedUser, ctx.Uid)'")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if isBanned {
		ctx.Logger.Error("get-photo: forbidden action - 'ErrBannedRequestingUserError'")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// get photo from filesystem
	path := filesystem.FileSystemPath + strconv.FormatUint(requestedPhoto, 10) + ".png"
	photofile, err := os.Open(path)
	if err != nil {
		ctx.Logger.WithError(err).Error("get-photo: resouce not found - 'ErrPhotoNotPresentInFileSystem'")
		w.WriteHeader(http.StatusNotFound)
		return
	} else {
		w.Header().Set("Content-Type", "image/png")
		buf := bytes.NewBuffer(nil)
		_, err := io.Copy(buf, photofile)
		if err != nil {
			ctx.Logger.WithError(err).Error("get-photo: error copying in Buffer")
			w.WriteHeader(http.StatusInternalServerError)
			return
		} else {
			w.WriteHeader(http.StatusOK)
			_, err = w.Write(buf.Bytes())
			if err != nil {
				ctx.Logger.WithError(err).Error("get-photo: error writing response")
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			return
		}
	}
}
