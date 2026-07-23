package http

import (
	commentdto "api-for-shops-on-facebook-page/internal/module/comment_facebook/adapter/data_transfer_object"
	"api-for-shops-on-facebook-page/internal/module/comment_facebook/domain"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CommentFacebookHttp struct {
	usecase domain.CommentFacebookUsecase
}

func NewCommentFacebookHttp(usecase domain.CommentFacebookUsecase) *CommentFacebookHttp {
	return &CommentFacebookHttp{
		usecase: usecase,
	}
}

func (uc *CommentFacebookHttp) CommentList(ctx *gin.Context) {
	// feedId := ctx.Param("feedId")

	commentModelDTO := &commentdto.CommentFacebookDTO{}

	if err := ctx.ShouldBind(&commentModelDTO); err != nil {
		messageError := fmt.Errorf("comment_facebook_http.go(CommentList): %w", err)
		log.Fatalln(messageError)
		return
	}

	fmt.Printf("\n[test] %s\n", commentModelDTO.FeedId)

	// feedDomain := &domain.CommentFacebook{
	// FeedId: feedId,
	// }

	res, err := uc.usecase.CommentList(ctx, commentModelDTO.ToDomain())

	if err != nil {
		messageError := fmt.Errorf("comment_facebook_http.go(CommentList): %w", err)
		log.Fatalln(messageError)

		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"status":  false,
				"message": "error",
				"data":    messageError,
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
