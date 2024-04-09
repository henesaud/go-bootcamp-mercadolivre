package main

import "fmt"

func exercise1() {
	bootcamp := "Palavra de teste"
	fmt.Println("tamanho:", len(bootcamp))

	for _, r := range bootcamp {
		fmt.Println(string(r))
	}
}

func exercise2() {
	const idadeMinimaCliente int = 22
	const tempoMinimoAtividadeCliente int = 1
	const salarioMinimoParaAusenciaDeJuros int = 100000

	clientes := []map[string]interface{}{
		{
			"Nome":      "João",
			"Idade":     21,
			"Atividade": 1,
			"Salário":   100000,
		},
		{
			"Nome":      "José",
			"Idade":     25,
			"Atividade": 1,
			"Salário":   100000,
		},
		{
			"Nome":      "Carlos",
			"Idade":     27,
			"Atividade": 2,
			"Salário":   99000,
		},
		{
			"Nome":      "Antônio",
			"Idade":     35,
			"Atividade": 5,
			"Salário":   150000,
		},
	}

	for _, c := range clientes {
		nome := c["Nome"]
		fmt.Println("Cliente:", nome)

		idade := c["Idade"].(int)
		if idade <= idadeMinimaCliente {
			fmt.Println("Não possui empréstimo disponível (idade insuficiente)")
			continue
		}

		atividade := c["Atividade"].(int)
		if atividade <= tempoMinimoAtividadeCliente {
			fmt.Println("Não possui empréstimo disponível (atividade insuficiente)")
			continue
		}

		salario := c["Salário"].(int)
		fmt.Print("Possui empréstimo disponível ")
		if salario > salarioMinimoParaAusenciaDeJuros {
			fmt.Println("sem juros")
		} else {
			fmt.Println("com juros")
		}
	}
}

func exercise3() {
	//usando maps pois tem menor complexidade para buscar um valor
	months := map[int]string{
		1:  "Janeiro",
		2:  "Fevereiro",
		3:  "Março",
		4:  "Abril",
		5:  "Maio",
		6:  "Junho",
		7:  "Julho",
		8:  "Agosto",
		9:  "Setembro",
		10: "Outubro",
		11: "Novembro",
		12: "Dezembro",
	}

	mes10 := 10
	fmt.Println(months[mes10])
}

func exercise4() {
	var employees = map[string]int{"Benjamin": 20, "Manuel": 26, "Brenda": 19, "Dario": 44, "Pedro": 30}

	fmt.Println("idade de Benjamin:", employees["Benjamin"])

	staffsWithMoreThan21 := 0
	for _, v := range employees {
		if v > 21 {
			staffsWithMoreThan21++
		}
	}

	fmt.Printf("funcionários com mais de 21 anos: %d\n", staffsWithMoreThan21)

	employees["Frederico"] = 25
	delete(employees, "Pedro")
	fmt.Println(employees)
}

func main() {
	exercise1()
	exercise2()
	exercise3()
	exercise4()
}
