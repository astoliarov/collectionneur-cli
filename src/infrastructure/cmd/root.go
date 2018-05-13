package cmd

import (
	"github.com/spf13/cobra"
	"collectionneur-cli/src/domain/usecases"
)



func InitCLI(loadAndSendSpendInfoUseCase *usecases.LoadAndSendSpendInfoUseCase) *cobra.Command {
	var RootCmd = &cobra.Command{
		Use:   "ccli",
		Short: "Collectionneur CLI",
		Long: `This application is used to interact with collectionneur service`,
	}

	uploadCmd := InitUploadCommand(loadAndSendSpendInfoUseCase)

	RootCmd.AddCommand(uploadCmd)
	return RootCmd
}
