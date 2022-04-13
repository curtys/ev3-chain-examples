package platform

import (
	"log"
	"net"
)

type Port string
type Driver string

// in/out ports
const (
	EV3OutA Port = "ev3-ports:outA"
	EV3OutB Port = "ev3-ports:outB"
	EV3OutC Port = "ev3-ports:outC"
	EV3OutD Port = "ev3-ports:outD"
	EV3In1  Port = "ev3-ports:in1"
	EV3Int2 Port = "ev3-ports:in2"
	EV3In3  Port = "ev3-ports:in3"
	EV3Int4 Port = "ev3-ports:in4"
)

// drivers
const (
	EV3LServo      Driver = "lego-ev3-l-motor"
	EV3MServo      Driver = "lego-ev3-m-motor"
	EV3TouchSensor Driver = "lego-ev3-touch"
)

func (p Port) String() string {
	return string(p)
}

func (d Driver) String() string {
	return string(d)
}

func Id() string {
	as, err := getMacAddr()
	if err != nil {
		log.Fatal(err)
	}
	return as[0]
}

func getMacAddr() ([]string, error) {
	ifas, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	var as []string
	for _, ifa := range ifas {
		a := ifa.HardwareAddr.String()
		if a != "" {
			as = append(as, a)
		}
	}
	return as, nil
}
