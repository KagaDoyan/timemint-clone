package controller

import (
	"go-fiber/api/middleware"
	"go-fiber/data/services"
	"go-fiber/domain/models"
	"strings"

	"fmt"

	"github.com/gofiber/fiber/v2"
)

type employeeController struct {
	employeeService services.EmployeeService
}

type EmployeeController interface {
	FindAll(ctx *fiber.Ctx) error
	FindByID(ctx *fiber.Ctx) error
	Create(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
	FindByEmail(ctx *fiber.Ctx) error
	Login(ctx *fiber.Ctx) error
	WhoAmI(ctx *fiber.Ctx) error
}

func (c *employeeController) FindAll(ctx *fiber.Ctx) error {
	// Parse pagination parameters
	page := ctx.QueryInt("page", 1)
	limit := ctx.QueryInt("limit", 10)

	// Call service method with pagination
	result, total, err := c.employeeService.FindAll(page, limit)
	if err != nil {
		return middleware.NewErrorMessageResponse(ctx, err)
	}

	// Calculate total pages
	totalPages := (total + int64(limit) - 1) / int64(limit)

	// Return paginated response
	return middleware.NewSuccessResponse(ctx, fiber.Map{
		"page":       page,
		"limit":      limit,
		"totalRows":  total,
		"totalPages": totalPages,
		"rows":       result,
	})
}

func (c *employeeController) FindByID(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return middleware.NewErrorMessageResponse(ctx, err)
	}

	result, err := c.employeeService.FindByID(uint(id))
	if err != nil {
		return middleware.NewErrorMessageResponse(ctx, err)
	}
	return middleware.NewSuccessResponse(ctx, result)
}

func (c *employeeController) Create(ctx *fiber.Ctx) error {
	var employee models.Employee
	if err := ctx.BodyParser(&employee); err != nil {
		return middleware.NewErrorMessageResponse(ctx, err)
	}
	data, err := c.employeeService.Create(&employee)
	if err != nil {
		return middleware.NewErrorMessageResponse(ctx, err)
	}
	return middleware.NewSuccessResponse(ctx, data)
}

func (c *employeeController) Update(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return middleware.NewErrorMessageResponse(ctx, err)
	}

	var employee models.Employee
	if err := ctx.BodyParser(&employee); err != nil {
		return middleware.NewErrorMessageResponse(ctx, err)
	}

	// Ensure the ID from the path matches the body
	employee.ID = uint(id)
	data, err := c.employeeService.Update(&employee)
	if err != nil {
		return middleware.NewErrorMessageResponse(ctx, err)
	}
	return middleware.NewSuccessResponse(ctx, data)
}

func (c *employeeController) Delete(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return middleware.NewErrorMessageResponse(ctx, err)
	}

	if err := c.employeeService.Delete(uint(id)); err != nil {
		return middleware.NewErrorMessageResponse(ctx, err)
	}
	return middleware.NewSuccessResponse(ctx, nil)
}

func (c *employeeController) FindByEmail(ctx *fiber.Ctx) error {
	email := ctx.Query("email")
	if email == "" {
		return middleware.NewErrorMessageResponse(ctx, fiber.NewError(fiber.StatusBadRequest, "Email is required"))
	}

	result, err := c.employeeService.FindByEmail(email)
	if err != nil {
		return middleware.NewErrorMessageResponse(ctx, err)
	}
	return middleware.NewSuccessResponse(ctx, result)
}

func (c *employeeController) Login(ctx *fiber.Ctx) error {
	// Parse login request
	var loginRequest struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}

	if err := ctx.BodyParser(&loginRequest); err != nil {
		return middleware.NewErrorMessageResponse(ctx, err)
	}

	// Attempt login
	employee, err := c.employeeService.Login(loginRequest.Email, loginRequest.Password)
	if err != nil {
		return middleware.NewErrorMessageResponse(ctx, err)
	}

	// Generate JWT token
	tokenPair, err := middleware.GenerateJWTToken(fmt.Sprintf("%d", employee.ID))
	if err != nil {
		return middleware.NewErrorMessageResponse(ctx, err)
	}

	// Return login response
	return middleware.NewSuccessResponse(ctx, fiber.Map{
		"user": employee,
		"tokens": fiber.Map{
			"access_token":  strings.Replace(string(tokenPair.AccessToken), "\"", "", -1),
			"refresh_token": strings.Replace(string(tokenPair.RefreshToken), "\"", "", -1),
		},
	})
}

func (c *employeeController) WhoAmI(ctx *fiber.Ctx) error {
	// Extract user ID from the access token
	employeeID, err := middleware.GetOwnerAccessToken(ctx)
	if err != nil {
		return middleware.NewErrorMessageResponse(ctx, err)
	}
	// Find employee details
	employee, err := c.employeeService.FindByID(*employeeID)
	if err != nil {
		return middleware.NewErrorMessageResponse(ctx, err)
	}

	// Return user details
	return middleware.NewSuccessResponse(ctx, employee)
}

func NewEmployeeController(employeeService services.EmployeeService) EmployeeController {
	return &employeeController{employeeService: employeeService}
}
