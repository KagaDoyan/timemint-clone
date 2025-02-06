package route

import (
	"go-fiber/api/controller"
	"go-fiber/api/middleware"
	"go-fiber/data/repositories"
	"go-fiber/data/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func NewShiftRouter(router fiber.Router, db *gorm.DB) {
	shiftRepo := repositories.NewShiftRepository(db)
	shiftService := services.NewShiftService(shiftRepo)
	shiftController := controller.NewShiftController(shiftService)
	shiftRoute := router.Group("/shifts", middleware.AccessToken)
	{
		shiftRoute.Get("/all", shiftController.FindAll)
		shiftRoute.Get("/id/:id", shiftController.FindByID)
		shiftRoute.Post("/", middleware.WithRoles(middleware.RoleAdmin, middleware.RoleManager), shiftController.Create)
		shiftRoute.Put("/:id", middleware.WithRoles(middleware.RoleAdmin, middleware.RoleManager), shiftController.Update)
		shiftRoute.Delete("/:id", middleware.WithRoles(middleware.RoleAdmin, middleware.RoleManager), shiftController.Delete)
		shiftRoute.Get("/options", shiftController.Option)
		shiftRoute.Get("/report", shiftController.ShiftReport)
	}
}
