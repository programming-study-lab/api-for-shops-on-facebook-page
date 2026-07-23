package app

import (
	commentapp "api-for-shops-on-facebook-page/internal/app/comment_app"
	"api-for-shops-on-facebook-page/internal/config"
	feedAdapter "api-for-shops-on-facebook-page/internal/module/feed_facebook/adapter/feed_facebook_adapter.go"
	feedHTTP "api-for-shops-on-facebook-page/internal/module/feed_facebook/delivery/http"
	feedUsecase "api-for-shops-on-facebook-page/internal/module/feed_facebook/usecase"
	"api-for-shops-on-facebook-page/internal/router"

	"github.com/gin-gonic/gin"
)

func Run() {
	engine := gin.Default()

	fbConfig := config.LoadFacebookConfig()

	dependencies := router.NewDependencies() //dependencies

	// feed
	feedAPT := feedAdapter.NewFeedFacebookAdapter(fbConfig)
	feedUC := feedUsecase.NewFeedFacebookUsecase(feedAPT)
	feedHttp := feedHTTP.NewFeedFacebookHttp(feedUC)

	dependencies.Feed = feedHttp // dependencies

	// comment
	comment := commentapp.NewCommentFacebookApp(fbConfig)
	commentHttp := comment.Run("enable")
	dependencies.Comment = commentHttp // dependencies

	router := router.NewRouer(engine, dependencies)
	rEngine := router.Setup()

	rEngine.Run(":5000")

}

func RunTest(endine *gin.Engine) *gin.Engine {
	dependencies := router.NewDependencies()
	router := router.NewRouer(endine, dependencies)
	rEngine := router.Setup()

	return rEngine
}
