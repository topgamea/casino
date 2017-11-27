package casino

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCasino(t *testing.T) {
	tmpFileName := "./config_tmp.json"
	c, err := Create(tmpFileName)
	if err != nil {
		t.Errorf("create casino object error: %v", err)
	}
	node, err := c.NewNode(DefaultLineCompute, DefaultFrontendGears)
	if err != nil {
		t.Errorf("casino new node error: %v", err)
	}
	b, _ := node.BM.SwitchBoard("1")
	_, err = node.Play()
	if err != nil {
		t.Errorf("casino start node error: %v", err)
	}
	for _, slot := range b.Slots {
		assert.Contains(t, []int{1, 2, 3}, slot.GetSymbol())
	}
	reward, lines, linesItems, err := node.LC.Compute(b)
	if err != nil {
		t.Errorf("casino compute line error: %v", err)
	}
	assert.Equal(t, 300, reward, "reward not matched")
	assert.Equal(t, []int{0, 1, 2}, lines, "reward lines not matched")
	assert.Equal(t, [][]int{[]int{0}, []int{0}, []int{0}}, linesItems, "reward lines items not matched")
	//test frontend gears
	frontendGears := DefaultFrontendGears
	aa, bb := frontendGears.GetGearWithItems(b)
	assert.Equal(t, []int{1, 1, 1, 1, 1}, aa[0], "frontend gears not matched")
	assert.Contains(t, []int{0, 1, 2}, bb[0], "next position not matched")
}

func TestHookCasino(t *testing.T) {
	tmpFileName := "./config_tmp.json"
	c, err := Create(tmpFileName)
	if err != nil {
		t.Errorf("create casino object error: %v", err)
	}
	node, err := c.NewNode(DefaultLineCompute, DefaultFrontendGears)
	if err != nil {
		t.Errorf("casino new node error: %v", err)
	}
	node.RegisterDefaultHooks()
	err = node.Execute()
	if err != nil {
		t.Errorf("execute hook chain error: %v", err)
	}
}
