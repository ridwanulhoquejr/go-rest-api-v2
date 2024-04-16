package main

import (
	"fmt"

	"github.com/ridwanulhoquejr/go-rest-api-v2/cmd/internal/comment"
	"github.com/ridwanulhoquejr/go-rest-api-v2/cmd/internal/db"
	transportHttp "github.com/ridwanulhoquejr/go-rest-api-v2/cmd/internal/transport/http"
)

// RUN - is going to be responsible for
// the instantiation and startup of our
// GO Application
func Run() error {
	fmt.Println("starting up our application")

	db, err := db.NewDatabase()

	if err != nil {
		fmt.Println("failed to connect to the database")
		return err
	}

	if err := db.MigrateDB(); err != nil {
		fmt.Println("failed to migrate the database")
		return err
	}

	// creating a new instance of the comment service
	//? we are passing the repository as a dependency to the service
	//? so that the service can use the repository to interact with the database
	cmtService := comment.NewService(db)

	// entry point for our http server route handling
	httpHandler := transportHttp.NewHandler(cmtService)
	if err := httpHandler.Serve(); err != nil {
		fmt.Println("failed to start the server")
		return err
	}

	return nil
}

func main() {

	fmt.Println("Hello World!!")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
