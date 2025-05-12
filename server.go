package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/multiplication-table", multiplicationTableHandler)

	fmt.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, Go Server!")
}

func multiplicationTableHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("multiplicationTable").Parse(`
<!DOCTYPE html>
<html>
<head>
<title>Multiplication Table</title>
<style>
  table {
    border-collapse: collapse;
    width: 50%;
    margin: 20px auto;
  }
  th, td {
    border: 1px solid black;
    padding: 8px;
    text-align: center;
  }
  th {
    background-color: #f2f2f2;
  }
</style>
</head>
<body>

<h2>Multiplication Table (9x9)</h2>

<table>
  <tr>
    <th>x</th>
    {{range .Headers}}<th>{{.}}</th>{{end}}
  </tr>
  {{range $rowIndex, $row := .Rows}}
  <tr>
    <th>{{index $.Headers $rowIndex}}</th>
    {{range $colIndex, $cell := $row}}
    <td>{{$cell}}</td>
    {{end}}
  </tr>
  {{end}}
</table>

</body>
</html>
`)
	if err != nil {
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
		return
	}

	data := struct {
		Headers []int
		Rows    [][]int
	}{
		Headers: make([]int, 9),
		Rows:    make([][]int, 9),
	}

	for i := 0; i < 9; i++ {
		data.Headers[i] = i + 1
		data.Rows[i] = make([]int, 9)
		for j := 0; j < 9; j++ {
			data.Rows[i][j] = (i + 1) * (j + 1)
		}
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
		return
	}
}