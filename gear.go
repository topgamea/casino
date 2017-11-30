package casino

//DefaultFrontendGears Default FrontendGears
var DefaultFrontendGears FrontendGears

func init() {
	DefaultFrontendGears = new(NormalFrontendGears)
}

//FrontendGears Used to Creat Gear Data For Frontend
type FrontendGears interface {
	GetGearWithItems(b *Board) ([][]int, []int)
}

//NormalFrontendGears The Normal FrontendGears struct
type NormalFrontendGears struct {
}

//GetGearWithItems GearID=>GearItemsArray, nextPosition
func (nfg *NormalFrontendGears) GetGearWithItems(b *Board) ([][]int, []int) {
	gears := make(map[string]int)
	frontendGears := make([][]int, 0)
	nextPositions := make([]int, 0)
	for _, s := range b.Slots {
		if s.Runner == nil {
			continue
		}
		gearID := s.Runner.BindGear
		if _, ok := gears[gearID]; ok {
			continue
		}
		startPos := s.Runner.NowPos - 1*config.ExtraNum
		if startPos < 0 {
			startPos = len(config.GearsConfig[gearID].Symbols) + startPos
		}
		items := make([]int, s.Runner.Vision+2*config.ExtraNum)
		nextPosition := startPos
		for i := 0; i < s.Runner.Vision+2*config.ExtraNum; i++ {
			index := (startPos + i) % len(config.GearsConfig[gearID].Symbols)
			items[i] = config.GearsConfig[gearID].Symbols[index]
			nextPosition = (index + 1) % len(config.GearsConfig[gearID].Symbols)
		}
		gears[gearID] = 1
		frontendGears = append(frontendGears, items)
		nextPositions = append(nextPositions, nextPosition)
	}
	return frontendGears, nextPositions
}
