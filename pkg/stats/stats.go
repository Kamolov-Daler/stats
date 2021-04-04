package stats

import (
	"github.com/Kamolov-Daler/bank/v2/pkg/types"
)

// Avg рассчитывает среднюю сумму платежа.
func Avg(payments []types.Payment) types.Money {
	result := types.Money(0)
	iter := types.Money(0)

	for _, papayment := range payments {
		if papayment.Status != types.StatusFail {
			iter++
			result += papayment.Amount
		}
	}
	result = result / iter
	return result
}

// TotalInCategory находит сумму покупок в определённой категории.
func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	categoryResult := types.Money(0)

	for _, payment := range payments {
		if category == payment.Category {
			if payment.Status != types.StatusFail {
				categoryResult += payment.Amount
			}
		}
	}
	return categoryResult
}

func CategoriesAvg(payments []types.Payment) map[types.Category]types.Money {
	categoriesPayment := map[types.Category]types.Money{}
	avgPayment := map[types.Category]types.Money{}

	for _, payment := range payments {
		categoriesPayment[payment.Category] += payment.Amount
		avgPayment[payment.Category] += 1
	}

	for key, value := range categoriesPayment {
		avg := avgPayment[key]
		categoriesPayment[key] = value / avg
	}
	return categoriesPayment
}

func PeriodsDynamic(
	first map[types.Category]types.Money, second map[types.Category]types.Money,
) map[types.Category]types.Money {
	result := map[types.Category]types.Money{
		"auto":   0,
		"food":   0,
		"mobile": 0,
	}

	for key := range result {
		result[key] = second[key] - first[key]
	}
	return result
}
