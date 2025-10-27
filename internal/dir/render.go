package dir

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

const HtmlTemplate string = `
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
			font-size: 1.2rem;
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

type Path struct {
	Req  string
	Rel  string
	Root string
	Full string
}

func ResolvePath(c *gin.Context, root string) (*Path, error) {
	req := strings.TrimPrefix(c.Param("filepath"), "/")
	rel := path.Clean(req)
	absRoot, _ := filepath.Abs(root)
	full := filepath.Join(absRoot, rel)

	if !strings.HasPrefix(full, absRoot+string(os.PathSeparator)) && full != absRoot {
		return nil, fmt.Errorf("traversal")
	}

	return &Path{
		Req:  req,
		Rel:  rel,
		Root: absRoot,
		Full: full,
	}, nil
}

func HandlePath(c *gin.Context, servePath string) {
	p, err := ResolvePath(c, servePath)
	if err != nil {
		c.String(http.StatusNotFound, "file or dir not found")
		return
	}

	fi, err := os.Stat(p.Full)
	if err != nil {
		c.String(http.StatusNotFound, "file or dir not found")
		return
	}

	if fi.IsDir() {
		RenderDirHtml(c, p) // club penguin
	} else {
		c.File(p.Full)
	}
}

// TODO improve thiS whole ass html rendering
func RenderDirHtml(c *gin.Context, p *Path) {
	entries, err := os.ReadDir(p.Full)
	if err != nil {
		c.String(http.StatusInternalServerError, "cant read the dir :(")
		return
	}

	var body string

	if p.Full != p.Root {
		parentPath := path.Dir(p.Rel)

		body += fmt.Sprintf(`<a href="/%s">..</a><br>`, parentPath)
	}

	for _, e := range entries {
		name := e.Name()
		link := path.Join(p.Rel, name)
		class := ""
		if e.IsDir() {
			link += "/"
			name += "/"
			class += "dir"
		}
		body += fmt.Sprintf(`<a class="%s" href="/%s">%s</a><br>`, class, link, name)
	}

	c.HTML(http.StatusOK, "dirlist", gin.H{
		"Path": p.Req,
		"Body": template.HTML(body),
	})
}
