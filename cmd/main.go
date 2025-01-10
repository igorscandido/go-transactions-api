package main

import (
	"fmt"

	"github.com/igorscandido/go-transactions-api/pkg/configs"
	"github.com/igorscandido/go-transactions-api/pkg/database"
)

func main() {
	configs := configs.NewConfigs()

	postgresConnection, err := database.NewPostgresAdapter(configs)
	if err != nil {
		panic(err)
	}

	fmt.Println("API successfuly started!", postgresConnection)
}
