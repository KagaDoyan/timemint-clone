package route

import (
	"go-fiber/api/controller"
	"go-fiber/api/middleware"
	"go-fiber/data/repositories"
	"go-fiber/data/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func NewDayOfWorkRouter(router fiber.Router, db *gorm.DB) {
	dayOfWorkRepo := repositories.NewDayOfWorkRepository(db)
	dayOfWorkService := services.NewDayOfWorkService(dayOfWorkRepo)
	dayOfWorkController := controller.NewDayOfWorkController(dayOfWorkService)
	dayOfWorkRoute := router.Group("/day-of-works", middleware.AccessToken, middleware.WithRoles(middleware.RoleAdmin))
	{
		dayOfWorkRoute.Post("/", dayOfWorkController.Create)
		dayOfWorkRoute.Put("/:id", dayOfWorkController.Update)
		dayOfWorkRoute.Delete("/:id", dayOfWorkController.Delete)
		dayOfWorkRoute.Get("/all", dayOfWorkController.FindAll)
		dayOfWorkRoute.Get("/id/:id", dayOfWorkController.FindByID)
	}
}
