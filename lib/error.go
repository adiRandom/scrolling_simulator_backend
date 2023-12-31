package lib

import (
	"fmt"
	"log"
)

type Error struct {
	Msg    string `json:"message"`
	Reason string `json:"reason"`
}

func (e Error) Error() string {
	return fmt.Sprintf("Error: %s, reason: %s", e.Msg, e.Reason)
}

func PanicOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
