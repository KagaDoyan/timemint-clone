package route

import (
	"go-fiber/api/controller"
	"go-fiber/data/repositories"
	"go-fiber/data/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func NewLeaveRequestRoute(app fiber.Router, db *gorm.DB) {
	leaveRepository := repositories.NewLeaveRequestRepository(db)
	leaveService := services.NewLeaveRequestService(leaveRepository)
	leaveRequestController := controller.NewLeaveRequestController(leaveService)

	leaveRequest := app.Group("/leave-requests")
	leaveRequest.Get("/all", leaveRequestController.FindAll)
	leaveRequest.Get("/id/:id", leaveRequestController.FindByID)
	leaveRequest.Post("/create", leaveRequestController.AdminLeaveCraete)
	leaveRequest.Post("/request", leaveRequestController.EmpLeaveRequest)
	leaveRequest.Put("/:id", leaveRequestController.Update)
	leaveRequest.Delete("/:id", leaveRequestController.Delete)
}
