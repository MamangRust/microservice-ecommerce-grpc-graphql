package slidergraphqlmapper

import (
	"github.com/MamangRust/microservice-ecommerce-shared/pb"
	"github.com/MamangRust/monolith-graphql-ecommerce-apigateway/internal/model"
)

type SliderGraphqlMapper interface {
	ToGraphqlResponseSlider(res *pb.ApiResponseSlider) *model.APIResponseSlider
	ToGraphqlResponseSliderDeleteAt(res *pb.ApiResponseSliderDeleteAt) *model.APIResponseSliderDeleteAt
	ToGraphqlResponsesSlider(res *pb.ApiResponsesSlider) *model.APIResponsesSlider
	ToGraphqlResponseSliderDelete(res *pb.ApiResponseSliderDelete) *model.APIResponseSliderDelete
	ToGraphqlResponseSliderAll(res *pb.ApiResponseSliderAll) *model.APIResponseSliderAll
	ToGraphqlResponsePaginationSliderDeleteAt(res *pb.ApiResponsePaginationSliderDeleteAt) *model.APIResponsePaginationSliderDeleteAt
	ToGraphqlResponsePaginationSlider(res *pb.ApiResponsePaginationSlider) *model.APIResponsePaginationSlider
}
