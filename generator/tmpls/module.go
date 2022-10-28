package tmpls

var Module = `module {{.Vendor}}{{.Service}}{{if .Client}}-client{{end}}

go 1.19


`
