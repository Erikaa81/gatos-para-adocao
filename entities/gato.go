package entities

import (
	"strings"
)

type Gato struct {
	Nome                   string
	IdadeAproximada        int
	Vacinado               bool
	Castrado               bool
	Sexo                   string
	Vermifugado            bool
	Doenca                 string
	PossuiDoencaContagiosa bool
}

var doencasTransmissiveis = []string{"PIF", "FIV", "FELV"}

func LiberarConvivioComDemais(doencaTransmissivel string) bool {
	doencaTransmissivel = strings.ToUpper(doencaTransmissivel)

	for _, gatoPortadorDoencaTransmissivel := range doencasTransmissiveis {
		if gatoPortadorDoencaTransmissivel == doencaTransmissivel {
			return true
		}
	}
	return false
}
func CadastrarGato(nome string, idade int, vacinado bool, castrado bool, sexo string, vermifugado bool, doenca string) Gato {

	return Gato{
		Nome:                   nome,
		IdadeAproximada:        idade,
		Vacinado:               vacinado,
		Castrado:               castrado,
		Sexo:                   sexo,
		Vermifugado:            vermifugado,
		Doenca:                 doenca,
		PossuiDoencaContagiosa: LiberarConvivioComDemais(doenca),
	}

}
