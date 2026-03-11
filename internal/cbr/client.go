package cbr

import (
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/emtsv/exchange-rate-cli/internal/commands/timeutil"
)

const baseURL = "https://www.cbr.ru/scripts/XML_daily.asp?date_req=%s"

func ParseValues(date string) (*ValCurs, error) {

	parsedDate, err := timeutil.ParseDate(date)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf(baseURL, parsedDate)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data ValCurs

	if err := xml.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return &data, nil
}
