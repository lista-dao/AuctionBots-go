package jobs

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestGetFlashOrderMock(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "/order", r.URL.Path)
		require.Equal(t, "POST", r.Method)
		require.Equal(t, "test-key", r.Header.Get("LM-API-KEY"))

		resp := LiquidMeshResponse{
			Code: 0,
		}
		resp.Data.CallMsg = map[string]interface{}{
			"data": "0xabcdef",
		}
		resp.Data.OutputAmount = "12345"

		_ = json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	client := &LiquidMeshClient{
		APIKey:  "test-key",
		BaseURL: server.URL,
		Client:  server.Client(),
	}

	res, err := client.GetFlashOrder(context.Background(), FlashOrderParams{
		UserAddress: "0xUser",
		InputToken:  "0xIn",
		OutputToken: "0xOut",
		InputAmount: "100",
		SlippageBps: 300,
	})

	require.NoError(t, err)
	require.Equal(t, "12345", res["dstAmount"])
	require.Equal(t, "0xabcdef", res["data"])
}

func TestGetFlashOrderSuccess(t *testing.T) {
	apiKey := os.Getenv("LIQUIDMESH_API_KEY")
	if apiKey == "" {
		t.Skip("LIQUIDMESH_API_KEY not set, skip integration test")
	}

	client := &LiquidMeshClient{
		APIKey:  apiKey,
		BaseURL: "https://api.liquidmesh.io/v1/bsc",
		Client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
	res, err := client.GetFlashOrder(context.Background(), FlashOrderParams{
		UserAddress:     "0x38fA0963ad71D09B6CE01bf41B857078F9f8962b",
		InputToken:      "0x0782b6d8c4551b9760e74c0545a9bcd90bdc41e5",
		OutputToken:     "0x55d398326f99059ff775485246999027b3197955",
		InputAmount:     "100000000000000000",
		SlippageBps:     100,
		DisableSimulate: false,
	})

	require.NoError(t, err)
	t.Log(res)
}
