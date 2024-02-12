package api

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/luigi-pizza/wasaPhoto/service/api/reqcontext"
	"github.com/luigi-pizza/wasaPhoto/service/components/requests"
	"github.com/luigi-pizza/wasaPhoto/service/components/schema"
)

func (rt *_router) put_username(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	var request requests.Username

	// decode json
	err := json.NewDecoder(r.Body).Decode(&request)
	_ = r.Body.Close()

	// check JSON err
	if err != nil {
		ctx.Logger.WithError(err).Error("put-username: error while decoding - JSON")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// check correctness of input
	if !request.IsValid() {
		ctx.Logger.WithField("username", request.Username).Error("put-username: username not valid - JSON")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// get result
	err = rt.db.Update_username(ctx.Uid, request.Username)

	// check for DB errors
	if errors.Is(err, schema.ErrUsernameAlreadyInUse) {
		ctx.Logger.WithField("username", request.Username).Error("put-username: forbidden Action - UsernameAlreadyInUse")
		w.WriteHeader(http.StatusForbidden)
		return
	}
	if err != nil {
		ctx.Logger.WithError(err).Error("put-username: error in DB - 'Update_username(ctx.Uid, request.Username)'")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// send result
	ctx.Logger.Debug("put-username: 204")
	w.WriteHeader(http.StatusNoContent)
}
