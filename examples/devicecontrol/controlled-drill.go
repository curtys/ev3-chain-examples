package main

import (
	"ev3-eth/chain"
	"ev3-eth/examples/devicecontrol/contract"
	"ev3-eth/examples/devicecontrol/robot"
	"ev3-eth/platform"

	"flag"
	"log"
	"os"
)

func main() {
	configPtr := flag.String("config", "config.toml", "Path to program configuration file")
	deployPtr := flag.Bool("deploy", false, "Deploy smart contract")
	flag.Parse()

	config := chain.ReadConfig(*configPtr)
	if *deployPtr {
		log.Default().Println("Deploying contract...")
		contract.Deploy(config)
		os.Exit(1)
	}

	id := platform.Id()
	ctx := chain.BindContext(config)
	controller := contract.Instance(config, ctx)

	log.Default().Printf("Registering device %s", id)
	tx, err := controller.Register(chain.Auth(ctx), id, "drill")
	if err != nil {
		log.Printf("failed to register device in contract: %v", err)
	} else {
		log.Default().Printf("Device %s registered. %s", id, tx.Hash())
	}

	drill, err := robot.DrillFor(platform.EV3In1, platform.EV3OutA)
	if err != nil {
		log.Fatalf("failed to initialize drill: %v", err)
	}
	stop := make(chan int, 1)
	checkReadyState := func(d *robot.Drill) {
		rs, err := controller.GetReadyState(chain.Ops(ctx), id)
		if err != nil {
			log.Fatalf("failed to check ready state: %v", err)
			stop <- 1
		} else if !rs {
			log.Println("device not ready")
			stop <- 1
		}
	}

	// metricCh := make(chan int, 10)

	drill.AttachListener(robot.Start, checkReadyState)
	go drill.Start()
	// go drill.CollectMetric(metricCh, stop)
	// go monitorMetric(metricCh, stop, id, ctx, controller)
	<-stop
}

func monitorMetric(ch chan int, stop chan int, id string, ctx chain.ConnectionContext, controller *contract.DeviceController) {
	for {
		select {
		case <-stop:
			return
		default:
			val := <-ch
			_, err := controller.UpdateMetric(chain.Auth(ctx), id, uint32(val))
			if err != nil {
				log.Fatalf("Could not store metric in contract: %v", err)
				stop <- 1
			}
		}
	}
}
