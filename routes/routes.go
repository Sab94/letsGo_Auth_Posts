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
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/posts-api/controllers"
	"github.com/posts-api/helpers"
	"github.com/posts-api/middlewares"
	"github.com/posts-api/queries"
	"log"
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

	schemaConfig := graphql.SchemaConfig{Query: queries.RootQuery}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}
	h := handler.New(&handler.Config{
		Schema:     &schema,
		Pretty:     true,
		GraphiQL:   true,
		Playground: true,
	})

	g := r.Group("/graphql")
	{
		g.POST("/", func(c *gin.Context) {
			h.ServeHTTP(c.Writer, c.Request)
		})
		g.GET("/", func(c *gin.Context) {
			h.ServeHTTP(c.Writer, c.Request)
		})
	}

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

			auth.GET("/allUsers", controllers.GetAllUsers)

			auth.GET("/follow/:id", controllers.Follow)
			auth.GET("/unfollow/:id", controllers.Unfollow)
		}

		// websocket route
		r.GET("/ws", func(c *gin.Context) {
			controllers.ServeWebsocket(hub, c.Writer, c.Request)
		})

	}

	return r
}
