package api

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/pkg/errors"
	"gitlab.com/nextwavedevs/drop/dal"
	"gitlab.com/nextwavedevs/drop/mid"
	"gitlab.com/nextwavedevs/drop/web"
	"go.mongodb.org/mongo-driver/mongo"
)

/*todos:*/
/* move the api part to another file*/
/*work on the other api handlers and profile methods*/

// Options represent optional parameters.
type Options struct {
	corsOrigin string
}

// WithCORS provides configuration options for CORS.
func WithCORS(origin string) func(opts *Options) {
	return func(opts *Options) {
		opts.corsOrigin = origin
	}
}

//This method is responsible for starting the api in the main method http.server for handler
func StartApi(build string, shutdown chan os.Signal, log *log.Logger, db *mongo.Client, options ...func(opts *Options)) http.Handler {

	var opts Options
	for _, option := range options {
		option(&opts)
	}

	// Construct the web.App which holds all routes as well as common Middleware.
	app := web.NewApp(shutdown, mid.Errors(log), mid.Logger(log))

	//Register the profile managment endpoints
	pg := profileGroup{
		profile: dal.New(log, db),
	}
	
	//Endpoins for profiles
	app.Handle(http.MethodPost, "/create/user", pg.CreateUserHandler)

	return app
}


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

// func GetStudioHandler(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	w.WriteHeader(http.StatusOK)
// 	fmt.Fprintf(w, "Category: %v\n", vars["studioId"])
// }

// func GetProfileHandler(w http.ResponseWriter, r *http.Request) {
// 	getProfile(w, r)
// }
