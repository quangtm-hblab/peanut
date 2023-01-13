package infra

import (
	"net/http"
	"time"

	"peanut/controller"
	"peanut/middleware"

	_ "peanut/docs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

type Server struct {
	Router *gin.Engine
	Store  *gorm.DB
}

func SetupServer(store *gorm.DB) Server {
	// Init router
	r := gin.New()
	r.MaxMultipartMemory = 8 << 20

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	// Custom middleware
	r.Use(middleware.HandleError)
	r.NoRoute(middleware.HandleNoRoute)
	r.NoMethod(middleware.HandleNoMethod)

	// CORS config
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST", "GET", "PATCH", "PUT", "DELETE"},
		AllowHeaders:     []string{"Access-Control-Allow-Headers", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           24 * time.Hour,
	}))

	// Config route
	v1 := r.Group("api/v1")
	{
		userCtrl := controller.NewUserController(store)
		bookCrtl := controller.NewBookController(store)
		contentCtrl := controller.NewContentController(store)
		auth := v1.Group("/auth")
		{
			auth.POST("/login", userCtrl.Login)
			auth.POST("/signup", userCtrl.CreateUser)
		}
		users := v1.Group("/users")
		users.Use(middleware.Authentication)
		{
			users.GET("", userCtrl.GetUsers)
			users.POST("", userCtrl.CreateUser)
			users.GET("/:id", userCtrl.GetUser)
			// users.PATCH("/:id", userCtrl.UpdateUser)
			// users.DELETE("/:id", userCtrl.DeleteUserByID)
		}

		books := v1.Group("/books")
		books.Use(middleware.Authentication)
		{
			books.GET("", bookCrtl.GetBooks)
			books.POST("", bookCrtl.CreateBook)
			books.GET("/:id", bookCrtl.GetBook)
			books.PUT("/:id", bookCrtl.UpdateBook)
			books.DELETE("/:id", bookCrtl.DeleteBook)
		}

		contents := v1.Group("/contents")
		contents.Use(middleware.Authentication)
		{
			contents.GET("", contentCtrl.GetContents)
			contents.POST("", contentCtrl.CreateContent)
		}

	}

	// health check
	r.GET("api/health", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return Server{
		Store:  store,
		Router: r,
	}
}
