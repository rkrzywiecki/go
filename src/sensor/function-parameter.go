package main

import (
	"fmt"
	"math/rand"
	"time"
)

type kelvin2 float64

func measureTemperature(samples int, sensor func() kelvin2) {
	for i := 0; i < samples; i++ {
		k := sensor()
		fmt.Printf("%v C K\n", k)
		time.Sleep(time.Second)
	}

}

func fakeSensor2() kelvin2 {
	return kelvin2(rand.Intn(151) + 150)
}

func main() {
	measureTemperature(3, fakeSensor2)

}
