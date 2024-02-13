package api

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/luigi-pizza/wasaPhoto/service/api/reqcontext"
	"github.com/luigi-pizza/wasaPhoto/service/components/schema"
	"github.com/luigi-pizza/wasaPhoto/service/filesystem"
	"github.com/luigi-pizza/wasaPhoto/service/globaltime"
)

func (rt *_router) post_photo(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// parse request
	err := r.ParseMultipartForm(filesystem.MaxRequestSize)
	if err != nil {
		ctx.Logger.WithError(err).Error("post-photo: error while parsing multipart form")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// get contents of request
	caption := r.FormValue("caption")
	file, fileHeader, err := r.FormFile("photo")
	if err != nil {
		ctx.Logger.WithError(err).Error("post-photo: error while obtaining photo resource")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer file.Close()

	// validate request contents
	if fileHeader.Header.Get("Content-Type") != filesystem.AcceptedImageFormat {
		ctx.Logger.Error("post-photo: invalid request - Unsupported content Type")
		w.WriteHeader(http.StatusUnsupportedMediaType)
		return
	}
	if fileHeader.Size > filesystem.MaxPhotoSize {
		ctx.Logger.Error("post-photo: invalid request - PhotoSize too large")
		w.WriteHeader(http.StatusRequestEntityTooLarge)
		return
	}
	if fileHeader.Size < filesystem.MinPhotoSize {
		ctx.Logger.Error("post-photo: invalid request - PhotoSize too small")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if len(caption) > filesystem.MaxTextLength {
		ctx.Logger.Error("post-photo: invalid request - Caption too long")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// get info for insertion in db
	now := globaltime.UnixNow()
	author, err := rt.db.Select_reducedUser(ctx.Uid)
	if err != nil {
		ctx.Logger.WithError(err).Error("post-photo: error in DB - 'Select_reducedUser(ctx.Uid)'")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Save photo to DB
	photoId, err := rt.db.Insert_photo(ctx.Uid, caption, now)
	if err != nil {
		ctx.Logger.WithError(err).Error("post-photo: error in DB - 'Insert_photo(ctx.Uid, caption, now)'")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Save photo to file system
	err = filesystem.SavePhoto(file, photoId)
	if err != nil {
		err = rt.db.Delete_photo(photoId)
		if err != nil {
			ctx.Logger.WithError(err).WithField("photoId", photoId).Error("post-photo: error in filesystem-save, check photoId")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		ctx.Logger.WithError(err).WithField("photoId", photoId).Error("post-photo: error in filesystem-save, reverted")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Send response

	photo := schema.Post{
		Id:             photoId,
		Author:         author,
		Caption:        caption,
		Likes:          0,
		Comments:       0,
		TimeOfCreation: now,
		IsLiked:        false,
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(photo)
}
