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
	Rows          int
	Columns       int
	LinesConfig   []*LineConfig
	ObtainsConfig map[int]*ObtainConfig
	BoardsConfig  map[int]*BoardConfig
	GearsConfig   map[int]*GearConfig
}

//FrontendConfig TODO
type FrontendConfig struct {
	Rows          int                   `json:"rows,omitempty"`
	Columns       int                   `json:"columns,omitempty"`
	LinesConfig   [][]int               `json:"lines,omitempty"`
	ObtainsConfig map[int]*ObtainConfig `json:"obtains,omitempty"`
	BoardsConfig  map[int]*BoardConfig  `json:"boards,omitempty"`
	GearsConfig   map[int]*GearConfig   `json:"gears,omitempty"`
}

//LineConfig TODO
type LineConfig struct {
	Line []int
}

//ObtainConfig TODO
type ObtainConfig struct {
	ID     int
	Reward []int
}

//BoardConfig TODO
type BoardConfig struct {
	ID     int
	Btype  int
	Payout int
	Rows   int
	Colums int
	Slots  []int
}

//GearConfig TODO
type GearConfig struct {
	ID      int
	Symbols []int
}

type originCasinoConfig struct {
	Rows    int                 `json:"rows,omitempty"`
	Columns int                 `json:"columns,omitempty"`
	Lines   []string            `json:"lines,omitempty"`
	Obtains []string            `json:"obtains,omitempty"`
	Boards  []originBoardConfig `json:"boards,omitempty"`
	Gears   []originGearConfig  `json:"gears,omitempty"`
}

type originBoardConfig struct {
	ID     int    `json:"id,omitempty"`
	Btype  int    `json:"btype,omitempty"`
	Payout int    `json:"payout,omitempty"`
	Data   string `json:"data,omitempty"`
}

type originGearConfig struct {
	ID   int    `json:"id,omitempty"`
	Data string `json:"data,omitempty"`
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
	config := new(Config)
	config.Rows = originConfig.Rows
	config.Columns = originConfig.Columns
	config.LinesConfig = make([]*LineConfig, 0)
	config.ObtainsConfig = make(map[int]*ObtainConfig)
	config.BoardsConfig = make(map[int]*BoardConfig)
	config.GearsConfig = make(map[int]*GearConfig)

	frontendConfig := new(FrontendConfig)
	frontendConfig.Rows = originConfig.Rows
	frontendConfig.Columns = originConfig.Columns
	frontendConfig.LinesConfig = make([][]int, 0)
	frontendConfig.ObtainsConfig = make(map[int]*ObtainConfig)
	frontendConfig.BoardsConfig = make(map[int]*BoardConfig)
	frontendConfig.GearsConfig = make(map[int]*GearConfig)

	//add lines config
	for _, line := range originConfig.Lines {
		lc := new(LineConfig)
		lc.Line = make([]int, 0)
		for _, s := range strings.Split(line, ",") {
			i, err := strconv.Atoi(s)
			if err != nil {
				return nil, nil, err
			}
			lc.Line = append(lc.Line, i)
		}
		config.LinesConfig = append(config.LinesConfig, lc)
		frontendConfig.LinesConfig = append(frontendConfig.LinesConfig, lc.Line)
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
	//add boards config
	for _, board := range originConfig.Boards {
		bc := new(BoardConfig)
		bc.ID = board.ID
		bc.Btype = board.Btype
		bc.Payout = board.Payout
		bc.Rows = config.Rows
		bc.Colums = config.Columns
		slots := make([]int, config.Rows*config.Columns)
		for _, coorAndGear := range strings.Split(board.Data, ",") {
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
			geari, err := strconv.Atoi(gear)
			if err != nil {
				return nil, nil, err
			}
			slots[(coorXi-1)*bc.Colums+(coorYi-1)] = geari
		}
		bc.Slots = slots
		config.BoardsConfig[board.ID] = bc
		frontendConfig.BoardsConfig[board.ID] = bc
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

	return config, nil, nil
}
