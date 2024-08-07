package exercicios_ninja_nivel_2

import (
	"fmt"

	"github.com/fabianoflorentino/aprendago/pkg/format"
)

func NaPraticaExercicio1() {
	topic := format.OutlineContent{
		Title: "Na prática - Exercício #1",
		Content: `
- Escreva um programa que mostre um número em decimal, binário, e hexadecimal.
- Solução: (https://play.golang.org/p/X7qm3aWSa6) ou use no menu do programa (--na-pratica-exercicio-1 --nivel-2 --resolucao)
		`,
	}

	format.FormatOutlineTopic(topic)
}

func ResolucaoNaPraticaExercicio1() {
	numero := 1000019

	resoluca := fmt.Sprintf("Decimal: %d \nBinário: %b \nHexadecimal: %#x \n", numero, numero, numero)

	format.FormatResolucaoExercicios(resoluca)
}
