package variaveis_valores_tipos

import (
	"fmt"

	"github.com/fabianoflorentino/aprendago/pkg/format"
)

func HelpMeVariaveisValoresTipos() {

	hlp := []format.HelpMe{
		{Flag: "--go-playground", Description: "Exibe o Go Playground.", Width: 0},
		{Flag: "--hello-world", Description: "Exibe o Hello World.", Width: 0},
		{Flag: "--operador-curto-de-declaracao", Description: "Exibe o operador curto de declaração.", Width: 0},
		{Flag: "--a-palavra-reservada-var", Description: "Exibe a palavra reservada var.", Width: 0},
		{Flag: "--explorando-tipos", Description: "Exibe como explorar tipos.", Width: 0},
		{Flag: "--valor-zero", Description: "Exibe o valor zero.", Width: 0},
		{Flag: "--o-pacote-fmt", Description: "Exibe o pacote fmt.", Width: 0},
		{Flag: "--criando-seu-proprio-tipo", Description: "Exibe como criar seu próprio tipo.", Width: 0},
		{Flag: "--conversao-nao-coercao", Description: "Exibe a conversão não coerção.", Width: 0},
	}

	fmt.Println("\nCapítulo 2: Variáveis, Valores e Tipos")
	format.PrintHelpMe(hlp)
}
