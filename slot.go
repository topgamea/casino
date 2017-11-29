package casino

//Slot TODO
type Slot struct {
	X            int
	Y            int
	Runner       *Runner
	Which        int
	LockedSymbol int
	Locked       bool
}

//AddRunner TODO
func (slot *Slot) AddRunner(runner *Runner) error {
	slot.Runner = runner
	slot.Which = runner.Vision
	runner.Vision = runner.Vision + 1
	return nil
}

//Lock TODO
func (slot *Slot) Lock() error {
	slot.Locked = true
	slot.LockedSymbol = slot.GetSymbol()
	return nil
}

//GetSymbol Default: Return the Symbol of the watched runner, Special: Return the Locked Symbol or some Special Symbols
func (slot *Slot) GetSymbol() int {
	if slot.Runner == nil {
		return 0
	}
	if slot.Locked {
		return slot.LockedSymbol
	}
	return slot.Runner.GetSymbol(slot.Which)
}
