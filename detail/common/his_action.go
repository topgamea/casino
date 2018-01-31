package common


import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

var (
	HistO     orm.Ormer
)


func InitMysqlHisModels(dsn string, syncDb bool) error {
	orm.DefaultTimeLoc = time.Local

	var err error
	err = orm.RegisterDriver("mysql", orm.DRMySQL)
	if err != nil {
		return err
	}
	err = orm.RegisterDataBase("his_db", "mysql", dsn, 30, 30)
	if err != nil {
		return err
	}

	orm.RegisterModel(new(Game), new(Round), new(Spin), new(GenericReward))
	if syncDb {
		err = orm.RunSyncdb("his_db", false, true)
		if err != nil {
			return err
		}

	}
	HistO = orm.NewOrm()
	HistO.Using("his_db")
	return err
}


//dividingTime 之前的记录会被移动到历史库里,执行时刻，如果不设则直接移动，设置格式 上午6点就设为6 [0-23]
func GetOldData(dividingTime time.Time, O orm.Ormer) ([]*Round,error)  {
	container := []*Round{}
	n,err := O.QueryTable(&Round{}).Filter("CreatedAt__lt",dividingTime).RelatedSel().All(&container)

	if err != nil {
		logger.Error(err.Error())
		return nil,err
	}
	for i := 0; i < int(n); i++ {
		r := container[i]


		spins := []*Spin{}
		_,err := O.QueryTable(&Spin{}).Filter("Round",r.RoundId).All(&spins)
		if err != nil {
			logger.Error(err.Error())
			return nil,err
		}

		r.Spins = spins

		for _,s := range spins {
			gr := []*GenericReward{}
			_,err = O.QueryTable(&GenericReward{}).Filter("Spin",s.Id).All(&gr)
			if err != nil {
				logger.Error(err.Error())
				return nil,err
			}
			s.RewardDetails = gr
		}

		/*
		data,err := json.Marshal(r)
		if err != nil {
			logger.Error(err)
		}
		logger.Infof(string(data))
		*/
	}
	return container,nil
}

//dividingTime 之前的记录会被移动到历史库里,执行时刻，如果不设则直接移动，设置格式 上午6点就设为6 [0-23]
// Oc 当前库 Oh 历史库
func MoveOldDataToHis(dividingTime time.Time,execHour int, Oc,Oh orm.Ormer) error {
	tn := time.Now().Local()
	if execHour != -1 && tn.Hour() != execHour {
		return nil
	}

	rs,err := GetOldData(dividingTime,Oc)
	if err != nil {
		return err
	}

	for _,r := range rs {
		err = InsertRound(r,Oh)
		if err != nil {
			logger.Error(err)
			continue
		}
		err = deleteRound(r,Oc)
		if err != nil {
			logger.Error(err)
			continue
		}
	}

	return nil

}

func deleteRound(r *Round,O orm.Ormer) error {
	var err error
	for _,sp := range r.Spins {
		for _,gr := range sp.RewardDetails {
			_,err = O.Delete(gr)
			if err != nil {
				logger.Error(err.Error())
				continue
			}
		}
		_,err = O.Delete(sp)
		logger.Error(err.Error())
		continue
	}
	_, err = O.Delete(r)
	if err != nil {
		logger.Error(err.Error())
	}
	return err
}

