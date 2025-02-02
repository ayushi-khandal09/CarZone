package models

import (
	"errors"
	"strconv"
	"time"

	"github.com/google/uuid"
)

type Car struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Year      string    `json:"year"`
	Brand     string    `json:"brand"`
	FuelType  string    `json:"fuel_type"`
	Engine    Engine    `json:"engine"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CarRequest struct {
	Name     string  `json:"name"`
	Year     string  `json:"year"`
	Brand    string  `json:"brand"`
	FuelType string  `json:"fuel_type"`
	Engine   Engine  `json:"engine"`
	Price    float64 `json:"price"`
}

func ValidateRequest(carReq CarRequest) error {
	if err := validateName(carReq.Name); err != nil{
		return err
	}
	if err := validateYear(carReq.Year); err != nil{
		return err
	}
	if err := validateBrand(carReq.Brand); err != nil{
		return err
	}
	if err := validateFuelType(carReq.FuelType); err != nil{
		return err
	}
	if err := validateEngine(carReq.Engine); err != nil {
		return err
	}
	if err := validatePrice(carReq.Price); err!=nil {
		return err
	}
	return nil
}

func validateName(name string) error {
	if name == "" {
		return errors.New("Name is required")
	}
	return nil
}

func validateYear(year string) error {
	if year == "" {
		return errors.New("Year is required")
	}
    _, err := strconv.Atoi(year)
	if err != nil {
		return errors.New("year must be a vaid number")
	}
	currentYear := time.Now().Year()
	yearInt,_ := strconv.Atoi(year)
	if yearInt < 1886 || yearInt > currentYear {
		return errors.New("Year must be between 188 and current year")
	}
	return nil
}

func validateBrand(brand string) error {
	if brand == "" {
		return errors.New("Brand is required")
	}
	return nil
}

func validateFuelType(fuelType string) error{
	validateFuelType := []string{"Petrol", "Diesel", "Electric", "Hybrid"}
	for _, validType := range validateFuelType {
		if fuelType == validType {
			return nil
		}
	}
	return errors.New("Fuel Type must be one of: Petrol, Diesel, Electric, Hybrid")
}

func validateEngine(engine Engine) error{
	if engine.EngineID == uuid.Nil {
		return errors.New("EngineID is required")
	}
	if engine.Displacement <= 0 {
		return errors.New("displacement must be greater than zero")
	}
	if engine.NoOfCyclinders <= 0 {
		return errors.New("noOfCyclinders must be greater than zero")
	}
	if engine.CarRange <= 0 {
		return errors.New("carRange must be greater than zero")
	}
	return nil
}

func validatePrice(price float64) error{
	if price <= 0 {
		return errors.New("Price must be greater than zero")
	}
	return nil
}