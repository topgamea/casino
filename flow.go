package casino

import (
	"math/rand"
	"time"
)

func checkBoard(c *Context) error {
	err := c.AddPair("boardID", 1)
	if err != nil {
		return err
	}
	return nil
}

func play(c *Context) error {
	rand.Seed(time.Now().UnixNano())
	n := c.N
	//get board id
	boardID, err := c.GetValue("boardID")
	if err != nil {
		return err
	}
	//Choose an Appropriate Board, default 1# board
	b, err := n.BM.SwitchBoard(boardID.(int))
	if err != nil {
		return err
	}
	//Runner start to randomly select from the binded gear
	err = n.RM.Run()
	if err != nil {
		return err
	}
	//Compute line reward
	reward, lines, linesItems, err := n.LC.Compute(b)
	if err != nil {
		return err
	}
	c.AddPair("board", b)
	c.AddPair("reward", reward)
	c.AddPair("lines", lines)
	c.AddPair("linesItems", linesItems)
	return nil
}

func getGearItems(c *Context) error {
	//test frontend gears
	frontendGears := c.N.FG
	b, err := c.GetValue("board")
	if err != nil {
		return err
	}
	items, nextIndex := frontendGears.GetGearWithItems(b.(*Board))
	c.AddPair("items", items)
	c.AddPair("nextIndex", nextIndex)
	return nil
}
