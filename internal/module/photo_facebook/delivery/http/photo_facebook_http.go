package http

import (
	photofacebookdto "api-for-shops-on-facebook-page/internal/module/photo_facebook/adapter/data_transfer_object"
	"api-for-shops-on-facebook-page/internal/module/photo_facebook/domain"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PhotoFacebookHttp struct {
	usecase domain.PhotoFacebookUsecase
}

func NewPhotoFacebookHttp(usecase domain.PhotoFacebookUsecase) *PhotoFacebookHttp {
	return &PhotoFacebookHttp{
		usecase: usecase,
	}
}

func (uc *PhotoFacebookHttp) PhotoCreate(ctx *gin.Context) {

	photoFacebook := &photofacebookdto.PhotoFacebookDTO{}

	if err := ctx.ShouldBind(&photoFacebook); err != nil {
		errorMessage := fmt.Errorf("[WARNING] photo_facebook_http.go(PhotoCreate): %w", err)
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

	fmt.Printf("\n[test] %s\n", photoFacebook.Caption)
	fmt.Printf("\n[test 2] %s\n", *photoFacebook.ToDomain())
	// return
	// a := *photoFacebook.ToDomain()
	// res, err := uc.usecase.PhotoCreate(ctx, a)
	res, err := uc.usecase.PhotoCreate(ctx, photoFacebook.ToDomain())

	if err != nil {
		errorMessage := fmt.Errorf("[WARNING] photo_facebook_http.go(PhotoCreate): %w", err)
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
