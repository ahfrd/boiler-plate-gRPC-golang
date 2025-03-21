package main

import (
	"grpc-boiler-plate-go/cmd"
	"grpc-boiler-plate-go/env"
	"grpc-boiler-plate-go/infra/database"
	"log"
	"os"
)

func main() {
	di, err := env.NewENV("config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	if di.DB, di.Err = database.NewMySQLDB(di.Params); di.Err != nil {
		// Handle with middleware here upon error
		log.Fatal(di.Err)
	}

	// if di.ScyllaDb, di.Err = database.NewScyllaDB(di.Params); di.Err != nil {
	// 	log.Fatal(di.Err)
	// }

	app := cmd.NewCLI(di, os.Args)

	if app.Start(); app.Error() != nil {
		log.Fatal(app.Error())
	}
}
