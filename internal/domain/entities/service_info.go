package entities

type ServiceInfo struct {
	HairdresserID int64 `db:"hairdresser_id"`
	Id            int64 `db:"id"`
	Price         int64 `db:"price"`
	ServiceID     int64 `db:"service_id"`
}
