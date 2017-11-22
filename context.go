package casino

import "errors"

//Context TODO
type Context struct {
	N  *Node
	KV map[string]interface{}
}

//AddPair TODO
func (c *Context) AddPair(key string, value interface{}) error {
	c.KV[key] = value
	return nil
}

//RemovePair TODO
func (c *Context) RemovePair(key string, value interface{}) error {
	if _, ok := c.KV[key]; !ok {
		return errors.New("pair not exist")
	}
	delete(c.KV, key)
	return nil
}
