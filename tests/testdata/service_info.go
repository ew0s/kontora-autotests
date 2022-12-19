package testdata

import "github.com/ew0s/kontora-autotests/swagger"

var (
	BaseServiceInfos = []swagger.ServiceInfoDto{
		{
			Id:    1,
			Price: 1000,
			ServiceDto: &swagger.ServiceDto{
				Id:   1,
				Name: "Стрижка налысо",
			},
			HairdresserDto: &swagger.HairdresserDto{
				Id:       3,
				FullName: "Стас Угольников",
			},
		},
		{
			Id:    2,
			Price: 1000,
			ServiceDto: &swagger.ServiceDto{
				Id:   2,
				Name: "Барбер стрижка",
			},
			HairdresserDto: &swagger.HairdresserDto{
				Id:       4,
				FullName: "Леха Москвинов",
			},
		},
		{
			Id:    3,
			Price: 1000,
			ServiceDto: &swagger.ServiceDto{
				Id:   4,
				Name: "Стрижка + маникюр",
			},
			HairdresserDto: &swagger.HairdresserDto{
				Id:       2,
				FullName: "Илья Шевчук",
			},
		},
		{
			Id:    4,
			Price: 1000,
			ServiceDto: &swagger.ServiceDto{
				Id:   3,
				Name: "Маникюр",
			},
			HairdresserDto: &swagger.HairdresserDto{
				Id:       1,
				FullName: "Егор Савковский",
			},
		},
	}
)
