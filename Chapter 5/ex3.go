/*
- Crie um novo tipo: veículo
- O tipo subjacente deve ser struct
- Deve conter os campos: portas, cor
- Crie dois novos tipos: caminhonete e sedan
- Os tipos subjacentes devem ser struct
- Ambos devem conter "veículo" como struct embutido
- O tipo caminhonete deve conter um campo bool chamado "traçãoNasQuatro"
- O tipo sedan deve conter um campo bool chamado "modeloLuxo"
- Usando os structs veículo, caminhonete e sedan:
- Usando composite literal, crie um valor de tipo caminhonete e dê valores a seus campos
- Usando composite literal, crie um valor de tipo sedan e dê valores a seus campos
- Demonstre estes valores.
- Demonstre um único campo de cada um dos dois.
*/
package main

import ("fmt")

type veiculo struct {
	portas int
	cor string
}

type caminhonete struct {
	veiculo
	tracaoNasQuatro bool
}

type sedan struct {
	veiculo
	modeloLuxo bool
}

func main() {

	x := sedan {
		veiculo: veiculo {
			cor: "roxo",
			portas: 4,
		},
		modeloLuxo: true,
	}

	y := caminhonete {
		veiculo: veiculo {
			portas: 8,
			cor: "ferrugem",
		},
		tracaoNasQuatro: false,
	}
	
	fmt.Println(y)
	fmt.Println(x)
}