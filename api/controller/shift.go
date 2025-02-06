package controller

import (
	"errors"
	"go-fiber/api/middleware"
	"go-fiber/data/services"
	"go-fiber/domain/models"

	"github.com/gofiber/fiber/v2"
)

type shiftController struct {
	service services.ShiftService
}

type ShiftController interface {
	FindAll(ctx *fiber.Ctx) error
	FindByID(ctx *fiber.Ctx) error
	Create(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
	Option(ctx *fiber.Ctx) error
	ShiftReport(ctx *fiber.Ctx) error
}

func NewShiftController(service services.ShiftService) ShiftController {
	return &shiftController{
		service: service,
	}
}

func (c shiftController) Option(ctx *fiber.Ctx) error {
	datas, err := c.service.Option()
	if err != nil {
		return middleware.NewErrorResponses(ctx, err)
	}
	return middleware.NewSuccessResponse(ctx, datas)
}

func (c shiftController) FindAll(ctx *fiber.Ctx) error {
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

func (c shiftController) FindByID(ctx *fiber.Ctx) error {
	shiftID, err := ctx.ParamsInt("id")
	if err != nil {
		return middleware.NewErrorResponses(ctx, errors.New("shift ID is required"))
	}

	result, err := c.service.FindById(uint(shiftID))
	if err != nil {
		return middleware.NewErrorResponses(ctx, err)
	}
	return middleware.NewSuccessResponse(ctx, result)
}

func (c shiftController) Create(ctx *fiber.Ctx) error {
	var shift models.Shift
	if err := ctx.BodyParser(&shift); err != nil {
		return middleware.NewErrorResponses(ctx, err)
	}

	data, err := c.service.Create(shift)
	if err != nil {
		return middleware.NewErrorResponses(ctx, err)
	}
	return middleware.NewSuccessResponse(ctx, data)
}

func (c shiftController) Update(ctx *fiber.Ctx) error {
	shiftID, err := ctx.ParamsInt("id")
	if err != nil {
		return middleware.NewErrorResponses(ctx, errors.New("shift ID is required"))
	}
	var shift models.Shift
	if err := ctx.BodyParser(&shift); err != nil {
		return middleware.NewErrorResponses(ctx, err)
	}

	data, err := c.service.Update(uint(shiftID), shift)
	if err != nil {
		return middleware.NewErrorResponses(ctx, err)
	}
	return middleware.NewSuccessResponse(ctx, data)
}

func (c shiftController) Delete(ctx *fiber.Ctx) error {
	shiftID, err := ctx.ParamsInt("id")
	if err != nil {
		return middleware.NewErrorResponses(ctx, errors.New("shift ID is required"))
	}
	if err := c.service.Delete(uint(shiftID)); err != nil {
		return middleware.NewErrorResponses(ctx, err)
	}
	return middleware.NewSuccessResponse(ctx, nil)
}

func (c shiftController) ShiftReport(ctx *fiber.Ctx) error {
	result, err := c.service.ShiftReport()
	if err != nil {
		return middleware.NewErrorResponses(ctx, err)
	}
	return middleware.NewSuccessResponse(ctx, result)
}
