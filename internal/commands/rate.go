package commands

import (
	"fmt"

	"github.com/emtsv/exchange-rate-cli/internal/cbr"
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
				return fmt.Errorf("неправильная дата: %s", err)
			}

			data, err := cbr.ParseValues(pdate)
			if err != nil {
				return fmt.Errorf("ошибка загрузки курсов: %w", err)
			}

			var found bool
			for _, v := range data.Valutes {
				if v.CharCode == code {
					rate, err := v.RateRUB()
					if err != nil {
						return fmt.Errorf("не удалось рассчитать курс: %w", err)
					}
					fmt.Printf("Курс валюты %s на дату %s: %.3f RUB\n", code, pdate, rate)
					found = true
					break
				}
			}

			if !found {
				return fmt.Errorf("валюта с кодом %s не найдена", code)
			}

			return nil
		},
	}

	rateCmd.Flags().StringVarP(&code, "code", "c", "", "Код валюты (например, USD, EUR)")
	rateCmd.Flags().StringVarP(&date, "date", "d", "", "Дата в формате YYYY-MM-DD (опционально)")

	return rateCmd
}
