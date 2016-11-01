package models

type Model struct {
	Id       int64
	CreatedAt  int64  `db:"created_at"`
}

