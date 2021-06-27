// Package tests contains supporting code for running tests.
package dal

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"testing"

	"gitlab.com/nextwavedevs/drop/database"
	"gitlab.com/nextwavedevs/drop/scripts"
	"go.mongodb.org/mongo-driver/mongo"
)

// Success and failure markers.
const (
	Success = "\u2713"
	Failed  = "\u2717"
)

// DBContainer provides configuration for a container to run.
type DBContainer struct {
	Image string
	Port  string
	Args  []string
}

// NewUnit creates a test database inside a Docker container. It creates the
// required table structure but the database is otherwise empty. It returns
// the database to use as well as a function to call at the end of the test.
func NewUnit(t *testing.T, dbc DBContainer) (*log.Logger, *mongo.Client, func()) {
	r, w, _ := os.Pipe()
	old := os.Stdout
	os.Stdout = w

	c := scripts.StartContainer(t, dbc.Image, dbc.Port, dbc.Args...)

	t.Log("Waiting for database to be ready ...")

	log := log.New(nil, "Test", 1)

	db := database.Client

	// teardown is the function that should be invoked when the caller is done
	// with the database.
	teardown := func() {
		t.Helper()
		scripts.StopContainer(t, c.ID)

		w.Close()
		var buf bytes.Buffer
		io.Copy(&buf, r)
		os.Stdout = old
		fmt.Println("******************** LOGS ********************")
		fmt.Print(buf.String())
		fmt.Println("******************** LOGS ********************")
	}

	return log, db, teardown
}