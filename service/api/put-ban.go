package api

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/luigi-pizza/wasaPhoto/service/api/reqcontext"
)

func (rt *_router) put_ban(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// check params
	requestedUser, err := strconv.ParseUint(ps.ByName("userId"), 10, 64)

	// check correctness of parameter
	if err != nil {
		ctx.Logger.Error("put-ban: unable to parse parameter - 'userId'")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// check if requestedUser is a real userID
	isuid, err := rt.db.IsUserId(requestedUser)
	if err != nil {
		ctx.Logger.WithError(err).Error("put-ban: error in DB - 'IsUserId(requestedUserId)'")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if !isuid {
		ctx.Logger.WithField("requestedUserId", requestedUser).Error("put-ban: resource not found  - 'requestedUserId'")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Authorised -> get result
	err = rt.db.Insert_ban(ctx.Uid, requestedUser)

	// check for DB errors
	if err != nil {
		ctx.Logger.WithError(err).Error("put-ban: error in DB - 'Insert_ban(ctx.Uid, requestedUser)'")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// send result
	ctx.Logger.Debug("put-ban: 204")
	w.WriteHeader(http.StatusNoContent)
}
