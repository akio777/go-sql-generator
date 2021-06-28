package main

import (
	// "database/sql"
	"database/sql"
	"fmt"
	"reflect"
	"time"
)

// Created_date: sql.NullTime{Valid: true, Time: time.Now()},
// Updated_date: sql.NullTime{Valid: false, Time: time.Now()},
func GenTable(table Table) string {
	//* table_name, id, pk, etc, fk
	sql_stm := `CREATE TABLE IF NOT EXISTS ` + fmt.Sprintf("%v(", table.Name)
	if table.Autoincreament {
		sql_stm += ` id SERIAL NOT NULL,`
	}
	for index, ele := range table.Columns {
		temp := ``
		if (reflect.TypeOf(ele.Type) != reflect.TypeOf(sql.NullTime{})) && (reflect.TypeOf(ele.Type) != reflect.TypeOf(time.Time{})) {
			temp = fmt.Sprintf(" %v %v(%v)", ele.Name, ele.Type, ele.Size)
			if ele.Default != nil {
				temp += fmt.Sprintf(" DEFAULT %v", ele.Default)
			}
		} else {
			temp = fmt.Sprintf(" %v %v DEFAULT %v", ele.Name, "TIMESTAMP", ele.Default)
		}
		if !ele.CanNull {
			temp += ` NOT NULL`
		}
		if index < len(table.Columns)-1 {
			temp += ","
		}
		sql_stm += temp
	}
	if len(table.Primarykey) > 0 {
		temp := `, PRIMARY KEY (`
		for index, ele := range table.Primarykey {
			temp += fmt.Sprintf(" %v", ele)
			if index < len(table.Primarykey)-1 {
				temp += `,`
			}
		}
		sql_stm += temp + ")"
		if len(table.Foreignkey) > 0 {
			sql_stm += ","
		}
	}
	if len(table.Foreignkey) > 0 {
		for index, ele := range table.Foreignkey {
			temp := ` FOREIGN KEY ` + fmt.Sprintf("(%v) REFERENCES %v(%v)", ele.Key, ele.TableName, ele.TableName)
			if index < len(table.Foreignkey)-1 {
				temp += ","
			}
			sql_stm += temp
		}
	}
	sql_stm += ");"
	return sql_stm
}

func main() {
	fmt.Println("Running main.go")

}
