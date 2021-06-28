package main

import (
	"database/sql"
	"fmt"
)

func exmaple_create() {
	attri1 := []Attribute{
		Attribute{Name: "roles", Type: "varchar", Size: 30, CanNull: false},
		Attribute{Name: "created_date", Type: sql.NullTime{}, CanNull: true, Default: "CURRENT_TIMESTAMP"},
		Attribute{Name: "updated_date", Type: sql.NullTime{}, CanNull: true, Default: "CURRENT_TIMESTAMP"},
	}
	roles1 := Table{
		Name:           "roles",
		Autoincreament: true,
		Primarykey:     []string{"roles"},
		Columns:        attri1,
	}
	temp := GenTable(roles1)
	fmt.Println(temp)
	fmt.Println()

	attri2 := []Attribute{
		Attribute{Name: "username", Type: "varchar", Size: 255, CanNull: false},
		Attribute{Name: "password", Type: "varchar", Size: 30, CanNull: false},
		Attribute{Name: "roles", Type: "varchar", Size: 30, CanNull: false},
		Attribute{Name: "created_date", Type: sql.NullTime{}, CanNull: true, Default: "CURRENT_TIMESTAMP"},
		Attribute{Name: "updated_date", Type: sql.NullTime{}, CanNull: true, Default: "CURRENT_TIMESTAMP"},
	}
	users := Table{
		Name:           "users",
		Autoincreament: true,
		Primarykey:     []string{"username"},
		Columns:        attri2,
		Foreignkey: []Forienkey{
			Forienkey{TableName: "roles", Key: "roles", RefName: "roles"},
		},
	}
	temp = GenTable(users)
	fmt.Println(temp)
}
