package estruturasdedados

import "fmt"

func todasAsEstruturasDeDadosPrimitivas() {
	var inteiro int = -30
	var pontoFlutuante float32 = 0.123
	var booleano bool = 5 > 2
	var fraseString string = "Olá, mundo!"
	var inteiroUnsigned uint = 35 // 0 + positivos (conjunto dos números naturais)

	/* Procure também
	   - uintptr
	   - rune
	   - byte
	   - complex32-64
	*/

	fmt.Println(
		inteiro,
		pontoFlutuante,
		booleano,
		fraseString,
		inteiroUnsigned,
	)
}
