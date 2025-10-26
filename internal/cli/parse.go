package cli

import (
	"fmt"
	"io"
	"os"

	"github.com/spf13/pflag"
)

type cliArgs struct {
	Path string
	Port int
	Time int
}

func ParseArgs() (argObj cliArgs) {
	portFlag := pflag.IntP("port", "p", 8080, "port to serve")
	timeFlag := pflag.IntP("time", "t", 0, "time in minutes to serve for")

	pflag.Parse()
	args := pflag.Args()

	argObj.Path = "."
	if len(args) > 0 {
		argObj.Path = args[0]
	}

	argObj.Port = *portFlag
	argObj.Time = *timeFlag

	return argObj
}

func ReadStdin() (isPiped bool, data string) {
	fi, _ := os.Stdin.Stat()

	if fi.Mode()&os.ModeCharDevice == 0 {
		data, err := io.ReadAll(os.Stdin)
		if err != nil {
			fmt.Println("error reading from stdin : ", err)
		}

		return true, string(data)

	} else {
		return false, ""
	}
}
