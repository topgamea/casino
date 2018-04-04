package cq

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
	"gopkg.in/resty.v1"
)

//RandomTokenResponse TODO
type RandomTokenResponse struct {
	Data   CQData   `json:"data"`
	Status CQStatus `json:"status"`
}

//AuthUserResponse TODO
type AuthUserResponse struct {
	Data   CQUser   `json:"data"`
	Status CQStatus `json:"status"`
}

//CQUser TODO
type CQUser struct {
	Balance  float64 `json:"balance"`
	BetLevel int     `json:"betlevel"`
	Currency string  `json:"currency"`
	GameCode string  `json:"gamecode"`
	GameHall string  `json:"gamehall"`
	GamePlat string  `json:"gameplat"`
	GameTech string  `json:"gametech"`
	GameType string  `json:"gametype"`
	ID       string  `json:"id"`
	ParentID string  `json:"parentid"`
	WebID    int     `json:"webid"`
}

//CQBetRequest TODO
type CQBetRequest struct {
	ID        string  `json:"id"`
	GameToken string  `json:"gametoken"`
	GameCode  string  `json:"gamecode"`
	Round     string  `json:"round"`
	Amount    float64 `json:"amount"`
	MTCode    string  `json:"mtcode"`
	DateTime  string  `json:"datetime"`
}

//CQBetDataResponse TODO
type CQBetDataResponse struct {
	Balance  float64 `json:"balance"`
	Currency string  `json:"currency"`
}

//CQBetResponse TODO
type CQBetResponse struct {
	Data   CQBetDataResponse `json:"data"`
	Status CQStatus          `json:"status"`
}

//CQData TODO
type CQData struct {
	GameToken string `json:"gametoken"`
}

//CQStatus TODO
type CQStatus struct {
	Code     string `json:"code"`
	Message  string `json:"message"`
	DateTime string `json:"datetime"`
}

type CQOrderAuth struct {
	RoundId string `json:"roundid"`
	Account string `json:"account"`
	PAccount string `json:"paccount"`
}

//CQService TODO
type CQService struct {
	//Config *server_config.CqGameConfig

	Url  string
	Auth string
	Code string
}

func NewCqService(url, auth, code string) *CQService {
	return &CQService{
		Url:  url,
		Auth: auth,
		Code: code,
	}
}

//CreateRandomToken TODO
func (s *CQService) CreateRandomToken(gameCode string) (string, error) {
	//url := (s.Config.URL) + "dev/peace/gametoken?account=random"
	//auth := s.Config.Auth

	url := s.Url + "dev/peace/gametoken?account=random&gametype=slot&gamecode="
	auth := s.Auth

	if gameCode != "" {
		url += gameCode
	} else {
		url = s.Url + "dev/peace/gametoken?account=random"
	}


	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", auth)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var tokenResponse RandomTokenResponse
	err = json.Unmarshal(body, &tokenResponse)
	if err != nil {
		return "", err
	}
	if tokenResponse.Status.Code != "0" {
		return "", errors.New("get random token error:" + tokenResponse.Status.Code + tokenResponse.Status.Message)
	}
	return tokenResponse.Data.GameToken, nil
}

//AuthUser TODO
func (s *CQService) AuthUser(token string) (*AuthUserResponse, error) {
	//url := s.Config.URL + "gamepool/cq9/player/auth"
	//auth := s.Config.Auth

	url := s.Url + "gamepool/cq9/player/auth"
	auth := s.Auth

	//create post params
	params := "gametoken=" + token
	req, err := http.NewRequest("POST", url, strings.NewReader(params))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", auth)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var authResponse AuthUserResponse
	err = json.Unmarshal(body, &authResponse)
	if err != nil {
		return nil, err
	}
	if authResponse.Status.Code != "0" {
		return nil, errors.New("auth user error:" + authResponse.Status.Code + "|" + authResponse.Status.Message)
	}
	return &authResponse, nil
}

//Bet TODO
func (s *CQService) Bet(uid string, gameToken string, round string, amount float64, mtcode string) (*CQBetResponse, error) {
	//url := s.Config.URL + "gamepool/cq9/game/bet"
	//auth := s.Config.Auth
	//gameCode := s.Config.Code

	url := s.Url + "gamepool/cq9/game/bet"
	auth := s.Auth
	gameCode := s.Code

	loc := time.FixedZone("AST", -4*3600)
	RFC3339 := "2006-01-02T15:04:05"
	datetime := time.Now().In(loc).Format(RFC3339) + "-04:00"

	params := "id=" + uid + "&" +
		"gametoken=" + gameToken + "&" +
		"gamecode=" + gameCode + "&" +
		"round=" + round + "&" +
		"amount=" + fmt.Sprintf("%.2f", amount) + "&" +
		"mtcode=" + mtcode + "&" +
		"datetime=" + datetime

	/*
	req, err := http.NewRequest("POST", url, strings.NewReader(params))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", auth)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	*/

	resp, err := resty.SetTimeout(700 * time.Millisecond).R().SetBody(params).SetHeaders(
		map[string]string{
			"Content-Type": "application/x-www-form-urlencoded",
			"Authorization": auth,
		}).Post(url)
	if err != nil {
		return nil, err
	}


	var betResponse CQBetResponse
	err = json.Unmarshal(resp.Body(), &betResponse)
	if err != nil {
		return nil, err
	}
	if betResponse.Status.Code != "0" {
		return nil, errors.New("bet error:" + betResponse.Status.Code + "|" + betResponse.Status.Message + "|" + params)
	}
	return &betResponse, nil
}

