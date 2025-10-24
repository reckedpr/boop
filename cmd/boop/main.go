package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
)

// need to hit tha hay, head is soo broked right now

func ValidatePath(path string) (IsFile bool, IsDir bool) {
	fi, err := os.Stat(path)
	if err != nil {
		fmt.Println(err)
		IsFile = true
	}

	switch mode := fi.Mode(); {
	case mode.IsDir():
		IsDir = true
		return
	case mode.IsRegular():
		IsFile = true
		return
	}

	return
}

func main() {
	pathFlag := pflag.StringP("path", "p", ".", "file or dirrr to uhhhhh ya")
	pflag.Parse()

	r := gin.Default()

	isFile, isDir := ValidatePath(*pathFlag)

	if isDir {
		r.StaticFS("/", gin.Dir(*pathFlag, true))
	} else if isFile {
		r.GET("/", func(c *gin.Context) {
			c.File(*pathFlag)
		})

	} else {
		fmt.Println("yeah nah")
		os.Exit(1)
	}

	r.Run()
}
