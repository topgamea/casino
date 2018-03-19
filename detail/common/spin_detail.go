package common

import (
	"time"
)

type Game struct {
	GameId      string `json:"game_id" orm:"pk"`
	GameName    string `json:"game_name"`
	Rows        uint8   `json:"rows"`
	Columns     uint8   `json:"columns"`
	ColumnsInfo string `json:"columns_info"` // json数组 [int]

	CreatedAt time.Time `json:"-" orm:"auto_now_add;type(datetime)"`
	UpdatedAt time.Time `json:"-" orm:"auto_now;type(datetime)"`
}

func (u *Game) TableName() string {
	return "t_game"
}

type Round struct {
	RoundId string `json:"round_id" orm:"pk"`
	UserId  string `json:"user_id"`             //玩家账号
	ProxyId string `json:"proxy_id" orm:"null"` //代理账号
						    //GameName string `json:"game_name"`

	StartTime   time.Time `json:"start_time" orm:"type(datetime)"`
	EndTime     time.Time `json:"end_time" orm:"type(datetime)"`
	TotalReward uint64    `json:"total_reward"` //总赢钱
	TotalBet    uint      `json:"total_bet"`    //下注金额 主游戏有效，freespin时无效

	Game  *Game      `json:"game" orm:"null;rel(fk);on_delete(set_null)"`
	Spins []*SpinNew `json:"spin_details" orm:"reverse(many)"`

	GameStat *GameStat `json:"game_stat" orm:"reverse(one)"`

	CreatedAt time.Time `json:"-" orm:"auto_now_add;type(datetime)"`
	UpdatedAt time.Time `json:"-" orm:"auto_now;type(datetime)"`
}

func (u *Round) TableName() string {
	return "t_game_round"
}

type SpinNew struct {
	Id         uint64 `json:"-" orm:"pk;auto"`
	Round      *Round `json:"-" orm:"rel(fk)"`
	FreeGameId uint8   `json:"free_game_id"` // 主游戏填0 freespin填具体值第几次freespin
	RespinId   uint8   `json:"respin_id"`    //非resipin情况下填0， respin情况下填具体第几次respin，从1开始计
	SpinType   uint8   `json:"spin_type"`    //子游戏类型  0 主游戏， 1 freespin ，2 主游戏中的respin 3 freespin中现的espin

	SpinReward       uint      `json:"spin_reward"`                   //此次spin赢钱
	SpinBet          uint      `json:"spin_bet"`                      //此次下注金额 主游戏有效，freespin时无效,对于美人鱼 辣椒 和红唇，该值和TotalReward相同
	BetTime          time.Time `json:"bet_time" orm:"type(datetime)"` //下注时间 XXXX/XX/XX XX:XX:XX
	RewardLineNumber uint8      `json:"reward_line_number"`            //赢钱线数

	Items         string           `json:"items" orm:"size(10000)"`            //json数组 [[int]]                       //旋转结果，一个子数组字代表一列
	RewardDetails []*GenericReward `json:"reward_details" orm:"reverse(many)"` //包含中scatter bonus jackpot 线等各种中奖

	CreatedAt time.Time `json:"-" orm:"auto_now_add;type(datetime)"`
	UpdatedAt time.Time `json:"-" orm:"auto_now;type(datetime)"`
}

func (u *SpinNew) TableName() string {
	return "t_game_spin"
}

func (u *SpinNew) TableUnique() [][]string {
	return [][]string{
		[]string{"Round", "FreeGameId", "RespinId"},
	}
}

type GenericReward struct {
	Id int64 `json:"-" orm:"pk;auto"`
									  //Round         *Round   `json:"-" orm:"rel(fk)"`
	Spin          *SpinNew `json:"-" orm:"rel(fk)"`
	LineId        uint     `json:"line_id"`                          //中奖线id  如果是bonus 或者 scatter中奖，就把该值设为 bonus 和 scatter图标
	RewardType    uint8     `json:"reward_type"`                      // 0 item中奖 1 scatter  2 bonus 3 jackpot
	RewardItems   string   `json:"reward_items"`                     //json数组 [int] 中奖图标 对于美人鱼，只有一个图标；红唇会有多个图标
	BetMultiple   uint     `json:"bet_multiple"`                     //下注倍数
	Reward        uint     `json:"reward"`                           //此线赢钱
	Multiple      uint     `json:"multiple"`                         // 根据玩法，玩法中如果有乘倍则显示倍数，如果没有则显示1（现阶段只有财神机器有乘倍）
	ItemNumber    uint8     `json:"item_number"`                      //几连线
	ItemPositions string   `json:"item_positions" orm:"size(10000)"` // json数组 [[column,row]]

	CreatedAt time.Time `json:"-" orm:"auto_now_add;type(datetime)"`
	UpdatedAt time.Time `json:"-" orm:"auto_now;type(datetime)"`
}

func (u *GenericReward) TableName() string {
	return "t_game_spin_reward"
}

func (u *GenericReward) TableUnique() [][]string {
	return [][]string{
		[]string{"Spin", "LineId", "RewardType"},
	}
}

//需要在消息中标注当前spin是否触发了jackpot、freespin、respin或bonus游戏
type GameStat struct {
	Id         uint64 `json:"-" orm:"pk;auto"`
	Round      *Round `json:"-" orm:"null;rel(one);on_delete(set_null)"`
	HitJackpot  bool      `json:"hit_jackpot"`
	HitFreespin bool      `json:"hit_freespin"`
	HitRespin   bool      `json:"hit_respin"`
	HitBonus    bool      `json:"hit_bounus"`
	CreatedAt   time.Time `json:"-" orm:"auto_now_add;type(datetime)"`
	UpdatedAt   time.Time `json:"-" orm:"auto_now;type(datetime)"`
}

func (u *GameStat) TableName() string {
	return "t_game_stat_reward"
}


func (u *GameStat) TableUnique() [][]string {
	return [][]string{
		[]string{"Round"},
	}
}
