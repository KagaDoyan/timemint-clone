package route

import (
	"go-fiber/api/controller"
	"go-fiber/api/middleware"
	"go-fiber/data/repositories"
	"go-fiber/data/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func NewHolidayRouter(router fiber.Router, db *gorm.DB) {
	holidayRepo := repositories.NewHolidayRepository(db)
	holidayService := services.NewHolidayService(holidayRepo)
	holidayController := controller.NewHolidayController(holidayService)
	holidayRoute := router.Group("/holidays", middleware.AccessToken)

	holidayRoute.Post("/", holidayController.Create)
	holidayRoute.Put("/:id", holidayController.Update)
	holidayRoute.Delete("/:id", holidayController.Delete)
	holidayRoute.Get("/all", holidayController.FindAll)
	holidayRoute.Get("/is-holiday/:date", holidayController.IsHoliday)
}
