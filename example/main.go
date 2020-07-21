package main

import "log"

// NOTES:
// - i should not need to define the last state,
//	right now i need to define it so that it can be generated - OK
// - too many unused fields in StateConfig and TransitionConfig
// that are only used during actual core initialization - OK
// - i don't like the callback initialization
// - maybe we can make all params (except context) including
// - should we consider the callback action and enter the same?
// 	coz they seem to have similar function
//	from state as single struct
// - integrate with go generate
// - also, don't forget to add test
// - use type aliasing for type State so we don't need to type cast - OK
// - consider adding a get current state function so we don't need to always include
// 	the from state in the context, but make it optional
// - generating graphs

func main() {
	// generateSwitch()
	// generateMatter()
	// generateCart()
	cartExample()
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
