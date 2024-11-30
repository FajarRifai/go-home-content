package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go-home-content/bean"
	"go-home-content/models"
	"go-home-content/service"
	"log"
	"net/http"
	"strconv"
)

type SectionController struct {
	Service *services.SectionService
}

func (h *SectionController) CreateSection(w http.ResponseWriter, r *http.Request) {
	var request models.CreateSection

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		bean.ErrorResponse(w, http.StatusBadRequest, "Invalid JSON payload")
		return
	}

	_, err := h.Service.CreateSection(request)
	if err != nil {
		bean.ErrorResponse(w, http.StatusInternalServerError, "Failed to create section")
		return
	}

	bean.JsonResponse(w, http.StatusOK, "00", "Success", nil)
}

func (h *SectionController) GetSections(w http.ResponseWriter, r *http.Request) {
	sections, err := h.Service.GetSections()
	if err != nil {
		bean.ErrorResponse(w, http.StatusInternalServerError, "Failed to fetch sections")
		return
	}

	bean.JsonResponse(w, http.StatusOK, "00", "Success", sections)
}

func (h *SectionController) GetSectionById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	// Validate the ID
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		bean.ErrorResponse(w, http.StatusBadRequest, "Invalid section ID")
		return
	}

	section, err := h.Service.GetSectionById(id)
	if err != nil {
		bean.ErrorResponse(w, http.StatusNotFound, fmt.Sprintf("CreateSection not found: %v", err))
		return
	}

	// Fetch the product codes from section details
	detail, err := h.Service.GetSectionDetailById(id)
	if err != nil {
		bean.ErrorResponse(w, http.StatusInternalServerError, fmt.Sprintf("Error fetching section details: %v", err))
		return
	}

	// Fetch the products
	baseURL := "http://localhost:8081/api/product-codes"
	products, err := services.FetchProducts(baseURL, detail)
	if err != nil {
		log.Fatalf("Error fetching products: %v", err)
	}
	log.Println("Fetched Products:", products)

	var sectionProducts []models.Product // Slice to store products for the section
	for _, product := range products {
		sectionProduct := models.Product{
			Name:        product.Name,
			Code:        product.Code,
			Description: product.Description,
			Qty:         product.Qty,
			Active:      product.Active,
			Deleted:     product.Deleted,
		}
		sectionProducts = append(sectionProducts, sectionProduct)
	}

	// Map fetched products to the section
	section.Products = sectionProducts

	// Return a successful response
	bean.JsonResponse(w, http.StatusOK, "00", "Success", section)
}
