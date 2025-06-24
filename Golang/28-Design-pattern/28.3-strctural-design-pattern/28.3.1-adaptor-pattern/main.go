package main

import "fmt"

/*

	this pattern is called also a wrapper  , it is used so that unrelated objects can work togather using
	adapter
	The things ythat joins the the unrelated object is called as Adapter
*/

// target
type mobile interface {
	chargeAppleMbile()
}

// prototype implementation

type Apple struct {
}

func (a *Apple) chargeAppleMbile() {
	fmt.Println("Charginig Apple Mobile")
}

// adaptee (changed functionality requrired by client )
type android struct {
}

func (a *android) chargeAndroidMbile() {
	fmt.Println("Charginig Android Mobile")
}

// Adapter
type androidAdapter struct {
	android *android
}

func (ad *androidAdapter) chargeAppleMbile() {
	ad.android.chargeAndroidMbile()
}

// client
type clinet struct{}

func (a *clinet) chargeMobile(mob mobile) {
	mob.chargeAppleMbile()
}

//Adapter

func main() {

	//First/Initial Requriement
	apple := &Apple{}
	clinet := &clinet{}
	clinet.chargeMobile(apple)

	//exteneded requirement

	android := &android{}
	androidAdapter := &androidAdapter{
		android: android,
	}

	clinet.chargeMobile(androidAdapter)
}
