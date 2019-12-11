package delivery

import (
	"awesomeProject/domain"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

type ItemHandler struct {
	iUsecase domain.ItemUsecase
}

func NewItemHandler(r *httprouter.Router, iU domain.ItemUsecase) {
	handler := &ItemHandler{iU}

	r.GET("/list/:listID/items", handler.GetAll)
	r.GET("/list/:listID/items/:itemID", handler.GetByID)
	r.POST("/list/:listID/items", handler.Create)
	r.PUT("/list/:listID/items", handler.Update)
	r.DELETE("/list/:listID/items/:itemID", handler.Delete)
}

func (h *ItemHandler) GetAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	i64, err := strconv.ParseInt(p.ByName("listID"), 10, 32)

	if err != nil {
		http.Error(w, "Invalid list id", 500)
	}

	id := uint(i64)

	items, err := h.iUsecase.GetAll(id)

	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	_ = json.NewEncoder(w).Encode(&items)
}

func (h *ItemHandler) GetByID(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	l64, listIDErr := strconv.ParseInt(p.ByName("listID"), 10, 32)
	i64, itemIDErr := strconv.ParseInt(p.ByName("listID"), 10, 32)

	if listIDErr != nil || itemIDErr != nil {
		http.Error(w, "Invalid ID", 500)
	}

	listID := uint(l64)
	itemID := uint(i64)

	item, err := h.iUsecase.GetByID(listID, itemID)

	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	_ = json.NewEncoder(w).Encode(&item)
}

func (h *ItemHandler) Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var item domain.Item

	l64, err := strconv.ParseInt(p.ByName("listID"), 10, 32)
	id := uint(l64)

	err = json.NewDecoder(r.Body).Decode(&item)

	if err != nil {
		http.Error(w, "Invalid body", 500)
	}

	err = h.iUsecase.Create(id, item)

	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	_ = json.NewEncoder(w).Encode("Item was successfully created")
}

func (h *ItemHandler) Update(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var item domain.Item

	err := json.NewDecoder(r.Body).Decode(&item)

	if err != nil {
		http.Error(w, "Invalid body", 500)
	}

	err = h.iUsecase.Update(item)

	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	_ = json.NewEncoder(w).Encode("Item was successfully updated")
}

func (h *ItemHandler) Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	l64, listIDErr := strconv.ParseInt(p.ByName("listID"), 10, 32)
	i64, itemIDErr := strconv.ParseInt(p.ByName("listID"), 10, 32)

	if listIDErr != nil || itemIDErr != nil {
		http.Error(w, "Invalid ID", 500)
	}

	listID := uint(l64)
	itemID := uint(i64)

	err := h.iUsecase.Delete(listID, itemID)

	if err != nil {
		http.Error(w, "Invalid item ID", 500)
	}

	_ = json.NewEncoder(w).Encode("Item was successfully deleted")
}
