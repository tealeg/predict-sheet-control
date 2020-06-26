package league

import (
	"fmt"
	"sort"

	"github.com/tealeg/predict-sheet-control/result"
)

func PredictedLeaguesFromData(data [][]interface{}) {

	type score struct {
		team         string
		score        int
		played       int
		goalsFor     int
		goalsAgainst int
		goalDiff     int
	}

	type table map[string]score

	tables := map[string]table{}
	for _, row := range data {
		if len(row) > 0 {
			homeTeam := row[HomeTeamCol].(string)
			awayTeam := row[AwayTeamCol].(string)
			if homeTeam == "" || awayTeam == "" {
				// Skip empty rows
				continue
			}
			// Check we really got a result
			_, err := result.MakeResultFromStartCol(row, HomeScoreCol)
			if err != nil {
				// No result has been recorded, skip this one
				continue
			}

			for user, homePredCol := range Users {
				teams, ok := tables[user]
				if !ok {
					teams = make(table)
				}
				pred, err := result.MakeResultFromStartCol(row, homePredCol)
				if err != nil {
					// Assume anything not predicted is a 0-0 draw
					pred = result.MakeResult(0, 0)
				}

				homeScore := teams[homeTeam]
				awayScore := teams[awayTeam]
				homeScore.team = homeTeam
				awayScore.team = awayTeam
				homeScore.played++
				awayScore.played++
				homeScore.goalsFor += pred.Home
				homeScore.goalsAgainst += pred.Away
				homeScore.goalDiff = homeScore.goalsFor - homeScore.goalsAgainst
				awayScore.goalsFor += pred.Away
				awayScore.goalsAgainst += pred.Home
				awayScore.goalDiff = awayScore.goalsFor - awayScore.goalsAgainst

				switch pred.Type {
				case result.HomeWin:
					homeScore.score += 3
				case result.Draw:
					homeScore.score++
					awayScore.score++
				case result.AwayWin:
					awayScore.score += 3
				}
				teams[homeTeam] = homeScore
				teams[awayTeam] = awayScore
				tables[user] = teams
			}
		}
	}

	fmt.Println("\\newpage")
	for user, tbl := range tables {
		fmt.Printf("* %s's Predicted Leauge\n", user)
		var ts []score
		for _, s := range tbl {
			ts = append(ts, s)
		}

		sort.Slice(ts, func(i, j int) bool {
			if ts[i].score == ts[j].score {
				// In the event of even scores, order by goal differente
				if ts[i].goalDiff == ts[j].goalDiff {
					// in the event that goal difference is the same, order by goals scored
					if ts[i].goalsFor == ts[j].goalsFor {
						// If all else fails, alphabetical order is used.
						return ts[i].team < ts[j].team
					}
					return ts[i].goalsFor > ts[j].goalsFor
				}
				return ts[i].goalDiff > ts[j].goalDiff
			}
			return ts[i].score > ts[j].score
		})

		fmt.Printf("|*%s*|*%s*|*%s*|*%s*|*%s*|*%s*|*%s*|\n", "Pos", "Team", "Points", "Played", "For", "Against", "Diff")
		fmt.Println("|--|--|--|--|--|--|--|")
		for i, t := range ts {
			fmt.Printf("|%d|%s|%d|%d|%d|%d|%d|\n", i+1, t.team, t.score, t.played, t.goalsFor, t.goalsAgainst, t.goalDiff)
		}
		fmt.Println("\\newpage")
	}
}
