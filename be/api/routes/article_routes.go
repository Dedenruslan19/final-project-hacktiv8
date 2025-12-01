package routes

import (
	"milestone3/be/api/middleware" // import admin middleware
	"milestone3/be/internal/controller"
)

func (r *EchoRouter) RegisterArticleRoutes(articleCtrl *controller.ArticleController) {
	articleRoutes := r.echo.Group("/articles")

	// public endpoints
	articleRoutes.GET("", articleCtrl.GetAllArticles)
	articleRoutes.GET("/:id", articleCtrl.GetArticleByID)

	// admin-protected: require JWT auth then admin check
	articleRoutes.POST("", articleCtrl.CreateArticle, middleware.JWTMiddleware, middleware.RequireAdmin)
	articleRoutes.PUT("/:id", articleCtrl.UpdateArticle, middleware.JWTMiddleware, middleware.RequireAdmin)
	articleRoutes.DELETE("/:id", articleCtrl.DeleteArticle, middleware.JWTMiddleware, middleware.RequireAdmin)
}
