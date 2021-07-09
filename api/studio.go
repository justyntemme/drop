package api

import (
	"context"
	"net/http"

	"github.com/pkg/errors"
	"gitlab.com/nextwavedevs/drop/dal"
	"gitlab.com/nextwavedevs/drop/web"
)

func (pg profileGroup) GetAllListingsByCompanyIdHandler(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	//Get the parameters coming from the request
	v, ok := ctx.Value(web.KeyValues).(*web.Values)
	if !ok {
		return web.NewShutdownError("web value missing from context")
	}
	//take the id part from the request
	uid := r.URL.Query().Get("id")
	//Return 400 if no uid is provided
	if uid == "" {
		return web.Respond(ctx, w, nil, http.StatusBadRequest)
	}

	//usr, err := pg.profile.GetUserById(ctx, v.TraceID, uid)
	usr, err := pg.profile.GetAllListingsByCompanyId(ctx, v.TraceID, uid)
	if err != nil {
		switch errors.Cause(err) {
		case dal.ErrInvalidID:
			return dal.NewRequestError(err, http.StatusBadRequest)
		case dal.ErrNotFound:
			return dal.NewRequestError(err, http.StatusNotFound)
		default:
			return errors.Wrapf(err, "ID: %s", uid)
		}
	}
	return web.Respond(ctx, w, usr, http.StatusOK)

}
