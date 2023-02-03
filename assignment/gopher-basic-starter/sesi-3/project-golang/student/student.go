package student

type Student struct {
	NIM  string
	Name string
	Age  int
	IPK  float64
}

func (s *Student) GraduationStatus() (string, error) {
	//TODO: silahkan implementasikan fungsi ini
	/*
		- Jika IPK >= 3.5 maka statusnya "Summa Cumlaude"
		- Jika IPK >= 3.0 maka statusnya "Cumlaude"
		- Jika IPK >= 2.5 maka statusnya "Memuaskan"
		- Jika IPKA < 0 atau IPK > 4 maka kembali error dengan pesan "Nilai IPK tidak valid"
	*/

	//start_answer

	//end_answer

	return "", nil //TODO: replace this

}
