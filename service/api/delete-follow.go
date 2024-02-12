package api

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/luigi-pizza/wasaPhoto/service/api/reqcontext"
)

func (rt *_router) delete_follow(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// check params
	requestedUser, err := strconv.ParseUint(ps.ByName("userId"), 10, 64)

	// check correctness of parameter
	if err != nil {
		ctx.Logger.Error("delete-follow: unable to parse parameter - 'userId'")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// check if requestedUser is a real userID
	isuid, err := rt.db.IsUserId(requestedUser)
	if err != nil {
		ctx.Logger.WithError(err).Error("delete-follow: error in DB - 'IsUserId(requestedUserId)'")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if !isuid {
		ctx.Logger.WithField("requestedUserId", requestedUser).Error("delete-follow: resource not found  - 'requestedUserId'")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// check if requestedUser has banned requestingUser
	isBanned, err := rt.db.IsBanned(requestedUser, ctx.Uid)
	if err != nil {
		ctx.Logger.WithError(err).Error("delete-follow: error in DB - 'IsBanned(requestedUser, ctx.Uid)'")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if isBanned {
		ctx.Logger.Error("delete-follow: forbidden action - 'ErrBannedRequestingUserError'")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// Authorised -> get result
	err = rt.db.Delete_follow(ctx.Uid, requestedUser)

	// check for DB errors
	if err != nil {
		ctx.Logger.WithError(err).Error("delete-follow: error in DB - 'Delete_follow(ctx.Uid, requestedUser)'")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// send result
	ctx.Logger.Debug("delete-follow: 204")
	w.WriteHeader(http.StatusNoContent)
}
