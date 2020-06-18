package config

import (
	"fmt"

	seelog "github.com/cihub/seelog"
)

var Logger seelog.LoggerInterface

func loadAppConfig() {
	var max string = "10"
	appConfig := `
<seelog minlevel="warn">
    <outputs formatid="Info">
        <rollingfile type="size" filename="./seelog/Transaction.log" maxsize="1024" maxrolls= "` + max + `"/>
    </outputs>
    <formats>
        <format id="Info" format="%Msg%n" />
    </formats>
</seelog>
`
	logger, err := seelog.LoggerFromConfigAsBytes([]byte(appConfig))
	if err != nil {
		fmt.Println(err)
		return
	}
	UseLogger(logger)
}

func init() {
	DisableLog()
	loadAppConfig()
}

// DisableLog disables all library log output
func DisableLog() {
	Logger = seelog.Disabled
}

// UseLogger uses a specified seelog.LoggerInterface to output library log.
// Use this func if you are using Seelog logging system in your app.
func UseLogger(newLogger seelog.LoggerInterface) {
	Logger = newLogger
}
