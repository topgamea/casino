package casino

//Context TODO
type Context struct {
	N      *Node
	GotoHC string
	KV     map[string]interface{}
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
