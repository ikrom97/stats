package stats

import (
	"fmt"
	"github.com/ikrom97/bank/v2/pkg/types"
)
var payments = []types.Payment{
	{
		ID: 1234,
		Amount: 3_000,
		Category: "sportcar",
		Status: types.StatusOK,
	},
	{
		ID: 2345,
		Amount: 1_000,
		Category: "armchaire",
		Status: types.StatusFail,
	},
	{
		ID: 3456,
		Amount: 1_000,
		Category: "refrigerator",
		Status: types.StatusFail,
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
	//Output: 3000
}
