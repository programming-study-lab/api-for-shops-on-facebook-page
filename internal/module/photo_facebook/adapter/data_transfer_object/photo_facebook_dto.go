package photofacebookdto

import (
	"api-for-shops-on-facebook-page/internal/module/photo_facebook/domain"
	"fmt"
)

// "api-for-shops-on-facebook-page/internal/module/photo_facebook/domain"

type PhotoFacebookDTO struct {
	Caption   string `json:"caption" form:"caption"`
	Url       string `json:"url" form:"url"`
	Published string `json:"published" form:"published"`
}

type PhotoFacebookResponseDTO struct {
	Id *string `json:"id"`
}

func (photoDTO *PhotoFacebookDTO) ToDomain() *domain.PhotoFacebook {
	// photoModelDomain := &domain.PhotoFacebook{}

	// if photoDTO.Caption != nil {
	// 	photoModelDomain.Caption = photoDTO.Caption
	// }

	// if photoDTO.Published != nil {
	// 	photoModelDomain.Published = photoDTO.Published
	// }

	// if photoDTO.Url != nil {
	// 	photoModelDomain.Url = photoDTO.Url
	// }

	// return photoModelDomain
	fmt.Printf("\n[debug] %s\n", photoDTO)
	a := &domain.PhotoFacebook{
		Caption:   photoDTO.Caption,
		Url:       photoDTO.Url,
		Published: photoDTO.Published,
	}
	fmt.Printf("\n[debug 2] %s\n", a)
	return a
	// return &domain.PhotoFacebook{
	// Caption:   photoDTO.Caption,
	// Url:       photoDTO.Url,
	// Published: photoDTO.Published,
	// }
}

func (photoDTOModel *PhotoFacebookDTO) FromDomain(photoModel *domain.PhotoFacebook) *PhotoFacebookDTO {
	// if photoModel.Caption != nil {
	// 	photoDTOModel.Caption = photoModel.Caption
	// }

	// if photoModel.Published != nil {
	// 	photoDTOModel.Published = photoModel.Published
	// }

	// if photoModel.Url != nil {
	// 	photoDTOModel.Url = photoModel.Url
	// }
	// return photoDTOModel
	fmt.Printf("\n[photo_facebook_dto.go(FromDomain)] %s\n", photoDTOModel)
	return &PhotoFacebookDTO{
		Caption:   photoModel.Caption,
		Url:       photoModel.Url,
		Published: photoModel.Published,
	}
}
