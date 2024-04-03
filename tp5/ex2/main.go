package main

import (
	"fmt"
)

type loja struct {
	produtos []produto
}

const (
	pequeno = "pequeno"
	medio   = "medio"
	grande  = "grande"
)

type produto struct {
	nome  string
	preco float64
	tipo  string
}

type Produto interface {
	CalcularCusto() (float64, error)
}
type Ecommerce interface {
	Total() (float64, error)
	Adicionar(produto) (bool, error)
}

func (p produto) CalcularCusto() (float64, error) {
	if p.tipo == pequeno {
		return p.preco, nil
	}
	if p.tipo == medio {
		return p.preco * 1.03, nil
	}
	if p.tipo == grande {
		return p.preco*1.06 + 2500, nil
	}
	return 0.0, fmt.Errorf("invalid type: %s", p.tipo)
}

func (l loja) Total() (float64, error) {
	total := 0.0
	for _, p := range l.produtos {
		custo, _ := p.CalcularCusto()
		total += custo
	}
	return total, nil
}

func (l *loja) Adicionar(p produto) (bool, error) {
	l.produtos = append(l.produtos, p)
	return true, nil
}

func novoProduto(nome string, preco float64, tipo string) Produto {
	return produto{
		nome:  nome,
		preco: preco,
		tipo:  tipo,
	}
}

func novaLoja(produtos []produto) Ecommerce {
	return &loja{
		produtos: produtos,
	}
}

func main() {
	loja1 := novaLoja(
		[]produto{
			{
				nome:  "camiseta",
				preco: 100.0,
				tipo:  pequeno,
			},
			{
				nome:  "jaqueta",
				preco: 200.0,
				tipo:  medio,
			},
			{
				nome:  "casaco",
				preco: 300.0,
				tipo:  grande,
			},
		},
	)

	total, err := loja1.Total()
	if err != nil {
		print(total, err)
	}
	fmt.Println(total)

	fmt.Println("Adding new product...")
	loja1.Adicionar(novoProduto("calca", 100.0, pequeno).(produto))
	total, err = loja1.Total()
	if err != nil {
		print(total, err)
	}
	fmt.Println(total)
}
