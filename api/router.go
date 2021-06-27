package api

import (
	"context"
	"log"
	"net/http"
	"os"

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
	app.Handle(http.MethodPost, "/v1/create", pg.CreateUserHandler)
	app.Handle(http.MethodGet, "/v1/getall/{page}/{rows}", pg.GetAllProfile)
	app.Handle(http.MethodGet, "/v1/get/{id}", pg.GetProfileById)
	app.Handle(http.MethodPut, "/v1/update/{id}", pg.UpdateProfile)
	app.Handle(http.MethodDelete, "/v1/delete/{id}", pg.deleteUser)

	// Accept CORS 'OPTIONS' preflight requests if config has been provided.
	// Don't forget to apply the CORS middleware to the routes that need it.
	// Example Config: `conf:"default:https://MY_DOMAIN.COM"`
	if opts.corsOrigin != "" {
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
			return nil
		}
		app.Handle(http.MethodOptions, "/*", h)
	}

	return app
}
