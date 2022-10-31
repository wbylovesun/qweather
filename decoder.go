package qweather

import (
	"encoding/json"
	"fmt"
)

func decode[T responseType](respBytes []byte, r *T) error {
	err := json.Unmarshal(respBytes, r)
	if err != nil {
		return err
	}
	if !(*r).isOk() {
		return fmt.Errorf("invalid api code %s", (*r).GetCode())
	}
	return nil
}
