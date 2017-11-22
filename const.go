package casino

//NodeType :Type of Node
type NodeType int

const (
	_ NodeType = iota
	//Personal Node just for the Single Person
	Personal
	//Room Node just for the Room (Multi Players)
	Room
)
