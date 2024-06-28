package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/pedrohscramos/enube/internal/infra/database"
)

type SupplierHandler struct {
	SupplierDB database.SupplierInterface
}

func NewSupplierHandler(db database.SupplierInterface) *SupplierHandler {
	return &SupplierHandler{
		SupplierDB: db,
	}
}

// ListSuppliers godoc
// @Summary      List suppliers
// @Description  get all suppliers
// @Tags         suppliers
// @Accept       json
// @Produce      json
// @Param        page      query     string  false  "page number"
// @Param        limit     query     string  false  "limit"
// @Success      200       {array}   entity.Supplier
// @Failure      404       {object}  Error
// @Failure      500       {object}  Error
// @Router       /suppliers [get]
// @Security ApiKeyAuth
func (h *SupplierHandler) GetSuppliers(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	limit := r.URL.Query().Get("limit")

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		pageInt = 0
	}
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		limitInt = 0
	}
	sort := r.URL.Query().Get("sort")

	suppliers, err := h.SupplierDB.FindAll(pageInt, limitInt, sort)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(suppliers)
}

// GetSupplier godoc
// @Summary      Get a supplier
// @Description  Get a supplier
// @Tags         suppliers
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "supplier ID" Format(int)
// @Success      200  {object}  entity.Supplier
// @Failure      404
// @Failure      500  {object}  Error
// @Router       /suppliers/{id} [get]
// @Security ApiKeyAuth
func (h *SupplierHandler) GetSupplier(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	idInt, err := strconv.Atoi(id)

	if idInt == 0 || err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	supplier, err := h.SupplierDB.FindByID(idInt)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(supplier)
}
