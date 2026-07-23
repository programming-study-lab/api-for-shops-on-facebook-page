package http

import (
	feedfacebookdto "api-for-shops-on-facebook-page/internal/module/feed_facebook/adapter/data_transfer_object"
	"api-for-shops-on-facebook-page/internal/module/feed_facebook/domain"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FeetFacebookHttp struct {
	usecase domain.FeedFacebookUsecase
}

func NewFeedFacebookHttp(usecase domain.FeedFacebookUsecase) *FeetFacebookHttp {
	return &FeetFacebookHttp{
		usecase: usecase,
	}
}

func (uc *FeetFacebookHttp) FeedCreate(ctx *gin.Context) {

	feedFacebook := &feedfacebookdto.FeedFacebookDTO{}

	if err := ctx.ShouldBindBodyWithJSON(&feedFacebook); err != nil {
		errorMessage := fmt.Errorf("[WARNING] feed_facebook_http.go(FeedCreate): %w", err)
		log.Fatalln(errorMessage)

		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"status":  true,
				"message": "success",
				"data":    errorMessage,
			},
		)
		return
	}

	// fmt.Printf("\n[test] %s\n", feed)

	res, err := uc.usecase.FeedCreate(ctx, feedFacebook.ToDomain())

	if err != nil {
		errorMessage := fmt.Errorf("[WARNING] feed_facebook_http.go(FeedCreate): %w", err)
		log.Fatalln(errorMessage)

		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"status":  true,
				"message": "success",
				"data":    errorMessage,
			},
		)
		return
	}

	ctx.AbortWithStatusJSON(
		http.StatusOK,
		gin.H{
			"status":  true,
			"message": "success",
			"data":    &res,
		},
	)
}

func (uc *FeetFacebookHttp) FeedList(ctx *gin.Context) {
	res, err := uc.usecase.FeedList(ctx)

	if err != nil {

		errorMessage := fmt.Errorf("[WARNING] feed_facebook_http.go(FeedList): %w", err)
		log.Fatalln(errorMessage)

		return
	}

	ctx.AbortWithStatusJSON(
		http.StatusOK,
		gin.H{
			"status":  true,
			"message": "success",
			"data":    res,
		},
	)
}
