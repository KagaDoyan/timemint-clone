package controller

import (
	"go-fiber/api/middleware"
	"go-fiber/data/services"
	"go-fiber/domain/models"

	"github.com/gofiber/fiber/v2"
)

type ActiveLocationController interface {
	FindAll(ctx *fiber.Ctx) error
	Find(ctx *fiber.Ctx) error
	Create(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}

type activeLocationController struct {
	service services.ActiveLocationService
}

func NewActiveLocationController(service services.ActiveLocationService) ActiveLocationController {
	return &activeLocationController{
		service: service,
	}
}

func (c activeLocationController) FindAll(ctx *fiber.Ctx) error {
	limit := ctx.QueryInt("limit", 10)
	page := ctx.QueryInt("page", 1)

	result, total, err := c.service.FindAll(page, limit)
	if err != nil {
		return middleware.NewErrorResponses(ctx, err)
	}

	totalPages := (total + int64(limit) - 1) / int64(limit)
	return middleware.NewSuccessResponse(ctx, map[string]interface{}{
		"data":       result,
		"total":      total,
		"page":       page,
		"limit":      limit,
		"totalPages": totalPages,
	})
}

func (c activeLocationController) Find(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return middleware.NewErrorResponses(ctx, err)
	}
	result, err := c.service.FindByID(uint(id))
	if err != nil {
		return middleware.NewErrorResponses(ctx, err)
	}
	return middleware.NewSuccessResponse(ctx, result)
}

func (c activeLocationController) Create(ctx *fiber.Ctx) error {
	activeLocation := models.ActiveLocation{}
	if err := ctx.BodyParser(&activeLocation); err != nil {
		return middleware.NewErrorResponses(ctx, err)
	}
	data, err := c.service.Create(activeLocation)
	if err != nil {
		return middleware.NewErrorResponses(ctx, err)
	}
	return middleware.NewSuccessResponse(ctx, data)
}

func (c activeLocationController) Update(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return middleware.NewErrorResponses(ctx, err)
	}
	activeLocation := models.ActiveLocation{}
	if err := ctx.BodyParser(&activeLocation); err != nil {
		return middleware.NewErrorResponses(ctx, err)
	}
	activeLocation.ID = uint(id)
	data, err := c.service.Update(uint(id), activeLocation)
	if err != nil {
		return middleware.NewErrorResponses(ctx, err)
	}
	return middleware.NewSuccessResponse(ctx, data)
}

func (c activeLocationController) Delete(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return middleware.NewErrorResponses(ctx, err)
	}
	if err := c.service.Delete(uint(id)); err != nil {
		return middleware.NewErrorResponses(ctx, err)
	}
	return middleware.NewSuccessResponse(ctx, nil)
}
