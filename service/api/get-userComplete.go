package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/luigi-pizza/wasaPhoto/service/api/reqcontext"
)

func (rt *_router) get_userComplete(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// check params
	requestedUser, err := strconv.ParseUint(ps.ByName("userId"), 10, 64)

	// check correctness of input
	if err != nil {
		ctx.Logger.WithField("passedVal", ps.ByName("userId")).Error("get-userComplete: unable to parse parameter - 'userId'")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// check if requestedUser is a real userID
	isuid, err := rt.db.IsUserId(requestedUser)
	if err != nil {
		ctx.Logger.WithError(err).Error("get-userComplete: error in DB - 'IsUserId(requestedUserId)'")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if !isuid {
		ctx.Logger.WithField("requestedUserId", requestedUser).Error("get-userComplete: resource not found  - 'requestedUserId'")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// check if requestedUser has banned requestingUser
	isBanned, err := rt.db.IsBanned(requestedUser, ctx.Uid)
	if err != nil {
		ctx.Logger.WithError(err).Error("get-userComplete: error in DB - 'IsBanned(requestedUser, ctx.Uid)'")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if isBanned {
		ctx.Logger.Error("get-userComplete: forbidden action - 'ErrBannedRequestingUserError'")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// Authorised -> get result
	userComplete, err := rt.db.Select_completeUser(ctx.Uid, requestedUser)

	// check for DB errors
	if err != nil {
		ctx.Logger.WithError(err).Error("get-userComplete: error in DB - 'Select_completeUser(ctx.Uid, requestedUser)'")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// send result
	ctx.Logger.Debug("get-userComplete: 200")
	w.Header().Set("content-type", "application/json")
	_ = json.NewEncoder(w).Encode(userComplete)
}
