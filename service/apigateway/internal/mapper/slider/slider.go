package slidergraphqlmapper

import (
	graphqlmapper "github.com/MamangRust/microservice-ecommerce-grpc-apigateway/internal/mapper/pagination"
	"github.com/MamangRust/microservice-ecommerce-grpc-apigateway/internal/model"
	"github.com/MamangRust/microservice-ecommerce-grpc-pb/slider"
)

type sliederResponseMapper struct{}

func NewSliderResponseMapper() *sliederResponseMapper {
	return &sliederResponseMapper{}
}

func (s *sliederResponseMapper) ToGraphqlResponseSlider(res *pb_slider.ApiResponseSlider) *model.APIResponseSlider {
	return &model.APIResponseSlider{
		Status:  res.Status,
		Message: res.Message,
		Data:    s.mapResponseSlider(res.Data),
	}
}

func (s *sliederResponseMapper) ToGraphqlResponseSliderDeleteAt(res *pb_slider.ApiResponseSliderDeleteAt) *model.APIResponseSliderDeleteAt {
	return &model.APIResponseSliderDeleteAt{
		Status:  res.Status,
		Message: res.Message,
		Data:    s.mapResponseSliderDeleteAt(res.Data),
	}
}

func (s *sliederResponseMapper) ToGraphqlResponsesSlider(res *pb_slider.ApiResponsesSlider) *model.APIResponsesSlider {
	return &model.APIResponsesSlider{
		Status:  res.Status,
		Message: res.Message,
		Data:    s.mapResponsesSlider(res.Data),
	}
}

func (s *sliederResponseMapper) ToGraphqlResponseSliderDelete(res *pb_slider.ApiResponseSliderDelete) *model.APIResponseSliderDelete {
	return &model.APIResponseSliderDelete{
		Status:  res.Status,
		Message: res.Message,
	}
}

func (s *sliederResponseMapper) ToGraphqlResponsePaginationSliderDeleteAt(
	res *pb_slider.ApiResponsePaginationSliderDeleteAt,
) *model.APIResponsePaginationSliderDeleteAt {
	return &model.APIResponsePaginationSliderDeleteAt{
		Status:     res.Status,
		Message:    res.Message,
		Data:       s.mapResponsesSliderDeleteAt(res.Data),
		Pagination: graphqlmapper.MapPaginationMeta(res.Pagination),
	}
}

func (s *sliederResponseMapper) ToGraphqlResponseSliderAll(res *pb_slider.ApiResponseSliderAll) *model.APIResponseSliderAll {
	return &model.APIResponseSliderAll{
		Status:  res.Status,
		Message: res.Message,
	}
}

func (s *sliederResponseMapper) ToGraphqlResponsePaginationSlider(
	res *pb_slider.ApiResponsePaginationSlider,
) *model.APIResponsePaginationSlider {
	return &model.APIResponsePaginationSlider{
		Status:     res.Status,
		Message:    res.Message,
		Data:       s.mapResponsesSlider(res.Data),
		Pagination: graphqlmapper.MapPaginationMeta(res.Pagination),
	}
}

func (s *sliederResponseMapper) mapResponseSlider(slider *pb_slider.SliderResponse) *model.SliderResponse {
	if slider == nil {
		return nil
	}

	return &model.SliderResponse{
		ID:        int32(slider.Id),
		Name:      slider.Name,
		Image:     slider.Image,
		CreatedAt: slider.CreatedAt,
		UpdatedAt: slider.UpdatedAt,
	}
}

func (s *sliederResponseMapper) mapResponsesSlider(sliders []*pb_slider.SliderResponse) []*model.SliderResponse {
	mapped := make([]*model.SliderResponse, 0, len(sliders))
	for _, slider := range sliders {
		mapped = append(mapped, s.mapResponseSlider(slider))
	}
	return mapped
}

func (s *sliederResponseMapper) mapResponseSliderDeleteAt(slider *pb_slider.SliderResponseDeleteAt) *model.SliderResponseDeleteAt {
	var deletedAt string

	if slider.DeletedAt != nil {
		deletedAt = slider.DeletedAt.Value
	}

	return &model.SliderResponseDeleteAt{
		ID:        int32(slider.Id),
		Name:      slider.Name,
		Image:     slider.Image,
		CreatedAt: slider.CreatedAt,
		UpdatedAt: slider.UpdatedAt,
		DeletedAt: &deletedAt,
	}
}

func (s *sliederResponseMapper) mapResponsesSliderDeleteAt(sliders []*pb_slider.SliderResponseDeleteAt) []*model.SliderResponseDeleteAt {
	mapped := make([]*model.SliderResponseDeleteAt, 0, len(sliders))
	for _, slider := range sliders {
		mapped = append(mapped, s.mapResponseSliderDeleteAt(slider))
	}
	return mapped
}
