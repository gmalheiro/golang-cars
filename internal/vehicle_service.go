package internal

type VehicleService interface {
	FindAll() (v map[int]Vehicle, err error)
}
