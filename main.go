/*Criar um programa em Go que guarde informações de gatos a serem adotados. Os gatos devem ser registrados com idade aproximada,
se é vacinado ou não, se é castrado ou não, se é macho ou fêmea, se é vermifugado ou não e se possui alguma das seguintes doenças:
 PIF, Fiv ou Felv, pois deverão ser criados longe de outros gatos, por se tratar de doenças transmissíveis.
 Crie um construtor para esta struct e garanta que todos os dados estarão validados corretamente e teste a função. */
package main

import (
	"fmt"

	"github.com/Erikaa81/gatos-para-adocao/entities"
)

func main() {

	gato := entities.CadastrarGato("Olivia", 7, true, true, "femea", true, "felv")
	if gato.PossuiDoencaContagiosa{

		fmt.Printf("O gato(a) %s, deve ficar separado dos demais:", gato.Nome)
		
	} else {
		fmt.Printf("O gato(a) %s ,está lierado para convivio com os demais!", gato.Nome)

	}
}



