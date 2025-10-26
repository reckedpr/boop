package cli

import "fmt"

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

var BoopPrefix = fmt.Sprintf("%s%s boop %s", BgHiMagenta, FgBlack, Reset)

func BoopLog(format string, a ...any) {
	fmt.Printf("%s %s\n", BoopPrefix, fmt.Sprintf(format, a...))
}

func BoopLogNl(format string, a ...any) {
	fmt.Printf("\n%s %s\n", BoopPrefix, fmt.Sprintf(format, a...))
}

func Colorise(text string, attrs ...string) string {
	out := ""
	for _, a := range attrs {
		out += a
	}
	return fmt.Sprintf("%s%s%s", out, text, Reset)
}
