package contract

import (
	"ev3-eth/chain"
	"log"

	"github.com/ethereum/go-ethereum/common"
)

func Deploy(config chain.Config) *DeviceController {
	ctx := chain.BindContext(config)
	auth := chain.Auth(ctx)

	// input := "1.0"
	address, _, instance, err := DeployDeviceController(auth, ctx.Client())
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Contract deployed at " + address.Hex())
	config.Address = address.Hex()
	chain.WriteConfig(&config)

	return instance
}

func Instance(config chain.Config, ctx chain.ConnectionContext) *DeviceController {
	address := common.HexToAddress(config.Address)
	instance, err := NewDeviceController(address, ctx.Client())
	if err != nil {
		log.Fatal(err)
	}
	return instance
}
