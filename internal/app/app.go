package app

import (
	commentapp "api-for-shops-on-facebook-page/internal/app/comment_app"
	photoapp "api-for-shops-on-facebook-page/internal/app/photo_app"
	"api-for-shops-on-facebook-page/internal/common/config"
	"api-for-shops-on-facebook-page/internal/infrastructure/server"
	feedAdapter "api-for-shops-on-facebook-page/internal/module/feed_facebook/adapter/feed_facebook_adapter.go"
	feedHTTP "api-for-shops-on-facebook-page/internal/module/feed_facebook/delivery/http"
	feedUsecase "api-for-shops-on-facebook-page/internal/module/feed_facebook/usecase"
	"api-for-shops-on-facebook-page/internal/router"
	"api-for-shops-on-facebook-page/routers"

	"github.com/gin-gonic/gin"
)

func Run() {

	gin.SetMode(gin.ReleaseMode)

	engine := gin.Default()

	engine.Use(gin.Recovery())

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

	// photo
	photo := photoapp.NewPhotoFacebookApp(fbConfig)
	photoHttp := photo.Run("enable")
	dependencies.Photo = photoHttp

	// router
	router := router.NewRouer(engine, dependencies)
	rEngine := router.Setup()

	routers.FacebookAPI(rEngine) // router api

	appConfig := config.LoadAppConfig()
	srv := server.NewServer(&appConfig.AppAddress, &appConfig.AppPort, rEngine)

	server := srv.Run()
	srv.Close(server)

	// rEngine.Run(":5000")

}

func RunTest(endine *gin.Engine) *gin.Engine {
	dependencies := router.NewDependencies()
	router := router.NewRouer(endine, dependencies)
	rEngine := router.Setup()

	return rEngine
}
