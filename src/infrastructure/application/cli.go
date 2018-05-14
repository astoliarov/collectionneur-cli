package application

import (
	"collectionneur-cli/src/domain/interfaces"
	"collectionneur-cli/src/domain/usecases"
	"collectionneur-cli/src/infrastructure/cmd"
	"collectionneur-cli/src/infrastructure/config"
	"collectionneur-cli/src/infrastructure/dao"
	"collectionneur-cli/src/infrastructure/serviceapi"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
	"time"
)

type CLIApp struct {
	config *config.Config

	loadAndSendSpendInfoUseCase *usecases.LoadAndSendSpendInfoUseCase
	spendInfoDAO                interfaces.ISpendInfoDAO
	api                         interfaces.IAPI

	rootCmd *cobra.Command
}

func (a *CLIApp) Run() error {
	return a.rootCmd.Execute()
}

func NewCLIApp() (*CLIApp, error) {
	config, err := config.ReadConfig()
	if err != nil {
		return nil, err
	}

	location, err := time.LoadLocation("Local")
	if err != nil {
		return nil, err
	}

	app := &CLIApp{}

	db, err := sql.Open("sqlite3", config.Data.DBPath)
	if err != nil {
		return nil, err
	}

	api, err := serviceapi.NewAPI(config.Server.AuthToken, config.Server.URL)
	if err != nil {
		return nil, err
	}

	spendInfoDAO := dao.NewSpendInfoDAO(db, 4)
	loadAndSendSpendInfoUseCase := usecases.NewLoadAndSendSpendInfoUseCase(
		api,
		spendInfoDAO,
		location,
	)

	app.spendInfoDAO = spendInfoDAO
	app.api = api
	app.loadAndSendSpendInfoUseCase = loadAndSendSpendInfoUseCase

	app.rootCmd = cmd.InitCLI(app.loadAndSendSpendInfoUseCase)
	return app, nil
}
