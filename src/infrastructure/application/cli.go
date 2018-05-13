package application

import (
	"collectionneur-cli/src/domain/usecases"
	"collectionneur-cli/src/domain/interfaces"
	"collectionneur-cli/src/infrastructure/config"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"collectionneur-cli/src/infrastructure/serviceapi"
	"collectionneur-cli/src/infrastructure/dao"
	"github.com/spf13/cobra"
	"collectionneur-cli/src/infrastructure/cmd"
)

type CLIApp struct {
	config *config.Config

	loadAndSendSpendInfoUseCase *usecases.LoadAndSendSpendInfoUseCase
	spendInfoDAO interfaces.ISpendInfoDAO
	api interfaces.IAPI

	rootCmd *cobra.Command
}

func (a *CLIApp) Run() error{
	return a.rootCmd.Execute()
}

func NewCLIApp() (*CLIApp, error) {
	config, err := config.ReadConfig()
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
	)

	app.spendInfoDAO = spendInfoDAO
	app.api = api
	app.loadAndSendSpendInfoUseCase = loadAndSendSpendInfoUseCase

	app.rootCmd = cmd.InitCLI(app.loadAndSendSpendInfoUseCase)
	return app, nil
}
