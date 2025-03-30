package mappers

import (
	domain "composition-api/internal/domain/uzi"
	api "composition-api/internal/generated/http/api"
)

func Image(image domain.Image) api.Image {
	return api.Image{
		ID:    image.Id,
		UziID: image.UziID,
		Page:  image.Page,
	}
}

func SliceImage(images []domain.Image) []api.Image {
	result := make([]api.Image, 0, len(images))
	for _, image := range images {
		result = append(result, Image(image))
	}
	return result
}
