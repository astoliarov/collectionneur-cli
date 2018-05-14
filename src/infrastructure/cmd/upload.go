package cmd

import (
	"collectionneur-cli/src/domain/usecases"
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

func InitUploadCommand(loadAndSendSpendInfoUseCase *usecases.LoadAndSendSpendInfoUseCase) *cobra.Command {
	uploadCmd := &cobra.Command{
		Use:   "upload",
		Short: "send info to collectionneur",
		Long:  `Collect spend info from IMessage DB and send it to collectionneur`,
		Run: func(cmd *cobra.Command, args []string) {
			count, err := loadAndSendSpendInfoUseCase.Execute()
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(fmt.Sprintf("Uploaded %d items", count))
		},
	}
	return uploadCmd
}
