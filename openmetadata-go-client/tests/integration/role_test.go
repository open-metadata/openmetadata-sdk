package integration

import (
	"context"
	"testing"

	"github.com/open-metadata/openmetadata-sdk/openmetadata-go-client/pkg/ometa"
)

func createTestRole(t *testing.T, ctx context.Context, name string) *ometa.Role {
	t.Helper()

	policy := createTestPolicy(t, ctx, name+"_policy")

	role, err := client.Roles.Create(ctx, &ometa.CreateRole{
		Name:     name,
		Policies: []string{*policy.FullyQualifiedName},
	})
	if err != nil {
		t.Fatalf("failed to create role '%s': %v", name, err)
	}

	t.Cleanup(func() {
		_ = client.Roles.DeleteByName(ctx, *role.FullyQualifiedName, &ometa.DeleteRoleByNameParams{
			HardDelete: ometa.Bool(true),
		})
	})

	return role
}

func TestCreateRole(t *testing.T) {
	ctx := context.Background()
	role := createTestRole(t, ctx, "test_create_role")

	if role.Name != "test_create_role" {
		t.Errorf("expected name 'test_create_role', got '%s'", role.Name)
	}
}

func TestGetRoleByID(t *testing.T) {
	ctx := context.Background()
	role := createTestRole(t, ctx, "test_get_role_by_id")

	got, err := client.Roles.GetByID(ctx, role.Id.String(), nil)
	if err != nil {
		t.Fatalf("failed to get role by ID: %v", err)
	}

	if got.Id != role.Id {
		t.Errorf("expected ID '%s', got '%s'", role.Id, got.Id)
	}
}

func TestGetRoleByName(t *testing.T) {
	ctx := context.Background()
	role := createTestRole(t, ctx, "test_get_role_by_name")

	got, err := client.Roles.GetByName(ctx, *role.FullyQualifiedName, nil)
	if err != nil {
		t.Fatalf("failed to get role by name: %v", err)
	}

	if got.Id != role.Id {
		t.Errorf("expected ID '%s', got '%s'", role.Id, got.Id)
	}
}

func TestDeleteRole(t *testing.T) {
	ctx := context.Background()

	policy := createTestPolicy(t, ctx, "test_delete_role_policy")
	role, err := client.Roles.Create(ctx, &ometa.CreateRole{
		Name:     "test_delete_role",
		Policies: []string{*policy.FullyQualifiedName},
	})
	if err != nil {
		t.Fatalf("failed to create role: %v", err)
	}

	err = client.Roles.Delete(ctx, role.Id.String(), &ometa.DeleteRoleParams{
		HardDelete: ometa.Bool(true),
	})
	if err != nil {
		t.Fatalf("failed to delete role: %v", err)
	}

	_, err = client.Roles.GetByID(ctx, role.Id.String(), nil)
	if !ometa.IsNotFound(err) {
		t.Errorf("expected 404 after delete, got: %v", err)
	}
}
