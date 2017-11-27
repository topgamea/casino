package casino

//BoardManage Manage different boards
type BoardManage struct {
	Boards     map[string]*Board
	Parent     *Node
	NowBoardID string
}

//NewBoardManage New a Board Manager
func NewBoardManage(n *Node) (*BoardManage, error) {
	bm := new(BoardManage)
	bm.Boards = make(map[string]*Board)
	bm.Parent = n
	return bm, nil
}

//SwitchBoard Swith Board by different Board ID in the config file
func (bm *BoardManage) SwitchBoard(boardID string) (*Board, error) {
	bm.NowBoardID = boardID
	if _, ok := bm.Boards[boardID]; ok {
		return bm.Boards[boardID], nil
	}
	boardConfig := config.BoardsConfig[boardID]
	b, err := NewBoard(boardConfig.Rows, boardConfig.Colums)
	if err != nil {
		return nil, err
	}
	for i := 0; i < boardConfig.Rows*boardConfig.Colums; i++ {
		if boardConfig.Slots[i] == "" {
			continue
		}
		r, err := bm.Parent.RM.AddRunner(boardConfig.Slots[i])
		if err != nil {
			return nil, err
		}
		b.Slots[i].AddRunner(r)
	}
	bm.Boards[boardID] = b
	return b, nil
}
