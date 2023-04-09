package google

import (
	"context"
	"fmt"
	"log"
	"os"
	"regexp"
	"time"

	"github.com/chromedp/chromedp"
)

func GetResult(site string) {
	fmt.Println("Gonter")

	dir, err := os.MkdirTemp("", "chromedp-example")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(dir)

	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", true),
		chromedp.DisableGPU,
		chromedp.UserDataDir(dir),
		chromedp.Flag("ignore-certificate-errors", true),
		chromedp.Flag("hide-scrollbars", true),
	)

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	// also set up a custom logger
	ctx, cancel := chromedp.NewContext(allocCtx, chromedp.WithLogf(log.Printf))
	defer cancel()

	var res string
	err = chromedp.Run(
		ctx,
		chromedp.Navigate("https://www.google.com/search?q=site:"+site+""),
		chromedp.Sleep(2*time.Second),
		chromedp.WaitVisible("#result-stats", chromedp.ByID),
		chromedp.Text("#result-stats", &res, chromedp.ByID),
	)
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Printf("Resultado da pesquisa: %s\n", res)
	re := regexp.MustCompile(`\d+\.\d+`)

	// Encontra todas as correspondências na string de texto
	matches := re.FindAllString(res, -1)

	// Exibe o valor "1.010" encontrado na string de texto
	if len(matches) > 0 {
		fmt.Println(matches[0])
	}
}
