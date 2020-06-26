package main

import (
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
	league.PredictedLeaguesFromData(data)
}
