package student

type Student struct {
	Name  string
	Age   int
	Score int
}

type StudentInterface interface {
	GetName() string
	GetAge() int
	ScoreConvert() int
}

func (s *Student) GetName() string {
	return s.Name
}

func (s *Student) GetAge() int {
	return s.Age
}

func (s *Student) GetScore() string {
	if s.Score > 80 {
		return "Awesome!"
	} else {
		return "Good"
	}
}
