package route

import (
	"go-fiber/api/controller"
	"go-fiber/api/middleware"
	"go-fiber/data/repositories"
	"go-fiber/data/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func NewEmployeeRouter(router fiber.Router, db *gorm.DB) {
	// Create repositories, services, and controllers
	employeeRepo := repositories.NewEmployeeRepository(db)
	employeeService := services.NewEmployeeServices(employeeRepo)
	employeeController := controller.NewEmployeeController(employeeService)

	// Employee routes
	employee_route := router.Group("/employees")
	{
		// Login route (public, no authentication required)
		employee_route.Post("/login", employeeController.Login)

		// WhoAmI route (requires access token)
		employee_route.Get("/whoami",
			middleware.AccessToken,
			employeeController.WhoAmI)

		// Admin and Manager can view all employees
		employee_route.Get("/all",
			middleware.AccessToken,
			middleware.WithRoles(middleware.RoleAdmin, middleware.RoleManager),
			employeeController.FindAll)

		// Admin and Manager can view employee details
		employee_route.Get("/id/:id",
			middleware.AccessToken,
			middleware.WithRoles(middleware.RoleAdmin, middleware.RoleManager),
			employeeController.FindByID)

		// Only Admin can create new employees
		employee_route.Post("/",
			middleware.AccessToken,
			middleware.WithRoles(middleware.RoleAdmin),
			employeeController.Create)

		// Only Admin can update employees
		employee_route.Put("/:id",
			middleware.AccessToken,
			middleware.WithRoles(middleware.RoleAdmin),
			employeeController.Update)

		// Only Admin can delete employees
		employee_route.Delete("/:id",
			middleware.AccessToken,
			middleware.WithRoles(middleware.RoleAdmin),
			employeeController.Delete)
	}
}
