package delivery

import (
	"awesomeProject/domain"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

type ListHandler struct {
	lUsecase domain.ListUsecase
}

type Message struct {
	message string
	payload []domain.List
	err     error
}

func NewListHandler(r *httprouter.Router, lu domain.ListUsecase) {
	handler := &ListHandler{lu}

	r.GET("/list", handler.Fetch)
	r.GET("/list/:listID", handler.GetByID)
	r.POST("/list", handler.Create)
	r.DELETE("/list/:listID", handler.Delete)
}

func (h *ListHandler) Fetch(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	l, err := h.lUsecase.Fetch()

	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	_ = json.NewEncoder(w).Encode(&l)
}

func (h *ListHandler) GetByID(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	i64, err := strconv.ParseInt(p.ByName("listID"), 10, 32)

	if err != nil {
		http.Error(w, "Invalid list id", 500)
	}

	listID := uint(i64)

	l, err := h.lUsecase.GetByID(listID)

	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	_ = json.NewEncoder(w).Encode(&l)
}

func (h *ListHandler) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var list domain.List

	err := json.NewDecoder(r.Body).Decode(&list)

	if err != nil {
		fmt.Println("Error => ", err)
	}

	err = h.lUsecase.Create(list)

	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	_ = json.NewEncoder(w).Encode(&list)
}

func (h *ListHandler) Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	i64, err := strconv.ParseInt(p.ByName("listID"), 10, 32)

	if err != nil {
		http.Error(w, "Invalid list id", 500)
	}

	listID := uint(i64)

	list := domain.List{}

	err = h.lUsecase.Delete(listID)

	if err != nil {
		http.Error(w, "Invalid list ID", 500)
	}

	_ = json.NewEncoder(w).Encode(&list)
}
