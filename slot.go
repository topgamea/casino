package casino

//Slot TODO
type Slot struct {
	Runner *Runner
	Which  int
}

//AddRunner TODO
func (slot *Slot) AddRunner(runner *Runner) error {
	slot.Runner = runner
	slot.Which = runner.Vision
	runner.Vision = runner.Vision + 1
	return nil
}

//GetSymbol Default: Return the Symbol of the watched runner, Special: Return the Locked Symbol or some Special Symbols
func (slot *Slot) GetSymbol() int {
	return slot.Runner.GetSymbol(slot.Which)
}
