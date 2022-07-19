package car

import (
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
)

const (
	CarsURL = "/cars"
	CarURL  = "/cars/{id}"
)

type Handler struct {
	router *mux.Router
	logger *zap.SugaredLogger
}

func NewHandlerCar(s *mux.Router, l *zap.SugaredLogger) *Handler {
	h := Handler{s, l}
	h.registerRoutes()

	return &h
}

func (h *Handler) registerRoutes() {
	h.router.HandleFunc(CarsURL, h.GetListCar).Methods("Get")
	h.router.HandleFunc(CarURL, h.GetCarByID).Methods("Get")
	h.router.HandleFunc(CarsURL, h.CreateCar).Methods("Post")
	h.router.HandleFunc(CarURL, h.UpdateCar).Methods("Put")
	h.router.HandleFunc(CarURL, h.PartiallyUpdateCar).Methods("Patch")
	h.router.HandleFunc(CarURL, h.DeleteCar).Methods("Delete")
}

func (h *Handler) GetListCar(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("Start [GetListUser] Handler")
	w.WriteHeader(200)
	w.Write([]byte("Hello World"))
}

func (h *Handler) GetCarByID(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("Start [GetUserByID] Handler")
	w.WriteHeader(200)
	w.Write([]byte("Hello world 2"))
}

func (h *Handler) CreateCar(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("Start [CreateUser] Handler")
	w.Write([]byte("Test"))
}

func (h *Handler) UpdateCar(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("Start [UpdateUser] Handler")
	w.Write([]byte("Test"))
}

func (h *Handler) PartiallyUpdateCar(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("Start [PartiallyUpdateUser] Handler")
	w.Write([]byte("Test"))
}

func (h *Handler) DeleteCar(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("Start [DeleteUser] Handler")
	w.Write([]byte("Test"))
}
