package cmd

import (
	"collectionneur-cli/src/domain/usecases"
	"github.com/spf13/cobra"
	"log"
)

func InitUploadCommand(loadAndSendSpendInfoUseCase *usecases.LoadAndSendSpendInfoUseCase) *cobra.Command{
	uploadCmd := &cobra.Command{
		Use:   "upload",
		Short: "send info to collectionneur",
		Long:  `Collect spend info from IMessage DB and send it to collectionneur`,
		Run: func(cmd *cobra.Command, args []string) {
			_, err := loadAndSendSpendInfoUseCase.Execute()
			if err != nil {
				log.Fatal(err)
			}
		},
	}

	return uploadCmd
}
