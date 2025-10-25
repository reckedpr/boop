package main

import (
	"fmt"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

func DisplayDir(c *gin.Context) {
	reqPath := strings.TrimPrefix(c.Param("filepath"), "/")
	relPath := path.Clean(reqPath)
	absRoot, _ := filepath.Abs(".")

	fullPath := filepath.Join(absRoot, relPath)

	fi, err := os.Stat(fullPath)
	if err != nil {
		c.String(http.StatusNotFound, "file or direk aint exist ngl")
	}

	if fi.IsDir() {
		entries, _ := os.ReadDir(fullPath)
		html := "<html><body>"

		if fullPath != absRoot {
			parentPath := path.Dir(relPath)

			html += fmt.Sprintf(`<a href="/%s">..</a><br>`, parentPath)
		}

		for _, e := range entries {
			name := e.Name()
			link := path.Join(relPath, name)
			if e.IsDir() {
				link += "/"
			}
			html += fmt.Sprintf(`<a href="/%s">%s</a><br>`, link, name)
		}

		html += "</body></html>"
		c.Data(200, "text/html; charset=utf-8", []byte(html))
	} else {
		c.File(fullPath)
	}
}

func main() {
	r := gin.Default()

	r.GET("/*filepath", func(c *gin.Context) {
		DisplayDir(c)
	})

	r.Run()
}
