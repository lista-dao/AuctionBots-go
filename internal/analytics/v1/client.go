package v1

import (
	"context"
	"encoding/json"
	"fmt"
	httpcon "github.com/helio-money/auctionbot/pkg/httpconn"
	"github.com/pkg/errors"
	"net/http"
	"net/url"
)

type Client struct {
	cn *httpcon.Connector
}

func NewClient(
	base *url.URL,
) *Client {
	cli := &Client{
		cn: httpcon.NewConnector(base, nil),
	}

	return cli
}

func (cli *Client) GetRedUsers(ctx context.Context) ([]User, error) {
	code, body, err := cli.cn.Get(ctx, "/liquidations/red", nil)
	if err != nil {
		return nil, err
	}

	if *code == http.StatusOK {
		var resp LiquidationUsersResp

		if err := json.Unmarshal(body, &resp); err != nil {
			return nil, errors.Wrapf(err, "failed to unmarshal body %s", string(body))
		}

		return resp.Users, nil
	}

	return nil, errors.New(fmt.Sprintf("failed with %s", http.StatusText(*code)))
}

func (cli *Client) GetOrangeUsers(ctx context.Context) ([]User, error) {
	code, body, err := cli.cn.Get(ctx, "/liquidations/orange", nil)
	if err != nil {
		return nil, err
	}

	if *code == http.StatusOK {
		var resp LiquidationUsersResp

		if err := json.Unmarshal(body, &resp); err != nil {
			return nil, errors.Wrapf(err, "failed to unmarshal body %s", string(body))
		}

		return resp.Users, nil
	}

	return nil, errors.New(fmt.Sprintf("failed with %s", http.StatusText(*code)))
}
