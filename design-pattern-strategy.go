package main

import "fmt"

type FlyBehavior interface {
	fly()
}

type FlyWithWings struct {
}

func (f *FlyWithWings) fly() {
	fmt.Println("FlyWithWings.fly()")
}

type FlyNoWay struct {
}

func (f *FlyNoWay) fly() {
	fmt.Println("FlyNoWay.fly()")
}

type Duck struct {
	flyBehavior FlyBehavior
}

func (d *Duck) performFly() {
	d.flyBehavior.fly()
}

func testStrategy() {
	duckWithWings := &Duck{
		flyBehavior: &FlyWithWings{},
	}
	duckWithWings.performFly()

	toyDuck := Duck{
		flyBehavior: &FlyNoWay{},
	}
	toyDuck.performFly()
}
