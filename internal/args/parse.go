package args

import (
	"fmt"
	"io"
	"os"

	"github.com/spf13/pflag"
)

func ParsePath() (servePath string) {
	pathFlag := pflag.StringP("path", "p", "", "file or dirrr to uhhhhh ya")
	pflag.Parse()

	if *pathFlag != "" {
		servePath = *pathFlag
	} else if len(pflag.Args()) > 0 {
		servePath = pflag.Args()[0]
	} else {
		servePath = "."
	}

	return servePath
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
