package casino

import (
	"sync"
)

//RunnerManage TODO
type RunnerManage struct {
	Runners map[string]*Runner
	Parent  *Node
	Locker  sync.Mutex
}

//NewRunnerManage TODO
func NewRunnerManage(n *Node) (*RunnerManage, error) {
	rm := new(RunnerManage)
	rm.Runners = make(map[string]*Runner)
	rm.Parent = n
	return rm, nil
}

//AddRunner TODO
func (rm *RunnerManage) AddRunner(id string) (*Runner, error) {
	if _, ok := rm.Runners[id]; ok {
		return rm.Runners[id], nil
	}
	r := new(Runner)
	r.BindGear = id
	r.ID = id
	rm.Runners[id] = r
	return r, nil
}

//Run TODO
func (rm *RunnerManage) Run() error {
	for _, r := range rm.Runners {
		err := r.Run()
		if err != nil {
			return err
		}
	}
	return nil
}
