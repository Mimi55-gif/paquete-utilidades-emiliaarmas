package utilidades
import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Muestra el menú y devuelve la opción seleccionada
func mostrarMenu() int {
	var opc int
	fmt.Println("Selecciona la moneda a transformar")
	fmt.Println("1. Euro\n2. Libra\n3. Won\n4. BTC\n0. Salir")
	fmt.Scan(&opc)
	return opc
}

// Realiza la conversión según la opción seleccionada
func convertirMoneda(opc int) {
	var dol float64
	fmt.Println("Por favor ingrese la cantidad de dólares que desea convertir")
	fmt.Scan(&dol)

	switch opc {
	case 1:
		euro := dol * 0.86
		fmt.Printf("Dólares = $%.2f\nEuros = €%.2f\n", dol, euro)
	case 2:
		libra := dol * 0.75
		fmt.Printf("Dólares = $%.2f\nLibras = £%.2f\n", dol, libra)
	case 3:
		won := dol * 1411.17
		fmt.Printf("Dólares = $%.2f\nWones = ₩%.2f\n", dol, won)
	case 4:
		btc := dol * 0.0000091
		fmt.Printf("Dólares = $%.2f\nBTC = ₿%.8f\n", dol, btc)
	default:
		fmt.Println("Opción inválida.")
	}
}

// Cuenta las vocales en una frase ingresada por el usuario
func contarVocales() {
	lect := bufio.NewReader(os.Stdin)
	fmt.Println("\nContador de vocales.")
	fmt.Println("Ingresa una frase:")
	frase, err := lect.ReadString('\n')
	if err != nil {
		fmt.Println("Ha ocurrido un error al leer la frase, adiós.")
		return
	}

	frase = strings.ToLower(frase)
	a, e, i, o, u := 0, 0, 0, 0, 0

	for _, c := range frase {
		switch c {
		case 'a':
			a++
		case 'e':
			e++
		case 'i':
			i++
		case 'o':
			o++
		case 'u':
			u++
		}
	}

	fmt.Printf("\nCantidad de vocales en la frase: %s", frase)
	fmt.Printf("a: %d\ne: %d\ni: %d\no: %d\nu: %d\n", a, e, i, o, u)
}
