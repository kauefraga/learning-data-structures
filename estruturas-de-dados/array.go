package estruturasdedados

import "fmt"

// Em Go, temos array estático e dinâmico
func Array() {
	var numerosArray [4]int // estático = tamanho fixo
	numerosArray[0] = 1
	numerosArray[1] = 5
	numerosArray[2] = 4
	numerosArray[3] = 7

	var numerosSlice []int // dinâmico = tamanho variável
	numerosSlice = append(numerosSlice, 1, 2, 3, 4)

	fmt.Println("Array:", numerosArray)
	fmt.Println("Tamanho do array:", len(numerosArray))
	fmt.Println("slice:", numerosSlice)
	fmt.Println("Tamanho do slice:", len(numerosSlice))
}
