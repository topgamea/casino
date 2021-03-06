package casino

import (
	"math/rand"
	"time"
)

//CheckBoard TODO
func CheckBoard(c *Context) error {
	_, err := c.GetValue("boardID")
	if err != ErrPairNotExist {
		return err
	}
	err = c.AddPair("boardID", "0_9")
	if err != nil {
		return err
	}
	return nil
}

//Play TODO
func Play(c *Context) error {
	rand.Seed(time.Now().UnixNano())
	n := c.N
	//get board id
	boardID, err := c.GetValue("boardID")
	if err != nil {
		return err
	}
	//Choose an Appropriate Board, default 1# board
	b, err := n.BM.SwitchBoard(boardID.(string))
	if err != nil {
		return err
	}
	//Runner start to randomly select from the binded gear
	err = n.RM.Run()
	if err != nil {
		return err
	}
	c.AddPair("board", b)
	return nil
}

//ComputeLine TODO
func ComputeLine(c *Context) error {
	bInContext, err := c.GetValue("board")
	if err != nil {
		return err
	}
	b := bInContext.(*Board)
	//Compute line reward
	reward, lines, linesItemsPos, biLineRewards, err := c.N.LC.Compute(b)
	if err != nil {
		return err
	}
	c.AddPair("reward", reward)
	c.AddPair("lines", lines)
	c.AddPair("linesItemsPos", linesItemsPos)
	c.AddPair("biLineRewards", biLineRewards)
	return nil
}

//GetGearItems TODO
func GetGearItems(c *Context) error {
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
