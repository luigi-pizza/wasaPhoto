package api

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/luigi-pizza/wasaPhoto/service/api/reqcontext"
	"github.com/luigi-pizza/wasaPhoto/service/components/requests"
)

func (rt *_router) get_userList(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	var request requests.UsernameSearch

	ctx.Logger.Debug("get-userList: decoding - JSON")
	err := json.NewDecoder(r.Body).Decode(&request)
	_ = r.Body.Close()

	if err != nil {
		ctx.Logger.WithError(err).Error("get-userList: error while decoding - JSON")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !request.IsValid() {
		ctx.Logger.WithField("username", request.Username).Error("get-userList: invalid resource - 'username'")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userList, err := rt.db.Select_userList(ctx.Uid, request.Username)

	if err != nil {
		ctx.Logger.WithError(err).Error("get-userList: error in DB - 'Select_userList(ctx.Uid, request.Username)'")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	ctx.Logger.Debug("get-userList: 200")
	w.Header().Set("content-type", "application/json")
	_ = json.NewEncoder(w).Encode(userList)
}
