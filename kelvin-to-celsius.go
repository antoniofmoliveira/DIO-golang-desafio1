package main

import (
	"fmt"
)

func main() {

	const valorDiferenca = 273.15
	temperaturaKelvin := 373.15
	temperaturaCelsius := temperaturaKelvin - valorDiferenca

	fmt.Printf(" %.2f°K equivale a %.2f°C\n", temperaturaKelvin, temperaturaCelsius)

}
