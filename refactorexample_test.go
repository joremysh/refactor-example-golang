package main

import (
	"encoding/json"
	"os"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Refactor example", func() {
	var invoices []Invoice
	var plays map[string]Play

	BeforeEach(func() {
		invoices, plays = parseSampleData()
	})

	Describe("Statement", func() {
		It("Statement", func() {
			for _, invoice := range invoices {
				expected := `Statement for BigCo
 Hamlet: $650.00 (55 seats)
 As You Like It: $580.00 (35 seats)
 Othello: $500.00 (40 seats)
Amount owed is $1,730.00
You earned 47 credits
`
				actual := Statement(invoice, plays)
				Expect(actual).To(Equal(expected))
			}
		})
	})
})

func parseSampleData() ([]Invoice, map[string]Play) {
	byteValue, err := os.ReadFile("invoices.json")
	Expect(err).To(BeNil())
	invoice := make([]Invoice, 0)
	err = json.Unmarshal(byteValue, &invoice)
	Expect(err).To(BeNil())

	byteValue, err = os.ReadFile("plays.json")
	Expect(err).To(BeNil())
	plays := make(map[string]Play)
	err = json.Unmarshal(byteValue, &plays)
	Expect(err).To(BeNil())

	return invoice, plays
}

func TestRefactorExampleGolang(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "RefactorExampleGolang Suite")
}
