package main

import (
	"final-project/course"
	mahasiswa2 "final-project/mahasiswa"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var tmpl = template.Must(template.ParseFiles("./web/mahasiswa/list-mahasiswa.html"))

		mahasiswa := mahasiswa2.Mahasiswa{}
		listMahasiswa, err := mahasiswa.GetMahasiswaList()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Prepare the data to pass to the template
		tmpl.Execute(w, listMahasiswa)
	})

	http.HandleFunc("/add-mahasiswa", func(w http.ResponseWriter, r *http.Request) {
		var tmpl = template.Must(template.ParseFiles("./web/mahasiswa/add-mahasiswa.html"))
		if r.Method == http.MethodPost {

			// Parse the form
			err := r.ParseForm()
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			nim := r.FormValue("nim")
			name := r.FormValue("name")
			age := r.FormValue("age")
			courses := r.FormValue("courses")

			// Convert age string to int
			ageInt, err := strconv.Atoi(age)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			var listOfCourses []course.Course
			splitCourses := strings.Split(courses, ",")
			for _, courseName := range splitCourses {
				listOfCourses = append(listOfCourses, course.Course{
					ID:   string(rune(rand.Int())),
					Name: courseName,
				})
			}

			mahasiswa := mahasiswa2.Mahasiswa{NIM: nim, Name: name, Age: ageInt, Courses: listOfCourses}

			_, err = mahasiswa.AddMahasiswa(mahasiswa)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			// Prepare the data to pass to the template
			data := struct{ Success bool }{true}
			tmpl.Execute(w, data)
			return
		}
		tmpl.Execute(w, nil)

	})

	http.HandleFunc("/list-mahasiswa", func(w http.ResponseWriter, r *http.Request) {
		var tmpl = template.Must(template.ParseFiles("./web/mahasiswa/list-mahasiswa.html"))

		mahasiswa := mahasiswa2.Mahasiswa{}
		listMahasiswa, err := mahasiswa.GetMahasiswaList()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Prepare the data to pass to the template
		tmpl.Execute(w, listMahasiswa)
	})

	http.HandleFunc("/delete-mahasiswa", func(w http.ResponseWriter, r *http.Request) {
		mahasiswa2 := mahasiswa2.Mahasiswa{}

		_, err := mahasiswa2.DeleteMahasiswa(r.URL.Query().Get("nim"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Redirect to list mahasiswa
		http.Redirect(w, r, "/list-mahasiswa", http.StatusSeeOther)
	})

	http.HandleFunc("/update-mahasiswa", func(w http.ResponseWriter, r *http.Request) {
		var tmpl = template.Must(template.ParseFiles("./web/mahasiswa/update-mahasiswa.html"))
		if r.Method == http.MethodPost {
			err := r.ParseForm()
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			nim := r.URL.Query().Get("nim")
			name := r.FormValue("name")
			age := r.FormValue("age")
			courses := r.FormValue("courses")

			// Convert age string to int
			ageInt, err := strconv.Atoi(age)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			var listOfCourses []course.Course
			splitCourses := strings.Split(courses, ",")
			for _, courseName := range splitCourses {
				listOfCourses = append(listOfCourses, course.Course{
					ID:   string(rune(rand.Int())),
					Name: courseName,
				})
			}

			mahasiswa := mahasiswa2.Mahasiswa{NIM: nim, Name: name, Age: ageInt, Courses: listOfCourses}

			_, err = mahasiswa.UpdateMahasiswa(mahasiswa)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			// Redirect to list mahasiswa
			http.Redirect(w, r, "/list-mahasiswa", http.StatusSeeOther)

			// Prepare the data to pass to the template
			data := struct{ Success bool }{true}
			tmpl.Execute(w, data)
			return
		}

		nim := r.URL.Query().Get("nim")
		mahasiswa := mahasiswa2.Mahasiswa{}
		mahasiswaData, err := mahasiswa.GetMahasiswa(nim)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		tmpl.Execute(w, mahasiswaData)
	})

	log.Println("Server started at localhost:8000")
	http.ListenAndServe(":8000", nil)
}
