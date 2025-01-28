package route

import (
	"go-fiber/api/controller"
	"go-fiber/api/middleware"
	"go-fiber/data/repositories"
	"go-fiber/data/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func NewDepartmentRouter(router fiber.Router, db *gorm.DB) {
	departmentRepo := repositories.NewDepartmentRepository(db)
	departmentService := services.NewDepartmentService(departmentRepo)
	departmentController := controller.NewDepartmentController(departmentService)
	departmentRoute := router.Group("/departments", middleware.AccessToken)
	{
		departmentRoute.Get("/all", departmentController.FindAll)
		departmentRoute.Get("/id/:id", departmentController.FindById)
		departmentRoute.Post("/", middleware.WithRoles(middleware.RoleAdmin), departmentController.Create)
		departmentRoute.Put("/:id", middleware.WithRoles(middleware.RoleAdmin), departmentController.Update)
		departmentRoute.Delete("/:id", middleware.WithRoles(middleware.RoleAdmin), departmentController.Delete)
	}
}
