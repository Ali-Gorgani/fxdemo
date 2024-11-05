package handler

import (
	"strconv"

	"example.com/fxdemo/api-gateway/domain"
	"example.com/fxdemo/api-gateway/usecase"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	usecase usecase.IUsecase
}

func NewHandler(usecase usecase.IUsecase) *Handler {
	return &Handler{
		usecase: usecase,
	}
}

func (h *Handler) GetUserByID(ctx *fiber.Ctx) error {
	strID := ctx.Params("id")
	id, err := strconv.Atoi(strID)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid user ID",
		})
	}

	user := &domain.User{
		ID: id,
	}
	foundUser, err := h.usecase.GetUserByID(ctx.Context(), user)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(foundUser)
}

func (h *Handler) CreateUser(ctx *fiber.Ctx) error {
	user := new(domain.User)
	if err := ctx.BodyParser(user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	createdUser, err := h.usecase.CreateUser(ctx.Context(), user)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(createdUser)
}
