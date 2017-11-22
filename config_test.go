package casino

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	tmpFileName := "./config_tmp.json"
	c, _, err := ParseCasinoConfig(tmpFileName)
	if err != nil {
		t.Errorf("parse casino config error: %v", err)
	}
	//check rows columns
	assert.Equal(t, 3, c.Rows, "rows not matched")
	assert.Equal(t, 3, c.Columns, "columns not matched")
	//check bets
	assert.Equal(t, 5000, c.ScoreBase, "score base not matched")
	assert.Equal(t, []int{1, 5, 10, 15, 20}, c.Bets, "bets not matched")
	//check lines config
	assert.Equal(t, 3, len(c.LinesConfig), "line config length not matched")
	assert.Equal(t, []int{1, 1, 1}, c.LinesConfig[0].Line, "line config one line data not matched")
	assert.Equal(t, []int{2, 2, 2}, c.LinesConfig[1].Line, "line config one line data not matched")
	assert.Equal(t, []int{3, 3, 3}, c.LinesConfig[2].Line, "line config one line data not matched")
	//check obtains config
	assert.Equal(t, []int{0, 500, 1000, 2000, 5000}, c.ObtainsConfig[0].Reward, "obtain config one obtain not matched")
	assert.Equal(t, []int{100, 500, 1000, 2000, 5000}, c.ObtainsConfig[1].Reward, "obtain config one obtain not matched")
	assert.Equal(t, []int{200, 500, 1000, 2000, 5000}, c.ObtainsConfig[2].Reward, "obtain config one obtain not matched")
	assert.Equal(t, []int{300, 500, 1000, 2000, 5000}, c.ObtainsConfig[3].Reward, "obtain config one obtain not matched")
	//check boards config
	assert.Equal(t, []int{6, 7, 8, 6, 7, 8, 6, 7, 8}, c.BoardsConfig[1].Slots, "boards config one board slots not matched")
	assert.Equal(t, []int{1, 2, 3, 1, 2, 3, 1, 2, 4}, c.BoardsConfig[2].Slots, "boards config one board slots not matched")
	assert.Equal(t, []int{1, 2, 3, 1, 2, 3, 1, 2, 4}, c.BoardsConfig[3].Slots, "boards config one board slots not matched")
	assert.Equal(t, []int{1, 2, 3, 1, 2, 3, 1, 2, 4}, c.BoardsConfig[4].Slots, "boards config one board slots not matched")
	assert.Equal(t, 90, c.BoardsConfig[1].Payout, "boards config one board payout not matched")
	assert.Equal(t, 80, c.BoardsConfig[2].Payout, "boards config one board payout not matched")
	assert.Equal(t, 200, c.BoardsConfig[3].Payout, "boards config one board payout not matched")
	assert.Equal(t, 150, c.BoardsConfig[4].Payout, "boards config one board payout not matched")
	//check gears config
	assert.Equal(t, []int{8, 5, 5, 0, 9, 3, 7, 14, 9, 0, 2, 9, 20, 6, 6, 8, 20, 8, 7, 7, 3, 5, 4, 4, 4, 4, 8, 5, 6, 9, 1, 1, 1, 1, 5, 9, 8, 7, 0, 9, 1, 9, 8, 2, 5, 8, 9, 8, 20, 9, 6, 1, 8, 2, 2, 2, 2, 3, 4, 9, 6, 5, 0, 0, 0, 0, 7, 6, 7, 9, 4, 3, 3, 3, 7, 4, 8, 3, 9, 8, 6, 8, 0, 7, 4, 20, 8, 4, 6, 3, 8, 2, 20, 9, 4, 6, 1, 7, 2, 1, 3, 9, 0, 7, 9, 5, 5, 4, 7, 7, 6}, c.GearsConfig[1].Symbols, "gears config one gear not matched")
	assert.Equal(t, []int{8, 5, 5, 0, 10, 3, 7, 14, 9, 0, 2, 9, 20, 6, 6, 8, 20, 8, 7, 7, 3, 5, 4, 4, 4, 4, 8, 5, 6, 9, 1, 1, 1, 1, 5, 9, 8, 7, 0, 9, 1, 9, 8, 2, 5, 8, 9, 8, 20, 9, 6, 1, 8, 2, 2, 2, 2, 3, 4, 9, 6, 5, 0, 0, 0, 0, 7, 6, 7, 9, 4, 3, 3, 3, 7, 4, 8, 3, 9, 8, 6, 8, 0, 7, 4, 20, 8, 4, 6, 3, 8, 2, 20, 9, 4, 6, 1, 7, 2, 1, 3, 9, 0, 7, 9, 5, 5, 4, 7, 7, 6}, c.GearsConfig[2].Symbols, "gears config one gear not matched")
	assert.Equal(t, []int{8, 5, 5, 0, 9, 13, 7, 14, 9, 0, 2, 9, 20, 6, 6, 8, 20, 8, 7, 7, 3, 5, 4, 4, 4, 4, 8, 5, 6, 9, 1, 1, 1, 1, 5, 9, 8, 7, 0, 9, 1, 9, 8, 2, 5, 8, 9, 8, 20, 9, 6, 1, 8, 2, 2, 2, 2, 3, 4, 9, 6, 5, 0, 0, 0, 0, 7, 6, 7, 9, 4, 3, 3, 3, 7, 4, 8, 3, 9, 8, 6, 8, 0, 7, 4, 20, 8, 4, 6, 3, 8, 2, 20, 9, 4, 6, 1, 7, 2, 1, 3, 9, 0, 7, 9, 5, 5, 4, 7, 7, 6}, c.GearsConfig[3].Symbols, "gears config one gear not matched")
	assert.Equal(t, []int{8, 5, 5, 0, 9, 3, 17, 14, 9, 0, 2, 9, 20, 6, 6, 8, 20, 8, 7, 7, 3, 5, 4, 4, 4, 4, 8, 5, 6, 9, 1, 1, 1, 1, 5, 9, 8, 7, 0, 9, 1, 9, 8, 2, 5, 8, 9, 8, 20, 9, 6, 1, 8, 2, 2, 2, 2, 3, 4, 9, 6, 5, 0, 0, 0, 0, 7, 6, 7, 9, 4, 3, 3, 3, 7, 4, 8, 3, 9, 8, 6, 8, 0, 7, 4, 20, 8, 4, 6, 3, 8, 2, 20, 9, 4, 6, 1, 7, 2, 1, 3, 9, 0, 7, 9, 5, 5, 4, 7, 7, 6}, c.GearsConfig[4].Symbols, "gears config one gear not matched")
}
