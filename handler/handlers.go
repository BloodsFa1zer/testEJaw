package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"test/database"
	"test/response"
	"test/service"
)

type SellerHandler struct {
	sellerService service.SellerServiceInterface
}

func NewSellerHandler(service service.SellerServiceInterface) *SellerHandler {
	return &SellerHandler{sellerService: service}
}

type SellerHandlerInterface interface {
	GetAll(w http.ResponseWriter, r *http.Request)
	GetByID(w http.ResponseWriter, r *http.Request)
	Post(w http.ResponseWriter, r *http.Request)
	Put(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

func (sh *SellerHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	sellers, err := sh.sellerService.GetAllSellers()
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, "Failed to get sellers:"+err.Error(), nil)
		return
	}

	response.JSON(w, http.StatusOK, "data:", sellers)
}

func (sh *SellerHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		response.JSON(w, http.StatusBadRequest, "Invalid ID: "+err.Error(), nil)
		return
	}

	seller, err := sh.sellerService.GetByIDSeller(id)
	if err != nil {
		response.JSON(w, http.StatusBadRequest, "Invalid ID: "+err.Error(), nil)
		return
	}

	response.JSON(w, http.StatusOK, "data:", seller)
}

func (sh *SellerHandler) Post(w http.ResponseWriter, r *http.Request) {
	var Seller database.Seller
	if err := json.NewDecoder(r.Body).Decode(&Seller); err != nil {
		response.JSON(w, http.StatusBadRequest, "Invalid input: "+err.Error(), nil)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := sh.sellerService.CreateSeller(Seller)
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, "Cannot create seller: "+err.Error(), nil)
		return
	}

	response.JSON(w, http.StatusCreated, "Seller created successfully", map[string]int{"id": id})
}

func (sh *SellerHandler) Put(w http.ResponseWriter, r *http.Request) {
	var seller database.Seller
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		response.JSON(w, http.StatusBadRequest, "Invalid ID: "+err.Error(), nil)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&seller); err != nil {
		response.JSON(w, http.StatusBadRequest, "Invalid input: "+err.Error(), nil)
		return
	}

	seller.ID = id
	if err := sh.sellerService.UpdateSeller(seller); err != nil {
		response.JSON(w, http.StatusInternalServerError, "Cannot update seller: "+err.Error(), nil)
		return
	}

	response.JSON(w, http.StatusOK, "Seller updated successfully", nil)
}

func (sh *SellerHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		response.JSON(w, http.StatusBadRequest, "Invalid ID: "+err.Error(), nil)
		return
	}

	if err := sh.sellerService.DeleteSeller(id); err != nil {
		response.JSON(w, http.StatusInternalServerError, "Cannot delete seller: "+err.Error(), nil)
		return
	}

	response.JSON(w, http.StatusOK, "Seller deleted successfully", nil)
}
