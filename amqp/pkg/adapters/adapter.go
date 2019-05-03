package adapters

import "context"

type Adapter interface {
	Start(ctx context.Context) error
}
