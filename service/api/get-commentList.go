package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/luigi-pizza/wasaPhoto/service/api/reqcontext"
)

func (rt *_router) get_commentList(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// check params
	requestedPhoto, err := strconv.ParseUint(ps.ByName("postId"), 10, 64)
	if err != nil {
		ctx.Logger.Error("get-commentList: unable to parse parameter - 'postId'")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	requestedPageString := r.URL.Query().Get("page")
	var requestedPage uint64 = 0
	if requestedPageString != "" {
		requestedPage, err = strconv.ParseUint(requestedPageString, 10, 32)
		if err != nil {
			ctx.Logger.Error("get-commentList: unable to parse parameter - 'page'")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	// check if requestedPhoto is a real photoId
	isuid, requestedUser, err := rt.db.IsPhotoId(requestedPhoto)
	if err != nil {
		ctx.Logger.WithError(err).Error("get-commentList: error in DB - 'IsPhotoId(requestedPhotoId)'")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if !isuid {
		ctx.Logger.WithField("requestedPhotoId", requestedPhoto).Error("get-commentList: resource not found  - 'requestedPhotoId'")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// check if the author of the photo has banned requestingUser
	isBanned, err := rt.db.IsBanned(requestedUser, ctx.Uid)
	if err != nil {
		ctx.Logger.WithError(err).Error("get-commentList: error in DB - 'IsBanned(requestedUser, ctx.Uid)'")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if isBanned {
		ctx.Logger.Error("get-commentList: forbidden action - 'ErrBannedRequestingUserError'")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// Authorised -> get result
	commentList, err := rt.db.Select_commentList(ctx.Uid, requestedUser, requestedPage)
	commentList.PageNumber = requestedPage
	
	// check for DB errors
	if err != nil {
		ctx.Logger.WithError(err).Error("get-commentList: error in DB - 'Select_commentList(ctx.Uid, requestedUser, requestedPage)'")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// send result
	ctx.Logger.Debug("get-commentList: 200")
	w.Header().Set("content-type", "application/json")
	_ = json.NewEncoder(w).Encode(commentList)
}
