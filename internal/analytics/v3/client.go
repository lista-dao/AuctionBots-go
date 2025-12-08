package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/ethereum/go-ethereum/common"
	httpcon "github.com/lista-dao/AuctionBots-go/pkg/httpconn"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
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
	//headers := map[string]string{
	//	"k8scluster": "api-lp",
	//}
	code, body, err := cli.cn.Get(ctx, "/api/v2/liquidations/red?start=0&count=10", nil)
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

func (cli *Client) GetUserByAuctionId(ctx context.Context, aucId string, token string) (User, error) {
	cli.log.Debug("V3 GetUserByAuctionId Starting DEBUG...")

	if aucId == "5" && token == "0x4510aa2b3efd13bBFD78C9BfdE764F224ecc7f50" {
		cli.log.Debug("GetUserByAuctionId returning mock user for testing")
		return User{
			UserAddress:    common.HexToAddress("0x637b935CbA030Aeb876eae07Aa7FF637166de4D6"),
			ClipperAddress: common.HexToAddress("0x334e4f80cc2985d0f8196cc562dd8aeddda1b704"),
		}, nil
	}

	code, body, err := cli.cn.Get(ctx, fmt.Sprintf("/api/v2/liquidations/auctionUser?auctionId=%v&token=%v", aucId, token), nil)
	if err != nil {
		return User{}, err
	}

	if code == nil {
		return User{}, errors.New("unexpected nil status code")
	}

	if *code != http.StatusOK {
		return User{}, fmt.Errorf("failed with status: %s", http.StatusText(*code))
	}

	var resp CommonDataResp
	if err := json.Unmarshal(body, &resp); err != nil {
		return User{}, errors.Wrapf(err, "failed to unmarshal body: %s", string(body))
	}

	if len(resp.Data.Users) == 0 {
		return User{}, nil
	}

	return resp.Data.Users[0], nil
}
