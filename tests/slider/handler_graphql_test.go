package slider_test

import (
	"net/http"
	"testing"

	tests "github.com/MamangRust/microservice-ecommerce-test"
	"github.com/stretchr/testify/suite"
)

type SliderGraphqlTestSuite struct {
	tests.BaseTestSuite
	handler  http.Handler
	sliderID int
}

func (s *SliderGraphqlTestSuite) SetupSuite() {
	s.BaseTestSuite.SetupSuite()
	s.SetupSliderService()

	s.handler = s.GraphQLHandler()
}

func (s *SliderGraphqlTestSuite) TestSliderGraphqlLifecycle() {
	// 1. Create
	create := s.GQLMultipart(s.handler, `mutation CreateSlider($input: CreateSliderRequest!) {
		createSlider(input: $input) { status message data { id name image } }
	}`, map[string]interface{}{
		"input": map[string]interface{}{
			"name":  "Test Slider",
			"image": nil,
		},
	}, []tests.UploadFile{
		{VariablePath: "input.image", Filename: "slider.jpg", Content: []byte("dummy image content")},
	})
	createData := s.Obj(s.Obj(create, "createSlider"), "data")
	s.sliderID = int(createData["id"].(float64))
	s.Equal("Test Slider", createData["name"])

	// 2. FindAll
	all := s.GQL(s.handler, `query FindAllSliders($input: FindAllSliderRequest!) {
		findAllSliders(input: $input) { status message pagination { total_records } data { id name } }
	}`, map[string]interface{}{"input": map[string]interface{}{"page": 1, "page_size": 10}})
	s.NotEmpty(s.Arr(s.Obj(all, "findAllSliders"), "data"))

	// 3. FindByActive
	active := s.GQL(s.handler, `query FindActiveSliders($input: FindAllSliderRequest!) {
		findActiveSliders(input: $input) { status message pagination { total_records } data { id name } }
	}`, map[string]interface{}{"input": map[string]interface{}{"page": 1, "page_size": 10}})
	s.NotEmpty(s.Arr(s.Obj(active, "findActiveSliders"), "data"))

	// 4. Update
	s.Require().NotZero(s.sliderID)
	updated := s.GQLMultipart(s.handler, `mutation UpdateSlider($input: UpdateSliderRequest!) {
		updateSlider(input: $input) { status message data { id name } }
	}`, map[string]interface{}{
		"input": map[string]interface{}{
			"id":    s.sliderID,
			"name":  "Updated Test Slider",
			"image": nil,
		},
	}, []tests.UploadFile{
		{VariablePath: "input.image", Filename: "slider_updated.jpg", Content: []byte("dummy image content")},
	})
	s.Equal("Updated Test Slider", s.Obj(s.Obj(updated, "updateSlider"), "data")["name"])

	// 5. Trash
	s.GQL(s.handler, `mutation TrashedSlider($input: FindByIdSliderRequest!) {
		trashedSlider(input: $input) { status message data { id name } }
	}`, map[string]interface{}{"input": map[string]interface{}{"id": s.sliderID}})

	// 6. FindByTrashed
	trashed := s.GQL(s.handler, `query FindTrashedSliders($input: FindAllSliderRequest!) {
		findTrashedSliders(input: $input) { status message pagination { total_records } data { id name } }
	}`, map[string]interface{}{"input": map[string]interface{}{"page": 1, "page_size": 10}})
	s.NotEmpty(s.Arr(s.Obj(trashed, "findTrashedSliders"), "data"))

	// 7. Restore
	s.GQL(s.handler, `mutation RestoreSlider($input: FindByIdSliderRequest!) {
		restoreSlider(input: $input) { status message data { id name } }
	}`, map[string]interface{}{"input": map[string]interface{}{"id": s.sliderID}})

	// 8. DeletePermanent (trash first so the row is eligible)
	s.GQL(s.handler, `mutation TrashedSlider($input: FindByIdSliderRequest!) {
		trashedSlider(input: $input) { status message data { id } }
	}`, map[string]interface{}{"input": map[string]interface{}{"id": s.sliderID}})
	s.GQL(s.handler, `mutation DeleteSliderPermanent($input: FindByIdSliderRequest!) {
		deleteSliderPermanent(input: $input) { status message }
	}`, map[string]interface{}{"input": map[string]interface{}{"id": s.sliderID}})

	// 9. RestoreAll
	s.GQL(s.handler, `mutation { restoreAllSliders { status message } }`, nil)

	// 10. DeleteAll
	s.GQL(s.handler, `mutation { deleteAllSlidersPermanent { status message } }`, nil)
}

func TestSliderGraphqlSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	suite.Run(t, new(SliderGraphqlTestSuite))
}
