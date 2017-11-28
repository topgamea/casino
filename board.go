package casino

//Board TODO
type Board struct {
	Rows    int
	Columns int
	Slots   []*Slot
}

//NewBoard TODO
func NewBoard(rows int, columns int) (*Board, error) {
	b := new(Board)
	b.Rows = rows
	b.Columns = columns
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			s := new(Slot)
			s.X = i
			s.Y = j
			b.Slots = append(b.Slots, s)
		}
	}
	return b, nil
}
