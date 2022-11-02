package main

import "fmt"

func main() {
	words := []string{"a", "cow", "smile", "gopher", "octopus", "anthropologist"}

	for _, word := range words {
		switch size := len(word); size {
		case 1, 2, 3, 4:
			fmt.Println(word, "is a short word!")
		case 5:
			wordLen := len(word)
			fmt.Println(word, "is exactly the right length:", wordLen)
		case 6, 7, 8, 9:
			// there is a keyword fallthrough thats not common
		default:
			fmt.Println(word, "is a long word!")
		}
	}

loop:
	for i := 0; i < 10; i++ {
		switch { // blank switch
		case i%2 == 0:
			fmt.Println(i, "is even")
		case i%3 == 0:
			fmt.Println(i, "is divisible by 3 but not 2")
		case i%7 == 0:
			fmt.Println("exit the loop!")
			break loop //if without the label, breaks only the case and not the for
		default:
			fmt.Println(i, "is boring")
		}
	}
}
