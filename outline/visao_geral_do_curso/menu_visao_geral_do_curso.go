package outline

import format "github.com/fabianoflorentino/aprendago/outline/format"

func MenuVisaoGeralDoCurso(args []string) []format.MenuOptions {
	return []format.MenuOptions{
		{Options: "--bem-vindo", ExecFunc: func() { BemVindo() }},
		{Options: "--porque-go", ExecFunc: func() { PorQueGo() }},
		{Options: "--sucesso", ExecFunc: func() { Sucesso() }},
		{Options: "--recursos", ExecFunc: func() { Recursos() }},
		{Options: "--como-esse-curso-funciona", ExecFunc: func() { ComoEsseCursoFunciona() }},
	}
}