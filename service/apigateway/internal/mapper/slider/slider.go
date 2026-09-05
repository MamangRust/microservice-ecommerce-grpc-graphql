package slidergraphqlmapper

import (
	"github.com/MamangRust/microservice-ecommerce-shared/pb"
	graphqlmapper "github.com/MamangRust/monolith-graphql-ecommerce-apigateway/internal/mapper/pagination"
	"github.com/MamangRust/monolith-graphql-ecommerce-apigateway/internal/model"
)

type sliederResponseMapper struct{}

func NewSliderResponseMapper() *sliederResponseMapper {
	return &sliederResponseMapper{}
}

func (s *sliederResponseMapper) ToGraphqlResponseSlider(res *pb.ApiResponseSlider) *model.APIResponseSlider {
	return &model.APIResponseSlider{
		Status:  res.Status,
		Message: res.Message,
		Data:    s.mapResponseSlider(res.Data),
	}
}

func (s *sliederResponseMapper) ToGraphqlResponseSliderDeleteAt(res *pb.ApiResponseSliderDeleteAt) *model.APIResponseSliderDeleteAt {
	return &model.APIResponseSliderDeleteAt{
		Status:  res.Status,
		Message: res.Message,
		Data:    s.mapResponseSliderDeleteAt(res.Data),
	}
}

func (s *sliederResponseMapper) ToGraphqlResponsesSlider(res *pb.ApiResponsesSlider) *model.APIResponsesSlider {
	return &model.APIResponsesSlider{
		Status:  res.Status,
		Message: res.Message,
		Data:    s.mapResponsesSlider(res.Data),
	}
}

func (s *sliederResponseMapper) ToGraphqlResponseSliderDelete(res *pb.ApiResponseSliderDelete) *model.APIResponseSliderDelete {
	return &model.APIResponseSliderDelete{
		Status:  res.Status,
		Message: res.Message,
	}
}

func (s *sliederResponseMapper) ToGraphqlResponsePaginationSliderDeleteAt(
	res *pb.ApiResponsePaginationSliderDeleteAt,
) *model.APIResponsePaginationSliderDeleteAt {
	return &model.APIResponsePaginationSliderDeleteAt{
		Status:     res.Status,
		Message:    res.Message,
		Data:       s.mapResponsesSliderDeleteAt(res.Data),
		Pagination: graphqlmapper.MapPaginationMeta(res.Pagination),
	}
}

func (s *sliederResponseMapper) ToGraphqlResponseSliderAll(res *pb.ApiResponseSliderAll) *model.APIResponseSliderAll {
	return &model.APIResponseSliderAll{
		Status:  res.Status,
		Message: res.Message,
	}
}

func (s *sliederResponseMapper) ToGraphqlResponsePaginationSlider(
	res *pb.ApiResponsePaginationSlider,
) *model.APIResponsePaginationSlider {
	return &model.APIResponsePaginationSlider{
		Status:     res.Status,
		Message:    res.Message,
		Data:       s.mapResponsesSlider(res.Data),
		Pagination: graphqlmapper.MapPaginationMeta(res.Pagination),
	}
}

func (s *sliederResponseMapper) mapResponseSlider(slider *pb.SliderResponse) *model.SliderResponse {
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

func (s *sliederResponseMapper) mapResponsesSlider(sliders []*pb.SliderResponse) []*model.SliderResponse {
	mapped := make([]*model.SliderResponse, 0, len(sliders))
	for _, slider := range sliders {
		mapped = append(mapped, s.mapResponseSlider(slider))
	}
	return mapped
}

func (s *sliederResponseMapper) mapResponseSliderDeleteAt(slider *pb.SliderResponseDeleteAt) *model.SliderResponseDeleteAt {
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

func (s *sliederResponseMapper) mapResponsesSliderDeleteAt(sliders []*pb.SliderResponseDeleteAt) []*model.SliderResponseDeleteAt {
	mapped := make([]*model.SliderResponseDeleteAt, 0, len(sliders))
	for _, slider := range sliders {
		mapped = append(mapped, s.mapResponseSliderDeleteAt(slider))
	}
	return mapped
}
