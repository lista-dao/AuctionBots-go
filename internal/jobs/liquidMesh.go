package jobs

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

type FlashOrderParams struct {
	UserAddress     string `json:"userAddress"`
	SlippageBps     int64  `json:"slippageBps"`
	InputToken      string `json:"inputToken"`
	OutputToken     string `json:"outputToken"`
	InputAmount     string `json:"inputAmount"`
	DisableSimulate bool   `json:"disableSimulate,omitempty"`
	ChainID         string `json:"chainId,omitempty"`
}

type LiquidMeshResponse struct {
	Code int `json:"code"`
	Data struct {
		CallMsg      map[string]interface{} `json:"callMsg"`
		OutputAmount string                 `json:"outputAmount"`
	} `json:"data"`
}

type LiquidMeshClient struct {
	APIKey  string
	BaseURL string
	Client  *http.Client
}

func (c *LiquidMeshClient) GetFlashOrder(
	ctx context.Context,
	params FlashOrderParams,
) (map[string]interface{}, error) {

	if params.ChainID == "" {
		params.ChainID = "56"
	}

	url := c.BaseURL + "/order"

	headers := map[string]string{
		"Content-Type": "application/json",
		"LM-API-KEY":   c.APIKey,
	}

	body, _ := json.Marshal(params)

	const (
		Timeout = 5 * time.Second
		Retry   = 2
	)

	var err error

	for attempt := 1; attempt <= Retry+1; attempt++ {

		var res LiquidMeshResponse

		reqCtx, cancel := context.WithTimeout(ctx, Timeout)

		req, reqErr := http.NewRequestWithContext(
			reqCtx,
			http.MethodPost,
			url,
			bytes.NewReader(body),
		)
		if reqErr != nil {
			cancel()
			return nil, reqErr
		}

		for k, v := range headers {
			req.Header.Set(k, v)
		}

		resp, reqErr := c.Client.Do(req)
		if reqErr != nil {
			err = reqErr
			cancel()
			goto RETRY
		}

		if resp.StatusCode != http.StatusOK {
			b, _ := io.ReadAll(resp.Body)
			resp.Body.Close()
			err = fmt.Errorf("status %d: %s", resp.StatusCode, string(b))
			cancel()
			goto RETRY
		}

		if err = json.NewDecoder(resp.Body).Decode(&res); err != nil {
			resp.Body.Close()
			cancel()
			goto RETRY
		}

		resp.Body.Close()
		cancel()

		if res.Code != 0 {
			err = fmt.Errorf("api code %d", res.Code)
			goto RETRY
		}

		res.Data.CallMsg["dstAmount"] = res.Data.OutputAmount
		return res.Data.CallMsg, nil

	RETRY:
		if attempt > Retry {
			return nil, err
		}
		time.Sleep(500 * time.Millisecond)
	}

	return nil, errors.New("[LiquidMesh] unreachable")
}
