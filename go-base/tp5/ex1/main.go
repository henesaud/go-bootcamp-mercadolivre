package main

import (
	"fmt"
	"time"
)

type aluno struct {
	Nome         string
	Sobrenome    string
	RG           string
	DataAdmissao time.Time
}

func (a aluno) printDetails() {
	fmt.Println("Nome:", a.Nome)
	fmt.Println("Sobrenome:", a.Sobrenome)
	fmt.Println("RG:", a.RG)
	fmt.Println("DataAdmissao:", a.DataAdmissao)
}

func main() {
	student1 := aluno{
		Nome:         "Hene",
		Sobrenome:    "SS",
		RG:           "12345678",
		DataAdmissao: time.Date(2024, 03, 18, 0, 0, 0, 0, time.UTC),
	}
	student1.printDetails()

	student2 := aluno{
		Nome:         "Carlos",
		Sobrenome:    "Lacerda",
		RG:           "44324125",
		DataAdmissao: time.Now().UTC(),
	}
	student2.printDetails()
}
