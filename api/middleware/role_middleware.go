package middleware

import (
	"fmt"
	"go-fiber/bootstrap"
	"go-fiber/core/logs"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/kataras/jwt"
)

// Role constants for easier management
const (
	RoleAdmin    = "ADMIN"
	RoleManager  = "MANAGER"
	RoleEmployee = "USER"
)

// RolePermission defines the allowed roles for specific routes
type RolePermission struct {
	Roles []string
}

// WithRoles creates a middleware that checks if the user has the required roles
func WithRoles(allowedRoles ...string) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		// First, verify access token
		err := AccessToken(ctx)
		if err != nil {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"status":  false,
				"message": "Unauthorized",
			})
		}

		// Extract user role from token
		userRole, err := extractUserRole(ctx)
		if err != nil {
			logs.Error(fmt.Sprintf("Role extraction failed: %v", err))
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"status":  false,
				"message": "Unauthorized",
			})
		}

		// Check if user's role is in the allowed roles
		if !roleAllowed(userRole, allowedRoles) {
			logs.Error(fmt.Sprintf("Role access denied. User role: %s, Allowed roles: %v", userRole, allowedRoles))
			return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"status":  false,
				"message": "Forbidden: Insufficient permissions",
			})
		}
		// Important: Always call Next() to pass control to the next handler
		return nil
	}
}

// extractUserRole extracts the user's role from the access token
func extractUserRole(ctx *fiber.Ctx) (string, error) {
	auth := ctx.Get("Authorization")
	if auth == "" {
		return "", fmt.Errorf("no authorization token")
	}

	jwtFromHeader := strings.TrimSpace(auth[7:])
	_, decrypt, _ := jwt.GCM([]byte(bootstrap.GlobalEnv.JWT.AccessToken), nil)
	verifiedToken, err := jwt.VerifyEncrypted(jwt.HS256, []byte(bootstrap.GlobalEnv.JWT.AccessToken), decrypt, []byte(jwtFromHeader))
	if err != nil {
		return "", err
	}

	var claims ClaimsToken
	err = verifiedToken.Claims(&claims)
	if err != nil {
		return "", err
	}

	// Log the extracted role for debugging
	// logs.Info(fmt.Sprintf("Extracted role: %s", claims.Role))

	return claims.Role, nil
}

// roleAllowed checks if the user's role is in the list of allowed roles
func roleAllowed(userRole string, allowedRoles []string) bool {
	// If no roles are specified, allow access
	if len(allowedRoles) == 0 {
		return true
	}

	// Check if user's role is in the allowed roles
	for _, role := range allowedRoles {
		if strings.EqualFold(userRole, role) {
			return true
		}
	}

	return false
}

// Predefined role permission sets
var (
	AdminOnly    = WithRoles(RoleAdmin)
	ManagerOnly  = WithRoles(RoleManager)
	EmployeeOnly = WithRoles(RoleEmployee)

	// Combinations
	AdminOrManager  = WithRoles(RoleAdmin, RoleManager)
	AdminOrEmployee = WithRoles(RoleAdmin, RoleEmployee)
)
