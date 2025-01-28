package controller

import (
	"go-fiber/api/middleware"
	"go-fiber/data/services"
	"go-fiber/domain/models"

	"github.com/gofiber/fiber/v2"
)

type RoleController interface {
	FindAll(ctx *fiber.Ctx) error
	FindByID(ctx *fiber.Ctx) error
	Create(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}

type roleController struct {
	roleService services.RoleService
}

func NewRoleController(roleService services.RoleService) RoleController {
	return &roleController{roleService: roleService}
}

func (c roleController) FindAll(ctx *fiber.Ctx) error {
	page := ctx.QueryInt("page", 1)
	limit := ctx.QueryInt("limit", 10)

	result, total, err := c.roleService.FindAll(page, limit)
	if err != nil {
		return middleware.NewErrorMessageResponse(ctx, err)
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

func (c roleController) FindByID(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return middleware.NewErrorMessageResponse(ctx, err)
	}

	result, err := c.roleService.FindByID(uint(id))
	if err != nil {
		return middleware.NewErrorMessageResponse(ctx, err)
	}
	return middleware.NewSuccessResponse(ctx, result)
}

func (c roleController) Create(ctx *fiber.Ctx) error {
	var role models.Role
	if err := ctx.BodyParser(&role); err != nil {
		return middleware.NewErrorMessageResponse(ctx, err)
	}
	data, err := c.roleService.Create(role)
	if err != nil {
		return middleware.NewErrorMessageResponse(ctx, err)
	}
	return middleware.NewSuccessResponse(ctx, data)
}

func (c roleController) Update(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return middleware.NewErrorMessageResponse(ctx, err)
	}

	var role models.Role
	if err := ctx.BodyParser(&role); err != nil {
		return middleware.NewErrorMessageResponse(ctx, err)
	}

	// Ensure the ID from the path matches the body
	role.ID = uint(id)
	data, err := c.roleService.Update(uint(id), role)
	if err != nil {
		return middleware.NewErrorMessageResponse(ctx, err)
	}
	return middleware.NewSuccessResponse(ctx, data)
}

func (c roleController) Delete(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return middleware.NewErrorMessageResponse(ctx, err)
	}
	if err := c.roleService.Delete(uint(id)); err != nil {
		return middleware.NewErrorMessageResponse(ctx, err)
	}
	return middleware.NewSuccessResponse(ctx, nil)
}
