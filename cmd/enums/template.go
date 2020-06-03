package enums

const Template = `package {{ .Package }}

import "errors"

type {{ title .Type }} uint8

const ({{ range $i, $v :=  .Values }} {{ if eq $i 0 }} 
   {{ withType $v }} = iota + 1 {{ else }} 
   {{ withType $v }} {{ end }}	{{ end }}
)

var {{ lower .Type }}Table = map[string]{{ .Type }}{ {{ range .Values }}
    "{{ upper . }}": {{ withType . }}, {{ end }}
}

func (t {{ title .Type }}) String() string {
	switch t { {{ range .Values }}    
    case {{ withType . }}:
        return "{{ upper . }}"	{{ end }}
    }
	return "unknown"
}

var ErrUnknown{{ title .Type }} = errors.New("unknown {{ lower .Type }}")

func New{{ title .Type }}(str string) ({{ title .Type }}, error) {
	if v, ok := {{ lower .Type }}Table[str]; ok {
		return v, nil
	} else {
		return {{ title .Type }}(0), ErrUnknown{{ title .Type }}
	}
}`
