package api

const AppTemplate = `NAME:
  {{.Name}} - {{.Usage}}

GLOBAL OPTIONS:
  {{if .Commands}}
  {{range .VisibleFlags}}{{.}}
  {{end}}{{end}}{{if .Copyright }}

COPYRIGHT:
  {{.Copyright}}
  {{end}}
 `

const NewTemplate = `NAME:
   {{.Name}} - {{.Usage}}
{{if .Version}}
VERSION:
   {{.Version}}
{{end}}
{{if .Commands}}
COMMANDS:
{{range .Commands}}{{if not .HideHelp}}   {{join .Names ", "}}{{ "\t"}}{{.Usage}}{{ "\n" }}{{end}}{{end}}{{end}}{{if .VisibleFlags}}
GLOBAL OPTIONS:
   {{range .VisibleFlags}}{{.}}
   {{end}}{{end}}
{{if .Copyright }}
COPYRIGHT:
   {{.Copyright}}
{{end}}
`
