package main

import (
	"fmt"
	"math"

	"github.com/leekchan/accounting"
)

func Statement(invoice Invoice, plays map[string]Play) string {
	totalAmount := 0.0
	volumeCredits := 0
	result := fmt.Sprintf("Statement for %s\n", invoice.Customer)
	usd := func(f float64) string {
		ac := accounting.Accounting{Symbol: "$", Precision: 2}
		return ac.FormatMoneyFloat64(f / 100)
	}

	var playFor = func(performance Performance) Play {
		return plays[performance.PlayID]
	}

	var volumeCreditsFor = func(performance Performance) int {
		result := 0
		result += Max(performance.Audience-30, 0)
		if "comedy" == playFor(performance).Type {
			result += int(math.Floor(float64(performance.Audience / 5)))
		}
		return result
	}

	var amountFor = func(perf Performance) float64 {
		result := 0.0
		switch playFor(perf).Type {
		case "tragedy":
			result = 40000
			if perf.Audience > 30 {
				result += 1000 * float64(perf.Audience-30)
			}
		case "comedy":
			result = 30000
			if perf.Audience > 20 {
				result += 10000 + 500*float64(perf.Audience-20)
			}
			result += 300 * float64(perf.Audience)
		default:
			panic(fmt.Sprintf("unknown type: %s", playFor(perf).Type))
		}

		return result
	}

	for _, perf := range invoice.Performances {

		result += fmt.Sprintf(" %s: %s (%d seats)\n", playFor(perf).Name, usd(amountFor(perf)), perf.Audience)
		totalAmount += amountFor(perf)
	}
	for _, perf := range invoice.Performances {
		volumeCredits += volumeCreditsFor(perf)
	}
	result += fmt.Sprintf("Amount owed is %s\n", usd(totalAmount))
	result += fmt.Sprintf("You earned %d credits\n", volumeCredits)
	return result
}
