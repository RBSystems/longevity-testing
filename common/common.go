package common

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/byuoitav/av-api/base"
	"github.com/byuoitav/common/log"
)

var addr = "http://ITB-1108M-CP1:8000/buildings/%s/rooms/%s"

func GetState(building, room string) error {
	resp, err := http.Get(fmt.Sprintf(addr, building, room))
	if err != nil {
		msg := fmt.Sprintf("error making request: %v", err.Error())
		log.L.Warn(msg)
		return errors.New(msg)
	}
	defer resp.Body.Close()

	log.L.Debugf("Done getting state")
	return nil
}

func SetState(building, room string, body base.PublicRoom) error {
	b, err := json.Marshal(body)
	if err != nil {
		msg := fmt.Sprintf("error marshalling room: %v", err.Error())
		log.L.Warn(msg)
		return errors.New(msg)
	}

	resp, err := http.Post(fmt.Sprintf(addr, building, room), "application/json", bytes.NewBuffer(b))
	if err != nil {
		msg := fmt.Sprintf("error making request: %v", err.Error())
		log.L.Warn(msg)
		return errors.New(msg)
	}
	defer resp.Body.Close()

	log.L.Debugf("Done setting state")
	return nil
}
