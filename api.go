package main

import (
	"os"
	"fmt"
	"github.com/AlienVault-OTX/OTX-Go-SDK/src/otxapi"
)

const API_KEY = "324e82368c01483f3e9e010e3adb9ed5620a9a388f0c0656853be09546e70690"

func apiInit() {
	os.Setenv("X_OTX_API_KEY", API_KEY)
	fmt.Println("Set environment key X_OTX_API_KEY to " + API_KEY)
}

func toString(str interface{}) string {
	return otxapi.Stringify(str)
}

func initClient() *otxapi.Client {
	client := otxapi.NewClient(nil)
	userDetail, _, err := client.UserDetail.Get()

	if err != nil {
		fmt.Println("API key invalid")
		os.Exit(1)
	}
	fmt.Println("Login as " + toString(userDetail.Username))

	return client
}

func getPulse(client *otxapi.Client) {
	pulseDetail, _, err := client.PulseDetail.Get("5aec1ab38170f445dbd22f2b")

	if err != nil {
		fmt.Println("Invalid Pulse ID")
		return
	}

	indicator := pulseDetail.Indicators

	for i, v := range indicator {
		fmt.Printf("Indicator no %d : %s\n", i+1, toString(v))
	}

}

func getList(client *otxapi.Client) {
	// pulseList, _, err := client.
}

func main() {
	apiInit()
	client := initClient()
	getPulse(client)

}
