package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-park-mail-ru/2023_2_Umlaut/model"
	"github.com/go-park-mail-ru/2023_2_Umlaut/static"
)

// @Summary get all complaint types
// @Tags complaint
// @ID complaintTypes
// @Accept  json
// @Produce  json
// @Success 200 {object} ClientResponseDto[[]model.ComplaintType]
// @Failure 401,500 {object} ClientResponseDto[string]
// @Router /api/v1/complaint_types [get]
func (h *Handler) getAllComplaintTypes(w http.ResponseWriter, r *http.Request) {

	complaintTypes, err := h.services.Complaint.GetComplaintTypes(r.Context())
	if err != nil {
		newErrorClientResponseDto(r.Context(), w, http.StatusInternalServerError, err.Error())
		return
	}

	NewSuccessClientResponseArrayDto(r.Context(), w, complaintTypes)
}

// @Summary create complaint
// @Tags complaint
// @ID complaint
// @Accept  json
// @Produce  json
// @Param input body model.Complaint true "Complaint data to create"
// @Success 200 {object} ClientResponseDto[string]
// @Failure 400, 401, 409, 500 {object} ClientResponseDto[string]
// @Router /api/v1/complaint [post]
func (h *Handler) createComplaint(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var complaint model.Complaint
	if err := decoder.Decode(&complaint); err != nil {
		newErrorClientResponseDto(r.Context(), w, http.StatusBadRequest, "invalid input body")
		return
	}

	complaint.ReporterUserId = r.Context().Value(keyUserID).(int)

	_, err := h.services.Complaint.CreateComplaint(r.Context(), complaint)
	if err != nil {
		if errors.Is(err, static.ErrAlreadyExists) {
			newErrorClientResponseDto(r.Context(), w, http.StatusConflict, "complaint already exists")
			return
		}
		newErrorClientResponseDto(r.Context(), w, http.StatusInternalServerError, err.Error())
		return
	}

	NewSuccessClientResponseDto(r.Context(), w, "")
}
