package model

type ID string

type Document struct {
	ID        ID
	Name      string
	Documents []ID
}
