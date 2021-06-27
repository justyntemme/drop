package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/pkg/errors"
	"gitlab.com/nextwavedevs/drop/dal"
	"gitlab.com/nextwavedevs/drop/web"
)

//Calling the instance of the Profile struct holding all the profile methods
type profileGroup struct {
	profile dal.Profile
}

//profileAPI calling the createUser/Profile profile method
func (pg profileGroup) CreateUserHandler(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	log.Print("Create Profile Endpoint Hit")
	v, ok := ctx.Value(web.KeyValues).(*web.Values)
	if !ok {
		return web.NewShutdownError("web value missing from context")
	}

	var u dal.User
	//Decode user input
	if err := web.Decode(r, &u); err != nil {
		return errors.Wrap(err, "unable to decode payload")
	}

	usr, err := pg.profile.CreateProfile(ctx, v.TraceID, u)
	if err != nil {
		return errors.Wrapf(err, "Profile: %+v", &usr)
	}

	//Respond to the client
	return web.Respond(ctx, w, usr, http.StatusCreated)
}

// GetProfileByID gets the specified user.
func (pg profileGroup) GetProfileById(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

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

	usr, err := pg.profile.GetUserById(ctx, v.TraceID, uid)
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

// GetAllProfile retrieves a list of existing users.
func (pg profileGroup) GetAllProfile(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	log.Print("GetAll Profile Endpoint Hit")
	v, ok := ctx.Value(web.KeyValues).(*web.Values)
	if !ok {
		return web.NewShutdownError("web value missing from context")
	}

	page := web.Param(r, "page")
	pageNumber, err := strconv.Atoi(page)
	if err != nil {
		return dal.NewRequestError(fmt.Errorf("invalid page format: %s", page), http.StatusBadRequest)
	}
	rows := web.Param(r, "rows")
	rowsPerPage, err := strconv.Atoi(rows)
	if err != nil {
		return dal.NewRequestError(fmt.Errorf("invalid rows format: %s", rows), http.StatusBadRequest)
	}

	products, err := pg.profile.GetAllUsers(ctx, v.TraceID, pageNumber, rowsPerPage)
	if err != nil {
		return errors.Wrap(err, "unable to query for profile")
	}

	return web.Respond(ctx, w, products, http.StatusOK)
}

// Update replaces a user document in the database.
func (pg profileGroup) UpdateProfile(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	v, ok := ctx.Value(web.KeyValues).(*web.Values)
	if !ok {
		return web.NewShutdownError("web value missing from context")
	}

	var upd dal.User
	if err := web.Decode(r, &upd); err != nil {
		return errors.Wrapf(err, "unable to decode payload")
	}

	id := web.Param(r, "id")
	if err := pg.profile.UpdateProfile(ctx, v.TraceID, id, upd); err != nil {
		switch errors.Cause(err) {
		case dal.ErrInvalidID:
			return dal.NewRequestError(err, http.StatusBadRequest)
		case dal.ErrNotFound:
			return dal.NewRequestError(err, http.StatusNotFound)
		case dal.ErrForbidden:
			return dal.NewRequestError(err, http.StatusForbidden)
		default:
			return errors.Wrapf(err, "ID: %s  User: %+v", id, &upd)
		}
	}

	return web.Respond(ctx, w, nil, http.StatusNoContent)
}

// Delete removes a user from the database.
func (pg profileGroup) deleteUser(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

	v, ok := ctx.Value(web.KeyValues).(*web.Values)
	if !ok {
		return web.NewShutdownError("web value missing from context")
	}

	id := web.Param(r, "id")
	if err := pg.profile.DeleteProfile(ctx, v.TraceID, id); err != nil {
		switch errors.Cause(err) {
		case dal.ErrInvalidID:
			return dal.NewRequestError(err, http.StatusBadRequest)
		default:
			return errors.Wrapf(err, "ID: %s", id)
		}
	}

	return web.Respond(ctx, w, nil, http.StatusNoContent)
}
