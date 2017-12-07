package casino

//Context TODO
type Context struct {
	N      *Node
	GotoHC string
	KV     map[string]interface{}
	BI     map[string]interface{}
}

//AddPair TODO
func (c *Context) AddPair(key string, value interface{}) error {
	c.KV[key] = value
	return nil
}

//RemovePair TODO
func (c *Context) RemovePair(key string, value interface{}) error {
	if _, ok := c.KV[key]; !ok {
		return ErrPairNotExist
	}
	delete(c.KV, key)
	return nil
}

//GetValue TODO
func (c *Context) GetValue(key string) (interface{}, error) {
	if _, ok := c.KV[key]; !ok {
		return nil, ErrPairNotExist
	}
	return c.KV[key], nil
}

//InjectBI TODO
func (c *Context) InjectBI(key string, value interface{}) error {
	if !c.N.Debug {
		return nil
	}
	c.BI[key] = value
	return nil
}

//AddBILineReward TODO
func (c *Context) AddBILineReward(id int, count int, reward int) error {
	if !c.N.Debug {
		return nil
	}
	biLineRewardsInContext, err := c.GetValue("biLineRewards")
	if err != nil {
		if err == ErrPairNotExist {
			biLineRewardsInContext = make([]BILineReward, 0)
		} else {
			return err
		}
	}
	biLineRewards := biLineRewardsInContext.([]BILineReward)

	lineReward := new(BILineReward)
	lineReward.ID = id
	lineReward.Count = count
	lineReward.Reward = reward
	biLineRewards = append(biLineRewards, *lineReward)

	err = c.AddPair("biLineRewards", biLineRewards)
	if err != nil {
		return err
	}
	return nil
}

//BILineReward TODO
type BILineReward struct {
	ID     int
	Count  int
	Reward int
}
