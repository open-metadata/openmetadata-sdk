package integration

import (
	"context"
	"testing"

	"github.com/open-metadata/openmetadata-sdk/openmetadata-go-client/pkg/ometa"
)

func createTestPolicy(t *testing.T, ctx context.Context, name string) *ometa.Policy {
	t.Helper()

	policy, err := client.Policies.Create(ctx, &ometa.CreatePolicy{
		Name: name,
		Rules: []ometa.Rule{
			{
				Name:       "test_rule",
				Effect:     "allow",
				Operations: []ometa.RuleOperations{"ViewAll"},
				Resources:  []string{"All"},
			},
		},
	})
	if err != nil {
		t.Fatalf("failed to create policy '%s': %v", name, err)
	}

	t.Cleanup(func() {
		_ = client.Policies.DeleteByName(ctx, *policy.FullyQualifiedName, &ometa.DeletePolicyByFQNParams{
			HardDelete: ometa.Bool(true),
		})
	})

	return policy
}

func TestCreatePolicy(t *testing.T) {
	ctx := context.Background()
	policy := createTestPolicy(t, ctx, "test_create_policy")

	if policy.Name != "test_create_policy" {
		t.Errorf("expected name 'test_create_policy', got '%s'", policy.Name)
	}
}

func TestGetPolicyByID(t *testing.T) {
	ctx := context.Background()
	policy := createTestPolicy(t, ctx, "test_get_policy_by_id")

	got, err := client.Policies.GetByID(ctx, policy.Id.String(), nil)
	if err != nil {
		t.Fatalf("failed to get policy by ID: %v", err)
	}

	if got.Id != policy.Id {
		t.Errorf("expected ID '%s', got '%s'", policy.Id, got.Id)
	}
}

func TestGetPolicyByName(t *testing.T) {
	ctx := context.Background()
	policy := createTestPolicy(t, ctx, "test_get_policy_by_name")

	got, err := client.Policies.GetByName(ctx, *policy.FullyQualifiedName, nil)
	if err != nil {
		t.Fatalf("failed to get policy by name: %v", err)
	}

	if got.Id != policy.Id {
		t.Errorf("expected ID '%s', got '%s'", policy.Id, got.Id)
	}
}

func TestDeletePolicy(t *testing.T) {
	ctx := context.Background()

	policy, err := client.Policies.Create(ctx, &ometa.CreatePolicy{
		Name: "test_delete_policy",
		Rules: []ometa.Rule{
			{
				Name:       "test_rule",
				Effect:     "allow",
				Operations: []ometa.RuleOperations{"ViewAll"},
				Resources:  []string{"All"},
			},
		},
	})
	if err != nil {
		t.Fatalf("failed to create policy: %v", err)
	}

	err = client.Policies.Delete(ctx, policy.Id.String(), &ometa.DeletePolicyParams{
		HardDelete: ometa.Bool(true),
	})
	if err != nil {
		t.Fatalf("failed to delete policy: %v", err)
	}

	_, err = client.Policies.GetByID(ctx, policy.Id.String(), nil)
	if !ometa.IsNotFound(err) {
		t.Errorf("expected 404 after delete, got: %v", err)
	}
}
