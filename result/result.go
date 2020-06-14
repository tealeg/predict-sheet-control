package result

import (
	"fmt"

	"github.com/tealeg/predict-sheet-control/sheet"
)

type ResultType int

const (
	HomeWin ResultType = iota
	Draw
	AwayWin
)

type Result struct {
	Home int
	Away int
	Type ResultType
}

func MakeResult(home, away int) *Result {
	res := &Result{Home: home, Away: away, Type: Draw}
	if res.Home > res.Away {
		res.Type = HomeWin
	}
	if res.Home < res.Away {
		res.Type = AwayWin
	}
	return res
}

func MakeResultFromStartCol(row []interface{}, col int) (*Result, error) {
	homeScore, err := sheet.IntFromCell(row, col)
	if err != nil {
		return nil, fmt.Errorf("HomeScore:  %w", err)
	}
	awayScore, err := sheet.IntFromCell(row, col+1)
	if err != nil {
		return nil, fmt.Errorf("AwayScore: %w", err)
	}
	return MakeResult(homeScore, awayScore), nil
}
