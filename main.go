package main

import (
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
)

func main() {
	fmt.Println("Gonter")

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	err := chromedp.Run(ctx, chromedp.Navigate("https://instagram.com/rogersilvasouza"))
	if err != nil {
		fmt.Println(err)
	}

	var spanText string
	err = chromedp.Run(ctx, chromedp.Text("span", &spanText, chromedp.NodeVisible))
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(spanText)
}
