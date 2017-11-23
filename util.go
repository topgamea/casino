package casino

//WildReplace TODO
func WildReplace(ids []int, except []int, i int, flag int) (in bool) {
	in = false
	for _, s := range ids {
		if s == i {
			in = true
			break
		}
	}
	for _, s := range except {
		if s == flag {
			in = false
		}
	}
	return in
}
