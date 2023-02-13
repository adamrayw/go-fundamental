package main

import (
	"html/template"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("webpage").Parse(`
		<html>
			<head>
				<title>Hello World</title>
			</head>
			<body>
				<h1>Welcome to my website</h1>
			</body>
		</html>
	`)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	name := r.FormValue("name")
	err = t.Execute(w, name)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