//BetWin TODO
func (s *CQService) BetWin(uid string, gameToken string, round string, amount float64, mtcode string) (*CQBetResponse, error) {
	//url := s.Config.URL + "gamepool/cq9/game/win"
	//auth := s.Config.Auth
	//gameCode := s.Config.Code

	url := s.Url + "gamepool/cq9/game/win"
	auth := s.Auth
	gameCode := s.Code


	loc := time.FixedZone("AST", -4*3600)
	RFC3339 := "2006-01-02T15:04:05"
	datetime := time.Now().In(loc).Format(RFC3339) + "-04:00"

	params := "id=" + uid + "&" +
		"gametoken=" + gameToken + "&" +
		"gamecode=" + gameCode + "&" +
		"round=" + round + "&" +
		"amount=" + fmt.Sprintf("%.2f", amount) + "&" +
		"mtcode=" + mtcode + "&" +
		"datetime=" + datetime

	/*
	req, err := http.NewRequest("POST", url, strings.NewReader(params))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", auth)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	*/
	resp, err := resty.SetTimeout(700 * time.Millisecond).R().SetBody(params).SetHeaders(
		map[string]string{
			"Content-Type": "application/x-www-form-urlencoded",
			"Authorization": auth,
		}).Post(url)
	if err != nil {
		return nil, err
	}


	var betResponse CQBetResponse
	err = json.Unmarshal(resp.Body(), &betResponse)
	if err != nil {
		return nil, err
	}
	if betResponse.Status.Code != "0" {
		return nil, errors.New("bet win error:" + betResponse.Status.Code + "|" + betResponse.Status.Message + "|" + params)
	}
	return &betResponse, nil
}

//BetEnd TODO
func (s *CQService) BetEnd(uid string, gameToken string, round string, amount float64, mtcode string) (*CQBetResponse, error) {
	//url := s.Config.URL + "gamepool/cq9/game/end"
	//auth := s.Config.Auth
	//gameCode := s.Config.Code

	url := s.Url + "gamepool/cq9/game/end"
	auth := s.Auth
	gameCode := s.Code

	loc := time.FixedZone("AST", -4*3600)
	RFC3339 := "2006-01-02T15:04:05"
	datetime := time.Now().In(loc).Format(RFC3339) + "-04:00"

	params := "id=" + uid + "&" +
		"gametoken=" + gameToken + "&" +
		"gamecode=" + gameCode + "&" +
		"round=" + round + "&" +
		"amount=" + fmt.Sprintf("%.2f", amount) + "&" +
		"mtcode=" + mtcode + "&" +
		"datetime=" + datetime

	/*
	req, err := http.NewRequest("POST", url, strings.NewReader(params))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", auth)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	*/

	resp, err := resty.SetTimeout(700 * time.Millisecond).R().SetBody(params).SetHeaders(
		map[string]string{
			"Content-Type": "application/x-www-form-urlencoded",
			"Authorization": auth,
		}).Post(url)
	if err != nil {
		return nil, err
	}


	var betResponse CQBetResponse
	err = json.Unmarshal(resp.Body(), &betResponse)
	if err != nil {
		return nil, err
	}
	if betResponse.Status.Code != "0" {
		return nil, errors.New("bet end error:" + betResponse.Status.Code + "|" + betResponse.Status.Message + "|" + params)
	}
	return &betResponse, nil
}

//Balance TODO
func (s *CQService) Balance(uid string) (float64, error) {
	//url := (s.Config.URL) + "gamepool/cq9/player/balance/" + uid
	//auth := s.Config.Auth


	url := s.Url + "gamepool/cq9/player/balance/" + uid
	auth := s.Auth

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return 0, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", auth)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	var betResponse CQBetResponse
	err = json.Unmarshal(body, &betResponse)
	if err != nil {
		return 0, err
	}
	if betResponse.Status.Code != "0" {
		return 0, errors.New("get balance error:" + betResponse.Status.Code + betResponse.Status.Message)
	}
	return betResponse.Data.Balance, nil
}

func (s *CQService) DetailAuth(ordertoken string) (*CQOrderAuth, error) {

	url := s.Url + "/gamepool/cq9/game/detailtoken"
	auth := s.Auth

	params := "token=" + ordertoken

	req, err := http.NewRequest("POST", url, strings.NewReader(params))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", auth)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	res := &struct {
		Data CQOrderAuth `json:"data"`
		Status CQStatus `json:"status"`
	}{}

	err = json.Unmarshal(body, res)
	if err != nil {
		return nil, err
	}
	if res.Status.Code != "0" {
		return nil, errors.New("auth detail token error:" + res.Status.Code + res.Status.Message)
	}
	return &res.Data,nil
}