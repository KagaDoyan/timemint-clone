package controller

import (
	"go-fiber/api/middleware"
	"go-fiber/data/services"
	"go-fiber/domain/models"

	"github.com/gofiber/fiber/v2"
)

type leaveTypeController struct {
	service services.LeaveTypeService
}

type LeaveTypeController interface {
	FindAll(ctx *fiber.Ctx) error
	FindByID(ctx *fiber.Ctx) error
	Create(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}

func NewLeaveTypeController(service services.LeaveTypeService) LeaveTypeController {
	return &leaveTypeController{
		service: service,
	}
}

func (c leaveTypeController) FindAll(ctx *fiber.Ctx) error {
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

func (c leaveTypeController) FindByID(ctx *fiber.Ctx) error {
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

func (c leaveTypeController) Create(ctx *fiber.Ctx) error {
	var leaveType models.LeaveType
	if err := ctx.BodyParser(&leaveType); err != nil {
		return middleware.NewErrorResponses(ctx, err)
	}

	result, err := c.service.Create(leaveType)
	if err != nil {
		return middleware.NewErrorResponses(ctx, err)
	}
	return middleware.NewSuccessResponse(ctx, result)
}

func (c leaveTypeController) Update(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return middleware.NewErrorResponses(ctx, err)
	}

	var leaveType models.LeaveType
	if err := ctx.BodyParser(&leaveType); err != nil {
		return middleware.NewErrorResponses(ctx, err)
	}

	result, err := c.service.Update(uint(id), leaveType)
	if err != nil {
		return middleware.NewErrorResponses(ctx, err)
	}
	return middleware.NewSuccessResponse(ctx, result)
}

func (c leaveTypeController) Delete(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return middleware.NewErrorResponses(ctx, err)
	}
	if err := c.service.Delete(uint(id)); err != nil {
		return middleware.NewErrorResponses(ctx, err)
	}
	return middleware.NewSuccessResponse(ctx, nil)
}
