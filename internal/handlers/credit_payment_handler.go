package handlers

import (
	"banksystem/internal/services"
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

type CreditPaymentHandler struct {
	service *services.CreditPaymentService
}

func NewCreditPaymentHandler(service *services.CreditPaymentService) *CreditPaymentHandler {
	return &CreditPaymentHandler{service: service}
}

func (h *CreditPaymentHandler) CreatePayment(w http.ResponseWriter, r *http.Request) {
	var request struct {
		CreditID    int     `json:"credit_id"`
		Amount      float64 `json:"amount"`
		PaymentDate string  `json:"payment_date"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	paymentDate, err := time.Parse(time.RFC3339, request.PaymentDate)
	if err != nil {
		http.Error(w, "Invalid payment date format", http.StatusBadRequest)
		return
	}

	if err := h.service.CreatePayment(request.CreditID, request.Amount, paymentDate); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *CreditPaymentHandler) ProcessPayment(w http.ResponseWriter, r *http.Request) {
	paymentIDStr := r.URL.Query().Get("payment_id")
	paymentID, err := strconv.Atoi(paymentIDStr)
	if err != nil {
		http.Error(w, "Invalid payment ID", http.StatusBadRequest)
		return
	}

	if err := h.service.ProcessPayment(paymentID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *CreditPaymentHandler) GetPaymentsByCreditID(w http.ResponseWriter, r *http.Request) {
	creditIDStr := r.URL.Query().Get("credit_id")
	creditID, err := strconv.Atoi(creditIDStr)
	if err != nil {
		http.Error(w, "Invalid credit ID", http.StatusBadRequest)
		return
	}

	payments, err := h.service.GetPaymentsByCreditID(creditID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(payments)
}

func (h *CreditPaymentHandler) GetPendingPayments(w http.ResponseWriter, r *http.Request) {
	payments, err := h.service.GetPendingPayments()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(payments)
} 