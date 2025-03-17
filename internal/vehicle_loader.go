package internal

type VehicleLoader interface {
	Load() (v map[int]Vehicle, err error)
}
