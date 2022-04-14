package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"final/controllers"
	"final/middlewares"

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// set db to gin context
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	PassMiddlewareRoute := r.Group("/ChangePassword")
	PassMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
	PassMiddlewareRoute.PATCH("/:username", controllers.ChangePassword)

	r.GET("/Catagories", controllers.GetAllCat)
	r.GET("/Catagories/:id", controllers.GetCatById)

	CatMiddlewareRoute := r.Group("/Catagories")
	CatMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())

	CatMiddlewareRoute.POST("/", controllers.CreateCat)
	CatMiddlewareRoute.PATCH("/:id", controllers.UpdateCat)
	CatMiddlewareRoute.DELETE("/:id", controllers.DeleteCat)

	r.GET("/Games", controllers.GetAllGame)
	r.GET("/Games/:id", controllers.GetGameById)
	r.GET("/Games/:id/Categories", controllers.GetGameByCategory)

	GameMiddlewareRoute := r.Group("/Games")
	GameMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())

	GameMiddlewareRoute.POST("/", controllers.CreateGame)
	GameMiddlewareRoute.PATCH("/:id", controllers.UpdateGame)
	GameMiddlewareRoute.DELETE("/:id", controllers.DeleteGame)

	r.GET("/HistoryRatings", controllers.GetAllHistoryRating)
	r.GET("/HistoryRatings/:id", controllers.GetHistoryRatingById)
	r.POST("/HistoryRatings", controllers.CreateHistoryRating)
	r.GET("/HistoryRatings/:id/Games", controllers.GetHistoryRatingByGame)
	r.PATCH("/HistoryRatings/:id", controllers.UpdateHistoryRating)
	r.DELETE("HistoryRatings/:id", controllers.DeleteHistoryRating)

	r.GET("/Rating", controllers.GetAllRating)

	RatingMiddlewareRoute := r.Group("/Rating")
	RatingMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())

	RatingMiddlewareRoute.POST("/", controllers.CreateRating)
	RatingMiddlewareRoute.PATCH("/:id", controllers.UpdateRating)
	RatingMiddlewareRoute.DELETE("/:id", controllers.DeleteRating)

	r.GET("/Reviews", controllers.GetAllReview)
	r.GET("/Reviews/:id", controllers.GetReviewById)
	r.GET("/Reviews/:id/Games", controllers.GetReviewByGame)
	r.GET("/Reviews/:id/Users", controllers.GetReviewByUser)

	ReviewMiddlewareRoute := r.Group("/Review")
	ReviewMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())

	ReviewMiddlewareRoute.POST("/", controllers.CreateReview)
	ReviewMiddlewareRoute.PATCH("/:id", controllers.UpdateReview)
	ReviewMiddlewareRoute.DELETE("/:id", controllers.DeleteReview)

	r.GET("/TopReview", controllers.GetAllTopR)

	TopMiddlewareRoute := r.Group("/TopRiview")
	TopMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())

	TopMiddlewareRoute.POST("/TopReview", controllers.CreateTopR)
	TopMiddlewareRoute.PATCH("/TopReview/:id", controllers.UpdateTopR)
	TopMiddlewareRoute.DELETE("TopReview/:id", controllers.DeleteTopR)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
