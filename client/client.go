package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/someError/mservice-example/types"
)

type Client struct {
	baseEndpoint string
}

func NewClient(endpoint string) *Client {
	return &Client{
		baseEndpoint: endpoint,
	}
}

func (c *Client) FetchPrice(ctx context.Context, ticker string) (*types.PriceResponse, error) {
	endpoint := fmt.Sprintf("%s?ticker=%s", c.baseEndpoint, ticker)

	request, err := http.NewRequest("GET", endpoint, nil)

	if err != nil {
		return nil, err
	}

	responce, err := http.DefaultClient.Do(request)

	if err != nil {
		return nil, err
	}

	if responce.StatusCode != http.StatusOK {
		httpErr := map[string]any{}
		err := json.NewDecoder(responce.Body).Decode(&httpErr)

		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("ticker is not supported, %s", httpErr["error"])
	}

	priceResponce := &types.PriceResponse{}

	decoderError := json.NewDecoder(responce.Body).Decode(priceResponce)

	if decoderError != nil {
		return nil, decoderError
	}

	return priceResponce, nil

}
