package league

import (
	"testing"

	qt "github.com/frankban/quicktest"
	"github.com/tealeg/predict-sheet-control/result"
)

func TestLeague(t *testing.T) {
	c := qt.New(t)

	var data = [][]interface{}{
		[]interface{}{"2020-03-07", "12:30:00", "Liverpool FC", "AFC Bournemouth", "2", "1", "2", "1", "40", "2", "1", "40", "3", "0", "10", "", "", "", "2", "1", "40", "3", "1", "10", "0", "5", "0", "2", "0", "10", "1", "2", "0", "2", "0", "10"},
		[]interface{}{"2020-03-07", "15:00:00", "Arsenal FC", "West Ham United FC", "1", "0", "2", "1", "10", "2", "1", "10", "3", "1", "10", "", "", "", "1", "1", "0", "2", "1", "10", "3", "2", "10", "2", "1", "10", "1", "1", "0", "2", "0", "10"},
		[]interface{}{"2020-03-07", "15:00:00", "Crystal Palace FC", "Watford FC", "1", "0", "1", "1", "0", "1", "1", "0", "0", "0", "0", "", "", "", "2", "1", "10", "1", "2", "0", "0", "3", "0", "2", "1", "10", "1", "1", "0", "1", "0", "40"},
		[]interface{}{"2020-03-07", "15:00:00", "Sheffield United FC", "Norwich City FC", "1", "0", "1", "1", "0", "1", "2", "0", "2", "1", "10", "", "", "", "3", "1", "10", "2", "3", "0", "4", "0", "10", "2", "0", "10", "2", "2", "0", "2", "2", "0"},
		[]interface{}{"2020-03-07", "15:00:00", "Southampton FC", "Newcastle United FC", "0", "1", "2", "2", "0", "1", "1", "0", "1", "2", "10", "", "", "", "3", "1", "0", "2", "2", "0", "3", "0", "0", "2", "1", "0", "2", "1", "0", "0", "2", "10"},
		[]interface{}{"2020-03-07", "15:00:00", "Wolverhampton Wanderers FC", "Brighton & Hove Albion FC", "0", "0", "2", "1", "0", "2", "1", "0", "2", "2", "10", "", "", "", "2", "0", "0", "3", "0", "0", "2", "0", "0", "1", "1", "10", "2", "1", "0", "1", "2", "0"},
		[]interface{}{"2020-03-07", "17:30:00", "Burnley FC", "Tottenham Hotspur FC", "1", "1", "2", "1", "0", "1", "2", "0", "2", "1", "0", "", "", "", "1", "1", "40", "2", "0", "0", "3", "0", "0", "1", "2", "0", "1", "2", "0", "1", "2", "0"},
		[]interface{}{"2020-03-08", "14:00:00", "Chelsea FC", "Everton FC", "4", "0", "2", "1", "10", "2", "2", "0", "1", "1", "0", "", "", "", "2", "2", "0", "2", "3", "0", "0", "4", "0", "1", "1", "0", "1", "2", "0", "3", "1", "10"},
		[]interface{}{"2020-03-08", "16:30:00", "Manchester United FC", "Manchester City FC", "2", "0", "2", "2", "0", "1", "3", "0", "2", "1", "10", "", "", "", "2", "3", "0", "2", "4", "0", "2", "1", "10", "1", "2", "0", "1", "1", "0", "3", "2", "10"},
		[]interface{}{"2020-03-09", "20:00:00", "Leicester City FC", "Aston Villa FC", "4", "0", "2", "1", "10", "1", "0", "10", "4", "1", "10", "", "", "", "2", "0", "10", "1", "1", "0", "5", "0", "10", "2", "1", "10", "2", "1", "10", "2", "1", "10"},
		[]interface{}{"", "", "", "", "", "", "", "", "70", "", "", "50", "", "", "70", "", "", "", "", "", "110", "", "", "30", "", "", "40", "", "", "60", "", "", "10", "", "", "100"},
		[]interface{}{"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", ""},
		[]interface{}{"2020-06-17", "18:00:00", "Aston Villa FC", "Sheffield United FC", "", "", "1", "1", "", "1", "1", "", "", "", "", "", "", "", "", "", "", "", "", "", "0", "4", "", "", "", "", "", "", "", "", "", ""},
		[]interface{}{"2020-06-17", "20:15:00", "Manchester City FC", "Arsenal FC", "", "", "2", "1", "", "3", "2", "", "", "", "", "", "", "", "", "", "", "", "", "", "0", "n2", "", "", "", "", "", "", "", "", "", ""},
		[]interface{}{"2020-06-19", "18:00:00", "Norwich City FC", "Southampton FC", "", "", "1", "2", "", "1", "1", "", "", "", "", "", "", "", "", "", "", "", "", "", "3", "3", "", "", "", "", "", "", "", "", "", ""},
		[]interface{}{"2020-06-19", "20:15:00", "Tottenham Hotspur FC", "Manchester United FC", "", "", "2", "1", "", "2", "2", "", "", "", "", "", "", "", "", "", "", "", "", "", "0", "1", "", "", "", "", "", "", "", "", "", ""},
		[]interface{}{"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", ""}}

	c.Run("NewLeagueFromData", func(c *qt.C) {
		l := NewLeagueFromData(data)
		score, err := l.Score("Abi")
		c.Assert(err, qt.Equals, nil)
		c.Assert(score, qt.Equals, 60)
		score, err = l.Score("Ann")
		c.Assert(err, qt.Equals, nil)
		c.Assert(score, qt.Equals, 100)
		score, err = l.Score("Dad")
		c.Assert(err, qt.Equals, nil)
		c.Assert(score, qt.Equals, 70)
		score, err = l.Score("Geoff")
		c.Assert(err, qt.Equals, nil)
		c.Assert(score, qt.Equals, 70)
		score, err = l.Score("Katy")
		c.Assert(err, qt.Equals, nil)
		c.Assert(score, qt.Equals, 20)
		score, err = l.Score("Louise")
		c.Assert(err, qt.Equals, nil)
		c.Assert(score, qt.Equals, 60)
		score, err = l.Score("Olly")
		c.Assert(err, qt.Equals, nil)
		c.Assert(score, qt.Equals, 40)
		score, err = l.Score("Steve")
		c.Assert(err, qt.Equals, nil)
		c.Assert(score, qt.Equals, 110)
		score, err = l.Score("Will")
		c.Assert(err, qt.Equals, nil)
		c.Assert(score, qt.Equals, 10)
	})

	c.Run("String", func(c *qt.C) {
		l := NewLeagueFromData(data)
		c.Assert(l.String(), qt.Equals, `Steve, 110
Ann, 100
Dad, 70
Geoff, 70
Abi, 60
Louise, 60
Olly, 40
Katy, 20
Will, 10
Joe, 0
`)

	})

	c.Run("AddPoints", func(c *qt.C) {
		l := NewLeague()
		for user, _ := range Users {
			c.Run(user, func(c *qt.C) {
				score, err := l.Score(user)
				c.Assert(err, qt.Equals, nil)
				c.Assert(score, qt.Equals, 0)
			})
		}
	})
}

func TestCalculatePoints(t *testing.T) {
	c := qt.New(t)

	res := result.MakeResult(1, 0)
	pred := result.MakeResult(1, 0)
	c.Assert(calculatePoints(res, pred), qt.Equals, 40)

	pred = result.MakeResult(3, 2)
	c.Assert(calculatePoints(res, pred), qt.Equals, 10)

	pred = result.MakeResult(0, 1)
	c.Assert(calculatePoints(res, pred), qt.Equals, 0)

}
