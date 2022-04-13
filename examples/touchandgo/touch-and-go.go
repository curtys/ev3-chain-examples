package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/ev3go/ev3dev"
)

type MotorUpdate func(bool)

func main() {
	fmt.Println("Start")
	in1, err := ev3dev.SensorFor("ev3-ports:in1", "lego-ev3-touch")
	if err != nil {
		log.Fatalf("failed to find button on in1: %v", err)
	}
	outA, err := ev3dev.TachoMotorFor("ev3-ports:outA", "lego-ev3-l-motor")
	if err != nil {
		log.Fatalf("failed to find motor on outA: %v", err)
	}

	update := func(run bool) {
		state, _ := outA.State()
		if run && state == 0 {
			maxSpeed := outA.MaxSpeed()
			outA.SetSpeedSetpoint(-maxSpeed)
			outA.Command("run-forever")
		} else if !run {
			outA.Command("stop")
		}
	}

	ch := make(chan int, 1)
	signal := make(chan int, 1)
	go btnMonitor(in1, ch, signal, update)
	<-signal
}

func btnMonitor(btn *ev3dev.Sensor, ch chan int, signal chan int, update MotorUpdate) {
	for {
		select {
		case <-signal:
			return
		default:
			strVal, _ := btn.Value(0)
			intVal, _ := strconv.Atoi(strVal)
			update(intVal == 1)
		}
	}
}
