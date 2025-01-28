package route

import (
	"go-fiber/api/controller"
	"go-fiber/data/repositories"
	"go-fiber/data/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func NewLeaveTypeRouter(router fiber.Router, db *gorm.DB) {
	leaveTypeRepo := repositories.NewLeaveTypeRepository(db)
	leaveTypeService := services.NewLeaveTypeService(leaveTypeRepo)
	leaveTypeController := controller.NewLeaveTypeController(leaveTypeService)
	leaveTypeRoute := router.Group("/leave-types")
	{
		leaveTypeRoute.Post("/", leaveTypeController.Create)
		leaveTypeRoute.Put("/:id", leaveTypeController.Update)
		leaveTypeRoute.Delete("/:id", leaveTypeController.Delete)
		leaveTypeRoute.Get("/all", leaveTypeController.FindAll)
		leaveTypeRoute.Get("/id/:id", leaveTypeController.FindByID)
	}
}
