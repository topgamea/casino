package casino

import (
	"math/rand"
	"time"
)

func play(c *Context) error {
	rand.Seed(time.Now().UnixNano())
	n := c.N
	//Choose an Appropriate Board, default 1# board
	b, err := n.BM.SwitchBoard(1)
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
	// c.AddPair("board", b)
	c.AddPair("reward", reward)
	c.AddPair("lines", lines)
	c.AddPair("linesItems", linesItems)
	return nil
}
