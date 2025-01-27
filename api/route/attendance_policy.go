package route

import (
	"go-fiber/api/controller"
	"go-fiber/api/middleware"
	"go-fiber/data/repositories"
	"go-fiber/data/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func NewAttendancePolicyRouter(router fiber.Router, db *gorm.DB) {
	attendancePolicyRepo := repositories.NewAttendancePolicyRepository(db)
	attendancePolicyService := services.NewAttendancePolicyService(attendancePolicyRepo)
	attendancePolicyController := controller.NewAttendancePolicyController(attendancePolicyService)
	attendancePolicyRoute := router.Group("/attendance-policies", middleware.AccessToken)
	{
		attendancePolicyRoute.Get("/", attendancePolicyController.Find)
		attendancePolicyRoute.Put("/:id", middleware.WithRoles(middleware.RoleAdmin), attendancePolicyController.Update)
	}
}
