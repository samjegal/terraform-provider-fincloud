package webdriver

import (
	"context"
	"log"

	"github.com/chromedp/chromedp"
	cdp "github.com/chromedp/chromedp"
	"github.com/samjegal/go-fincloud-helpers/fincloud"
)

type Builder struct {
	Subaccount string
	Username   string
	Password   string
	Config     *fincloud.Config
	ConfigData *fincloud.ConfigData
}

type Config struct {
	Headless   bool
	DisableGpu bool
}

func (b Builder) Build() error {
	// TODO: 설정파일에서 읽어서 세팅할 수 있을까?
	config := Config{
		Headless:   true,
		DisableGpu: true,
	}

	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		cdp.Flag("headless", config.Headless),
		cdp.Flag("disable-gpu", config.DisableGpu))

	ctx, cancel := cdp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel = cdp.NewContext(ctx, cdp.WithLogf(log.Printf))
	defer cancel()

	webdriver := []webdriver{
		webBasicAuthTasks{},
		webTwofactorAuthTasks{},
	}

	for _, w := range webdriver {
		driver, err := w.build(b)
		if err != nil {
			return err
		}

		driver.do(ctx)
	}

	return nil
}
