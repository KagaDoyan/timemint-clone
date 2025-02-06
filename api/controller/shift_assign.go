package controller

import (
	"go-fiber/api/middleware"
	"go-fiber/data/services"
	"go-fiber/domain/models"

	"github.com/gofiber/fiber/v2"
)

type shiftAssignController struct {
	service services.ShiftAssignService
}

type ShiftAssignController interface {
	FindById(ctx *fiber.Ctx) error
	FindAll(ctx *fiber.Ctx) error
	Create(ctx *fiber.Ctx) error
	CreateBatch(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
	CalendarShift(ctx *fiber.Ctx) error
	ShiftAssignmentReport(ctx *fiber.Ctx) error
}

func NewShiftAssignController(service services.ShiftAssignService) ShiftAssignController {
	return &shiftAssignController{
		service: service,
	}
}

func (c shiftAssignController) CreateBatch(ctx *fiber.Ctx) error {
	var shiftAssigns []models.ShiftAssignment

	if err := ctx.BodyParser(&shiftAssigns); err != nil {
		return middleware.NewErrorResponses(ctx, err)
	}
	userID, err := middleware.GetOwnerAccessToken(ctx)
	if err != nil {
		return middleware.NewErrorResponses(ctx, err)
	}
	created_by := *userID

	totalrecords, inserted, failed, failures := c.service.CreateBatch(created_by, shiftAssigns)
	return middleware.NewSuccessResponse(ctx, map[string]interface{}{
		"data": map[string]interface{}{
			"totalrecords": totalrecords,
			"inserted":     inserted,
			"failed":       failed,
			"failures":     failures,
		},
	})
}

func (c shiftAssignController) FindAll(ctx *fiber.Ctx) error {
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

func (c shiftAssignController) FindById(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return middleware.NewErrorResponses(ctx, err)
	}

	result, err := c.service.FindById(uint(id))
	if err != nil {
		return middleware.NewErrorResponses(ctx, err)
	}
	return middleware.NewSuccessResponse(ctx, result)
}

func (c shiftAssignController) Create(ctx *fiber.Ctx) error {
	var shiftAssign models.ShiftAssignment
	if err := ctx.BodyParser(&shiftAssign); err != nil {
		return middleware.NewErrorResponses(ctx, err)
	}
	userID, err := middleware.GetOwnerAccessToken(ctx)
	if err != nil {
		return middleware.NewErrorResponses(ctx, err)
	}
	shiftAssign.CreatedBy = *userID
	data, err := c.service.Create(shiftAssign)
	if err != nil {
		return middleware.NewErrorResponses(ctx, err)
	}
	return middleware.NewSuccessResponse(ctx, data)
}

func (c shiftAssignController) Delete(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return middleware.NewErrorResponses(ctx, err)
	}
	if err := c.service.Delete(uint(id)); err != nil {
		return middleware.NewErrorResponses(ctx, err)
	}
	return middleware.NewSuccessResponse(ctx, nil)
}

func (c shiftAssignController) CalendarShift(ctx *fiber.Ctx) error {
	month, err := ctx.ParamsInt("month")
	if err != nil {
		return middleware.NewErrorResponses(ctx, err)
	}
	year, err := ctx.ParamsInt("year")
	if err != nil {
		return middleware.NewErrorResponses(ctx, err)
	}
	result, err := c.service.CalendarShift(month, year)
	if err != nil {
		return middleware.NewErrorResponses(ctx, err)
	}
	return middleware.NewSuccessResponse(ctx, result)
}

func (c shiftAssignController) ShiftAssignmentReport(ctx *fiber.Ctx) error {
	start := ctx.Query("start")
	end := ctx.Query("end")
	result, err := c.service.ShiftAssignmentReport(start, end)
	if err != nil {
		return middleware.NewErrorResponses(ctx, err)
	}
	return middleware.NewSuccessResponse(ctx, result)
}
