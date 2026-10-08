package slidergraphqlmapper

import (
	"github.com/MamangRust/microservice-ecommerce-grpc-apigateway/internal/model"
	"github.com/MamangRust/microservice-ecommerce-grpc-pb/slider"
)

type SliderGraphqlMapper interface {
	ToGraphqlResponseSlider(res *pb_slider.ApiResponseSlider) *model.APIResponseSlider
	ToGraphqlResponseSliderDeleteAt(res *pb_slider.ApiResponseSliderDeleteAt) *model.APIResponseSliderDeleteAt
	ToGraphqlResponsesSlider(res *pb_slider.ApiResponsesSlider) *model.APIResponsesSlider
	ToGraphqlResponseSliderDelete(res *pb_slider.ApiResponseSliderDelete) *model.APIResponseSliderDelete
	ToGraphqlResponseSliderAll(res *pb_slider.ApiResponseSliderAll) *model.APIResponseSliderAll
	ToGraphqlResponsePaginationSliderDeleteAt(res *pb_slider.ApiResponsePaginationSliderDeleteAt) *model.APIResponsePaginationSliderDeleteAt
	ToGraphqlResponsePaginationSlider(res *pb_slider.ApiResponsePaginationSlider) *model.APIResponsePaginationSlider
}
