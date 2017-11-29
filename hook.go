package casino

//HookFunc TODO
type HookFunc func(*Context) error

//HookChain TODO
type HookChain struct {
	Hooks []HookFunc
}

func (hc *HookChain) addHook(hookF HookFunc) error {
	hc.Hooks = append(hc.Hooks, hookF)
	return nil
}

func (hc *HookChain) execute(c *Context) error {
	for _, hookFunc := range hc.Hooks {
		err := hookFunc(c)
		if err != nil {
			return err
		}
		if c.GotoHC != "" {
			break
		}
	}
	return nil
}
