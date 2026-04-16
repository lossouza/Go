package main

import (
	"errors"
	"fmt"
)

type Student struct {
	Name   string
	Scores []float64
}

type Report struct {
	Student string
	Average float64
	Status  string
}

func main() {
	students := []Student{
		{Name: "Ana", Scores: []float64{8.5, 7.0, 9.2}},
		{Name: "Bruno", Scores: []float64{6.0, 5.5, 7.0}},
		{Name: "Carla", Scores: []float64{}},
	}

	reports := make([]Report, 0, len(students))

	for _, student := range students {
		report, err := buildReport(student)
		if err != nil {
			fmt.Printf("erro ao processar %s: %v\n", student.Name, err)
			continue
		}

		reports = append(reports, report)
		fmt.Printf("%s teve media %.2f e ficou %s\n", report.Student, report.Average, report.Status)
	}

	fmt.Println("\nResumo final")
	for index, report := range reports {
		fmt.Printf("%d. %+v\n", index+1, report)
	}
}

func buildReport(student Student) (Report, error) {
	average, err := calculateAverage(student.Scores)
	if err != nil {
		return Report{}, err
	}

	return Report{
		Student: student.Name,
		Average: average,
		Status:  evaluateStatus(average),
	}, nil
}

func calculateAverage(scores []float64) (float64, error) {
	if len(scores) == 0 {
		return 0, errors.New("o aluno nao possui notas")
	}

	total := 0.0
	for _, score := range scores {
		total += score
	}

	return total / float64(len(scores)), nil
}

func evaluateStatus(average float64) string {
	switch {
	case average >= 7:
		return "aprovado"
	case average >= 5:
		return "recuperacao"
	default:
		return "reprovado"
	}
}
