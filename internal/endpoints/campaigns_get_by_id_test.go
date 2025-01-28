package endpoints

import (
	"emailn/internal/contract"
	internalmock "emailn/internal/test/mock"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_CampaignsGetById_should_return_campaign(t *testing.T) {
	assert := assert.New(t)
	campaign := contract.CampaignResponse{
		ID:      "343",
		Name:    "Test",
		Content: "HI",
		Status:  "Pending",
	}
	service := new(internalmock.CampaignServiceMock)
	service.On(("GetBy"), mock.Anything).Return(&campaign, nil)
	handler := Handler{CampaignService: service}

	req, _ := http.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()

	response, status, _ := handler.CampaignGetById(rr, req)

	assert.Equal(200, status)
	assert.Equal(campaign.ID, response.(*contract.CampaignResponse).ID)
	assert.Equal(campaign.Name, response.(*contract.CampaignResponse).Name)
}
