package funcoes

import (
	"fmt"

	"github.com/fabianoflorentino/aprendago/pkg/format"
)

func Callback() {
	topic := format.OutlineContent{
		Title: "Callback",
		Content: `
- Primeiro veja se você entende isso: https://play.golang.org/p/QkAtwMZU-z
- Callback é passar uma função como argumento.
- Exemplo:
  - Criando uma função que retorna uma função e um []int, e usa somente os números pares como argumentos para a função.
  - Go Playground:
- Desafio: Crie uma função no programa acima que utilize somente os números *ímpares.*
- Solução:
    `,
	}

	format.FormatOutlineTopic(topic)
}

func ResolucaoDesafioCallback() {
	result := somaImapres(soma, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}...)

	explicacao := format.OutlineContent{
		Title: "\nExplicação",
		Content: `
A função somaImpares() recebe uma função e um slice de inteiros e retorna a soma dos números pares.
A função passada recebe um número variatico de inteiros e retorna um inteiro.
Os números recebidos pela função são variaticos e serão passados para a função que soma os números somente se forem ímpares.
    `,
	}

	fmt.Printf("Soma dos números ímpares é: %d\n", result)
	format.FormatOutlineTopic(explicacao)
}

// soma recebe um número variatico de inteiros e retorna a soma de todos os números.
func soma(numeros ...int) int {
	somaNum := 0

	for _, num := range numeros {
		somaNum += num
	}

	return somaNum
}

// somaImpares recebe uma função e um slice de inteiros e retorna a soma dos números pares.
// A função passada recebe um número variatico de inteiros e retorna um inteiro.
// Os números recebidos pela função são variaticos e serão passados para a função que soma os números somente se forem ímpares.
func somaImapres(f func(...int) int, numeros ...int) int {
	soma := 0

	for _, num := range numeros {
		if num%2 != 0 {
			soma += f(num)
		}
	}
	return soma
}
