package main

type Table struct {
	Name           string
	Autoincreament bool
	Primarykey     []string
	Foreignkey     []Forienkey
	Columns        []Attribute
}

type Attribute struct {
	Name    string
	Type    interface{}
	Size    int
	Default interface{}
	CanNull bool
	Foreign string
}

type Forienkey struct {
	Key       string
	TableName string
	RefName   string
}
