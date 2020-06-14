package sheet

import (
	"fmt"
	"log"
	"strconv"

	"google.golang.org/api/sheets/v4"
)

func GetPredictionData(srv *sheets.Service) ([][]interface{}, error) {
	sheet, err := getPredictionsSheet(srv)
	if err != nil {
		log.Fatalf("Unable to retrieve Sheet")
	}

	rowCount := sheet.Properties.GridProperties.RowCount

	resp, err := srv.Spreadsheets.Values.Get("1gJlMak-dmc1LSbbNKlLf7jm8WHk-iXyccLV6CqGaTB8", fmt.Sprintf("A3:AJ%d", rowCount)).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet: %w", err)
	}

	if len(resp.Values) == 0 {
		return resp.Values, fmt.Errorf("No data returned")
	}
	return resp.Values, nil
}

func IntFromCell(row []interface{}, col int) (int, error) {
	if col >= len(row) || col < 0 {
		return -1, fmt.Errorf("intFromCell: col index %d out of range\n", col)
	}
	str, ok := row[col].(string)
	if !ok {
		return -1, fmt.Errorf("intFromCell: cannot type assert %v to string, it's a %T\n", row[col], row[col])
	}
	val, err := strconv.Atoi(str)
	if err != nil {
		return -1, fmt.Errorf("intFromCell: %w\n", err)
	}
	return val, nil
}

func getPredictionsSheet(srv *sheets.Service) (*sheets.Sheet, error) {
	spread, err := srv.Spreadsheets.Get("1gJlMak-dmc1LSbbNKlLf7jm8WHk-iXyccLV6CqGaTB8").Do()
	if err != nil {
		return nil, err
	}
	if len(spread.Sheets) == 0 {
		return nil, fmt.Errorf("No sheets in document")
	}
	return spread.Sheets[0], nil
}
