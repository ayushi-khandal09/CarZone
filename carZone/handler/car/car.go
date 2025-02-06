package car

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/ayushi-khandal09/carZone/models"
	"github.com/ayushi-khandal09/carZone/service"
	"github.com/gorilla/mux"
)

type CarHandler struct {
	service service.CarServiceInterface
}

func NewCarHandler(service service.CarServiceInterface) *CarHandler {
	return &CarHandler{
		service: service,
	}
}

func (h *CarHandler) GetCarById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)
	id := vars["id"]

	resp, err := h.service.GetCarById(ctx, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error : ", err)
		return
	}

	body, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error : ", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, err = w.Write(body)
	if err != nil {
		log.Println("Error Writing Response : ", err)
	}
}

func (h *CarHandler) GetCarByBrand(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	brand := r.URL.Query().Get("brand")
	isEngine := r.URL.Query().Get("isEngine") == "true"

	resp, err := h.service.GetCarsByBrand(ctx, brand, isEngine)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error :", err)
		return
	}

	body, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error : ", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, err = w.Write(body)
	if err != nil {
		log.Println("Error Writing Response : ", err)
	}
}

func (h *CarHandler) CreateCar(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Error while reading request body :", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var carReq models.CarRequest
	err = json.Unmarshal(body, &carReq)
	if err != nil {
		log.Println("Error while Unmarshalling Request body", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	createdCar, err := h.service.CreateCar(ctx, &carReq)
	if err != nil {
		log.Println("Error creating car :", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	responseBody, err := json.Marshal(createdCar)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error while marshalling:", err)
		return 
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	_, err = w.Write(responseBody)
	if err != nil {
		log.Println("Error Writing Response : ", err)
	}
}

func (h *CarHandler) UpdateCar(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	params := mux.Vars(r)
	id := params["id"]

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Error while reading request body :", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var carReq models.CarRequest
	err = json.Unmarshal(body, &carReq)
	if err != nil {
		log.Println("Error while Unmarshalling Request body", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	updatedCar, err := h.service.UpdateCar(ctx, id, &carReq)
	if err != nil {
		log.Println("Error updating car :", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	responseBody, err := json.Marshal(updatedCar)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error while marshalling:", err)
		return 
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, err = w.Write(responseBody)
	if err != nil {
		log.Println("Error Writing Response : ", err)
	}
}

func (h *CarHandler) DeleteCar(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	params := mux.Vars(r)
	id := params["id"]

	deleteCar, err := h.service.DeleteCar(ctx, id)
	if err != nil {
		log.Println("Error deleting car :", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	body, err := json.Marshal(deleteCar)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error while marshalling:", err)
		return 
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, err = w.Write(body)
	if err != nil {
		log.Println("Error Writing Response : ", err)
	}
}
