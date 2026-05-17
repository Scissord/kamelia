package handlers

import (
	"encoding/json"
	"net/http"
	"task-app/internal/models"
	"task-app/internal/repository"

	"github.com/jackc/pgx/v5"
)

type ClientHandler struct {
	repository *repository.ClientRepository
}

func NewClientHandler(repo *repository.ClientRepository) *ClientHandler {
	return &ClientHandler{repository: repo}
}

// 1) GET CLIENTS
func (h *ClientHandler) GetClients(w http.ResponseWriter, r *http.Request) {
	// 1. TAKE ROWS FROM REPOSITORY
	clients, err := h.repository.GetClients()

	// 2. CHECK FOR ERRORS
	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to fetch clients")
		return
	}

	// 3. RETURN SUCCESS
	writeJSON(w, http.StatusOK, clients)
}

// 2) CREATE CLIENT
func (h *ClientHandler) CreateClient(w http.ResponseWriter, r *http.Request) {
	// 1. CLOSE BODY CONNECTION (закрыть поток)
	defer r.Body.Close()

	// 2. CHECK FOR CONTENT-TYPE APPLICATION/JSON
	if r.Header.Get("Content-Type") != "application/json" {
		writeError(w, http.StatusUnsupportedMediaType, "content type must be application/json")
		return
	}

	// 3. CREATE VARIABLE TO FILL BODY
	var input models.CreateClientInput

	// 4. FILL BODY
	err := json.NewDecoder(r.Body).Decode(&input)

	// 5. CHECK FOR ERROR
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid json")
		return
	}

	// 6. CHECK FOR INPUT VALIDATION
	if input.FirstName == "" {
		writeError(w, http.StatusBadRequest, "first_name is required")
		return
	}

	if input.LastName == "" {
		writeError(w, http.StatusBadRequest, "last_name is required")
		return
	}

	if input.MiddleName == "" {
		writeError(w, http.StatusBadRequest, "middle_name is required")
		return
	}

	if input.Phone == "" {
		writeError(w, http.StatusBadRequest, "phone is required")
		return
	}

	if input.Email == "" {
		writeError(w, http.StatusBadRequest, "email is required")
		return
	}

	// 7. CREATE ROW
	client, err := h.repository.CreateClient(input)

	// 8. CHECK FOR ERROR FROM DB
	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to create client")
		return
	}

	// 9. RETURN SUCCESS
	writeJSON(w, http.StatusCreated, client)
}

// 3) GET CLIENT
func (h *ClientHandler) GetClientById(w http.ResponseWriter, r *http.Request) {
	// 1. PARSE ID
	id, err := getIDFromChi(r)

	// 2. CHECK IF ID IS VALID
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid id")
		return
	}

	// 3. USE REPO TO GET CLIENT
	client, err := h.repository.GetClientById(id)

	// 4. CHECK FOR ERRORS
	if err != nil {
		// 5. IF ROWS IS EMPTY THAN THROW STATUS NOT FOUND
		if err == pgx.ErrNoRows {
			writeError(w, http.StatusNotFound, "client not found")
			return
		}
		// 6. IF OTHER ERROR THROW INTERNAL SERVER ERROR
		writeError(w, http.StatusInternalServerError, "failed to get client")
		return
	}

	// 7. SUCCESS
	writeJSON(w, http.StatusOK, client)
}

// 4) DELETE CLIENT
func (h *ClientHandler) DeleteClient(w http.ResponseWriter, r *http.Request) {
	// 1. PARSE ID
	id, err := getIDFromChi(r)

	// 2. CHECK IF ID IS VALID
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid id")
		return
	}

	// 3. USE REPOSITORY TO DELETE CLIENT
	err = h.repository.DeleteClient(id)

	// 3. CHECK FOR ERRORS
	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to delete client")
		return
	}

	// 4. SUCCESS
	writeJSON(w, http.StatusOK, map[string]string{
		"message": "client deleted",
	})
}

// 5) UPDATE CLIENT
func (h *ClientHandler) UpdateClient(w http.ResponseWriter, r *http.Request) {
	// 1. CLOSE BODY CONNECTION
	defer r.Body.Close()

	// 2. PARSE ID
	id, err := getIDFromChi(r)

	// 3. CHECK IF ID IS VALID
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid id")
		return
	}

	// 4. CREATE VARIABLE FOR INPUT
	var input models.UpdateClientInput

	// 5. FILL VARIABLE
	err = json.NewDecoder(r.Body).Decode(&input)

	// 6. CHECK FOR ERROR
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid json")
		return
	}

	// 7. GET CLIENT FROM DB
	client, err := h.repository.GetClientById(id)

	// 8. CHECK FOR ERRORS
	if err != nil {
		// 9. IF ROWS IS EMPTY THAN THROW STATUS NOT FOUND
		if err == pgx.ErrNoRows {
			writeError(w, http.StatusNotFound, "client not found")
			return
		}
		// 10. IF OTHER ERROR THROW INTERNAL SERVER ERROR
		writeError(w, http.StatusInternalServerError, "failed to get client")
		return
	}

	// 11. REPLACE CLIENT FIELDS VIA INPUT FIELDS
	if input.FirstName != nil {
		client.FirstName = *input.FirstName
	}
	if input.FirstName != nil {
		client.LastName = *input.LastName
	}
	if input.FirstName != nil {
		client.MiddleName = *input.MiddleName
	}
	if input.FirstName != nil {
		client.Phone = *input.Phone
	}
	if input.FirstName != nil {
		client.Email = *input.Email
	}

	// 12. USE REPO TO UPDATE CLIENT
	err = h.repository.UpdateClient(client)

	// 13. CHECK FOR ERRORS FROM DB
	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to update client")
		return
	}

	// 14. SUCCESS
	writeJSON(w, http.StatusOK, client)
}
