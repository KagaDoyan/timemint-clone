package route

import (
	"go-fiber/api/controller"
	"go-fiber/api/middleware"
	"go-fiber/data/repositories"
	"go-fiber/data/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func NewEventRouter(router fiber.Router, db *gorm.DB) {
	eventRepo := repositories.NewEventRepository(db)
	eventService := services.NewEventService(eventRepo)
	eventController := controller.NewEventController(eventService)
	eventRoute := router.Group("/events", middleware.AccessToken)

	eventRoute.Post("/", middleware.WithRoles(middleware.RoleAdmin, middleware.RoleManager), eventController.Create)
	eventRoute.Put("/:id", middleware.WithRoles(middleware.RoleAdmin, middleware.RoleManager), eventController.Update)
	eventRoute.Delete("/:id", middleware.WithRoles(middleware.RoleAdmin, middleware.RoleManager), eventController.Delete)
	eventRoute.Get("/all", eventController.FindAll)
	eventRoute.Get("/is-event/:date", eventController.IsEvent)
	eventRoute.Get("/calendar/:month/:year", eventController.CalendarEvent)
}
