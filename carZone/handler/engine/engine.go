package engine

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/ayushi-khandal09/carZone/models"
	"github.com/ayushi-khandal09/carZone/service"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type EngineHandler struct {
	service service.EngineServiceInterface
}

func NewEngineHandler(service service.EngineServiceInterface) *EngineHandler {
	return &EngineHandler{
		service: service,
	}
}

func (e *EngineHandler) GetEngineById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)
	id := vars["id"]

	resp, err := e.service.GetEngineById(ctx, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	body, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Print(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, err = w.Write(body)
	if err != nil {
		log.Println("Error writing Response:", err)
	}
}

func (e *EngineHandler) CreateEngine(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Error reading request body:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var engineReq models.EngineRequest
	err = json.Unmarshal(body, &engineReq)
	if err != nil {
		log.Println("Error Unmarshalling the engine request body:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	createdEngine, err := e.service.CreateEngine(ctx, &engineReq)
	if err != nil {
		log.Println("Error while creating Engine:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	respBody, err := json.Marshal(createdEngine)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Print(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, _ = w.Write(respBody)
}

func (e *EngineHandler) UpdateEngine(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	params := mux.Vars(r)
	id := params["id"]
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Error reading request body:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var engineReq models.EngineRequest
	err = json.Unmarshal(body, &engineReq)
	if err != nil {
		log.Println("Error Unmarshalling the engine request body:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	updatedEngine, err := e.service.UpdateEngine(ctx, id, &engineReq)
	if err != nil {
		log.Println("Error while updating Engine:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	respBody, err := json.Marshal(updatedEngine)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Print(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, _ = w.Write(respBody)
}

func (e *EngineHandler) DeleteEngine(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	params := mux.Vars(r)
	id := params["id"]

	deleteEngine, err := e.service.DeleteEngine(ctx, id)
	if err != nil {
		log.Println("Error while deleting Engine:", err)
		w.WriteHeader(http.StatusInternalServerError)
		response := map[string]string{"error": "Invalid ID or Engine not Found"}
		jsonResponse, _ := json.Marshal(response)
		_, _ = w.Write(jsonResponse)
		return
	}

	if deleteEngine.EngineID == uuid.Nil {
		w.WriteHeader(http.StatusNotFound)
		response := map[string]string{"error": "Engine Not Found"}
		jsonResponse, _ := json.Marshal(response)
		_, _ = w.Write(jsonResponse)
		return 
	}
	jsonResponse, err := json.Marshal(deleteEngine)
	if err != nil {
		log.Print("Error while marshalling delete engine response:", err)
		w.WriteHeader(http.StatusInternalServerError)
		response := map[string]string {"error": "Internal server error"}
		jsonResponse, _ := json.Marshal(response)
		_, _ = w.Write(jsonResponse)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, _ = w.Write(jsonResponse)
}
