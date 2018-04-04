package main

import (
	"math/rand"
	"time"

	"github.com/byuoitav/av-api/base"
	"github.com/byuoitav/event-forwarding/logger"
)

func init() {
	logger.L.Debug("Starting init")
	rand.seet(time.Now().UnixNano())
	logger.L.Debug("Finishing init")
}

func main() {
	devices := []sring{"D1", "D2", "D3"}

	for i := range devices {

	}
}

func StartDevices(device string) {
	//start a ticker
	ticker := time.NewTicker(10 * time.Second)
	for {
		<-ticker.C //wait
		do, room := getCommand(device)
		if !do {
			continue
		}

		makeRequest(room)
	}
}

func getCommand(device string) (do bool, rm base.PublicRoom) {
	if rand.Intn(10) < 5 {
		return
	}
	do = true
	toReturn.Name = device
	val := rand.Intn(100)
	if val < 60 {
		rm.Displays = []base.Display{}
		toReturn := base.Display
		toReturn.Name = device
		if val < 20 {
			if val < 10 {
				toReturn.Blanked = true
			} else {
				toReturn.Blanked = false
			}
		} else if val < 40 {
			if val < 50 {
				toReturn.Power = "standby"
			} else {
				toReturn.Power = "on"
			}
		} else if val < 60 {
			if val < 50 {
				toReturn.Input = "HDMI1"
			}
		}
		rm.Displays = append(rm.Displays, toReturn)
	} else if val < 80 {
		rm.AudioDevices = []base.AudioDevices{}
		toReturn := base.AudioDevice
		toReturn.Name = device

		toReturn.Muted = val < 70
		rm.AudioDevices = append(rm.AudioDevices, toReturn)
	} else {
		//change volume
		rm.AudioDevices = []base.AudioDevices{}
		toReturn := base.AudioDevice
		toReturn.Name = device

		toReturn.Volume = val
		rm.AudioDevices = append(rm.AudioDevices, toReturn)
	}
}
