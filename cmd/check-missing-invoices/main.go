package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/stripe/stripe-go/v87"
	"io/fs"
	"log"
	"math"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type StripePayment struct {
	ID       string
	Amount   int
	Currency stripe.Currency
}

func main() {
	_, err := getSuccessfulPaymentsFromStripe()

	if err != nil {
		log.Fatal(err)
	}

	// check whether the invoice was already created in ~/Documents/Taxes/Invoices/2026
	homeUserDir, err := os.UserHomeDir()

	if err != nil {
		log.Fatalf("get user home dir: %v", err)
	}

	invoicesDir := filepath.Join(homeUserDir, "Documents/Taxes/Invoices/2026")

	entries, err := os.ReadDir(invoicesDir)

	if err != nil {
		log.Fatalf("read invoices dir: %v", err)
	}

	nextnativeInvoices := make([]fs.DirEntry, 0)

	for _, entry := range entries {
		re := regexp.MustCompile(`^Invoice \d{4}-\d{3}\.pdf$`)

		if re.MatchString(entry.Name()) {
			// fmt.Printf("%#v\n", entry)
			nextnativeInvoices = append(nextnativeInvoices, entry)
		}
	}
}

func getSuccessfulPaymentsFromStripe() ([]*StripePayment, error) {
	err := godotenv.Load()

	if err != nil {
		return nil, fmt.Errorf("read .env file: %w", err)
	}

	stripeApiKey := os.Getenv("STRIPE_API_KEY")

	// get transactions from stripe
	sc := stripe.NewClient(stripeApiKey)

	params := &stripe.PaymentIntentListParams{}
	params.Limit = stripe.Int64(3)

	ctx := context.Background()

	result := sc.V1PaymentIntents.List(context.TODO(), params)

	payments := make([]*StripePayment, 0)

	for pi, err := range result.All(ctx) {
		if err != nil {
			return nil, fmt.Errorf("list payment intents: %w", err)
		}

		if pi.Status != stripe.PaymentIntentStatusSucceeded {
			continue
		}

		fmt.Printf("%#v\n\n", pi.LatestCharge)

		payments = append(payments, &StripePayment{
			ID:       pi.ID,
			Amount:   int(pi.Amount),
			Currency: pi.Currency,
		})

		fmt.Printf("Transaction #%s: %g %s 💰\n", pi.ID, math.Ceil(float64(pi.Amount/100)), strings.ToUpper(string(pi.Currency)))
	}

	return payments, nil
}
