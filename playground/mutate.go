package main


func main() {


	// Declare variable of type int with a value of 10.
	x := 10

	// Display the "value of" and "address of" count.
	
	println("x:\tValue Of[", x, "]\tAddr Of[", &x, "]")


	// Pass the "value of" the count.
	inc(&x)

	println("x:\tValue Of[", x, "]\tAddr Of[", &x, "]")
}

func inc(x *int) {

	// Increment the "value of" inc.
	*x++
	println("inc:\tValue Of[", *x, "]\tAddr Of[", &x, "]")
}
	