package loader

import (
	"encoding/json"
	"os"

	"github.com/gmalheiro/playground/internal"
)

type VehicleJSONFile struct {
	path string
}

type VehicleJSON struct {
	Id              int     `json:"id"`
	Brand           string  `json:"brand"`
	Model           string  `json:"model"`
	Registration    string  `json:"registration"`
	Color           string  `json:"color"`
	FabricationYear int     `json:"year"`
	Capacity        int     `json:"passengers"`
	MaxSpeed        float64 `json:"max_speed"`
	FuelType        string  `json:"fuel_type"`
	Transmission    string  `json:"transmission"`
	Weight          float64 `json:"weight"`
	Height          float64 `json:"height"`
	Length          float64 `json:"length"`
	Width           float64 `json:"width"`
}

func NewVehicleJSONFile(path string) *VehicleJSONFile {
	return &VehicleJSONFile{
		path: path,
	}
}

func (f *VehicleJSONFile) Load() (vehicles map[int]internal.Vehicle, err error) {
	file, err := os.Open(f.path)
	if err != nil {
		return
	}
	defer file.Close()

	var vehiclesJson []VehicleJSON
	err = json.NewDecoder(file).Decode(&vehiclesJson)

	if err != nil {
		return
	}
	vehicles = make(map[int]internal.Vehicle)
	for _, vehicle := range vehiclesJson {
		vehicles[vehicle.Id] = internal.Vehicle{
			Id: vehicle.Id,
			VehicleAttributes: internal.VehicleAttributes{
				Brand:           vehicle.Brand,
				Model:           vehicle.Model,
				Registration:    vehicle.Registration,
				Color:           vehicle.Color,
				FabricationYear: vehicle.FabricationYear,
				Capacity:        vehicle.Capacity,
				MaxSpeed:        vehicle.MaxSpeed,
				FuelType:        vehicle.FuelType,
				Transmission:    vehicle.Transmission,
				Weight:          vehicle.Weight,
				Dimensions: internal.Dimensions{
					Height: vehicle.Height,
					Length: vehicle.Length,
					Width:  vehicle.Width,
				},
			},
		}
	}
	return
}
