package entities

func InitiliaseRowId() func() uint16 {
	var rowId uint16 = 0
	return func() uint16 {
		rowId++
		return rowId
	}
}

var GetNextRowId func() uint16

func init() {
	GetNextRowId = InitiliaseTableId()
}

type Row struct {
	rowId uint16
	data  []*RowData
}

func NewRow() *Row {
	return &Row{rowId: GetNextRowId(),
		data: make([]*RowData, 0),
	}
}

func (row *Row) AddRowData(rowData *RowData) *Row {
	row.data = append(row.data, rowData)
	return row
}

type RowData struct {
	column *Column
	value  interface{}
}

func NewRowData(column *Column, value interface{}) *RowData {
	return &RowData{column: column, value: value}
}

func (row *Row) GetRowData() []*RowData {
	return row.data
}
