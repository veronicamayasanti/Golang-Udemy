package main

import "fmt"

func main() {
	//	Operasi Matematika
	//	+		penjumlahan
	//	-		pengurangan
	//	*		perkalian
	//	/		pembagian
	//	%		sisa pembagian

	a := 5
	b := 5
	c := a + b
	d := a - b
	e := a * b
	f := a / b

	fmt.Println("nilai c d e f = ", c, d, e, f)

	//	Augmented Assignments
	//	operasi matematika		augmented assignments
	//	a = a + 10				a += 10
	//	a = a - 10				a -= 10
	//	a = a * 10				a *= 10
	//	a = a / 10				a /= 10
	//	a = a % 10				a %= 10

	i := 10
	i += 10

	fmt.Println("nilai i = ", i)

	//	unary operatorr
	//	operator			keterangan
	//	++					a = a + 1
	//	--					a = a - 1
	//	-					negative
	//	+					positive
	//	!					Boolean kebalikan

	j := 1
	j++
	fmt.Println("nilai j pertama ", j)
	j++
	fmt.Println("nilai j kedua ", j)
	g := 10
	g--
	fmt.Println("nilai g = ", g)
	g--
	fmt.Println("nilai g ke2 = ", g)
}
