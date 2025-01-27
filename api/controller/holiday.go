package controller

import (
	"go-fiber/api/middleware"
	"go-fiber/data/services"
	"go-fiber/domain/models"

	"github.com/gofiber/fiber/v2"
)

type holidayController struct {
	service services.HolidayService
}

type HolidayController interface {
	FindAll(ctx *fiber.Ctx) error
	Create(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
	IsHoliday(ctx *fiber.Ctx) error
}

func NewHolidayController(service services.HolidayService) HolidayController {
	return &holidayController{service}
}

func (c holidayController) IsHoliday(ctx *fiber.Ctx) error {
	date := ctx.Params("date")
	isHoliday, err := c.service.IsHoliday(date)
	if err != nil {
		return middleware.NewErrorMessageResponse(ctx, err)
	}
	return middleware.NewSuccessResponse(ctx, isHoliday)
}

func (c holidayController) FindAll(ctx *fiber.Ctx) error {
	limit := ctx.QueryInt("limit", 10)
	page := ctx.QueryInt("page", 1)

	result, total, err := c.service.FindAll(page, limit)
	if err != nil {
		return middleware.NewErrorMessageResponse(ctx, err)
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

func (c holidayController) Create(ctx *fiber.Ctx) error {
	var holiday models.Holiday
	if err := ctx.BodyParser(&holiday); err != nil {
		return middleware.NewErrorMessageResponse(ctx, err)
	}
	userID, err := middleware.GetOwnerAccessToken(ctx)
	if err != nil {
		return middleware.NewErrorMessageResponse(ctx, err)
	}
	holiday.CreatedBy = *userID
	date, err := c.service.Create(holiday)
	if err != nil {
		return middleware.NewErrorMessageResponse(ctx, err)
	}
	return middleware.NewSuccessResponse(ctx, date)
}

func (c holidayController) Update(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return middleware.NewErrorMessageResponse(ctx, err)
	}
	var holiday models.Holiday
	if err := ctx.BodyParser(&holiday); err != nil {
		return middleware.NewErrorMessageResponse(ctx, err)
	}
	data, err := c.service.Update(uint(id), holiday)
	if err != nil {
		return middleware.NewErrorMessageResponse(ctx, err)
	}
	return middleware.NewSuccessResponse(ctx, data)
}

func (c holidayController) Delete(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return middleware.NewErrorMessageResponse(ctx, err)
	}
	if err := c.service.Delete(uint(id)); err != nil {
		return middleware.NewErrorMessageResponse(ctx, err)
	}
	return middleware.NewSuccessResponse(ctx, nil)
}
