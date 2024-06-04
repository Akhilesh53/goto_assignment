package main

import (
	"goto/entities"
)

/*
Task is to implement a simplified version of an in-memory database. Plan your design according to the level specifications below:

In-memory database should support basic operations to manipulate records, fields, and values within fields.
It should be possible to create or delete tables in a database.
The supported column types are string and int.
It should be possible to insert records in a table.
It should be possible to print all records in a table.
It should be possible to display records if we query based on column values
*/

func main() {
	// create a db
	personsDB := entities.CreateDB("Persons DB")
	//fmt.Println(personsDB.GetName())

	// // create a table with int he db defined above
	personsTable := personsDB.AddTable("Persons")
	//fmt.Println(personsTable.GetName())

	// // add columns to table
	nameColumn := entities.NewColumn("Name", entities.STRING).AddConstraint(entities.NOT_NULL).AddConstraint(entities.UNIQUE)
	ageColumn := entities.NewColumn("Age", entities.INT)
	//testColumn := entities.NewColumn("Default", entities.DEFAULT)
	personsTable.AddColumn(nameColumn).AddColumn(ageColumn) //.AddColumn(testColumn)

	// add rows to the table
	row1 := entities.NewRow().AddRowData(entities.NewRowData(nameColumn, "Akhilesh")).AddRowData(entities.NewRowData(ageColumn, 30))
	row2 := entities.NewRow().AddRowData(entities.NewRowData(nameColumn, "Test User")).AddRowData(entities.NewRowData(ageColumn, 10))
	row3 := entities.NewRow().AddRowData(entities.NewRowData(nameColumn, "Test User")).AddRowData(entities.NewRowData(ageColumn, 10))

	personsTable.AddRow(row1).AddRow(row2).AddRow(row3)

	// print all the rows in the table by table name
	//personsDB.PrintAllTableRows(personsTable.GetName())

	// get records for a particular column value
	//personsDB.GetRecordsForColumnValue(personsTable.GetName(), "Name", "Akhilesh")

	// // delete table from db
	// personsDB.DeleteTable(personsTable.GetName())
}
