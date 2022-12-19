package entities

type Service struct {
	Id   int64  `db:"id"`
	Name string `db:"name"`
}
