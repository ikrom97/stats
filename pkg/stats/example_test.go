package stats

import (
	"fmt"
	"github.com/ikrom97/bank/pkg/types"
)
var payments = []types.Payment{
	{
		ID: 1234,
		Amount: 1_000,
		Category: "sportcar",
	},
	{
		ID: 2345,
		Amount: 1_000,
		Category: "armchaire",
	},
	{
		ID: 3456,
		Amount: 1_000,
		Category: "refrigerator",
	},
}
func ExampleAvg() {
	avgAmount := Avg(payments)
	fmt.Println(avgAmount)
	//Output: 1000
}
func ExampleTotalInCategory() {
	categoryAmount := TotalInCategory(payments, "sportcar")
	fmt.Println(categoryAmount)
	//Output: 1000
}
