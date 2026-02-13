package integration

import (
	"context"
	"testing"

	"github.com/open-metadata/openmetadata-sdk/openmetadata-go-client/pkg/ometa"
)

func createTestUser(t *testing.T, ctx context.Context, name string) *ometa.User {
	t.Helper()

	user, err := client.Users.Create(ctx, &ometa.CreateUser{
		Name:  name,
		Email: name + "@test.openmetadata.org",
	})
	if err != nil {
		t.Fatalf("failed to create user '%s': %v", name, err)
	}

	t.Cleanup(func() {
		_ = client.Users.DeleteByName(ctx, *user.FullyQualifiedName, &ometa.DeleteUserByNameParams{
			HardDelete: ometa.Bool(true),
		})
	})

	return user
}

func TestCreateUser(t *testing.T) {
	ctx := context.Background()
	user := createTestUser(t, ctx, "test_create_user")

	if user.Name != "test_create_user" {
		t.Errorf("expected name 'test_create_user', got '%s'", user.Name)
	}
}

func TestGetUserByID(t *testing.T) {
	ctx := context.Background()
	user := createTestUser(t, ctx, "test_get_user_by_id")

	got, err := client.Users.GetByID(ctx, user.Id.String(), nil)
	if err != nil {
		t.Fatalf("failed to get user by ID: %v", err)
	}

	if got.Id != user.Id {
		t.Errorf("expected ID '%s', got '%s'", user.Id, got.Id)
	}
}

func TestGetUserByName(t *testing.T) {
	ctx := context.Background()
	user := createTestUser(t, ctx, "test_get_user_by_name")

	got, err := client.Users.GetByName(ctx, *user.FullyQualifiedName, nil)
	if err != nil {
		t.Fatalf("failed to get user by name: %v", err)
	}

	if got.Id != user.Id {
		t.Errorf("expected ID '%s', got '%s'", user.Id, got.Id)
	}
}

func TestListUsers(t *testing.T) {
	ctx := context.Background()
	user := createTestUser(t, ctx, "test_list_user")

	found := false
	for u, err := range client.Users.List(ctx, &ometa.ListUsersParams{
		Limit: ometa.Int32(100),
	}) {
		if err != nil {
			t.Fatalf("error during list: %v", err)
		}
		if u.Id == user.Id {
			found = true
			break
		}
	}

	if !found {
		t.Errorf("created user not found in list results")
	}
}

func TestDeleteUser(t *testing.T) {
	ctx := context.Background()

	user, err := client.Users.Create(ctx, &ometa.CreateUser{
		Name:  "test_delete_user",
		Email: "test_delete_user@test.openmetadata.org",
	})
	if err != nil {
		t.Fatalf("failed to create user: %v", err)
	}

	err = client.Users.Delete(ctx, user.Id.String(), &ometa.DeleteUserParams{
		HardDelete: ometa.Bool(true),
	})
	if err != nil {
		t.Fatalf("failed to delete user: %v", err)
	}

	_, err = client.Users.GetByID(ctx, user.Id.String(), nil)
	if !ometa.IsNotFound(err) {
		t.Errorf("expected 404 after delete, got: %v", err)
	}
}
