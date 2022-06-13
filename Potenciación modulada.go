package main

//Importo para imprimir con alias p
import p "fmt"

/**
* Función que dados los números enteros positivos a, b y c
* tal que c>=2, calcula la potencia modulada a^b mod c
* Argumentos:
* a (Int): Entero al que se le sacará la potencia modulada
* b (Int): Potencia a la que se elevará a
* c (Int): Entero mayor a 2 con el que se sacará el módulo.
* Retorna: Entero producto de la operación
 */
func potenciaModulada(a int, b int, c int) int {
	if b == 0 {
		return 1
	}
	return ((a % c) * (potenciaModulada(a, b-1, c))) % c
}

// Función principal con algunos casos de prueba
func main() {
	p.Println(potenciaModulada(4, 7, 9))
	p.Println(potenciaModulada(40, 23, 2))
	p.Println(potenciaModulada(32, 45, 9))
	p.Println(potenciaModulada(43, 4, 45))
}
