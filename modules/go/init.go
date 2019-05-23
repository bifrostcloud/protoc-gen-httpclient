package module

import (
	"io"
	"strings"
	templatex "text/template"

	"github.com/gobuffalo/packr/v2"
)

var template *templatex.Template

func init() {
	template = templatex.New("plugin")
	box := packr.New("pluginBox", "../../templates/go")
	err := box.Walk(walkFN)
	if err != nil {
		panic(err)
	}
}

var walkFN = func(s string, file packr.File) error {
	var sb strings.Builder
	if _, err := io.Copy(&sb, file); err != nil {
		return err
	}
	var err error
	template, err = template.Parse(sb.String())
	return err
}
