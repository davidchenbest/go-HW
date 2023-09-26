package main

import (
	"encoding/json"
	"fmt"

	"github.com/stripe/stripe-go/v75"
	"github.com/stripe/stripe-go/v75/checkout/session"
)

func main() {
	stripe.Key = "sk_test_4eC39HqLyjWDarjtT1zdp7dc"

	params := &stripe.CheckoutSessionParams{
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				// Price:    stripe.String("price_1NufQF2eZvKYlo2CyY5zjs5S"),
				PriceData: &stripe.CheckoutSessionLineItemPriceDataParams{
					Currency:   stripe.String("usd"),
					UnitAmount: stripe.Int64(100),
					ProductData: &stripe.CheckoutSessionLineItemPriceDataProductDataParams{
						Name:   stripe.String("Your Custom Product Name"),
						Images: stripe.StringSlice([]string{"https://picsum.photos/id/3/200/300", "https://picsum.photos/id/1/200/300", "https://picsum.photos/id/2/200/300"}),
					},
				},
				Quantity: stripe.Int64(2),
			},
		},
		Mode:       stripe.String("payment"),
		SuccessURL: stripe.String("https://example.com/success"),
	}
	s, _ := session.New(params)

	jsonData, err := json.Marshal(s)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the JSON data
	fmt.Println(string(jsonData))

}
