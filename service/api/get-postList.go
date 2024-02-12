package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/luigi-pizza/wasaPhoto/service/api/reqcontext"
)

func (rt *_router) get_postList(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// check params
	requestedUser, err := strconv.ParseUint(ps.ByName("userId"), 10, 64)
	if err != nil {
		ctx.Logger.Error("get-postList: unable to parse parameter - 'userId'")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	requestedPageString := r.URL.Query().Get("page")
	var requestedPage uint64 = 0
	if requestedPageString != "" {
		requestedPage, err = strconv.ParseUint(requestedPageString, 10, 32)
		if err != nil {
			ctx.Logger.Error("get-postList: unable to parse parameter - 'page'")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	// check if requestedUser is a real userId
	isUID, err := rt.db.IsUserId(requestedUser)
	if err != nil {
		ctx.Logger.WithError(err).Error("get-postList: error in DB - 'IsUserId(requestedUserId)'")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if !isUID {
		ctx.Logger.WithField("requestedUserId", requestedUser).Error("get-postList: resource not found  - 'requestedUserId'")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// check if requestedUser has banned requestingUser
	isBanned, err := rt.db.IsBanned(requestedUser, ctx.Uid)
	if err != nil {
		ctx.Logger.WithError(err).Error("get-postList: error in DB - 'IsBanned(requestedUser, ctx.Uid)'")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if isBanned {
		ctx.Logger.Error("get-postList: forbidden action - 'ErrBannedRequestingUserError'")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// Authorised -> get result
	photoList, err := rt.db.Select_postList(ctx.Uid, requestedUser, requestedPage)
	photoList.PageNumber = requestedPage

	// check for DB errors
	if err != nil {
		ctx.Logger.WithError(err).Error("get-postList: error in DB - 'Select_postList(ctx.Uid, requestedUser, requestedPage)'")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// send result
	ctx.Logger.Debug("get-postList: 200")
	w.Header().Set("content-type", "application/json")
	_ = json.NewEncoder(w).Encode(photoList)
}
