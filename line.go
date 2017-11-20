package casino

//LineCompute TODO
type LineCompute interface {
	Compute(b *Board) (int, error)
}

//NormalLine TODO
type NormalLine struct {
}

//Compute TODO
func (l *NormalLine) Compute(b *Board) (int, error) {
	lineSlots := make([]*Slot, b.Columns)
	reward := 0
	for _, lineConfig := range config.LinesConfig {
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
			if s.GetSymbol() == firstSymbol {
				totalCount++
			} else {
				break
			}
		}
		obtainConfig := config.ObtainsConfig[firstSymbol]
		if obtainConfig.Reward[totalCount-1] != 0 {
			reward += obtainConfig.Reward[totalCount-1]
		}
	}
	return reward, nil
}
