package casino

//LineCompute TODO
type LineCompute interface {
	Compute(b *Board) (reward int, lines []int, linesItemsIndex [][][]int, err error)
	SetNode(n *Node) error
}

//DefaultLineCompute Default Line Computer
var DefaultLineCompute LineCompute

func init() {
	DefaultLineCompute = new(NormalLine)
}

//NormalLine TODO
type NormalLine struct {
	Node *Node
}

//SetNode TODO
func (l *NormalLine) SetNode(n *Node) error {
	l.Node = n
	return nil
}

//Compute TODO
func (l *NormalLine) Compute(b *Board) (int, []int, [][][]int, error) {
	biLineRewards := make([]BILineReward, 0)

	lineSlots := make([]*Slot, b.Columns)
	lines := make([]int, 0)
	lineItems := make([][]int, 0)
	// lineItemsIndex := make([][]int, 0)
	lineItemsPos := make([][][]int, 0)
	reward := 0
	for lineIndex, lineConfig := range config.LinesConfig {
		// itemsIndex := make([]int, 0)
		itemsPos := make([][]int, 0)
		items := make([]int, 0)
		column := 1
		for _, row := range lineConfig.Line {
			lineSlots[column-1] = b.Slots[(row-1)*b.Columns+column-1]
			column++
		}
		//from the left first symbol of the line to the end
		firstSymbol := lineSlots[0].GetSymbol()
		//Check whether in the special symbol
		if _, ok := config.ObtainsConfig[firstSymbol]; !ok {
			continue
		}
		//Analyse how many times the first symbol show
		totalCount := 0
		for _, s := range lineSlots {
			if s.GetSymbol() == firstSymbol || WildReplace(config.WildConfig.IDs, config.WildConfig.Except, s.GetSymbol(), firstSymbol) {
				// itemsIndex = append(itemsIndex, totalCount)
				// items = append(items, firstSymbol)
				pos := []int{0, 0}
				pos[0] = s.Y
				pos[1] = s.X
				itemsPos = append(itemsPos, pos)
				totalCount++
			} else {
				break
			}
		}
		if totalCount == 1 {
			continue
		}
		obtainConfig := config.ObtainsConfig[firstSymbol]
		if obtainConfig.Reward[totalCount-2] != 0 {
			reward += obtainConfig.Reward[totalCount-2]
			lines = append(lines, lineIndex)
			// lineItemsIndex = append(lineItemsIndex, itemsPos)
			lineItemsPos = append(lineItemsPos, itemsPos)
			lineItems = append(lineItems, items)
			//debug for bi
			lineReward := new(BILineReward)
			lineReward.ID = firstSymbol
			lineReward.Count = totalCount
			lineReward.Reward = obtainConfig.Reward[totalCount-2]
			biLineRewards = append(biLineRewards, *lineReward)
		}
	}
	//debug for bi
	// err := l.Node.C.AddPair("biLineRewards", biLineRewards)
	// if err != nil {
	// 	return 0, nil, nil, err
	// }
	return reward, lines, lineItemsPos, nil
}
