package casino

import "errors"

//NodeType :Type of Node
type NodeType int

const (
	_ NodeType = iota
	//Personal Node just for the Single Person
	Personal
	//Room Node just for the Room (Multi Players)
	Room
)

var (
	//ErrPairAlreadyExist context already include this kv pair
	ErrPairAlreadyExist = errors.New("pair already exists")
	//ErrPairNotExist context not include this kv pair
	ErrPairNotExist = errors.New("pair not exist")
)
