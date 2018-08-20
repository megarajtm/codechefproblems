package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }
func scanf(f string, a ...interface{})  { fmt.Fscanf(reader, f, a...) }

func main() {
	defer writer.Flush()
	var testCase int
	scanf("%d\n", &testCase)
	for i := 0; i < testCase; i++ {
		var loss float64
		var noRecipies int
		scanf("%d\n", &noRecipies)
		for j := 0; j < noRecipies; j++ {
			var unitPrice, noUnits, dis float64
			scanf("%f %f %f\n", &unitPrice, &noUnits, &dis)
			dis = dis / 100
			sellingPrice := unitPrice * (1 + dis)
			lossPerUnit := unitPrice - (sellingPrice * (1 - dis))
			loss = loss + (lossPerUnit * noUnits)
		}
		printf("%f\n", loss)
	}
}
