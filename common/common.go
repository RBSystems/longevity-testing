package common

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/byuoitav/av-api/base"
	"github.com/byuoitav/event-forwarding/logger"
)

var addr = "http://ITB-1108B-CP1:8000/buildings/%s/rooms/%s"

func GetState(building, room string) error {
	_, err := http.Get(fmt.Sprintf(addr, building, room))
	if err != nil {
		msg := fmt.Sprintf("error making request: %v", err.Error())
		logger.L.Warn(msg)
		return errors.New(msg)
	}
	logger.L.Debugf("Done getting state")
	return nil
}

func SetState(building, room string, body base.PublicRoom) error {

	b, err := json.Marshal(body)
	if err != nil {
		msg := fmt.Sprintf(
			"error marshalling room: %v", err.Error())
		logger.L.Warn(msg)
		return errors.New(msg)
	}
	_, err = http.Post(
		fmt.Sprintf(addr, building, room),
		"application/json",
		bytes.NewBuffer(b),
	)

	if err != nil {
		msg := fmt.Sprintf("error making request: %v", err.Error())
		logger.L.Warn(msg)
		return errors.New(msg)
	}
	logger.L.Debugf("Done setting state")
	return nil
}
