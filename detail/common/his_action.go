package common


import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
	"github.com/gin-gonic/gin/json"
)

var (
	HO     orm.Ormer
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
	HO = orm.NewOrm()
	HO.Using("his_db")
	return err
}


//dividingTime 之前的记录会被移动到历史库里,执行时刻，如果不设则直接移动，设置格式 上午6点就设为6 [0-23]
func MoveOldDataToHis(dividingTime time.Time,execHour int) error {
	tn := time.Now().Local()
	if execHour != -1 && tn.Hour() != execHour {
		return nil
	}
	container := []*Round{}
	n,err := O.QueryTable(&Round{}).Filter("CreatedAt__lt",dividingTime).RelatedSel().All(&container)

	if err != nil {
		logger.Error(err.Error())
		return err
	}
	logger.Infof("=========================")
	logger.Infof("%v",n)
	for i := 0; i < int(n); i++ {
		r := container[i]

		spins := []*Spin{}
		sn,err := O.QueryTable(&Spin{}).Filter("Round",r.RoundId).All(&spins)
		if err != nil {
			logger.Error(err.Error())
			return err
		}
		logger.Infof("%v",sn)

		r.Spins = spins

		for _,s := range spins {
			gr := []*GenericReward{}
			_,err = O.QueryTable(&GenericReward{}).Filter("Round",r.RoundId,"Spin",s.Id).All(&gr)
			if err != nil {
				logger.Error(err.Error())
				return err
			}
			s.RewardDetails = gr
		}
		data,err := json.Marshal(r)
		if err != nil {
			logger.Error(err)
		}
		logger.Infof(string(data))


	}
	return nil
}

