package entities

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"text/template"
)

type ExifRecord struct {
	Filename  string `json:"filename"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

func (r *ExifRecord) String() string {
	return fmt.Sprintf("%s %s %s", r.Filename, r.Latitude, r.Longitude)
}

func (r *ExifRecord) Json() (jsonRecord string, err error) {
	jsonBArr, err := json.Marshal(r)
	if err != nil {
		return "", err
	}
	jsonRecord = string(jsonBArr)
	return
}

func (r *ExifRecord) ToArray() []string {
	return []string{r.Filename, r.Latitude, r.Longitude}
}

func PrintRaw(records []ExifRecord) {
	for _, record := range records {
		fmt.Fprintf(os.Stdout, "%v\n", record.String())
	}
}

func PrintCsv(records []ExifRecord) {
	w := csv.NewWriter(os.Stdout)
	w.Write([]string{"filename", "latitude", "longitude"})
	for _, record := range records {
		w.Write(record.ToArray())
	}
	w.Flush()
}

func PrintHtml(records []ExifRecord) {
	var recordsStr [][]string
	for _, record := range records {
		recordsStr = append(recordsStr, record.ToArray())
	}
	printHTMLTable(recordsStr)
}

func printHTMLTable(data [][]string) {
	const tpl = `
	<table>
	<tr>
		<th>filename</th>
		<th>latitude</th>
		<th>longitude</th>
	</tr>
	{{range .}}
		<tr>
		{{range .}}
			<td>{{.}}</td>
		{{end}}
		</tr>
	{{end}}
	</table>
	`
	t := template.Must(template.New("table").Parse(tpl))
	err := t.Execute(os.Stdout, data)
	if err != nil {
		fmt.Println("Error al generar la tabla HTML:", err)
	}
}

func PrintJson(records []ExifRecord) {
	for _, record := range records {
		r, err := record.Json()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: %s\n", record.Filename, err.Error())
			return
		}
		fmt.Fprintf(os.Stdout, "%v\n", r)
	}
}
