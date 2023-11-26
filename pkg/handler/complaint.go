package handler

import (
	"net/http"
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
