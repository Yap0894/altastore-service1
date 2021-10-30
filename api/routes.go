package api

import (
	"AltaStore/api/middleware"
	"AltaStore/api/v1/admin"
	"AltaStore/api/v1/adminauth"
	"AltaStore/api/v1/user"
	"AltaStore/api/v1/userauth"
	"net/http"

	echo "github.com/labstack/echo/v4"
)

func RegisterPath(e *echo.Echo,
	userController *user.Controller,
	adminController *admin.Controller,
	userAuthController *userauth.Controller,
	adminAuthController *adminauth.Controller,
) {
	if userController == nil ||
		userAuthController == nil ||
		adminController == nil ||
		adminAuthController == nil {
		panic("Invalid parameter")
	}

	// Add logger
	e.Use(middleware.MiddlewareLogger)

	// Custome response
	e.HTTPErrorHandler = func(e error, c echo.Context) {
		type Response struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
		}
		var response Response
		response.Code = http.StatusInternalServerError // defaul 500
		response.Message = "Internal Server Error"

		if he, ok := e.(*echo.HTTPError); ok {
			response.Code = he.Code
			response.Message = http.StatusText(he.Code)
		}

		c.Logger().Error(e)

		_ = c.JSON(response.Code, response)
	}

	// Routing
	regis := e.Group("v1/register")
	regis.POST("", userController.InsertUser)
	regis.POST("/admin", adminController.InsertAdmin)

	login := e.Group("v1/login")
	login.POST("", userAuthController.UserLogin)
	login.POST("/admin", adminAuthController.AdminLogin)

	admin := e.Group("v1/admins")
	admin.Use(middleware.JWTMiddleware())
	admin.PUT("/:id", adminController.UpdateAdmin)
	admin.DELETE("/:id", adminController.DeleteAdmin)
	admin.GET("/:id", adminController.FindAdminByID)
	admin.PUT("/:id/password", adminController.UpdateAdminPassword)

	user := e.Group("v1/users")
	user.Use(middleware.JWTMiddleware())
	user.PUT("/:id", userController.UpdateUser)
	user.DELETE("/:id", userController.DeleteUser)
	user.GET("/:id", userController.FindUserByID)
	user.PUT("/:id/password", userController.UpdateUserPassword)
}
