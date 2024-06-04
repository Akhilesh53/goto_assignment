package entities

import (
	"fmt"
	"reflect"
)

func InitiliaseTableId() func() uint16 {
	var tableId uint16 = 0
	return func() uint16 {
		tableId++
		return tableId
	}
}

var GetNextTableId func() uint16

func init() {
	GetNextTableId = InitiliaseTableId()
}

type Table struct {
	tableId uint16
	dbID    uint16
	name    string
	rows    []*Row
	columns []*Column
}

func CreateTable(dbId uint16, name string) *Table {
	return &Table{
		tableId: GetNextTableId(),
		name:    name,
		dbID:    dbId,
		rows:    make([]*Row, 0),
		columns: make([]*Column, 0),
	}
}

func (table *Table) GetName() string {
	return table.name
}

// get columns
func (table *Table) GetColumns() []*Column {
	return table.columns
}

func (table *Table) AddColumn(column *Column) *Table {
	//check if column types are of sring and int
	if column.columnType != STRING && column.columnType != INT {
		fmt.Println("Invalid column type :: ", column.name)
		return table
	}

	table.columns = append(table.columns, column)
	return table
}

func (table *Table) AddRow(row *Row) *Table {
	// if no columns are there, nmo data can be inserted
	if len(table.columns) == 0 {
		fmt.Println("No columns are there in the table")
		return table
	}

	// if with in the row all the data are of proper data types
	for _, rowData := range row.GetRowData() {
		if rowData.column.columnType == STRING && reflect.TypeOf(rowData.value).Kind() != reflect.String {
			fmt.Println("Invalid data type for column ", rowData.column.name)
			return table
		} else if rowData.column.columnType == INT && reflect.TypeOf(rowData.value).Kind() != reflect.Int {
			fmt.Println("Invalid data type for column ", rowData.column.name)
			return table
		}
	}

	// check for the constraints
	for _, column := range table.columns {
		for _, rowData := range row.GetRowData() {
			if column.GetName() == rowData.column.GetName() {
				for _, constraint := range column.GetConstraints() {
					// check for null and unique constraint
					if constraint == NOT_NULL {
						if rowData.value == nil {
							fmt.Println("Null constraint failed for column ", column.GetName())
							return table
						}
					} else if constraint == UNIQUE {
						for _, r := range table.rows {
							for _, d := range r.GetRowData() {
								if d.column.GetName() == column.GetName() && d.value == rowData.value {
									fmt.Println("Unique constraint failed for column ", column.GetName())
									return table
								}
							}
						}
					}
				}
			}
		}

	}

	table.rows = append(table.rows, row)
	return table
}

// get rows
func (table *Table) GetRows() []*Row {
	return table.rows
}
