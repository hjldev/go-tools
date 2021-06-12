package model

type Table struct {
	Name    string   `json:"name"`
	Comment string   `json:"comment"`
	Columns []Column `json:"columns,omitempty"`
}
