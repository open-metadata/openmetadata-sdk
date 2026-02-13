package ometa

import (
	"context"
	"encoding/json"
	"fmt"
	"iter"
	"net/url"
)

type listPage[T any] struct {
	Data []T `json:"data"`
	Paging *Paging `json:"paging,omitempty"`
}

func newEntityIterator[T any](
	ctx context.Context,
	b Backend,
	path string,
	params any) iter.Seq2[T, error] {
		return func(yield func(T, error) bool) {
			qp := EncodeParams(params)
			if qp == nil {
				qp = url.Values{}
			}
			for {
				raw, err := b.Call(ctx, "GET", path, nil, qp)
				if err != nil {
					var zero T
					yield(zero, err)
					return
				}

				var page listPage[T]
				if err := json.Unmarshal(raw, &page); err != nil {
					var zero T
					yield(zero, fmt.Errorf("openmetadata: failed to decode response %w", err))
					return
				}

				for _, entity := range page.Data {
					if !yield(entity, nil) {
						return
					}
				}

				if page.Paging == nil || page.Paging.After == nil || *page.Paging.After == "" {
					return
				}
				qp.Set("after", *page.Paging.After)
			}
		}
}

