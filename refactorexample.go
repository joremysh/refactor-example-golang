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

	for _, perf := range invoice.Performances {
		play := plays[perf.PlayID]
		thisAmount := 0.0

		switch play.Type {
		case "tragedy":
			thisAmount = 40000
			if perf.Audience > 30 {
				thisAmount += 1000 * float64(perf.Audience-30)
			}
		case "comedy":
			thisAmount = 30000
			if perf.Audience > 20 {
				thisAmount += 10000 + 500*float64(perf.Audience-20)
			}
			thisAmount += 300 * float64(perf.Audience)
		default:
			panic(fmt.Sprintf("unknown type: %s", play.Type))
		}

		// add volume credit
		volumeCredits += Max(perf.Audience-30, 0)

		// every 10 comedy audiences could get extra volume credits
		if "comedy" == play.Type {
			volumeCredits += int(math.Floor(float64(perf.Audience / 5)))
		}

		result += fmt.Sprintf(" %s: %s (%d seats)\n", play.Name, format(thisAmount/(100)), perf.Audience)
		totalAmount += thisAmount
	}
	result += fmt.Sprintf("Amount owed is %s\n", format(totalAmount/100))
	result += fmt.Sprintf("You earned %d credits\n", volumeCredits)
	return result
}
