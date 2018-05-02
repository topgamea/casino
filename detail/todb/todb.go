package todb

import (
	"casino/detail/common"
	"encoding/json"
	"errors"
	"gopkg.in/resty.v1"
	"time"
)

type status struct {
	Code     int       `json:"code"`
	Message  string    `json:"message"`
	Datetime time.Time `json:"datetime"`
}
type JsonResp struct {
	Data   interface{} `json:"data"`
	Status *status     `json:"status"`
}

func SaveToDb(r *common.Round, toUrl string) error {

	res, err := resty.DefaultClient.R().SetQueryParams(map[string]string{
		"user":    r.UserId,
		"game_id": r.Game.GameId,
	}).SetHeader("Content-Type", "application/json").SetBody(r).Post(toUrl)
	if err != nil {
		//logger.Error(err.Error())
		return err
	}

	js := &JsonResp{}
	err = json.Unmarshal(res.Body(), js)
	if err != nil {
		return err
	}

	if js.Status.Code != 0 {
		return errors.New(js.Status.Message)
	}
	return nil
}
