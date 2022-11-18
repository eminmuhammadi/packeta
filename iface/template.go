package iface

import (
	"bytes"
	"text/tabwriter"
	"text/template"

	"github.com/google/gopacket/pcap"
)

type Template struct {
	Header string
	Items  []pcap.Interface
}

func GenerateTable(ifaces []pcap.Interface) (string, error) {
	t := template.New("iface")

	text := `{{.Header}}{{range .Items}}
{{.Name}}	{{.Description}}{{end}}`

	t, _ = t.Parse(text)

	data := Template{
		Header: "NAME	DESCRIPTION",
		Items:  ifaces,
	}

	var result bytes.Buffer
	w := tabwriter.NewWriter(&result, 8, 8, 8, ' ', 0)
	if err := t.Execute(w, data); err != nil {
		return "", err
	}
	w.Flush()

	return result.String(), nil
}
