package timeouts

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/samjegal/terraform-provider-fincloud/fincloud/internal/features"
)

func ForCreate(ctx context.Context, d *schema.ResourceData) (context.Context, context.CancelFunc) {
	return buildWithTimeout(ctx, d.Timeout(schema.TimeoutCreate))
}

func ForCreateUpdate(ctx context.Context, d *schema.ResourceData) (context.Context, context.CancelFunc) {
	if d.IsNewResource() {
		return ForCreate(ctx, d)
	}

	return ForUpdate(ctx, d)
}

func ForDelete(ctx context.Context, d *schema.ResourceData) (context.Context, context.CancelFunc) {
	return buildWithTimeout(ctx, d.Timeout(schema.TimeoutDelete))
}

func ForRead(ctx context.Context, d *schema.ResourceData) (context.Context, context.CancelFunc) {
	return buildWithTimeout(ctx, d.Timeout(schema.TimeoutRead))
}

func ForUpdate(ctx context.Context, d *schema.ResourceData) (context.Context, context.CancelFunc) {
	return buildWithTimeout(ctx, d.Timeout(schema.TimeoutUpdate))
}

func buildWithTimeout(ctx context.Context, timeout time.Duration) (context.Context, context.CancelFunc) {
	if features.SupportsCustomTimeouts() {
		return context.WithTimeout(ctx, timeout)
	}

	nullFunc := func() {}
	return ctx, nullFunc
}
