package main

import (
    "time"

    "gobot.io/x/gobot"
    "gobot.io/x/gobot/platforms/sphero"
)

var spheroPort string = "/dev/tty.Sphero-RRY-AMP-SPP"

func main() {
    adaptor := sphero.NewAdaptor(spheroPort)
    driver := sphero.NewSpheroDriver(adaptor)

    work := func() {
        gobot.Every(3*time.Second, func() {
            driver.Roll(30, uint16(gobot.Rand(360)))
        })
    }

    robot := gobot.NewRobot("sphero",
        []gobot.Connection{adaptor},
        []gobot.Device{driver},
        work,
    )

    robot.Start()
}