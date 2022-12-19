package entities

type Hairdresser struct {
	Description string `db:"description"`
	FullName    string `db:"full_name"`
	Id          int64  `db:"id"`
	Phone       string `db:"phone"`
}
