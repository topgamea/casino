package casino

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCasino(t *testing.T) {
	tmpFileName := "./config_tmp.json"
	c, err := Create(tmpFileName, new(NormalLine))
	if err != nil {
		t.Errorf("create casino object error: %v", err)
	}
	node, err := c.NewNode()
	if err != nil {
		t.Errorf("casino new node error: %v", err)
	}
	b, _ := node.BM.SwitchBoard(1)
	_, err = node.Play()
	if err != nil {
		t.Errorf("casino start node error: %v", err)
	}
	for _, slot := range b.Slots {
		assert.Contains(t, []int{1}, slot.GetSymbol())
	}
	reward, err := node.LC.Compute(b)
	if err != nil {
		t.Errorf("casino compute line error: %v", err)
	}
	assert.Equal(t, 3000, reward, "reward not matched")
}
