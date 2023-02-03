package mahasiswa

import (
	"final-project/course"
	"fmt"
)

type Mahasiswa struct {
	NIM     string
	Name    string
	Age     int
	Courses []course.Course
}

type MahasiswaInterface interface {
	AddMahasiswa(mahasiswa Mahasiswa) ([]Mahasiswa, error)
	DeleteMahasiswa(nim string) ([]Mahasiswa, error)
	GetMahasiswa(nim string) (*Mahasiswa, error)
	GetMahasiswaByCourse(courseName string) ([]Mahasiswa, error)
	GetMahasiswaList() ([]Mahasiswa, error)
	UpdateMahasiswa(mahasiswa Mahasiswa) ([]Mahasiswa, error)
}

var DummyMahasiswa = []Mahasiswa{
	{
		NIM:  "123",
		Name: "John Doe",
		Age:  20,
		Courses: []course.Course{
			{
				ID:   "1",
				Name: "Math",
			},
		},
	},
	{
		NIM:  "456",
		Name: "Jane Doe",
		Age:  20,
		Courses: []course.Course{
			{
				ID:   "1",
				Name: "Math",
			},
			{
				ID:   "2",
				Name: "English",
			},
		},
	},
}

func (m *Mahasiswa) GetMahasiswaList() ([]Mahasiswa, error) {
	if len(DummyMahasiswa) == 0 {
		return nil, fmt.Errorf("Tidak ada data mahasiswa")
	}
	return DummyMahasiswa, nil
}

func (m *Mahasiswa) UpdateMahasiswa(mahasiswa Mahasiswa) ([]Mahasiswa, error) {
	var mahasiswaIndex int = -1

	for i, m := range DummyMahasiswa {
		if m.NIM == mahasiswa.NIM {
			mahasiswaIndex = i
		}
	}

	if mahasiswaIndex == -1 {
		return nil, fmt.Errorf("Mahasiswa dengan NIM %s tidak ditemukan", mahasiswa.NIM)
	}

	DummyMahasiswa[mahasiswaIndex] = mahasiswa

	return DummyMahasiswa, nil
}

func (m *Mahasiswa) AddMahasiswa(mahasiswa Mahasiswa) ([]Mahasiswa, error) {
	//TODO: silahkan implementasikan fungsi ini untuk menambahkan mahasiswa
	/*
		Hint:
		- Gunakan slice DummyMahasiswa untuk menyimpan data mahasiswa
		- Gunakan slice DummyMahasiswa sebagai return value
		- Jika NIM sudah ada, kembalikan error dengan pesan "Mahasiswa dengan NIM <NIM> sudah ada"
		- Jika NIM belum ada, tambahkan mahasiswa ke slice DummyMahasiswa dan kembalikan slice DummyMahasiswa
	*/

	//start_answer

	//end_answer

	return DummyMahasiswa, nil //TODO: replace this
}

func (m *Mahasiswa) DeleteMahasiswa(nim string) ([]Mahasiswa, error) {
	//TODO: silahkan implementasikan fungsi ini untuk menghapus mahasiswa
	/*
		Hint:
		- Gunakan slice DummyMahasiswa sebagai sumber data
		- Gunakan slice DummyMahasiswa sebagai return value
		- Jika NIM tidak ditemukan, kembalikan error dengan pesan "Mahasiswa dengan NIM <NIM> tidak ditemukan" dan nil sebagai return value
		- Jika NIM ditemukan, hapus mahasiswa dari slice DummyMahasiswa dan kembalikan slice DummyMahasiswa
		- Untuk menghapus mahasiswa dari slice, gunakan perulangan dan slice slicing (slice[start:end])
		  - ex : slice = append(slice[:index], slice[index+1:]...)
	*/

	//start_answer

	// end_answer

	return nil, nil //TODO: replace this
}

func (m *Mahasiswa) GetMahasiswa(nim string) (*Mahasiswa, error) {
	//TODO: silahkan implementasikan fungsi ini untuk mendapatkan mahasiswa
	/*
		Hint:
		- Gunakan slice DummyMahasiswa sebagai sumber data
		- Jika NIM tidak ditemukan, kembalikan error dengan pesan "Mahasiswa dengan NIM <NIM> tidak ditemukan" dan nil sebagai return value
		- Jika NIM ditemukan, kembalikan mahasiswa yang ditemukan
	*/

	//start_answer

	//end_answer

	return nil, nil //TODO: replace this
}

func (m *Mahasiswa) GetMahasiswaByCourse(courseName string) ([]Mahasiswa, error) {
	//TODO: silahkan implementasikan fungsi ini untuk mendapatkan mahasiswa berdasarkan course
	/*
		Hint:
		- Gunakan slice DummyMahasiswa sebagai sumber data
		- Jika course tidak ditemukan, kembalikan error dengan pesan "Mahasiswa dengan course dengan nama <courseName> tidak ditemukan" dan nil sebagai return value dari slice mahasiswa
		- Jika course ditemukan, kembalikan mahasiswa yang mengambil course tersebut
	*/
	var mahasiswaByCourse []Mahasiswa
	//start_answer

	//end_answer

	return mahasiswaByCourse, nil //TODO: replace this
}
