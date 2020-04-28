package webdriver

import (
	"context"
	"time"

	"github.com/chromedp/chromedp"
)

type webBasicAuthTasks struct {
	Subaccount string
	Username   string
	Password   string
}

func (w webBasicAuthTasks) build(b Builder) (webdriver, error) {
	driver := webBasicAuthTasks{
		Subaccount: b.Subaccount,
		Username:   b.Username,
		Password:   b.Password,
	}
	return driver, nil
}

func (w webBasicAuthTasks) do(ctx context.Context) error {
	var host string
	if w.Subaccount != "" {
		host = "https://auth.fin-ncloud.com/nsa/" + w.Subaccount
	} else {
		host = "https://auth.fin-ncloud.com"
	}

	err := chromedp.Run(ctx, chromedp.Tasks{
		chromedp.Navigate(host),
		chromedp.SendKeys(`#username`, w.Username),
		chromedp.SendKeys(`#passwordPlain`, w.Password),
		chromedp.Click(`#loginForm > button`),
		chromedp.ActionFunc(func(context context.Context) error {
			time.Sleep(1 * time.Second)
			return nil
		}),
	})
	if err != nil {
		return err
	}

	return nil
}
