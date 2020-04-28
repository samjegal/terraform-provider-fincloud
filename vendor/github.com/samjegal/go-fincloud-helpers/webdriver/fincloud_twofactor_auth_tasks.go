package webdriver

import (
	"context"
	"syscall"
	"time"

	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/cdproto/page"
	cdp "github.com/chromedp/chromedp"
	"github.com/samjegal/go-fincloud-helpers/fincloud"
	"golang.org/x/crypto/ssh/terminal"
)

type webTwofactorAuthTasks struct {
	Subaccount string
	Config     *fincloud.Config
	ConfigData *fincloud.ConfigData
}

func (w webTwofactorAuthTasks) build(b Builder) (webdriver, error) {
	driver := webTwofactorAuthTasks{
		Subaccount: b.Subaccount,
		Config:     b.Config,
		ConfigData: b.ConfigData,
	}
	return driver, nil
}

func (w webTwofactorAuthTasks) do(ctx context.Context) error {
	err := listenDialog(ctx)
	if err != nil {
		return err
	}

	otp, err := otpPassword(ctx)
	if err != nil {
		return err
	}

	certKey, err := twofactorAuthenticate(ctx, otp)
	if err != nil {
		return err
	}

	var index int = 0
	for _, cert := range w.ConfigData.CertificateList {
		if cert.SubaccountName == w.Subaccount {
			break
		}
		index++
	}
	timeFormat := "2006-01-02 15:04:05 MST"
	w.ConfigData.CertificateList[index].Key = certKey
	w.ConfigData.CertificateList[index].CreateYmdt = time.Now().Format(timeFormat)

	return nil
}

func listenDialog(ctx context.Context) error {
	cdp.ListenTarget(ctx, func(ev interface{}) {
		if _, ok := ev.(*page.EventJavascriptDialogOpening); ok {
			go func() {
				if err := cdp.Run(ctx, page.HandleJavaScriptDialog(true)); err != nil {
					panic(err)
				}
			}()
		}
	})

	return nil
}

func otpPassword(ctx context.Context) (string, error) {
	err := cdp.Run(ctx, cdp.Tasks{
		cdp.Click(`#loginForm > div > a`),
		cdp.Click(`#app > div.popup > div.panel.certi > div.content > div:nth-child(3) > div.btn-wrap > a`),
	})
	if err != nil {
		return "", err
	}

	bp, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return "", err
	}
	return string(bp), nil
}

func twofactorAuthenticate(ctx context.Context, otpValue string) (string, error) {
	var certKey string
	err := cdp.Run(ctx, cdp.Tasks{
		cdp.SendKeys(`#loginForm > div > input[type=text]`, otpValue),
		cdp.Click(`#loginForm > a`),
		cdp.ActionFunc(func(context context.Context) error {
			time.Sleep(1 * time.Second)
			return nil
		}),
		cdp.ActionFunc(func(context context.Context) error {
			cookies, err := network.GetAllCookies().Do(context)
			if err != nil {
				return err
			}

			for _, cookie := range cookies {
				if cookie.Name == "ncp" {
					certKey = cookie.Value
					break
				}
			}

			return nil
		}),
	})
	if err != nil {
		return "", err
	}

	return certKey, nil
}
