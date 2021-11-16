package main

func getResult(nota float64) string {
	if nota >= 6 {
		return "Aprovado"
	}
	return "Reprovado"

}

func main() {
	println(getResult(6.0))
	println(getResult(4.9))
}
