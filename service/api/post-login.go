package api

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/luigi-pizza/wasaPhoto/service/api/reqcontext"
	"github.com/luigi-pizza/wasaPhoto/service/components/requests"
	"github.com/luigi-pizza/wasaPhoto/service/components/schema"
)

func (rt *_router) post_login(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	var request requests.Username

	// decode json
	err := json.NewDecoder(r.Body).Decode(&request)
	_ = r.Body.Close()

	// check JSON err
	if err != nil {
		ctx.Logger.WithError(err).Error("post-login: error while decoding - JSON")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// check correctness of input
	if !request.IsValid() {
		ctx.Logger.WithField("username", request.Username).Error("post-login: invalid resource - 'username'")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// get result
	uid, created, err := rt.db.Insert_user(request.Username)

	// check for DB errors
	if err != nil {
		ctx.Logger.WithError(err).Error("post-login: error in DB -'Insert_user(request.Username)'s ")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// send result
	ctx.Logger.Debug("post-login: OK")

	result := schema.ReducedUser{Id: uid, Username: request.Username}
	w.Header().Set("content-type", "application/json")
	if created { w.WriteHeader(http.StatusCreated) }

	_ = json.NewEncoder(w).Encode(result)
}
