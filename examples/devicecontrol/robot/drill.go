package robot

import (
	"ev3-eth/platform"
	"fmt"
	"strconv"

	"github.com/ev3go/ev3dev"
)

type MotorUpdate func(bool)

type Event uint
type EventListener func(*Drill)
type ControllSignal uint

type Drill struct {
	sensor    *ev3dev.Sensor
	motor     *ev3dev.TachoMotor
	listeners map[Event][]EventListener
	State     ControllSignal
}

const (
	Start Event = iota
	Stop
)

const (
	Started ControllSignal = iota
	Stopped
)

func DrillFor(sensorPort platform.Port, motorPort platform.Port) (*Drill, error) {
	in, err := ev3dev.SensorFor(sensorPort.String(), platform.EV3TouchSensor.String())
	if err != nil {
		return nil, err
	}
	out, err := ev3dev.TachoMotorFor(motorPort.String(), platform.EV3MServo.String())
	if err != nil {
		return nil, err
	}
	var d Drill
	d.motor = out
	d.sensor = in
	d.listeners = make(map[Event][]EventListener)
	d.State = Stopped

	return &d, nil
}

func (d *Drill) AttachListener(event Event, listener EventListener) {
	d.listeners[event] = append(d.listeners[event], listener)
}

func (d *Drill) Start() {
	if d.State == Started {
		return
	}
	fmt.Println("drill > starting drill")
	d.State = Started

	update := func(run bool) {
		state, _ := d.motor.State()
		if run && state == 0 {
			maxSpeed := d.motor.MaxSpeed()
			d.motor.SetSpeedSetpoint(-maxSpeed)
			d.motor.Command("run-forever")
			d.dispatch(Start)
		} else if !run {
			d.motor.Command("stop")
			d.dispatch(Stop)
		}
	}

	go d.monitorSensor(update)
}

func (d *Drill) Stop() {
	fmt.Println("drill > stopping drill")
	d.motor.Command("stop")
	d.State = Stopped
	d.dispatch(Stop)
}

func (d *Drill) monitorSensor(update MotorUpdate) {
	for d.State == Started {
		strVal, _ := d.sensor.Value(0)
		intVal, _ := strconv.Atoi(strVal)
		update(intVal == 1)
	}
}

func (d *Drill) dispatch(event Event) {
	for _, l := range d.listeners[event] {
		go l(d)
	}
}
