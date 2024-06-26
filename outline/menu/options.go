package menu

import (
	"fmt"

	"github.com/fabianoflorentino/aprendago/outline"
	"github.com/fabianoflorentino/aprendago/outline/agrupamento_de_dados"
	"github.com/fabianoflorentino/aprendago/outline/exercicios_ninja_nivel_1"
	"github.com/fabianoflorentino/aprendago/outline/exercicios_ninja_nivel_2"
	"github.com/fabianoflorentino/aprendago/outline/exercicios_ninja_nivel_3"
	"github.com/fabianoflorentino/aprendago/outline/exercicios_ninja_nivel_4"
	"github.com/fabianoflorentino/aprendago/outline/fluxo_de_controle"
	"github.com/fabianoflorentino/aprendago/outline/format"
	"github.com/fabianoflorentino/aprendago/outline/fundamentos_da_programacao"
	"github.com/fabianoflorentino/aprendago/outline/variaveis_valores_tipos"
	"github.com/fabianoflorentino/aprendago/outline/visao_geral_do_curso"
)

// Options é a função principal do pacote menu. Ela é responsável por exibir as opções disponíveis em cada menu.
// A função Options recebe uma lista de argumentos (args) como parâmetro. Esses argumentos são usados para determinar
func Options(args []string) {
	fmt.Printf("Aprenda GO\n\n")

	// A função buildOptions é responsável por construir as opções disponíveis em cada menu.
	// Ela recebe uma lista de argumentos (args) e uma lista de opções (options) como parâmetros.
	buildOptions(
		args,
		generalOptions(args),
		variaveis_valores_tipos.MenuVariaveisValoresTipos(args),
		visao_geral_do_curso.MenuVisaoGeralDoCurso(args),
		exercicios_ninja_nivel_1.MenuExerciciosNinjaNivel1(args),
		fundamentos_da_programacao.MenuFundamentosDaProgramcao(args),
		exercicios_ninja_nivel_2.MenuExerciciosNinjaNivel2(args),
		fluxo_de_controle.MenuFluxoDeControle(args),
		exercicios_ninja_nivel_3.MenuExerciciosNinjaNivel3(args),
		agrupamento_de_dados.MenuAgrupamentoDeDados(args),
		exercicios_ninja_nivel_4.MenuExerciciosNinjaNivel4(args),
	)
}

// A função generalOptions retorna uma lista de opções gerais disponíveis em todos os menus.
// Neste caso, a função retorna duas opções: --help e --outline.
func generalOptions([]string) []format.MenuOptions {
	return []format.MenuOptions{
		{Options: "--help", ExecFunc: func() { HelpMe() }},
		{Options: "--outline", ExecFunc: func() { outline.Outline() }},
	}
}

// Ela recebe dois parâmetros: args, que é uma lista de argumentos do menu, e options, que é uma lista de opções do menu.
// A função começa juntando os argumentos separados por espaço em uma única string usando a função strings.Join(args, " ").
// Isso é feito para facilitar a comparação com as opções disponíveis.

// Em seguida, a função percorre cada opção disponível usando um loop for range. Para cada opção, ela verifica se a
// string de argumentos (argStr) é igual à opção (opt.Options). Se houver uma correspondência, a função opt.ExecFunc() é executada.
// Essa função é responsável por executar a ação associada à opção do menu.

// Se nenhuma correspondência for encontrada, a função imprime o cabeçalho do menu usando fmt.Print(headerMenu).
// Isso pode ser útil para exibir uma mensagem de erro ou mostrar as opções disponíveis novamente.
func buildOptions(args []string, options ...[]format.MenuOptions) {
	var opt []format.MenuOptions

	// options ...[]format.MenuOptions: Aqui, options é um parâmetro de função que permite receber um número variável de argumentos do tipo []format.MenuOptions.
	// Isso significa que você pode passar zero ou mais argumentos do tipo []format.MenuOptions para a função buildMenu.
	// Esses argumentos serão agrupados em uma única fatia ([]format.MenuOptions) dentro da função.

	// o := range options: Aqui, o é uma variável que recebe cada elemento da fatia options à medida que o loop for itera sobre ela.
	// O uso de ... antes da variável o indica que o é uma fatia de elementos do tipo format.MenuOptions.
	// Isso permite que você itere sobre uma fatia de fatias de format.MenuOptions e extraia cada elemento individualmente.
	for _, o := range options {
		opt = append(opt, o...)
	}

	format.FormatMenuOptions(args, opt)
}
