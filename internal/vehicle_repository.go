package internal

type VehicleRepository interface {
	FindAll() (v map[int]Vehicle, err error)
	Create(vh Vehicle) (v Vehicle, err error)
}
