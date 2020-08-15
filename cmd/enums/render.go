package enums

import (
	"bytes"
	"github.com/pkg/errors"
	"os"
	"strings"
	"text/template"
)

type Enum struct {

}

func newWithTypeFunction(t string) func (v string) string{
	return func(v string) string {
		return strings.Title(t) + strings.Title(v)
	}
}

var funcMap = template.FuncMap{
	"lower": strings.ToLower,
	"title": strings.Title,
	"upper": strings.ToUpper,
}

type RenderArgs struct {
	Package string
	Type string
	Values []string
	Path string
}

func (e *Enum) Render(args RenderArgs) error {
	funcMap["withType"] = newWithTypeFunction(args.Type)

	tpl, err := template.New("enum").Funcs(funcMap).Parse(Template)
	if err != nil {
		return errors.Wrap(err, "render")
	}

	var buf bytes.Buffer
	if err := tpl.Execute(&buf, args); err != nil {
		return errors.Wrap(err, "render")
	}

	wd, err := os.Getwd()
	if err != nil {
		return errors.Wrap(err, "can not get working directory")
	}

	path := wd + "/" + args.Path
	f, err := os.Create(path)
	if err != nil {
		partsPath := strings.Split(path, "/")
		dir := strings.Join(partsPath[:len(partsPath) - 1], "/")
		if err := os.MkdirAll(dir, 0755); err != nil {
			return errors.Wrap(err, "can not create file")
		}

		f, err = os.Create(path)
		return errors.Wrap(err, "can not create file")
	}

	if _, err := f.Write(buf.Bytes()); err != nil {
		return errors.Wrap(err, "can not write data to file")
	}

	return nil
}
