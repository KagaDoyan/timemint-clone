package controller

import (
	"go-fiber/api/middleware"
	"go-fiber/data/services"
	"go-fiber/domain/models"

	"github.com/gofiber/fiber/v2"
)

type attendancePolicyController struct {
	attendancePolicyService services.AttendancePolicyService
}

type AttendancePolicyController interface {
	Find(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
}

func NewAttendancePolicyController(attendancePolicyService services.AttendancePolicyService) AttendancePolicyController {
	return &attendancePolicyController{
		attendancePolicyService: attendancePolicyService,
	}
}

func (c attendancePolicyController) Find(ctx *fiber.Ctx) error {
	attendancePolicy, err := c.attendancePolicyService.Find()
	if err != nil {
		return middleware.NewErrorMessageResponse(ctx, err)
	}
	return middleware.NewSuccessResponse(ctx, attendancePolicy)
}

func (c attendancePolicyController) Update(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return middleware.NewErrorMessageResponse(ctx, err)
	}
	attendancePolicy := new(models.AttendancePolicy)
	if err := ctx.BodyParser(attendancePolicy); err != nil {
		return middleware.NewErrorMessageResponse(ctx, err)
	}
	if err := c.attendancePolicyService.Update(uint(id), attendancePolicy); err != nil {
		return middleware.NewErrorMessageResponse(ctx, err)
	}
	return middleware.NewSuccessResponse(ctx, attendancePolicy)
}
