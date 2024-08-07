package agrupamento_de_dados

import "github.com/fabianoflorentino/aprendago/pkg/format"

func MenuAgrupamentoDeDados([]string) []format.MenuOptions {
	return []format.MenuOptions{
		{Options: "--array", ExecFunc: func() { Array() }},
		{Options: "--slice-literal-composta", ExecFunc: func() { SliceLiteralComposta() }},
		{Options: "--slice-for-range", ExecFunc: func() { SliceForRange() }},
		{Options: "--slice-fatiando-ou-deletando-de-uma-fatia", ExecFunc: func() { SliceFatiandoOuDeletandoDeUmaFatia() }},
		{Options: "--slice-fatiando-ou-deletando-de-uma-fatia --resolucao", ExecFunc: func() { ResolucaoFatiaOuDeletandoDeUmaFatia() }},
		{Options: "--slice-anexando-a-uma-slice", ExecFunc: func() { SliceAnexandoASlice() }},
		{Options: "--slice-make", ExecFunc: func() { SliceMake() }},
		{Options: "--slice-multi-dimensional", ExecFunc: func() { SliceMultiDimensional() }},
		{Options: "--slice-a-surpresa-do-array-subjacente", ExecFunc: func() { SliceASurpresaDoArraySubjacente() }},
		{Options: "--maps-introducao", ExecFunc: func() { MapsIntroducao() }},
		{Options: "--maps-range-e-deletando", ExecFunc: func() { MapRangeEDeletando() }},
	}
}
