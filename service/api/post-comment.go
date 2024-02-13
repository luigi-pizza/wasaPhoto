package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/luigi-pizza/wasaPhoto/service/api/reqcontext"
	"github.com/luigi-pizza/wasaPhoto/service/components/requests"
	"github.com/luigi-pizza/wasaPhoto/service/components/schema"
	"github.com/luigi-pizza/wasaPhoto/service/globaltime"
)

func (rt *_router) post_comment(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// check parameter
	requestedPhoto, err := strconv.ParseUint(ps.ByName("postId"), 10, 64)
	if err != nil {
		ctx.Logger.Error("post-comment: unable to parse parameter - 'postId'")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var text requests.Text
	err = json.NewDecoder(r.Body).Decode(&text)
	if err != nil {
		ctx.Logger.WithError(err).Error("post-comment: error while decoding - JSON")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	_ = r.Body.Close()

	// check if requestedPhoto is a real photo_id
	isPID, requestedUser, err := rt.db.IsPhotoId(requestedPhoto)
	if err != nil {
		ctx.Logger.WithError(err).Error("post-comment: error in DB - 'IsPhotoId(requestedPhotoId)'")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if !isPID {
		ctx.Logger.WithField("requestedPhotoId", requestedPhoto).Error("post-comment: resource not found  - 'requestedPhotoId'")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// check if requestedUser has banned requestingUser
	isBanned, err := rt.db.IsBanned(requestedUser, ctx.Uid)
	if err != nil {
		ctx.Logger.WithError(err).Error("post-comment: error in DB - 'IsBanned(requestedUser, ctx.Uid)'")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if isBanned {
		ctx.Logger.Error("post-comment: forbidden action - 'ErrBannedRequestingUserError'")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// Authorised -> get result

	author, err := rt.db.Select_reducedUser(ctx.Uid)
	if err != nil {
		ctx.Logger.WithError(err).Error("post-comment: error in DB - 'Select_reducedUser(ctx.Uid)'")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	now := globaltime.UnixNow()
	comment_id, err := rt.db.Insert_comment(requestedPhoto, ctx.Uid, text.Text, now)
	if err != nil {
		ctx.Logger.WithError(err).Error("post-comment: error in DB - 'Insert_comment(requestedPhoto, ctx.Uid, text.Text, now)'")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	comment := schema.Comment{
		Id:             comment_id,
		Author:         author,
		PhotoId:        requestedPhoto,
		CommentText:    text.Text,
		TimeOfCreation: now,
	}

	// send result
	ctx.Logger.Debug("post-comment: 201")
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(comment)
}
