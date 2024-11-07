package main

import (
	"fmt"
)

func main() {

	const temperaturaEbulicaoAguaEmKelvin = 373.15
	const valorDiferenca = 273.15
	temperaturaCelsius := temperaturaEbulicaoAguaEmKelvin - valorDiferenca

	fmt.Printf("O ponto de ebulicao da agua em %.2f°K equivale a %.2f°C\n", temperaturaEbulicaoAguaEmKelvin, temperaturaCelsius)

}
