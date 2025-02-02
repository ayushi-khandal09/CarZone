package models

import (
	"errors"

	"github.com/google/uuid"
)

type Engine struct {
	EngineID       uuid.UUID `json:"engine_id"`
	Displacement   int64     `json:"displacement"`
	NoOfCyclinders int64     `json:"noOfCyclinders"`
	CarRange       int64     `json:"carRange"`
}

type EngineRequest struct {
	Displacement int64 `json:"displacement"`
	NoOfCyclinders int64     `json:"noOfCyclinders"`
	CarRange       int64     `json:"carRange"`
}

func ValidateEngineRequest(EngineReq EngineRequest) error {
	if err := validateDisplacement(EngineReq.Displacement); err != nil {
		return err
	}
	if err := validateNoOfCylinders(EngineReq.NoOfCyclinders); err != nil {
		return err
	}
	if err := validateCarRange(EngineReq.CarRange); err != nil {
		return err
	}
	return nil
}

func validateDisplacement(displacement int64) error {
	if displacement <= 0 {
		return errors.New("displacement must be greater than zero")
	}
	return nil
}

func validateNoOfCylinders(noOfCyclinders int64) error {
	if noOfCyclinders <= 0 {
		return errors.New("noOfCylinders must be greater than zero")
	}
	return nil
}

func validateCarRange(carRange int64) error {
	if carRange <= 0 {
		return errors.New("carRange must be greater than zero")
	}
	return nil
}
