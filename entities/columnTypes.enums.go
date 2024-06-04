package entities

type ColumnType int

const (
	STRING ColumnType = iota
	INT
	DEFAULT
)

func (c ColumnType) String() string {
	return [...]string{"STRING", "INT", "DEFAULT"}[c]
}
