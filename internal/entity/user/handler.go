package user

import (
	"carsharing/internal/entity/user/model"
	"encoding/json"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
)

const (
	UsersURL = "/users"
	UserURL  = "/user/{id}"
)

type Handler struct {
	router *mux.Router
	logger *zap.SugaredLogger
}

func NewHandlerUser(s *mux.Router, l *zap.SugaredLogger) *Handler {
	h := Handler{s, l}
	h.registerRoutes()

	return &h
}

func (h *Handler) registerRoutes() {
	h.router.HandleFunc(UsersURL, h.GetListUser).Methods("Get")
	h.router.HandleFunc(UserURL, h.GetUserByID).Methods("Get")
	h.router.HandleFunc(UsersURL, h.CreateUser).Methods("Post")
	h.router.HandleFunc(UserURL, h.UpdateUser).Methods("Put")
	h.router.HandleFunc(UserURL, h.PartiallyUpdateUser).Methods("Patch")
	h.router.HandleFunc(UserURL, h.DeleteUser).Methods("Delete")
}

func (h *Handler) GetListUser(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("Start [GetListUser] Handler")
	w.WriteHeader(200)
	w.Write([]byte("Hello World"))
}

func (h *Handler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("Start [GetUserByID] Handler")
	w.WriteHeader(200)
	w.Write([]byte("Hello world 2"))
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("Start [CreateUser] Handler")
	defer r.Body.Close()

	h.logger.Info("Read to request body")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		h.logger.Fatal(err)
	}

	var user model.User
	json.Unmarshal(body, &user)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Created")
}

func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("Start [UpdateUser] Handler")
	w.Write([]byte("Test"))
}

func (h *Handler) PartiallyUpdateUser(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("Start [PartiallyUpdateUser] Handler")
	w.Write([]byte("Test"))
}

func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("Start [DeleteUser] Handler")
	w.Write([]byte("Test"))
}
