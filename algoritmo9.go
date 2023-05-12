package main

import "fmt"

type livro struct {
	autor  string
	preço  float64
	titulo string
}
p := livro{autor: "Fitzgerald" , titulo: "Gatsby" , preço: 19.99}
func imprimelivro(p livro) {
	var desconto *float64 = &p.preço
	fmt.Println("O autor é: ", p.autor)
	*desconto = (p.preço / 100) * 90
	fmt.Println("O preço é: ", p.preço)
	fmt.Println("O titulo é: ", p.titulo)
}
