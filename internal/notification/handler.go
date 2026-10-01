package notification

import (
	"encoding/json"
	"net/http"
)

type Handler struct {
	repository *Repository
}

func NewHandler(repository *Repository) *Handler {
	return &Handler{
		repository: repository,
	}
}

// @Summary Create a new notification
// @Description Create a new notification
// @Accept json
// @Produce json
// @Param notification body CreateNotificationRequest true "Notification to create"
// @Success 201 {object} Notification
// @Router /notifications [post]
// @Tags notifications
func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var input CreateNotificationRequest

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	notification := Notification{
		UserID:  input.UserID,
		Message: input.Message,
		Status:  "pending",
	}

	if err := h.repository.Create(&notification); err != nil {
		http.Error(
			w,
			"failed to create notification",
			http.StatusInternalServerError,
		)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(notification)
}

// @Summary List notifications
// @Description List notifications
// @Accept json
// @Produce json
// @Success 200 {object} ListNotificationsDto
// @Router /notifications [get]
// @Tags notifications
func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	notifications, err := h.repository.List()

	if err != nil {
		http.Error(
			w,
			"failed to list notifications",
			http.StatusInternalServerError,
		)
		return
	}

	total := int64(len(notifications))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(ListNotificationsDto{
		Notifications: notifications,
		Total:         total,
	})

}
