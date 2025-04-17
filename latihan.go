package main 
import "fmt"

func main(){
	const NMAX int = 100
	var A[NMAX]int
	var i, n, positif, negatif, nol int

	fmt.Scan(&n)
	for i = 0; i < n; i++{
		fmt.Scan(&A[i])
	}
	for i = 0; i < n; i++{
		if A[i] > 0{
			positif++
		}else if A[i] < 0{
			negatif++
		}else{
			nol++
		}
	}
	fmt.Println("Positif :", positif)
	fmt.Println("Negatif :", negatif)
	fmt.Println("Nol :", nol)
}