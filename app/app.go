package app

import (
	_ "encoding/json"
)

type app struct {

}

func New() app {
	return app{}
}