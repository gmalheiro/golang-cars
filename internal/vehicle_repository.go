package internal

type VehicleRepository interface {
	FindAll() (v map[int]Vehicle, err error)
}
