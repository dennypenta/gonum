package enums

// {{ if eq $i 0 }} 
// {{ withType $v }} = iota + 1 {{ else }}

const Template = `package {{ .Package }}

import "errors"

type {{ title .Type }} uint8

const (
	Invalid{{ title .Type }} {{ .Type }} = iota	{{ range $i, $v :=  .Values }}
	{{ title $v }}{{ end }}
)

var {{ lower .Type }}Table = map[string]{{ .Type }}{ {{ range .Values }}
    "{{ title . }}": {{ title . }}, {{ end }}
}

func (t {{ title .Type }}) String() string {
	switch t { {{ range .Values }}    
    case {{ title . }}:
        return "{{ title . }}"	{{ end }}
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
