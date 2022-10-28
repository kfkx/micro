package generator

import (
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type Generator interface {
	Generate([]File) error
	Options() Options
}

type generator struct {
	opts Options
}

type File struct {
	Path     string
	Template string
}

func (g *generator) Generate(files []File) error {

	for _, file := range files {

		fp := filepath.Join(g.opts.Directory, file.Path)
		dir := filepath.Dir(fp)

		if file.Template == "" {
			dir = fp
		}

		if _, err := os.Stat(dir); os.IsNotExist(err) {
			if err := os.MkdirAll(dir, 0755); err != nil {
				return err
			}
		}

		if file.Template == "" {
			continue
		}

		f, err := os.Create(fp)
		if err != nil {
			return err
		}

		fn := template.FuncMap{
			"dehyphen": func(s string) string {
				return strings.ReplaceAll(s, "-", "")
			},
			"lowerhyphen": func(s string) string {
				return strings.ReplaceAll(s, "-", "_")
			},
			"tohyphen": func(s string) string {
				return strings.ReplaceAll(s, "_", "-")
			},
			"gitorg": func(s string) string {
				list := strings.Split(s, "/")
				return strings.Join(list[:2], "/")
			},
			"lower": strings.ToLower,
			"title": func(s string) string {
				t := strings.ReplaceAll(strings.Title(s), "-", "")
				return strings.ReplaceAll(t, "_", "")
			},
		}

		t, err := template.New(fp).Funcs(fn).Parse(file.Template)
		if err != nil {
			return err
		}

		err = t.Execute(f, g.opts)
		if err != nil {
			return err
		}
	}
	return nil
}

func (g *generator) Options() Options {
	return g.opts
}

func New(opts ...Option) Generator {
	var options Options

	for _, o := range opts {
		o(&options)
	}

	return &generator{
		opts: options,
	}
}
