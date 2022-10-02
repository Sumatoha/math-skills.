package main

import (
	"math-skills/apply"
	"os"
)

func main() {
	temp := apply.Overflow(apply.Convert(apply.Read(apply.Check(os.Args))))
	apply.PrintAll(temp)
}
