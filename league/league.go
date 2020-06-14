package league

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/tealeg/predict-sheet-control/result"
)

const (
	HomeTeamCol  int = 2
	AwayTeamCol  int = 3
	HomeScoreCol int = 4
	AwayScoreCol int = 5
)

var (
	// Users maps usernames to the first colum of their predictions
	Users = map[string]int{
		"Geoff":  6,
		"Louise": 9,
		"Dad":    12,
		"Joe":    15,
		"Steve":  18,
		"Katy":   21,
		"Olly":   24,
		"Abi":    27,
		"Will":   30,
		"Ann":    33,
	}
)

type League struct {
	standings                  map[string]int
	shouldLogMissingResult     bool
	shouldLogMissedPredictions bool
}

type LeagueOption func(l *League)

func LogMissingResult(l *League) {
	l.shouldLogMissingResult = true
}
func LogMissedPredictions(l *League) {
	l.shouldLogMissedPredictions = true
}

func NewLeague() *League {
	l := &League{
		standings: make(map[string]int),
	}

	// Init user scores
	for user, _ := range Users {
		l.AddPoints(user, 0)
	}
	return l
}

func NewLeagueFromData(data [][]interface{}, options ...LeagueOption) *League {
	l := NewLeague()
	for rowNum, row := range data {
		if len(row) > 0 {
			homeTeam := row[HomeTeamCol].(string)
			awayTeam := row[AwayTeamCol].(string)
			if homeTeam == "" || awayTeam == "" {
				// Skip empty rows
				continue
			}
			res, err := result.MakeResultFromStartCol(row, HomeScoreCol)
			if err != nil {
				if l.shouldLogMissingResult {
					fmt.Fprintf(os.Stderr, "Row %d: %s\n", rowNum, err)
					fmt.Printf("Row %d: No result recorded for %s v %s, assigning zero points to everyone\n", rowNum, homeTeam, awayTeam)
				}
				continue
			}

			for user, homePredCol := range Users {
				pred, err := result.MakeResultFromStartCol(row, homePredCol)
				if err != nil {
					if l.shouldLogMissedPredictions {
						fmt.Fprintf(os.Stderr, "Row %d: Prediction User %s:  %s\n", rowNum, user, err)
						fmt.Printf("No prediction for %s for  %s v %s, zero points assigned\n", user, homeTeam, awayTeam)
					}
					continue
				}

				points := calculatePoints(res, pred)
				l.AddPoints(user, points)
			}
		}
	}
	return l
}

func (l *League) Score(user string) (int, error) {
	score, found := l.standings[user]
	if !found {
		return -1, fmt.Errorf("No score found for user %s", user)
	}
	return score, nil
}

func (l *League) String() string {
	type kv struct {
		Key   string
		Value int
	}

	var ss []kv
	for k, v := range l.standings {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		if ss[i].Value == ss[j].Value {
			// In the event of even scores, order alphabetically by name
			return ss[i].Key < ss[j].Key
		}
		return ss[i].Value > ss[j].Value
	})

	var result strings.Builder
	for _, kv := range ss {
		result.WriteString(fmt.Sprintf("%s, %d\n", kv.Key, kv.Value))
	}
	return result.String()
}

func (l *League) AddPoints(user string, points int) {
	total, ok := l.standings[user]
	if !ok {
		total = 0
	}
	total += points
	l.standings[user] = total
}

func calculatePoints(res, pred *result.Result) int {
	if pred.Type == res.Type {
		if pred.Home == res.Home && pred.Away == res.Away {
			return 40
		}
		return 10
	}
	return 0
}
