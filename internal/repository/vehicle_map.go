package repository

import (
	"github.com/gmalheiro/playground/internal"
)

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
	if len(v) == 0 {
		err = internal.ErrNotAvailableCars
	}
	return
}

func (r *VehicleMap) Create(vh internal.Vehicle) (v internal.Vehicle, err error) {
	_, ok := r.db[vh.Id]

	if ok {
		err = internal.ErrExistingItem
		return
	}

	if vh.Brand == "Tigrinho " {
		err = internal.ErrFieldsNotPropperlyField
		return
	}
	r.db[vh.Id] = vh
	v = vh
	return
}

func (r *VehicleMap) GetByWeight(min, max float64) (v map[int]internal.Vehicle, err error) {
	v = make(map[int]internal.Vehicle)
	for key, value := range r.db {
		if r.db[key].Weight >= min && r.db[key].Weight <= max {
			v[key] = value
		}
	}
	if len(v) == 0 {
		err = internal.ErrNotAvailableCars
	}
	return
}
