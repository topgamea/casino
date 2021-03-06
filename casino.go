package casino

import (
	"errors"
	"math/rand"
	"time"
)

/**
=====================================================Casino======================================================
===========Node-1======================Node-2========================Node-3=======================Node-4===================
==(RunnerManage+BoardManage)====(RunnerManage+BoardManage)====(RunnerManage+BoardManage)====(RunnerManage+BoardManage)==
**/

//Casino :Long Time Running Casino instance, Create New Node and Save some Global Constants or Variables or Context
//In the future, we can use a Pool of Nodes to improve the Performance
type Casino struct {
	Config         *Config
	FrontendConfig *FrontendConfig
	LC             LineCompute
}

//Create :New a Casino instance
func Create(configFile string) (*Casino, error) {
	casino := new(Casino)
	c, fc, err := ParseCasinoConfig(configFile)
	if err != nil {
		return nil, err
	}
	casino.Config = c
	casino.FrontendConfig = fc
	config = c
	frontendConfig = fc
	return casino, nil
}

//NewNode :Create a Node instance For Player or Room Used To Play Casino
//Default: NodeType == Personl
//In the future, we can Pass Node Type in the Params To the NewNode function
func (c *Casino) NewNode(lc LineCompute, nf FrontendGears) (*Node, error) {
	n := new(Node)
	n.Type = Personal
	rm, err := NewRunnerManage(n)
	if err != nil {
		return nil, err
	}
	n.RM = rm
	bm, err := NewBoardManage(n)
	if err != nil {
		return nil, err
	}
	n.BM = bm
	n.LC = lc
	n.FG = nf
	n.HCList = make(map[string]*HookChain)
	n.C = &Context{N: n, KV: make(map[string]interface{})}
	return n, nil
}

//Node :New a Node for every player or room to Play casino
type Node struct {
	Type   NodeType
	Debug  bool
	C      *Context
	HCList map[string]*HookChain
	RM     *RunnerManage
	BM     *BoardManage
	LC     LineCompute
	FG     FrontendGears
}

//RegisterDefaultHooks TODO
func (n *Node) RegisterDefaultHooks() {
	n.RegisterHook("default", CheckBoard)
	n.RegisterHook("default", Play)
	n.RegisterHook("default", GetGearItems)
}

//RegisterHook TODO
func (n *Node) RegisterHook(hcID string, hf HookFunc) {
	if _, ok := n.HCList[hcID]; !ok {
		n.HCList[hcID] = new(HookChain)
	}
	n.HCList[hcID].addHook(hf)
}

//Execute TODO
func (n *Node) Execute(id string) error {
	for {
		if _, ok := n.HCList[id]; !ok {
			return errors.New("hook chain not exist")
		}
		err := n.HCList[id].execute(n.C)
		if err != nil {
			return err
		}
		if n.C.GotoHC == "" {
			break
		}
		id = n.C.GotoHC
	}
	return nil
}

//Play :Start or Init the Node
func (n *Node) Play() (int, error) {
	rand.Seed(time.Now().UnixNano())
	//Choose an Appropriate Board, default 1# board
	b, err := n.BM.SwitchBoard("1")
	if err != nil {
		return 0, err
	}
	//Runner start to randomly select from the binded gear
	err = n.RM.Run()
	if err != nil {
		return 0, err
	}
	//Compute line reward
	reward, _, _, _, err := n.LC.Compute(b)
	if err != nil {
		return 0, err
	}
	return reward, nil
}
