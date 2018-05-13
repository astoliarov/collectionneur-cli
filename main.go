package main

import (
	"collectionneur-cli/src/domain/usecases"
	"collectionneur-cli/src/infrastructure"
	"collectionneur-cli/src/infrastructure/dao"
	"collectionneur-cli/src/infrastructure/serviceapi"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func main() {

	config, err := infrastructure.ReadConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("sqlite3", config.Data.DBPath)
	if err != nil {
		log.Fatalf("Error, when try to open database: %s", err)
	}

	api, err := serviceapi.NewAPI(config.Server.AuthToken, config.Server.URL)
	if err != nil {
		log.Fatal(err)
	}

	d := dao.NewSpendInfoDAO(db, 4)

	usecase := usecases.NewLoadAndSendSpendInfoUseCase(api, d)
	count, err := usecase.Execute()
	if err != nil {
		log.Fatal(err)
	}

	log.Println(fmt.Sprintf("Added %d items", count))
}
