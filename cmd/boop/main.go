package main

import (
	"fmt"
	"html/template"
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
		var body string

		if fullPath != absRoot {
			parentPath := path.Dir(relPath)

			body += fmt.Sprintf(`<a href="/%s">..</a><br>`, parentPath)
		}

		for _, e := range entries {
			name := e.Name()
			link := path.Join(relPath, name)
			class := ""
			if e.IsDir() {
				link += "/"
				name += "/"
				class += "dir"
			}
			body += fmt.Sprintf(`<a class="%s" href="/%s">%s</a><br>`, class, link, name)
		}

		c.HTML(http.StatusOK, "dirlist", gin.H{
			"Path": reqPath,
			"Body": template.HTML(body),
		})
	} else {
		c.File(fullPath)
	}
}

const html string = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>boop /{{.Path}}</title>
    <style>
        body{
            font-family: monospace;
			background-color: #303446;
			color: #c6d0f5;
			padding: 8px 8px;
        }
        .list{
            display: flex;
            flex-direction: column;
			font-size: 1.4rem;
        }
		a{
            color: #c6d0f5;
			text-decoration: none;
        }

		.dir { color: #a6d189; }
		.dir :visited{ color: #a6d189; }
    </style>
</head>
<body>
	<h1>/{{.Path}}</h1>
	
    <div class="list">
		{{.Body}}
    </div>
</body>
</html>
`

func main() {

	r := gin.Default()

	tpl := template.Must(template.New("dirlist").Parse(html))

	r.SetHTMLTemplate(tpl)

	r.GET("/*filepath", func(c *gin.Context) {
		DisplayDir(c)
	})

	r.Run()
}
