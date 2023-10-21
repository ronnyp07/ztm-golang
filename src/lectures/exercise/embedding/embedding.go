//--Summary:
//  Create a system monitoring dashboard using the existing dashboard
//  component structures. Each array element in the components represent
//  a 1-second sampling.
//
//--Requirements:
//* Create functions to calculate averages for each dashboard component
//* Using struct embedding, create a Dashboard structure that contains
//  each dashboard component
//* Print out a 5-second average from each component using promoted
//  methods on the Dashboard

package main

import (
	"fmt"
)

type Bytes int
type Celcius float32

type BandwidthUsage struct {
	amount []Bytes
}

type CpuTemp struct {
	temp []Celcius
}

type MemoryUsage struct {
	amount []Bytes
}

func (b *BandwidthUsage) AverageBandwidth() int {
	sum := 0
	for _, item := range b.amount {
		sum += int(item)
	}
	return sum / len(b.amount)
}

func (c *CpuTemp) AverageCpuTemp() int {
	sum := 0
	for _, item := range c.temp {
		sum += int(item)
	}
	return sum / len(c.temp)
}

func (m *MemoryUsage) AverageMemoryUsage() int {
	sum := 0
	for _, item := range m.amount {
		sum += int(item)
	}
	return sum / len(m.amount)
}

type Dashboard struct {
	BandwidthUsage
	CpuTemp
	MemoryUsage
}

func main() {
	bandwidth := BandwidthUsage{[]Bytes{50000, 100000, 130000, 80000, 90000}}
	temp := CpuTemp{[]Celcius{50, 51, 53, 51, 52}}
	memory := MemoryUsage{[]Bytes{800000, 800000, 810000, 820000, 800000}}

	dash := Dashboard{
		BandwidthUsage: bandwidth,
		CpuTemp:        temp,
		MemoryUsage:    memory,
	}

	fmt.Printf("Avarage bandwidth usage: %v \n", dash.AverageBandwidth())
	fmt.Printf("Avarage cpuTemp usage: %v \n", dash.AverageCpuTemp())
	fmt.Printf("Avarage memory usage: %v \n", dash.AverageMemoryUsage())
}
