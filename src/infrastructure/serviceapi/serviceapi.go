package serviceapi

import (
	"collectionneur-cli/src/domain/entities"
	"collectionneur-cli/src/domain/interfaces"
	"context"
	"fmt"
	api "github.com/astoliarov/collectionneur_api"
	"github.com/twitchtv/twirp"
	"net/http"
	"time"
)

func spendInfoToModel(info *entities.SpendInfo) *api.SpendInfo {
	model := api.SpendInfo{
		Card:        info.Card,
		Currency:    info.Currency,
		Date:        uint64(info.Date.Unix()),
		Description: info.Description,
		Sum:         info.Sum,
		Category:    "",
	}
	return &model
}

func modelToSpendInfo(model *api.SpendInfo) *entities.SpendInfo {
	return &entities.SpendInfo{
		Card:        model.Card,
		Currency:    model.Currency,
		Date:        time.Unix(int64(model.Date), 0),
		Description: model.Description,
		Sum:         model.Sum,
		Category:    "",
	}
}

type API struct {
	token     string
	serverURL string

	client api.CostService
}

func (a *API) SendSpendInfos(infos []*entities.SpendInfo) (int, error) {

	models := []*api.SpendInfo{}
	for _, info := range infos {
		models = append(models, spendInfoToModel(info))
	}

	request := api.AddInfosRequest{
		Infos: models,
	}

	ctx, err := a.setHeader(context.Background())
	if err != nil {
		return 0, err
	}

	resp, err := a.client.AddInfos(ctx, &request)
	return int(resp.CountOfAdded), err
}

func (a *API) GetLastSpendInfo() (*entities.SpendInfo, error) {
	request := api.GetInfosRequest{
		Meta: &api.RequestMeta{
			Limit:  1,
			Offset: 0,
		},
		OrderBy: &api.OrderBy{
			Field:     api.OrderField_date,
			Direction: api.OrderDirection_desc,
		},
	}

	ctx, err := a.setHeader(context.Background())
	if err != nil {
		return nil, err
	}

	response, err := a.client.GetInfos(ctx, &request)
	if err != nil {
		return nil, err
	}

	if len(response.Infos) < 1 {
		return nil, interfaces.ErrApiNoInfo
	}

	return modelToSpendInfo(response.Infos[0]), nil
}

func (a *API) setHeader(ctx context.Context) (context.Context, error) {
	header := make(http.Header)
	header.Set("Authorization", fmt.Sprintf("Bearer %s", a.token))
	return twirp.WithHTTPRequestHeaders(ctx, header)
}

func NewAPI(token, serverURL string) (*API, error) {
	client := api.NewCostServiceJSONClient(serverURL, &http.Client{})

	a := API{
		token:     token,
		serverURL: serverURL,
		client:    client,
	}
	return &a, nil
}
