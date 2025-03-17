package repository

import "github.com/gmalheiro/playground/internal"

type VehicleMap struct {
	db map[int]internal.Vehicle
}

func NewVehicleMap(db map[int]internal.Vehicle) *VehicleMap {
	defaultDb := make(map[int]internal.Vehicle)
	if db != nil {
		defaultDb = db
	}
	return &VehicleMap{db: defaultDb}
}

func (r *VehicleMap) FindAll() (v map[int]internal.Vehicle, err error) {
	v = make(map[int]internal.Vehicle)
	for key, value := range r.db {
		v[key] = value
	}
	return
}
