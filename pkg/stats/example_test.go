package stats

import (
	"fmt"

	"github.com/Kamolov-Daler/bank/v2/pkg/types"
)

func ExampleAvg() {
	payments := []types.Payment{
		{
			ID:       1,
			Amount:   500_00,
			Category: "phone",
		},
		{
			ID:       2,
			Amount:   500_00,
			Category: "phone",
		},
		{
			ID:       3,
			Amount:   500_00,
			Category: "phone",
		},
	}
	fmt.Println(Avg(payments))

	// Output:
	// 50000
}

func ExampleTotalInCategory() {
	payments := []types.Payment{
		{
			ID:       1,
			Amount:   500_00,
			Category: "phone",
		},
		{
			ID:       2,
			Amount:   500_00,
			Category: "phone",
		},
		{
			ID:       3,
			Amount:   500_00,
			Category: "phone",
		},
	}
	fmt.Println(TotalInCategory(payments, "phone"))

	// Output:
	// 150000
}
