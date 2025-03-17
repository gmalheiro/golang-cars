package service

import "github.com/gmalheiro/playground/internal"

type VehicleDefault struct {
	rp internal.VehicleRepository
}

func NewVehicleDefault(rp internal.VehicleRepository) *VehicleDefault {
	return &VehicleDefault{rp: rp}
}

func (s *VehicleDefault) FindAll() (v map[int]internal.Vehicle, err error) {
	v, err = s.rp.FindAll()
	return
}

func (s *VehicleDefault) Create(vh internal.Vehicle) (v internal.Vehicle, err error) {
	v, err = s.rp.Create(vh)
	return
}
