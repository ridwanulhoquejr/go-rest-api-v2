package main

import (
	"context"
	"fmt"

	"github.com/ridwanulhoquejr/go-rest-api-v2/cmd/internal/comment"
	"github.com/ridwanulhoquejr/go-rest-api-v2/cmd/internal/db"
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

	cmtService := comment.NewService(db)

	fmt.Println(cmtService.GetComment(context.Background(),
		"af7c1fe6-d669-414e-b066-e9733f0de7a8",
	))

	return nil
}

func main() {

	fmt.Println("Hello World!!")
	if err := Run(); err != nil {
		fmt.Println(err)
	}

}
