package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/bootcamp-go/web/request"
	"github.com/bootcamp-go/web/response"
	"github.com/gmalheiro/playground/internal"
	"github.com/go-chi/chi/v5"
)

type VehicleJSON struct {
	ID              int     `json:"id"`
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

type VehicleDefault struct {
	sv internal.VehicleService
}

func NewVehicleDefault(sv internal.VehicleService) *VehicleDefault {
	return &VehicleDefault{sv: sv}
}

func (h *VehicleDefault) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		v, err := h.sv.FindAll()
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, nil)
			return
		}

		if errors.Is(err, internal.ErrNotAvailableCars) {
			response.JSON(w, http.StatusNotFound, nil)
		}

		data := make(map[int]VehicleJSON)
		for key, value := range v {
			data[key] = VehicleJSON{
				ID:              value.Id,
				Brand:           value.Brand,
				Model:           value.Model,
				Registration:    value.Registration,
				Color:           value.Color,
				FabricationYear: value.FabricationYear,
				Capacity:        value.Capacity,
				MaxSpeed:        value.MaxSpeed,
				FuelType:        value.FuelType,
				Transmission:    value.Transmission,
				Weight:          value.Weight,
				Height:          value.Height,
				Length:          value.Length,
				Width:           value.Width,
			}
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})
	}
}

func (h *VehicleDefault) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var vehicle internal.Vehicle

		if err := request.JSON(r, &vehicle); err != nil {
			body := map[string]any{"message": "invalid request", "data": err.Error()}
			response.JSON(w, http.StatusUnprocessableEntity, body)
			return
		}

		v, err := h.sv.Create(vehicle)

		body := map[string]any{"message": "invalid request", "data": err.Error()}
		if errors.Is(err, internal.ErrFieldsNotPropperlyField) {
			response.JSON(w, http.StatusBadRequest, body)
			return
		}
		if errors.Is(err, internal.ErrExistingItem) {
			response.JSON(w, http.StatusConflict, body)
			return
		}

		if err != nil {
			response.JSON(w, http.StatusBadRequest, body)
			return
		}

		response.JSON(w, http.StatusCreated, v)
	}
}

func (h *VehicleDefault) GetByWeight() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		minStr := r.URL.Query().Get("min")
		maxStr := r.URL.Query().Get("max")

		if minStr == "" || maxStr == "" {
			body := map[string]any{"message": "invalid query params", "data": nil}
			response.JSON(w, http.StatusBadRequest, body)
			return
		}

		min, minErr := strconv.ParseFloat(minStr, 64)
		max, maxErr := strconv.ParseFloat(maxStr, 64)

		if maxErr != nil || minErr != nil {
			body := map[string]any{"message": "float conversion failed", "data": nil}
			response.JSON(w, http.StatusBadGateway, body)
			return
		}

		data, err := h.sv.GetByWeight(min, max)
		if errors.Is(err, internal.ErrNotAvailableCars) {
			body := map[string]any{"message": "Cars not found", "data": err}
			response.JSON(w, http.StatusNotFound, body)
		}
		if err != nil {
			body := map[string]any{"message": "error in finding by weight", "data": err}
			response.JSON(w, http.StatusBadRequest, body)
		}
		body := map[string]any{"message": "cars retrieved sucess", "data": data}
		response.JSON(w, http.StatusOK, body)
	}
}

func (h *VehicleDefault) UpdateFuel() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		if idStr == "" {
			body := map[string]any{"message": "Missing id field", "data": nil}
			response.JSON(w, http.StatusBadRequest, body)
		}

		id, idErr := strconv.Atoi(idStr)
		if idErr != nil {
			body := map[string]any{"message": "error while parsing string to int", "data": nil}
			response.JSON(w, http.StatusBadGateway, body)
			return
		}

		var vehicle internal.Vehicle
		if err := request.JSON(r, &vehicle); err != nil {
			body := map[string]any{"message": "invalid request body", "data": nil}
			response.JSON(w, http.StatusUnprocessableEntity, body)
			return
		}
		v, err := h.sv.UpdateFuel(vehicle, id)
		if errors.Is(err, internal.ErrItemNotFound) {
			body := map[string]any{"message": "Vehicle not found", "data": nil}
			response.JSON(w, http.StatusNotFound, body)
			return
		}
		if err != nil {

		}
		response.JSON(w, http.StatusOK, v)
		return
	}
}
