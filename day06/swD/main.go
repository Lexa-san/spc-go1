package main

import (
	"fmt"
)

func main() {

	var (
		bullPrice   int = 20
		cowPrice    int = 10
		calfPrice   int = 5
		budgetLimit int
		qtyLimit    int
	)
	fmt.Scan(&budgetLimit, &qtyLimit)

	var maxBullQty int = int(budgetLimit / bullPrice)
	var maxCowQty int
	var maxCalfQty int
loopI:
	for i := 1; i <= maxBullQty; i++ {

		if i > qtyLimit {
			break loopI
		}
		maxCowQty = int((budgetLimit - (bullPrice * i)) / cowPrice)

	loopJ:
		for j := 0; j <= maxCowQty; j++ {
			if i+j > qtyLimit {
				break loopJ
			}

			maxCalfQty = int((budgetLimit - (bullPrice * i) - (cowPrice * j)) / calfPrice)
		loopZ:
			for z := 0; z <= maxCalfQty; z++ {
				if i+j+z > qtyLimit {
					break loopZ
				}

				if i*bullPrice+j*cowPrice+z*calfPrice == budgetLimit && i+j+z == qtyLimit {
					fmt.Println(i, j, z)
				}

			}

		}

	}

}
