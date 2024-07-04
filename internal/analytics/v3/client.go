package v1

import (
	"context"
	"encoding/json"
	"fmt"
	httpcon "github.com/lista-dao/AuctionBots-go/pkg/httpconn"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"net/http"
	"net/url"
)

type Client struct {
	cn  *httpcon.Connector
	log *logrus.Logger
}

func NewClient(
	base *url.URL,
	log *logrus.Logger,
) *Client {
	cli := &Client{
		cn:  httpcon.NewConnector(base, nil),
		log: log,
	}
	cli.log.Debug("V3 New Client ", base.Host)
	return cli
}

func (cli *Client) GetRedUsers(ctx context.Context) ([]User, error) {
	cli.log.Debug("V3 GetRedUsers  Starting...")
	code, body, err := cli.cn.Get(ctx, "/api/v2/liquidations/red", nil)
	if err != nil {
		return nil, err
	}

	if *code == http.StatusOK {
		var resp CommonDataResp

		if err := json.Unmarshal(body, &resp); err != nil {
			return nil, errors.Wrapf(err, "failed to unmarshal body %s", string(body))
		}

		return resp.Data.Users, nil
	}

	return nil, errors.New(fmt.Sprintf("failed with %s", http.StatusText(*code)))
}
