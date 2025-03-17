package internal

type VehicleRepository interface {
	FindAll() (v map[int]Vehicle, err error)
	Create(vh Vehicle) (v Vehicle, err error)
	GetByWeight(min, max float64) (v map[int]Vehicle, err error)
	UpdateFuel(vh Vehicle, id int) (v Vehicle, err error)
}
