// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package internal // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/internal"

import (
	"context"
	"fmt"

	"go.opentelemetry.io/collector/pdata/pcommon"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
)

func GetSliceValue[K any](ctx context.Context, tCtx K, s pcommon.Slice, key ottl.Key[K]) (any, error) {
	if key == nil {
		return nil, fmt.Errorf("cannot get slice value without key")
	}

	i, err := key.Int(ctx, tCtx)
	if err != nil {
		return nil, err
	}
	if i == nil {
		return nil, fmt.Errorf("non-integer indexing is not supported")
	}

	idx := int(*i)

	if idx < 0 || idx >= s.Len() {
		return nil, fmt.Errorf("index %d out of bounds", idx)
	}

	return getIndexableValue[K](ctx, tCtx, s.At(idx), key.Next())
}

func SetSliceValue[K any](ctx context.Context, tCtx K, s pcommon.Slice, key ottl.Key[K], val any) error {
	if key == nil {
		return fmt.Errorf("cannot set slice value without key")
	}

	i, err := key.Int(ctx, tCtx)
	if err != nil {
		return err
	}
	if i == nil {
		return fmt.Errorf("non-integer indexing is not supported")
	}

	idx := int(*i)

	if idx < 0 || idx >= s.Len() {
		return fmt.Errorf("index %d out of bounds", idx)
	}

	return setIndexableValue[K](ctx, tCtx, s.At(idx), val, key.Next())
}
