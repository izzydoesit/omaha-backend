package handlers

import (
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/izzydoesit/omaha-backend/internal/models"
	"github.com/izzydoesit/omaha-backend/internal/services"
	"gorm.io/gorm"
)

// HandsHandler handles HTTP requests related to hands
type HandsHandler struct {
	DB *gorm.DB
}

// CreateHand handles POST /api/hands - Save a new hand
// @Summary Save a new hand
// @Description Save a new hand
// @Tags hands
// @Accept json
// @Produce json
// @Param hand body models.Hand true "Hand object"
// @Success 200 {object} models.Hand
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/hands [post]
func (h *HandsHandler) CreateHand(c *fiber.Ctx) error {
	var req struct {
		UserID string `json:"user_id"`
		Cards []string `json:"cards"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}
	if len(req.Cards) != 4 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Omaha hand must have exactly 4 cards"})
	}

	hand := models.Hand{
		UserID: req.UserID,
		Cards: strings.Join(req.Cards, ","),  // e.g. "As,Kd,7c,3h"
		CreatedAt: time.Now(),
	}

	if err := services.SaveHand(h.DB, &hand); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to save hand"})
	}
	return c.Status(fiber.StatusCreated).JSON(hand)
}

// ListHands handles GET /api/hands?user_id=... List hands for a user
// @Summary List hands for a user
// @Description List hands for a user
// @Tags hands
// @Accept json
// @Produce json
// @Param user_id query string false "User ID"
// @Success 200 {array} models.Hand
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/hands [get]
func (h *HandsHandler) ListHands(c *fiber.Ctx) error {
	userID := c.Query("user_id")
	if userID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "user_id is required"})
	}
	hands, err := services.GetHandsByUser(h.DB, userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to get hands"})
	}
	return c.JSON(hands)
}
