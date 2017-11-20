package casino

import (
	"math/rand"
)

//Runner TODO
type Runner struct {
	ID       int
	BindGear int
	NowPos   int
	Vision   int
}

//Run TODO
func (r *Runner) Run() error {
	gearConfig := config.GearsConfig[r.BindGear]
	randPos := rand.Intn(len(gearConfig.Symbols))
	r.NowPos = randPos
	return nil
}

//GetRunnerPos TODO
func (r *Runner) GetRunnerPos() int {
	return r.NowPos
}

//GetSymbol TODO
func (r *Runner) GetSymbol(which int) int {
	gearConfig := config.GearsConfig[r.BindGear]
	index := (r.NowPos + which + 1) / len(gearConfig.Symbols)
	if index == 0 {
		index = len(gearConfig.Symbols)
	}
	index = index - 1
	return gearConfig.Symbols[index]
}
