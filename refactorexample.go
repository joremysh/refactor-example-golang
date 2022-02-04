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
	format := func(f float64) string {
		ac := accounting.Accounting{Symbol: "$", Precision: 2}
		return ac.FormatMoneyFloat64(f)
	}

	var playFor = func(performance Performance) Play {
		return plays[performance.PlayID]
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
		// add volume credit
		volumeCredits += Max(perf.Audience-30, 0)

		// every 10 comedy audiences could get extra volume credits
		if "comedy" == playFor(perf).Type {
			volumeCredits += int(math.Floor(float64(perf.Audience / 5)))
		}

		result += fmt.Sprintf(" %s: %s (%d seats)\n", playFor(perf).Name, format(amountFor(perf)/(100)), perf.Audience)
		totalAmount += amountFor(perf)
	}
	result += fmt.Sprintf("Amount owed is %s\n", format(totalAmount/100))
	result += fmt.Sprintf("You earned %d credits\n", volumeCredits)
	return result
}
