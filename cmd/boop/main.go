package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
)

type PathType struct {
	IsFile bool
	IsDir  bool
}

func ValidatePath(path string) PathType {
	t := PathType{}

	fi, err := os.Stat(path)
	if err != nil {
		fmt.Println(err)
		return t
	}

	switch mode := fi.Mode(); {
	case mode.IsDir():
		t.IsDir = true
		return t
	case mode.IsRegular():
		t.IsFile = true
		return t
	}

	return t
}

func main() {
	pathFlag := pflag.StringP("path", "p", "", "file or dirrr to uhhhhh ya")
	pflag.Parse()

	r := gin.Default()

	test := ValidatePath(*pathFlag)

	if test.IsDir {
		r.StaticFS("/", gin.Dir(*pathFlag, true))
	} else if test.IsFile {
		r.GET("/", func(c *gin.Context) {
			c.File(*pathFlag)
		})

	} else {
		fmt.Println("NEY")
		os.Exit(1)
	}

	r.Run()
}
