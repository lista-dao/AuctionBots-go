package wallet

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"math/big"
)

type Walleter interface {
	Opts(ctx context.Context) (*bind.TransactOpts, error)
	Address() common.Address
}

func NewWallet(ethCli *ethclient.Client, key *ecdsa.PrivateKey) Walleter {
	publicKeyECDSA, ok := key.Public().(*ecdsa.PublicKey)
	if !ok {
		panic("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	chainID, err := ethCli.ChainID(context.Background())
	if err != nil {
		panic("failed to get chain id")
	}

	auth, err := bind.NewKeyedTransactorWithChainID(key, chainID)
	if err != nil {
		panic("err failed to create keyed transactor with chainID")
	}

	return wallet{
		privKey: key,
		ethCli:  ethCli,
		address: crypto.PubkeyToAddress(*publicKeyECDSA),
		auth:    auth,
	}
}

type wallet struct {
	privKey *ecdsa.PrivateKey
	ethCli  *ethclient.Client
	auth    *bind.TransactOpts
	address common.Address
}

func (w wallet) Opts(ctx context.Context) (*bind.TransactOpts, error) {
	auth := *w.auth
	nonce, err := w.ethCli.PendingNonceAt(ctx, w.address)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get nonce")
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)       // in wei
	auth.GasLimit = uint64(10000000) // in units
	auth.Context = ctx

	gasPrice, err := w.ethCli.SuggestGasPrice(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get gas price")
	}

	auth.GasPrice = gasPrice

	return &auth, nil
}

func (w wallet) Address() common.Address {
	return w.address
}
