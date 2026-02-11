package commands

import (
	"fmt"

	"github.com/emtsv/exchange-rate-cli/internal/commands/timeutil"
	"github.com/spf13/cobra"
)

func NewRateCMD() *cobra.Command {
	var code string
	var date string
	rateCmd := &cobra.Command{
		Use:   "rate",
		Short: "Показать курс выбранной валюты (например, USD)",
		RunE: func(cmd *cobra.Command, args []string) error {
			if code == "" {
				return fmt.Errorf("нужно указать валюту через флаг --code, например: fx rate --code USD")
			}

			pdate, err := timeutil.ParseDate(date)
			if err != nil {
				return err
			}

			fmt.Printf("Курс валюты %s на дату %s (заглушка)\n", code, pdate)
			return nil
		},
	}

	rateCmd.Flags().StringVarP(&code, "code", "c", "", "Код валюты (например, USD, EUR)")
	rateCmd.Flags().StringVarP(&date, "date", "d", "", "Дата в формате YYYY-MM-DD (опционально)")

	return rateCmd
}
