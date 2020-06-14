// calculate-league is a simple program to calculate a league table from
// a Google Sheets spreadsheet full of premier league football
// predictions made by my family.
package main

import (
	"fmt"
	"log"

	"github.com/tealeg/predict-sheet-control/client"
	"github.com/tealeg/predict-sheet-control/league"
	"github.com/tealeg/predict-sheet-control/sheet"
)

func main() {
	srv, err := client.MakeClient()
	if err != nil {
		log.Fatalf(err.Error())
	}
	data, err := sheet.GetPredictionData(srv)
	if err != nil {
		log.Fatalf(err.Error())
	}
	l := league.NewLeagueFromData(data)
	fmt.Println(l.String())
}
