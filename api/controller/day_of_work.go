package controller

import (
	"go-fiber/api/middleware"
	"go-fiber/data/services"
	"go-fiber/domain/models"

	"github.com/gofiber/fiber/v2"
)

type dayOfWorkController struct {
	service services.DayOfWorkService
}

type DayOfWorkController interface {
	FindAll(ctx *fiber.Ctx) error
	FindByID(ctx *fiber.Ctx) error
	Create(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}

func NewDayOfWorkController(service services.DayOfWorkService) DayOfWorkController {
	return &dayOfWorkController{
		service: service,
	}
}

func (c dayOfWorkController) FindAll(ctx *fiber.Ctx) error {
	limit := ctx.QueryInt("limit", 10)
	page := ctx.QueryInt("page", 1)

	result, total, err := c.service.FindAll(page, limit)
	if err != nil {
		return middleware.NewErrorResponses(ctx, err)
	}

	totalPages := (total + int64(limit) - 1) / int64(limit)

	return middleware.NewSuccessResponse(ctx, fiber.Map{
		"page":       page,
		"limit":      limit,
		"totalRows":  total,
		"totalPages": totalPages,
		"data":       result,
	})
}

func (c dayOfWorkController) FindByID(ctx *fiber.Ctx) error {
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

func (c dayOfWorkController) Create(ctx *fiber.Ctx) error {
	var dayOfWork models.DayOfWork
	if err := ctx.BodyParser(&dayOfWork); err != nil {
		return middleware.NewErrorResponses(ctx, err)
	}
	data, err := c.service.Create(dayOfWork)
	if err != nil {
		return middleware.NewErrorResponses(ctx, err)
	}
	return middleware.NewSuccessResponse(ctx, data)
}

func (c dayOfWorkController) Update(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return middleware.NewErrorResponses(ctx, err)
	}

	var dayOfWork models.DayOfWork
	if err := ctx.BodyParser(&dayOfWork); err != nil {
		return middleware.NewErrorResponses(ctx, err)
	}

	// Ensure the ID from the path matches the body
	dayOfWork.ID = uint(id)
	data, err := c.service.Update(uint(id), dayOfWork)
	if err != nil {
		return middleware.NewErrorResponses(ctx, err)
	}
	return middleware.NewSuccessResponse(ctx, data)
}

func (c dayOfWorkController) Delete(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return middleware.NewErrorResponses(ctx, err)
	}
	if err := c.service.Delete(uint(id)); err != nil {
		return middleware.NewErrorResponses(ctx, err)
	}
	return middleware.NewSuccessResponse(ctx, nil)
}
