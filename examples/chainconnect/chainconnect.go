package main

import (
	"context"
	"ev3-eth/chain"
	"flag"
	"log"
)

func main() {
	configPtr := flag.String("config", "config.toml", "Path to program configuration file")
	flag.Parse()

	config := chain.ReadConfig(*configPtr)
	ctx := chain.BindContext(config)
	num, err := ctx.Client().BlockNumber(context.Background())
	if err != nil {
		log.Fatalf("Could not get block number: %v\n", err)
	} else {
		log.Println(num)
	}
}
