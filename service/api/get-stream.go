package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/luigi-pizza/wasaPhoto/service/api/reqcontext"
)

func (rt *_router) get_stream(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// check params
	requestedPage, err := strconv.ParseUint(ps.ByName("page"), 10, 32)
	if err != nil {
		ctx.Logger.Error("get-stream: unable to parse parameter - 'page'")
		w.WriteHeader(http.StatusBadRequest)
		return
	}


	// Authorised -> get result
	photoList, err := rt.db.Select_stream(ctx.Uid, requestedPage)

	// check for DB errors
	if err != nil {
		ctx.Logger.WithError(err).Error("get-stream: error in DB - 'Select_stream(ctx.Uid, requestedPage)'")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// send result
	ctx.Logger.Debug("get-stream: 200")
	w.Header().Set("content-type", "application/json")
	_ = json.NewEncoder(w).Encode(photoList)
}
