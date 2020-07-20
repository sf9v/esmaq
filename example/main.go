package main

func main() {
	cartExample()
	// generateCart()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
