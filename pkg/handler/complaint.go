package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-park-mail-ru/2023_2_Umlaut/model"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/microservices/admin/proto"
	"github.com/go-park-mail-ru/2023_2_Umlaut/static"
	"github.com/go-park-mail-ru/2023_2_Umlaut/utils"
	"github.com/golang/protobuf/ptypes"
	"github.com/gorilla/mux"
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
// @Failure 400,401,409,500 {object} ClientResponseDto[string]
// @Router /api/v1/complaint [post]
func (h *Handler) createComplaint(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var complaint model.Complaint
	if err := decoder.Decode(&complaint); err != nil {
		newErrorClientResponseDto(r.Context(), w, http.StatusBadRequest, "invalid input body")
		return
	}

	complaint.ReporterUserId = r.Context().Value(static.KeyUserID).(int)

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

// @Summary get next complaint
// @Tags complaint
// @Accept  json
// @Produce  json
// @Success 200 {object} ClientResponseDto[model.Complaint]
// @Failure 401,500 {object} ClientResponseDto[string]
// @Router /api/v1/admin/complaint [get]
func (h *Handler) getNextComplaint(w http.ResponseWriter, r *http.Request) {
	complaint, err := h.adminMicroservice.GetNextComplaint(r.Context(), &proto.AdminEmpty{})
	if err != nil {
		statusCode, message := utils.ParseError(err)
		newErrorClientResponseDto(r.Context(), w, statusCode, message)
		return
	}
	createdAt, err := ptypes.Timestamp(complaint.CreatedAt)
	if err != nil {
		newErrorClientResponseDto(r.Context(), w, http.StatusInternalServerError, err.Error())
		return
	}

	NewSuccessClientResponseDto(r.Context(), w, model.Complaint{
		Id: int(complaint.Id),
		ReporterUserId: int(complaint.ReporterUserId),
		ReportedUserId: int(complaint.ReportedUserId),
		CreatedAt: &createdAt,
	})
}

// @Summary delete complaint
// @Tags complaint
// @Accept  json
// @Param id path integer true "complaint ID"
// @Produce  json
// @Success 200 {object} ClientResponseDto[string]
// @Failure 401,500 {object} ClientResponseDto[string]
// @Router /api/v1/admin/complaint/{id} [delete]
func (h *Handler) deleteComplaint(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		newErrorClientResponseDto(r.Context(), w, http.StatusBadRequest, "invalid params")
		return
	}
	_, err = h.adminMicroservice.DeleteComplaint(r.Context(), &proto.Complaint{Id: int32(id)})
	if err != nil {
		statusCode, message := utils.ParseError(err)
		newErrorClientResponseDto(r.Context(), w, statusCode, message)
	}

	NewSuccessClientResponseDto(r.Context(), w, "")
}

// @Summary accept complaint
// @Tags complaint
// @Accept  json
// @Param id path integer true "complaint ID"
// @Produce  json
// @Success 200 {object} ClientResponseDto[string]
// @Failure 401,500 {object} ClientResponseDto[string]
// @Router /api/v1/admin/complaint/{id} [get]
func (h *Handler) acceptComplaint(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		newErrorClientResponseDto(r.Context(), w, http.StatusBadRequest, "invalid params")
		return
	}
	_, err = h.adminMicroservice.AcceptComplaint(r.Context(), &proto.Complaint{Id: int32(id)})
	if err != nil {
		statusCode, message := utils.ParseError(err)
		newErrorClientResponseDto(r.Context(), w, statusCode, message)
	}

	NewSuccessClientResponseDto(r.Context(), w, "")
}