package ometa

import (
	"context"
	"encoding/json"
	"fmt"
)

type JSONPatchOp struct {
	Op    string      `json:"op"`
	Path  string      `json:"path"`
	Value interface{} `json:"value"`
	From string      `json:"from,omitempty"`
}

func get[T any](
	ctx context.Context,
	b Backend,
	path string,
	params any) (*T, error) {
		qp := EncodeParams(params)
		raw, err := b.Call(ctx, "GET", path, nil, qp)
		if err != nil {
			return nil, err
		}
		if raw == nil {
			return nil, nil
		}
		var result T
		if err := json.Unmarshal(raw, &result); err != nil {
			return nil, fmt.Errorf("openmetadata: failed to decode response %w", err)
		}

		return &result, nil
}

func put[T any](
	ctx context.Context,
	b Backend,
	path string,
	body any) (*T, error) {
		data, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("openmetadata: failed to encode request body: %w", err)
		}
		raw, err := b.Call(ctx, "PUT", path, data, nil)
		if err != nil {
			return nil, err
		}
		if raw == nil {
			return nil, nil
		}
		var result T
		if err := json.Unmarshal(raw, &result); err != nil {
			return nil, fmt.Errorf("openmetadata: failed to decode response %w", err)
		}
		
		return &result, nil
}

func create[T any](
	ctx context.Context,
	b Backend,
	path string,
	body any) (*T, error) {
		data, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("openmetadata: failed to encode request body: %w", err)
		}
		raw, err := b.Call(ctx, "POST", path, data, nil)
		if err != nil {
			return nil, err
		}
		if raw == nil {
			return nil, nil
		}
		var result T
		if err := json.Unmarshal(raw, &result); err != nil {
			return nil, fmt.Errorf("openmetadata: failed to decode response %w", err)
		}
		
		return &result, nil
}

func patch[T any](
	ctx context.Context,
	b Backend,
	path string,
	ops []JSONPatchOp) (*T, error) {
		data, err := json.Marshal(ops)
		if err != nil {
			return nil, fmt.Errorf("openmetadata: failed to encode request body: %w", err)
		}
		raw, err := b.Call(ctx, "PATCH", path, data, nil)
		if err != nil {
			return nil, err
		}
		if raw == nil {
			return nil, nil
		}
		var result T
		if err := json.Unmarshal(raw, &result); err != nil {
			return nil, fmt.Errorf("openmetadata: failed to decode response %w", err)
		}
		
		return &result, nil
}

func delete(
	ctx context.Context,
	b Backend,
	path string) error {
		_, err := b.Call(ctx, "DELETE", path, nil, nil)
		return err
}
