package controller

import (
	"go-fiber/api/middleware"
	"go-fiber/data/services"
	"go-fiber/domain/models"

	"github.com/gofiber/fiber/v2"
)

type leaveRequestController struct {
	services.LeaveRequestService
}

type LeaveRequestController interface {
	EmpLeaveRequest(ctx *fiber.Ctx) error
	AdminLeaveCraete(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
	FindAll(ctx *fiber.Ctx) error
	FindByID(ctx *fiber.Ctx) error
}

func NewLeaveRequestController(service services.LeaveRequestService) LeaveRequestController {
	return leaveRequestController{service}
}

func (c leaveRequestController) EmpLeaveRequest(ctx *fiber.Ctx) error {
	leaverequest := models.LeaveRequest{}
	if err := ctx.BodyParser(&leaverequest); err != nil {
		return err
	}
	userID, err := middleware.GetOwnerAccessToken(ctx)
	if err != nil {
		return err
	}
	employeeID := *userID
	leaverequest.EmployeeID = employeeID
	result, err := c.LeaveRequestService.EmpLeaveRequests(employeeID, leaverequest)
	if err != nil {
		return err
	}
	return middleware.NewSuccessResponse(ctx, result)
}

func (c leaveRequestController) AdminLeaveCraete(ctx *fiber.Ctx) error {
	leaverequest := models.LeaveRequest{}
	if err := ctx.BodyParser(&leaverequest); err != nil {
		return err
	}
	userID, err := middleware.GetOwnerAccessToken(ctx)
	if err != nil {
		return err
	}
	employeeID := *userID
	leaverequest.EmployeeID = employeeID
	result, err := c.LeaveRequestService.CraeteLeaveRequests(employeeID, leaverequest)
	if err != nil {
		return err
	}
	return middleware.NewSuccessResponse(ctx, result)
}

func (c leaveRequestController) Update(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return err
	}
	leaverequest := models.LeaveRequest{}
	if err := ctx.BodyParser(&leaverequest); err != nil {
		return err
	}
	result, err := c.LeaveRequestService.Update(uint(id), leaverequest)
	if err != nil {
		return err
	}
	return middleware.NewSuccessResponse(ctx, result)
}

func (c leaveRequestController) Delete(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return err
	}
	if err := c.LeaveRequestService.Delete(uint(id)); err != nil {
		return err
	}
	return middleware.NewSuccessResponse(ctx, nil)
}

func (c leaveRequestController) FindAll(ctx *fiber.Ctx) error {
	limit := ctx.QueryInt("limit", 10)
	page := ctx.QueryInt("page", 1)
	status := ctx.Query("status")
	empID := ctx.QueryInt("empID")
	result, total, err := c.LeaveRequestService.FindAll(page, limit, status, uint(empID))
	if err != nil {
		return err
	}
	//find total page
	totalPages := (total + int64(limit) - 1) / int64(limit)
	return middleware.NewSuccessResponse(ctx, fiber.Map{
		"page":       page,
		"limit":      limit,
		"totalRows":  total,
		"totalPages": totalPages,
		"data":       result,
	})
}

func (c leaveRequestController) FindByID(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return err
	}
	result, err := c.LeaveRequestService.FindById(uint(id))
	if err != nil {
		return err
	}
	return middleware.NewSuccessResponse(ctx, result)
}
