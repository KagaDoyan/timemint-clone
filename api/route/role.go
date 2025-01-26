package route

import (
	"go-fiber/api/controller"
	"go-fiber/data/repositories"
	"go-fiber/data/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func NewRoleRouter(router fiber.Router, db *gorm.DB) {
	roleRepo := repositories.NewRoleRepository(db)
	roleService := services.NewRoleService(roleRepo)
	roleController := controller.NewRoleController(roleService)
	roleRoute := router.Group("/roles")
	{
		roleRoute.Post("/", roleController.Create)
		roleRoute.Put("/:id", roleController.Update)
		roleRoute.Delete("/:id", roleController.Delete)
		roleRoute.Get("/all", roleController.FindAll)
		roleRoute.Get("/id/:id", roleController.FindByID)
	}
}
