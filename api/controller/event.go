package controller

import (
	"fmt"
	"go-fiber/api/middleware"
	"go-fiber/data/services"
	"go-fiber/domain/models"

	"github.com/gofiber/fiber/v2"
)

type eventController struct {
	service services.EventService
}

type EventController interface {
	FindAll(ctx *fiber.Ctx) error
	Create(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
	IsEvent(ctx *fiber.Ctx) error
	CalendarEvent(ctx *fiber.Ctx) error
}

func NewEventController(service services.EventService) EventController {
	return &eventController{service}
}

func (c eventController) IsEvent(ctx *fiber.Ctx) error {
	date := ctx.Params("date")
	isEvent, err := c.service.IsEvent(date)
	if err != nil {
		return middleware.NewErrorMessageResponse(ctx, err)
	}
	return middleware.NewSuccessResponse(ctx, isEvent)
}

func (c eventController) FindAll(ctx *fiber.Ctx) error {
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

func (c eventController) Create(ctx *fiber.Ctx) error {
	var event models.Event
	if err := ctx.BodyParser(&event); err != nil {
		return middleware.NewErrorMessageResponse(ctx, err)
	}
	userID, err := middleware.GetOwnerAccessToken(ctx)
	if err != nil {
		return middleware.NewErrorMessageResponse(ctx, err)
	}
	event.CreatedBy = *userID
	inserted, err := c.service.Create(event)
	if err != nil {
		return middleware.NewErrorMessageResponse(ctx, err)
	}
	return middleware.NewSuccessMessageResponse(ctx, fmt.Sprintf("Event created %d", inserted))
}

func (c eventController) Update(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return middleware.NewErrorMessageResponse(ctx, err)
	}
	var event models.Event
	if err := ctx.BodyParser(&event); err != nil {
		return middleware.NewErrorMessageResponse(ctx, err)
	}
	data, err := c.service.Update(uint(id), event)
	if err != nil {
		return middleware.NewErrorMessageResponse(ctx, err)
	}
	return middleware.NewSuccessResponse(ctx, data)
}

func (c eventController) Delete(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return middleware.NewErrorMessageResponse(ctx, err)
	}
	if err := c.service.Delete(uint(id)); err != nil {
		return middleware.NewErrorMessageResponse(ctx, err)
	}
	return middleware.NewSuccessResponse(ctx, nil)
}

func (c eventController) CalendarEvent(ctx *fiber.Ctx) error {
	month, err := ctx.ParamsInt("month")
	if err != nil {
		return middleware.NewErrorMessageResponse(ctx, err)
	}
	year, err := ctx.ParamsInt("year")
	if err != nil {
		return middleware.NewErrorMessageResponse(ctx, err)
	}
	result, err := c.service.CalendarEvent(month, year)
	if err != nil {
		return middleware.NewErrorMessageResponse(ctx, err)
	}
	return middleware.NewSuccessResponse(ctx, result)
}
