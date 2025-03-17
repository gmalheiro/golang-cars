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

func (s *VehicleDefault) GetByWeight(min, max float64) (v map[int]internal.Vehicle, err error) {
	v, err = s.rp.GetByWeight(min, max)
	return
}

func (s *VehicleDefault) UpdateFuel(vh internal.Vehicle, id int) (v internal.Vehicle, err error) {
	v, err = s.rp.UpdateFuel(vh, id)
	return
}
