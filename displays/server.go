package main

import (
	"math/rand"
	"sync"
	"time"

	"github.com/byuoitav/av-api/base"
	"github.com/byuoitav/event-forwarding/logger"
	"github.com/byuoitav/longevity-testing/common"
)

var building = "ITB"
var room = "1108B"

func init() {
	logger.L.Debug("Starting init")
	rand.Seed(time.Now().UnixNano())
	logger.L.Debug("Finishing init")
}

func main() {
	devices := []string{"D1", "D2", "D3"}

	//don't stop! Believin!
	wg := sync.WaitGroup{}
	wg.Add(1)

	for i := range devices {
		go StartDevice(devices[i])
	}

	//I want to cry right now, too.
	wg.Wait()

}

func StartDevice(device string) {
	//start a ticker
	ticker := time.NewTicker(10 * time.Second)
	for {
		<-ticker.C //wait
		logger.L.Debugf("%v Starting.", device)

		//check if we're gonna get state or set it - we have a 40% chance of getting status vs setting
		if rand.Intn(10) < 4 {
			common.GetState(building, room)
			continue
		}
		logger.L.Debugf("%v Trying for a set state", device)

		do, body := getCommand(device)
		if !do {
			continue
			logger.L.Debugf("%v not this time", device)
		}

		logger.L.Debugf("%v setting state", device)
		common.SetState(building, room, body)
	}
}

func getCommand(device string) (do bool, rm base.PublicRoom) {
	//50% chance of actually setting state
	if rand.Intn(10) < 5 {
		return
	}
	action := ""
	do = true
	val := rand.Intn(100)
	if val < 70 {

		rm.Displays = []base.Display{}
		toReturn := base.Display{}
		toReturn.Name = device

		if val < 20 {
			if val < 10 {
				toReturn.Blanked = getBoolPointer(true)
				action = "blank"
			} else {
				toReturn.Blanked = getBoolPointer(false)
				action = "unblank"
			}
		} else if val < 60 {
			if val < 30 {
				toReturn.Power = "standby"
				action = "standby"
			} else {
				toReturn.Power = "on"
				action = "on"
			}
		} else if val < 70 {
			toReturn.Input = "HDMI1"
			action = "input"
		}
		rm.Displays = append(rm.Displays, toReturn)
	} else if val < 85 {
		rm.AudioDevices = []base.AudioDevice{}
		toReturn := base.AudioDevice{}
		toReturn.Name = device

		toReturn.Muted = getBoolPointer(val < 70)
		rm.AudioDevices = append(rm.AudioDevices, toReturn)
		action = "mute"
	} else {
		//change volume
		rm.AudioDevices = []base.AudioDevice{}
		toReturn := base.AudioDevice{}
		toReturn.Name = device

		a := val
		toReturn.Volume = &a
		rm.AudioDevices = append(rm.AudioDevices, toReturn)
		action = "volume"
	}
	logger.L.Debugf("%v Num was %v, so action is %v", device, val, action)
	return
}

//have i mentioned that I want to cry yet?
//we're doing this so omitempty will acutally omit, I assume?
func getBoolPointer(b bool) *bool {
	t := true
	f := false
	if b {
		return &t
	}
	return &f

}
