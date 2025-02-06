package controller

import (
	"go-fiber/api/middleware"
	"go-fiber/data/services"
	"go-fiber/domain/models"

	"github.com/gofiber/fiber/v2"
)

type departmentController struct {
	service services.DepartmentService
}

type DepartmentController interface {
	FindAll(ctx *fiber.Ctx) error
	FindById(ctx *fiber.Ctx) error
	Create(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}

func NewDepartmentController(service services.DepartmentService) DepartmentController {
	return &departmentController{
		service: service,
	}
}

func (c departmentController) FindAll(ctx *fiber.Ctx) error {
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

func (c departmentController) FindById(ctx *fiber.Ctx) error {
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

func (c departmentController) Create(ctx *fiber.Ctx) error {
	var department models.Department
	if err := ctx.BodyParser(&department); err != nil {
		return middleware.NewErrorResponses(ctx, err)
	}

	result, err := c.service.Create(department)
	if err != nil {
		return middleware.NewErrorResponses(ctx, err)
	}
	return middleware.NewSuccessResponse(ctx, result)
}

func (c departmentController) Update(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return middleware.NewErrorResponses(ctx, err)
	}

	var department models.Department
	if err := ctx.BodyParser(&department); err != nil {
		return middleware.NewErrorResponses(ctx, err)
	}

	result, err := c.service.Update(uint(id), department)
	if err != nil {
		return middleware.NewErrorResponses(ctx, err)
	}
	return middleware.NewSuccessResponse(ctx, result)
}

func (c departmentController) Delete(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return middleware.NewErrorResponses(ctx, err)
	}
	if err := c.service.Delete(uint(id)); err != nil {
		return middleware.NewErrorResponses(ctx, err)
	}
	return middleware.NewSuccessResponse(ctx, nil)
}
