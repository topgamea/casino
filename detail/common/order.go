package common

type Order struct {
	RoundId string `json:"round_id" orm:"pk"`
	Uid     string `json:"uid" orm:"uid"`
}
