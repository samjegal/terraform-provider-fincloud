package webdriver

import "context"

type webdriver interface {
	build(b Builder) (webdriver, error)

	do(ctx context.Context) error
}
