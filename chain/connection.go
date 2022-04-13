package chain

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type ConnectionContext struct {
	privateKey  *ecdsa.PrivateKey
	client      *ethclient.Client
	fromAddress common.Address
}

func (ctx *ConnectionContext) Client() *ethclient.Client {
	return ctx.client
}

func BindContext(config Config) ConnectionContext {
	privateKey, err := crypto.HexToECDSA(config.Secret)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	client, err := ethclient.Dial(config.RpcEndpoint)

	if err != nil {
		log.Fatalf("Failed to connect to chain: %v", err)
	} else {
		fmt.Println("Success! you are connected to the Ethereum Network")
	}

	return ConnectionContext{privateKey: privateKey, client: client, fromAddress: fromAddress}
}

func Auth(connCtx ConnectionContext) *bind.TransactOpts {

	nonce, err := connCtx.client.PendingNonceAt(context.Background(), connCtx.fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := connCtx.client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(connCtx.privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = gasPrice

	return auth
}

func Ops(connCtx ConnectionContext) *bind.CallOpts {
	blockNum, err := connCtx.client.BlockNumber(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	ops := new(bind.CallOpts)
	ops.BlockNumber = big.NewInt(0).SetUint64(blockNum)
	ops.From = connCtx.fromAddress
	ops.Context = context.Background()
	return ops
}
