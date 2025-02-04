package route

import (
	"go-fiber/api/controller"
	"go-fiber/api/middleware"
	"go-fiber/data/repositories"
	"go-fiber/data/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func NewRoleRouter(router fiber.Router, db *gorm.DB) {
	roleRepo := repositories.NewRoleRepository(db)
	roleService := services.NewRoleService(roleRepo)
	roleController := controller.NewRoleController(roleService)
	roleRoute := router.Group("/roles", middleware.AccessToken)
	{
		roleRoute.Post("/", middleware.WithRoles(middleware.RoleAdmin, middleware.RoleManager), roleController.Create)
		roleRoute.Put("/:id", middleware.WithRoles(middleware.RoleAdmin, middleware.RoleManager), roleController.Update)
		roleRoute.Delete("/:id", middleware.WithRoles(middleware.RoleAdmin, middleware.RoleManager), roleController.Delete)
		roleRoute.Get("/all", roleController.FindAll)
		roleRoute.Get("/id/:id", roleController.FindByID)
	}
}
