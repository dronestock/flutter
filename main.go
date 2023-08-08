package main

import (
	"github.com/dronestock/drone"
	"github.com/dronestock/flutter/internal"
)

func main() {
	drone.New(internal.NewPlugin).Boot()
}
