package entities

type Column struct {
	name        string
	columnType  ColumnType
	constraints []Constraints
}

func NewColumn(name string, columnType ColumnType) *Column {
	return &Column{name: name, columnType: columnType}
}

func (column *Column) GetName() string {
	return column.name
}

// add constraint
func (column *Column) AddConstraint(constraint Constraints) *Column {
	column.constraints = append(column.constraints, constraint)
	return column
}

// get constraints
func (column *Column) GetConstraints() []Constraints {
	return column.constraints
}

