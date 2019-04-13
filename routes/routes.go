/*
|--------------------------------------------------------------------------
| API Routes
|--------------------------------------------------------------------------
|
| Here is where you can register API routes for your application.
| TokenAuthMiddleware middleware is used for X_API_KEY authentication.
| Enjoy building your API!
|
*/

package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-oauth2/gin-server"
	"github.com/posts-api/controllers"
	"github.com/posts-api/helpers"
	"github.com/posts-api/middlewares"
)

func PaveRoutes() *gin.Engine {
	r := gin.Default()

	// websocket setup
	hub := controllers.NewHub()
	go hub.Run()

	r.Use(cors.Default())

	config := ginserver.Config{
		ErrorHandleFunc: func(ctx *gin.Context, err error) {
			helpers.RespondWithError(ctx, 401, "invalid access_token")
		},
		TokenKey: "github.com/go-oauth2/gin-server/access-token",
		Skipper: func(_ *gin.Context) bool {
			return false
		},
	}

	controllers.AuthInit()

	// Grouped api
	v1 := r.Group("/api/v1")
	{
		v1.GET("/", controllers.Home)
		v1.POST("/register", controllers.Register)
		v1.POST("/login", controllers.Login)
		v1.GET("/credentials", controllers.GetCredentials)
		v1.GET("/token", controllers.GetToken)
		auth := v1.Group("auth")
		{
			auth.Use(ginserver.HandleTokenVerify(config))
			auth.Use(middlewares.AuthMiddleware())
			auth.GET("/", controllers.Verify)
			auth.GET("post", controllers.Posts)
			auth.POST("post", controllers.Post)
			auth.GET("/whoami", controllers.WhoAmI)
		}

		// websocket route
		r.GET("/ws", func(c *gin.Context) {
			controllers.ServeWebsocket(hub, c.Writer, c.Request)
		})

	}

	return r
}
