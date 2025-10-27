package cli

import (
	"fmt"
	"os"
)

// i hate everything in this file, SO tacky but we ball

const (
	Reset = "\033[0m"
	Bold  = "\033[1m"

	FgBlack   = "\033[30m"
	FgRed     = "\033[31m"
	FgGreen   = "\033[32m"
	FgYellow  = "\033[33m"
	FgBlue    = "\033[34m"
	FgMagenta = "\033[35m"
	FgCyan    = "\033[36m"
	FgWhite   = "\033[37m"

	BgBlack   = "\033[40m"
	BgRed     = "\033[41m"
	BgGreen   = "\033[42m"
	BgYellow  = "\033[43m"
	BgBlue    = "\033[44m"
	BgMagenta = "\033[45m"
	BgCyan    = "\033[46m"
	BgWhite   = "\033[47m"

	BgHiBlack   = "\033[100m"
	BgHiRed     = "\033[101m"
	BgHiGreen   = "\033[102m"
	BgHiYellow  = "\033[103m"
	BgHiBlue    = "\033[104m"
	BgHiMagenta = "\033[105m"
	BgHiCyan    = "\033[106m"
	BgHiWhite   = "\033[107m"
)

var (
	Verbose     = false
	BoopPrefix  = fmt.Sprintf("%s%s boop %s", BgHiMagenta, FgBlack, Reset)
	ErrorPrefix = fmt.Sprintf("%s%s  err %s", BgHiRed, FgBlack, Reset)
	InfoPrefix  = fmt.Sprintf("%s%s info %s", BgHiYellow, FgBlack, Reset)
)

func BoopLog(text string, a ...any) {
	fmt.Printf("%s %s\n", BoopPrefix, fmt.Sprintf(text, a...))
}

func BoopLogNl(text string, a ...any) {
	fmt.Printf("\n%s %s\n", BoopPrefix, fmt.Sprintf(text, a...))
}

func BoopErr(text string) {
	fmt.Printf("%s %s\n", ErrorPrefix, text)
	os.Exit(1)
}

func BoopInfo(text string) {
	if Verbose {
		fmt.Printf("%s %s\n", InfoPrefix, text)
	}
}

func BoopHttp(code int, text string) {
	if !Verbose {
		return
	}

	color := StatusToColor(code)
	status := fmt.Sprintf("%s %d %s", color, code, Reset)

	fmt.Printf("%s %s %s\n", BoopPrefix, status, text)
}

func StatusToColor(code int) string {
	switch {
	case code >= 200 && code < 300:
		return FgBlack + BgHiGreen
	case code >= 300 && code < 400:
		return FgBlack + BgHiYellow
	case code >= 400:
		return FgBlack + BgHiRed
	default:
		return FgBlack + BgHiCyan
	}
}

func Colorise(text string, attrs ...string) string {
	out := ""
	for _, a := range attrs {
		out += a
	}
	return fmt.Sprintf("%s%s%s", out, text, Reset)
}
