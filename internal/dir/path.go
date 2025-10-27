package dir

import (
	"fmt"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/reckedpr/boop/internal/cli"
)

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

func HandlePath(c *gin.Context, args *cli.CliArgs) {
	p, err := ResolvePath(c, args.Path)
	if err != nil {
		c.String(404, "file or dir not found")
		return
	}

	fi, err := os.Stat(p.Full)
	if err != nil {
		c.String(http.StatusNotFound, "file or dir not found")
		return
	}

	if fi.IsDir() {
		RenderDirHtml(c, p)
	} else {
		ServeFile(c, p.Full, args)
	}
}

func ServeFile(c *gin.Context, path string, args *cli.CliArgs) {
	if args.Dload {
		c.FileAttachment(path, filepath.Base(path))
	} else {
		c.File(path)
	}
}
