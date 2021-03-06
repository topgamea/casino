package casino

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

var config *Config
var frontendConfig *FrontendConfig

//Config TODO
type Config struct {
	Rows          int                     `json:"rows,omitempty"`
	Columns       int                     `json:"columns,omitempty"`
	ExtraNum      int                     `json:"extraNum,omitempty"`
	Bets          []int                   `json:"bets,omitempty"`
	ScoreBase     int                     `json:"score_base,omitempty"`
	LinesConfig   []*LineConfig           `json:"lines,omitempty"`
	ObtainsConfig map[int]*ObtainConfig   `json:"obtains,omitempty"`
	BonusConfig   map[int]*BonusConfig    `json:"bonus,omitempty"`
	BoardsConfig  map[string]*BoardConfig `json:"boards,omitempty"`
	GearsConfig   map[string]*GearConfig  `json:"gears,omitempty"`
	WildConfig    *WildConfig             `json:"wild,omitempty"`
}

//FrontendConfig TODO
type FrontendConfig struct {
	Rows          int                   `json:"rows,omitempty"`
	Columns       int                   `json:"columns,omitempty"`
	ExtraNum      int                   `json:"extraNum,omitempty"`
	Bets          []int                 `json:"bets,omitempty"`
	ScoreBase     int                   `json:"score_base,omitempty"`
	LinesConfig   [][]int               `json:"lines,omitempty"`
	ObtainsConfig map[int]*ObtainConfig `json:"obtains,omitempty"`
	BonusConfig   map[int]*BonusConfig  `json:"bonus,omitempty"`
	//	BoardsConfig  map[int]*BoardConfig   `json:"boards,omitempty"`
	GearsConfig map[string]*GearConfig `json:"gears,omitempty"`
	WildConfig  *WildConfig            `json:"wild,omitempty"`
}

//LineConfig TODO
type LineConfig struct {
	Line []int `json:"line,omitempty"`
}

//ObtainConfig TODO
type ObtainConfig struct {
	ID     int   `json:"id"`
	Reward []int `json:"reward,omitempty"`
}

//BonusConfig TODO
type BonusConfig struct {
	ID     int   `json:"id"`
	Reward []int `json:"reward,omitempty"`
}

//BoardConfig TODO
type BoardConfig struct {
	ID     string   `json:"id"`
	Btype  int      `json:"btype"`
	Payout int      `json:"payout"`
	Rows   int      `json:"rows,omitempty"`
	Colums int      `json:"colums,omitempty"`
	Gears  []string `json:"gears,omitempty"`
	Slots  []string `json:"slots,omitempty"`
}

//GearConfig TODO
type GearConfig struct {
	ID      string `json:"id"`
	Symbols []int  `json:"symbols,omitempty"`
}

//WildConfig TODO
type WildConfig struct {
	IDs    []int `json:"ids"`
	Except []int `json:"except"`
}

type originCasinoConfig struct {
	Rows      int                 `json:"rows,omitempty"`
	Columns   int                 `json:"columns,omitempty"`
	ExtraNum  int                 `json:"extraNum,omitempty"`
	Bets      []int               `json:"bets,omitempty"`
	ScoreBase int                 `json:"score_base,omitempty"`
	Lines     []string            `json:"lines,omitempty"`
	Obtains   []string            `json:"obtains,omitempty"`
	Bonus     []string            `json:"bonus,omitempty"`
	Boards    []originBoardConfig `json:"boards,omitempty"`
	Gears     []originGearConfig  `json:"gears,omitempty"`
	Wild      originWildConfig    `json:"wild,omitempty"`
}

type originBoardConfig struct {
	ID     string   `json:"id"`
	Btype  int      `json:"btype"`
	Payout int      `json:"payout"`
	Data   []string `json:"data,omitempty"`
}

type originGearConfig struct {
	ID   string `json:"id"`
	Data string `json:"data,omitempty"`
}

type originWildConfig struct {
	IDs    string `json:"ids"`
	Except string `json:"except"`
}

