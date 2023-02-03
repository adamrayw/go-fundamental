package course

type CourseInterface interface {
	GetCourseByName(name string) (*Course, error)
	AddCourse(course Course) ([]Course, error)
}

type Course struct {
	ID   string
	Name string
}

var DummyCourses = []Course{
	{
		ID:   "1",
		Name: "Math",
	},
	{
		ID:   "2",
		Name: "English",
	},
}

func (c *Course) GetCourseByName(name string) (*Course, error) {
	// TODO: implementasi fungsi ini untuk mendapatkan course berdasarkan nama
	/*
		HINT :
		- Gunakan perulangan untuk mencari course yang sesuai dari slice DummyCourses
		- Jika tidak ditemukan course yang sesuai, kembalikan error dengan pesan "Course not found"
		- Jika ditemukan course yang sesuai, kembalikan course tersebut
	*/
	// start_answer

	// end_answer
	return nil, nil // TODO: replace this
}

func (c *Course) AddCourse(course Course) ([]Course, error) {
	// TODO: implementasi fungsi ini untuk menambahkan course baru
	/*
		HINT:
		- Gunakan append untuk menambahkan course baru ke slice DummyCourses
		- Jika course yang akan ditambahkan sudah ada di slice DummyCourses, kembalikan error dengan pesan "Course already exist"
		- Jika course yang akan ditambahkan tidak ada di slice DummyCourses, kembalikan slice DummyCourses yang sudah ditambahkan course baru
	*/
	// start_answer

	// end_answer

	return DummyCourses, nil // TODO: replace this
}
