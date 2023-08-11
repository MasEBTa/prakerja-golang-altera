package main

import (
	"fmt"
	"math"
)

type Student struct {
	Name  string
	Score []int
}

func (s *Student) AddScore(score int) {
	s.Score = append(s.Score, score)
}

func (s Student) AverageScore() float64 {
	sum := 0
	for _, score := range s.Score {
		sum += score
	}
	return float64(sum) / float64(len(s.Score))
}

func (s Student) MinMaxScore() (int, int) {
	minScore := math.MaxInt32
	maxScore := math.MinInt32
	for _, score := range s.Score {
		if score < minScore {
			minScore = score
		}
		if score > maxScore {
			maxScore = score
		}
	}
	return minScore, maxScore
}

func main() {
	var students [5]Student

	for i := 0; i < 5; i++ {
		var name string
		fmt.Printf("Masukkan nama siswa ke-%d: ", i+1)
		fmt.Scan(&name)

		var scores []int
		for j := 0; j < 3; j++ {
			var score int
			fmt.Printf("Masukkan skor ke-%d untuk siswa %s: ", j+1, name)
			fmt.Scan(&score)
			scores = append(scores, score)
		}

		students[i] = Student{Name: name, Score: scores}
	}

	totalAverage := 0.0
	for _, student := range students {
		average := student.AverageScore()
		fmt.Printf("Siswa %s memiliki rata-rata skor %.2f\n", student.Name, average)
		totalAverage += average
	}

	fmt.Printf("Rata-rata skor seluruh siswa: %.2f\n", totalAverage/float64(len(students)))

	for _, student := range students {
		minScore, maxScore := student.MinMaxScore()
		fmt.Printf("Siswa %s memiliki skor minimum %d dan skor maksimum %d\n", student.Name, minScore, maxScore)
	}
}
