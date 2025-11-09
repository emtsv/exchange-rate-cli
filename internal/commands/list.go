package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewListCMD() *cobra.Command {
	var listDate string
	listCmd := &cobra.Command{
		Use:   "list",
		Short: "Показать список всех валют за дату",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Printf("Список валют на дату %s (заглушка)\n", listDate)
			return nil
		},
	}
	listCmd.Flags().StringVarP(&listDate, "date", "d", "", "Дата в формате YYYY-MM-DD (опционально)")

	return listCmd
}
