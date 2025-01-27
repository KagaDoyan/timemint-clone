package route

import (
	"go-fiber/api/controller"
	"go-fiber/api/middleware"
	"go-fiber/data/repositories"
	"go-fiber/data/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func NewActiveLocationRouter(router fiber.Router, db *gorm.DB) {
	activeLocationRepo := repositories.NewActiveLocationRepository(db)
	activeLocationService := services.NewActiveLocationService(activeLocationRepo)
	activeLocationController := controller.NewActiveLocationController(activeLocationService)
	activeLocationRoute := router.Group("/active-locations", middleware.AccessToken)
	{
		activeLocationRoute.Get("/all", middleware.WithRoles(middleware.RoleAdmin, middleware.RoleManager, middleware.RoleEmployee), activeLocationController.FindAll)
		activeLocationRoute.Get("/id/:id", middleware.WithRoles(middleware.RoleAdmin, middleware.RoleManager, middleware.RoleEmployee), activeLocationController.Find)
		activeLocationRoute.Post("/", middleware.WithRoles(middleware.RoleAdmin), activeLocationController.Create)
		activeLocationRoute.Put("/:id", middleware.WithRoles(middleware.RoleAdmin), activeLocationController.Update)
		activeLocationRoute.Delete("/:id", middleware.WithRoles(middleware.RoleAdmin), activeLocationController.Delete)
	}
}
