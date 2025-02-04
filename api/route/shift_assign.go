package route

import (
	"go-fiber/api/controller"
	"go-fiber/api/middleware"
	"go-fiber/data/repositories"
	"go-fiber/data/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func NewShiftAssignRouter(router fiber.Router, db *gorm.DB) {
	shiftAssignRepo := repositories.NewShiftAssignRepository(db)
	shiftAssignService := services.NewShiftAssignService(shiftAssignRepo)
	shiftAssignController := controller.NewShiftAssignController(shiftAssignService)
	shiftAssignRoute := router.Group("/shift-assign", middleware.AccessToken)
	{
		shiftAssignRoute.Get("/id/:id", shiftAssignController.FindById)
		shiftAssignRoute.Get("/all", shiftAssignController.FindAll)
		shiftAssignRoute.Post("/", middleware.WithRoles(middleware.RoleAdmin, middleware.RoleManager), shiftAssignController.Create)
		shiftAssignRoute.Post("/batch", middleware.WithRoles(middleware.RoleAdmin, middleware.RoleManager), shiftAssignController.CreateBatch)
		shiftAssignRoute.Delete("/:id", middleware.WithRoles(middleware.RoleAdmin, middleware.RoleManager), shiftAssignController.Delete)
		shiftAssignRoute.Get("/calendar/:month/:year", shiftAssignController.CalendarShift)
	}
}
