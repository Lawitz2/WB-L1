package main

import "fmt"

type mobile interface {
	chargeAppleDevice()
}

type appleDevice struct{}

func (a *appleDevice) chargeAppleDevice() {
	fmt.Println("charging apple device")
}

type client struct{}

func (c *client) chargeMobile(mob mobile) {
	mob.chargeAppleDevice()
}

// expanding functionality of existing code

type androidDevice struct{}

func (a *androidDevice) chargeAndroidDevice() {
	fmt.Println("charging android device")
}

type androidAdapter struct {
	android *androidDevice
}

func (andAd *androidAdapter) chargeAppleDevice() {
	andAd.android.chargeAndroidDevice()
}

func Ex21_1() {
	apple := &appleDevice{}
	cl := &client{}
	cl.chargeMobile(apple)

	andPhone := &androidDevice{}
	androidAd := &androidAdapter{android: andPhone}
	cl.chargeMobile(androidAd)
}
