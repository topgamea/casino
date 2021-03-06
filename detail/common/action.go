package common

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"go.pkg.wesai.com/p/base_lib/log"
	"time"
)

var (
	CurO   orm.Ormer
	logger = log.DLogger()
)

func InitMysqlModels(dsn string, syncDb bool) error {
	orm.DefaultTimeLoc = time.Local

	var err error
	err = orm.RegisterDriver("mysql", orm.DRMySQL)
	if err != nil {
		return err
	}
	err = orm.RegisterDataBase("default", "mysql", dsn, 30, 30)
	if err != nil {
		return err
	}

	orm.RegisterModel(new(Game), new(Round), new(SpinNew), new(GenericReward))
	if syncDb {
		err = orm.RunSyncdb("default", false, true)
		if err != nil {
			return err
		}

	}
	CurO = orm.NewOrm()
	CurO.Using("default")
	return err
}

func CreateGame(g *Game, O orm.Ormer) error {
	rows, err := O.InsertOrUpdate(g)
	if err != nil {
		//logger.Error(err.Error())
		return err
	}
	if rows != 1 {
		//return errors.New("unexpected insert game")
	}

	return nil
}

func InsertMultiRound(rs []*Round, O orm.Ormer) (int64, error) {
	return O.InsertMulti(len(rs), rs)
}

func InsertMultiSpin(spins []*SpinNew, O orm.Ormer) (int64, error) {
	return O.InsertMulti(len(spins), spins)
}

func InsertMultiReward(rewards []*GenericReward, O orm.Ormer) (int64, error) {
	return O.InsertMulti(len(rewards), rewards)
}

func InsertFreeSpin(r *Round, O orm.Ormer) error {
	_, err := O.InsertOrUpdate(r, "EndTime", "TotalReward")
	if err != nil {
		logger.Error(err)
		return err
	}
	for _, sp := range r.Spins {
		err := InsertSpin(sp, r, O)
		if err != nil {
			logger.Error(err)
			return err
		}
	}

	return nil

}

func InsertRound(r *Round, O orm.Ormer) error {

	_, err := O.Insert(r)
	if err != nil {
		logger.Error(err)
		return err
	}
	for _, sp := range r.Spins {
		err := InsertSpin(sp, r, O)
		if err != nil {
			logger.Error(err)
			return err
		}
	}

	return nil
}

func InsertSpin(s *SpinNew, parentRound *Round, O orm.Ormer) error {
	if s.Round == nil {
		s.Round = parentRound
	}
	_, err := O.Insert(s)
	if err != nil {
		logger.Error(err)
		return err
	}

	for _, rw := range s.RewardDetails {
		err := InsertGenericReward(rw, s, parentRound, O)
		if err != nil {
			logger.Error(err)
			return err
		}
	}

	return nil
}

func InsertGenericReward(gr *GenericReward, parentSpin *SpinNew, parentRound *Round, O orm.Ormer) error {

	if gr.Spin == nil {
		gr.Spin = parentSpin
	}

	_, err := O.Insert(gr)
	if err != nil {
		return err
	}

	return nil
}

func GetRound(round string, O orm.Ormer) (*Round, error) {
	var err error
	r := &Round{
		RoundId: round,
	}

	err = O.Read(r)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	_, err = O.LoadRelated(r, "Game", true)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	_, err = O.LoadRelated(r, "Spins", true)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	for _, sp := range r.Spins {
		_, err := O.LoadRelated(sp, "RewardDetails", true)
		if err != nil {
			logger.Error(err)
			return nil, err
		}
	}

	return r, nil

}

func GetRoundOnly(round string, O orm.Ormer) (*Round, error) {
	var err error
	r := &Round{
		RoundId: round,
	}

	err = O.Read(r)
	if err != nil {
		return nil, err
	}

	_, err = O.LoadRelated(r, "Game", true)
	if err != nil {
		return nil, err
	}

	return r, nil

}
