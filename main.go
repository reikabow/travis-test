package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"travis-test/mathutils"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())
		nfloat := float64(n)
		if err != nil {
			fmt.Println("Not an integer")
			continue
		}
		pow := mathutils.Pow(2, nfloat)

		if nfloat == 12 {
			fmt.Println("YOU HIT THE JACKPOT!~")
			os.Exit(0)
		}

		fmt.Printf("2^%f = %f\n", nfloat, pow)
	}
}
