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

	//fileName := e.filename(args.Path, args.FileName)
	if !strings.HasSuffix(args.Path, ".go") {
		args.Path += ".go"
	}
	f, err := os.Create(args.Path)
	if err != nil {
		return errors.Wrap(err, "can not create file")
	}

	if _, err := f.Write(buf.Bytes()); err != nil {
		return errors.Wrap(err, "can not write data to file")
	}

	return nil
}

//func (e Enum) filename(path, filename string) string {
//	if strings.HasSuffix(path, "/") {
//		return path + filename
//	}
//
//	return path + "/" + filename
//}