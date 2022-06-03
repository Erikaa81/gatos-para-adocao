package entities

import (
	"testing"
)

func TestGatos(t *testing.T) {

	gato1 := Gato{Nome: "Luma", IdadeAproximada: 10, Vacinado: true, Castrado: true, Sexo: "femea", Vermifugado: true, Doenca: "aquela"}

	t.Run("Esperado gato não portador de doença transmisivel, liberado para conviver com os demais", func(t *testing.T) {
		informacaoRecebida := LiberarConvivioComDemais(gato1.Doenca)

		informacaoEsperada := false
		if informacaoEsperada != informacaoRecebida {
			t.Errorf(" esperado , %t, resultado  %t", informacaoEsperada, informacaoRecebida)
		}

	})

	gato2 := Gato{Nome: "Luma", IdadeAproximada: 10, Vacinado: true, Castrado: true, Sexo: "femea", Vermifugado: true, Doenca: "Fiv"}

	t.Run("Esperado gato portador de PIF, não liberado para convivio com os demais", func(t *testing.T) {
		informacaoRecebida := LiberarConvivioComDemais(gato2.Doenca)

		informacaoEsperada := true
		if informacaoEsperada != informacaoRecebida {
			t.Errorf(" esperado , %t, resultado  %t", informacaoEsperada, informacaoRecebida)
		}
	})

	gato3 := Gato{Nome: "Benedito", IdadeAproximada: 5, Vacinado: true, Castrado: true, Sexo: "macho", Vermifugado: true, Doenca: "Felv"}

	t.Run("Esperado gato portador de Fiv,  não liberado para convivio com os demais", func(t *testing.T) {
		informacaoRecebida := LiberarConvivioComDemais(gato3.Doenca)

		informacaoEsperada := true
		if informacaoEsperada != informacaoRecebida {
			t.Errorf(" esperado , %t, resultado  %t", informacaoEsperada, informacaoRecebida)
		}
	})

	gato4 := Gato{Nome: "Barão", IdadeAproximada: 12, Vacinado: true, Castrado: true, Sexo: "macho", Vermifugado: true, Doenca: "PIF"}

	t.Run("Esperado gato portador de Felv,  não liberado para convivio com os demais", func(t *testing.T) {
		informacaoRecebida := LiberarConvivioComDemais(gato4.Doenca)

		informacaoEsperada := true
		if informacaoEsperada != informacaoRecebida {
			t.Errorf(" esperado , %t, resultado  %t", informacaoEsperada, informacaoRecebida)
		}
	})

	gato5 := Gato{Nome: "Joana", IdadeAproximada: 10, Vacinado: true, Castrado: true, Sexo: "feminino", Vermifugado: true, Doenca: "PIF", PossuiDoencaContagiosa: true}

	t.Run("Esperado que o atributo PossuiDoencaContagiosa seja true porque o animal possui doença contagiosa", func(t *testing.T) {
		informacaoRecebida := LiberarConvivioComDemais(gato5.Doenca)

		informacaoEsperada := gato5.PossuiDoencaContagiosa
		if informacaoEsperada != informacaoRecebida {
			t.Errorf(" esperado , %v, resultado  %v", informacaoEsperada, informacaoRecebida)
		}
	})

	gato6 := Gato{Nome: "Joana", IdadeAproximada: 5, Vacinado: true, Castrado: true, Sexo: "feminino", Vermifugado: true, Doenca: "Sarna", PossuiDoencaContagiosa: false}

	t.Run("Esperado que o atributo PossuiDoencaContagiosa seja false porque o animal nao possui doença contagiosa", func(t *testing.T) {
		informacaoRecebida := LiberarConvivioComDemais(gato6.Doenca)

		informacaoEsperada := gato6.PossuiDoencaContagiosa
		if informacaoEsperada != informacaoRecebida {
			t.Errorf(" esperado , %v, resultado  %v", informacaoEsperada, informacaoRecebida)
		}
	})

	gato7 := Gato{Nome: "Juju", IdadeAproximada: 6, Vacinado: true, Castrado: false, Sexo: "feminino", Vermifugado: true, Doenca: "PIF", PossuiDoencaContagiosa: true}

	t.Run("Chamando o construtor, esperado que o atributo PossuiDoencaContagiosa seja true porque o animal possui doença contagiosa", func(t *testing.T) {

		informacaoRecebida := CadastrarGato(gato7.Nome, gato7.IdadeAproximada, gato7.Vacinado, gato7.Castrado, gato7.Sexo, gato7.Vermifugado, gato7.Doenca)

		informacaoEsperada := gato7
		if informacaoEsperada != informacaoRecebida {
			t.Errorf(" esperado , %v, resultado  %v", informacaoEsperada, informacaoRecebida)
		}
	})

	gato8 := Gato{Nome: "Cacau", IdadeAproximada: 10, Vacinado: true, Castrado: true, Sexo: "feminino", Vermifugado: true, Doenca: "gripe", PossuiDoencaContagiosa: false}

	t.Run("Chamando o construtor, esperado que o atributo PossuiDoencaContagiosa seja false porque o animal nao possui doença contagiosa", func(t *testing.T) {

		informacaoRecebida := CadastrarGato(gato8.Nome, gato8.IdadeAproximada, gato8.Vacinado, gato8.Castrado, gato8.Sexo, gato8.Vermifugado, gato8.Doenca)

		informacaoEsperada := gato8
		if informacaoEsperada != informacaoRecebida {
			t.Errorf(" esperado , %v, resultado  %v", informacaoEsperada, informacaoRecebida)
		}
	})

}
