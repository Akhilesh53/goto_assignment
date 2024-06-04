package entities

import "fmt"

func InitiliaseDBId() func() uint16 {
	var dbId uint16 = 0
	return func() uint16 {
		dbId++
		return dbId
	}
}

var GetNextDBId func() uint16

func init() {
	GetNextDBId = InitiliaseDBId()

}

type DB struct {
	dbId   uint16
	name   string
	tables map[string]*Table
}

func CreateDB(name string) *DB {
	return &DB{
		dbId:   GetNextDBId(),
		name:   name,
		tables: make(map[string]*Table),
	}
}

// get set methods for DB
func (db *DB) GetDBId() uint16 {
	return db.dbId
}

func (db *DB) SetDBId(dbId uint16) {
	db.dbId = dbId
}

func (db *DB) GetName() string {
	return db.name
}

func (db *DB) SetName(name string) {
	db.name = name
}

func (db *DB) GetTables() map[string]*Table {
	return db.tables
}

func (db *DB) SetTables(tables map[string]*Table) {
	db.tables = tables
}

// create table
func (db *DB) AddTable(tableName string) *Table {
	table := CreateTable(db.GetDBId(), tableName)
	db.tables[tableName] = table
	return table
}

func (db *DB) PrintAllTableRows(tableName string) {
	// get table
	if table, ok := db.tables[tableName]; ok {
		for _, row := range table.GetRows() {
			fmt.Print("row id :: ", row.rowId, " ")
			for _, rowData := range row.GetRowData() {
				fmt.Print(rowData.column.name, " :: ", rowData.value, " ")
			}
			fmt.Println()
		}
	} else {
		fmt.Println("Table not found")
	}
}

// delete table
func (db *DB) DeleteTable(tableName string) {
	if _, ok := db.tables[tableName]; ok {
		delete(db.tables, tableName)
	} else {
		fmt.Println("Table not found")
	}
}


// get records for a particular column value
func (db *DB) GetRecordsForColumnValue(tableName string, columnName string, value interface{}){
	if table, ok := db.tables[tableName]; ok {
		for _, row := range table.GetRows() {
			for _, rowData := range row.GetRowData() {
				if rowData.column.name == columnName && rowData.value == value {
					fmt.Print("row id :: ", row.rowId, " ")
					for _, rowData := range row.GetRowData() {
						fmt.Print(rowData.column.name, " :: ", rowData.value, " ")
					}
				}
			}
		}
	} else {
		fmt.Println("Table not found")
	}
}