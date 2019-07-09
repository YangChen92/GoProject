package model

type student struct{
	name string
	score float64
}

func NewStudent(n string,s float64) *student{
	return &student{
		Name : n,
		Score : s,
	}
}

func (s *student) GetScore() float64{
	return s.score
}