//ParseCasinoConfig TODO
func ParseCasinoConfig(file string) (*Config, *FrontendConfig, error) {
	_, err := os.Stat(file)
	if err != nil {
		return nil, nil, err
	}
	bytesInFile, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, nil, err
	}
	var originConfig originCasinoConfig
	err = json.Unmarshal(bytesInFile, &originConfig)
	if err != nil {
		return nil, nil, err
	}
	wildConfig := new(WildConfig)
	config := new(Config)
	config.Rows = originConfig.Rows
	config.Columns = originConfig.Columns
	config.ExtraNum = originConfig.ExtraNum
	config.Bets = originConfig.Bets
	config.ScoreBase = originConfig.ScoreBase
	config.LinesConfig = make([]*LineConfig, 0)
	config.ObtainsConfig = make(map[int]*ObtainConfig)
	config.BonusConfig = make(map[int]*BonusConfig)
	config.BoardsConfig = make(map[string]*BoardConfig)
	config.GearsConfig = make(map[string]*GearConfig)
	config.WildConfig = wildConfig

	frontendConfig := new(FrontendConfig)
	frontendConfig.Rows = originConfig.Rows
	frontendConfig.Columns = originConfig.Columns
	frontendConfig.ExtraNum = originConfig.ExtraNum
	frontendConfig.Bets = originConfig.Bets
	frontendConfig.ScoreBase = originConfig.ScoreBase
	frontendConfig.LinesConfig = make([][]int, 0)
	frontendConfig.ObtainsConfig = make(map[int]*ObtainConfig)
	frontendConfig.BonusConfig = make(map[int]*BonusConfig)
	//frontendConfig.BoardsConfig = make(map[string]*BoardConfig)
	frontendConfig.GearsConfig = make(map[string]*GearConfig)
	frontendConfig.WildConfig = wildConfig

	//add lines config
	for _, line := range originConfig.Lines {
		lc := new(LineConfig)
		lc1 := new(LineConfig)
		lc.Line = make([]int, 0)
		lc1.Line = make([]int, 0)
		for _, s := range strings.Split(line, ",") {
			i, err := strconv.Atoi(s)
			if err != nil {
				return nil, nil, err
			}
			lc.Line = append(lc.Line, i+1)
			lc1.Line = append(lc1.Line, i)
		}
		config.LinesConfig = append(config.LinesConfig, lc)
		frontendConfig.LinesConfig = append(frontendConfig.LinesConfig, lc1.Line)
	}
	//add obtains config
	for _, obtain := range originConfig.Obtains {
		oc := new(ObtainConfig)
		oc.Reward = make([]int, 0)
		idAndReward := strings.Split(obtain, ":")
		id, err := strconv.Atoi(idAndReward[0])
		if err != nil {
			return nil, nil, err
		}
		oc.ID = id
		for _, reward := range strings.Split(idAndReward[1], ",") {
			i, err := strconv.Atoi(reward)
			if err != nil {
				return nil, nil, err
			}
			oc.Reward = append(oc.Reward, i)
		}
		config.ObtainsConfig[id] = oc
		frontendConfig.ObtainsConfig[id] = oc
	}
	//add bonus config
	for _, bonus := range originConfig.Bonus {
		oc := new(BonusConfig)
		oc.Reward = make([]int, 0)
		idAndReward := strings.Split(bonus, ":")
		id, err := strconv.Atoi(idAndReward[0])
		if err != nil {
			return nil, nil, err
		}
		oc.ID = id
		for _, reward := range strings.Split(idAndReward[1], ",") {
			i, err := strconv.Atoi(reward)
			if err != nil {
				return nil, nil, err
			}
			oc.Reward = append(oc.Reward, i)
		}
		config.BonusConfig[id] = oc
		frontendConfig.BonusConfig[id] = oc
	}
	//add boards config
	for _, board := range originConfig.Boards {
		bc := new(BoardConfig)
		bc.ID = board.ID
		bc.Btype = board.Btype
		bc.Payout = board.Payout
		bc.Rows = config.Rows
		bc.Colums = config.Columns
		slots := make([]string, config.Rows*config.Columns)
		bc.Gears = make([]string, 0)
		checkGears := make(map[string]int)
		for _, coorAndGear := range board.Data {
			coor := strings.Split(coorAndGear, ":")[0]
			gear := strings.Split(coorAndGear, ":")[1]
			coorX := strings.Split(coor, "-")[0]
			coorY := strings.Split(coor, "-")[1]
			coorXi, err := strconv.Atoi(coorX)
			if err != nil {
				return nil, nil, err
			}
			coorYi, err := strconv.Atoi(coorY)
			if err != nil {
				return nil, nil, err
			}
			// geari, err := strconv.Atoi(gear)
			// if err != nil {
			// 	return nil, nil, err
			// }
			slots[coorXi*bc.Colums+coorYi] = gear
			if _, ok := checkGears[gear]; !ok {
				checkGears[gear] = 1
				bc.Gears = append(bc.Gears, gear)
			}
		}
		bc.Slots = slots
		config.BoardsConfig[board.ID] = bc
		//frontendConfig.BoardsConfig[board.ID] = bc
	}
	//add gears config
	for _, gear := range originConfig.Gears {
		gc := new(GearConfig)
		gc.ID = gear.ID
		gc.Symbols = make([]int, 0)
		for _, symbol := range strings.Split(gear.Data, ",") {
			symboli, err := strconv.Atoi(symbol)
			if err != nil {
				return nil, nil, err
			}
			gc.Symbols = append(gc.Symbols, symboli)
		}
		config.GearsConfig[gc.ID] = gc
		frontendConfig.GearsConfig[gc.ID] = gc
	}
	//add wild config
	owc := originConfig.Wild
	ids := make([]int, 0)
	except := make([]int, 0)
	for _, idStr := range strings.Split(owc.IDs, ",") {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return nil, nil, err
		}
		ids = append(ids, id)
	}
	for _, exceptStr := range strings.Split(owc.Except, ",") {
		exceptInt, err := strconv.Atoi(exceptStr)
		if err != nil {
			return nil, nil, err
		}
		except = append(except, exceptInt)
	}
	wildConfig.IDs = ids
	wildConfig.Except = except

	return config, frontendConfig, nil
}
