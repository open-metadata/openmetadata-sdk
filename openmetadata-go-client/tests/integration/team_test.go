package integration

import (
	"context"
	"testing"

	"github.com/open-metadata/openmetadata-sdk/openmetadata-go-client/pkg/ometa"
)

func createTestTeam(t *testing.T, ctx context.Context, name string) *ometa.Team {
	t.Helper()

	team, err := client.Teams.Create(ctx, &ometa.CreateTeam{
		Name:     name,
		TeamType: "Group",
	})
	if err != nil {
		t.Fatalf("failed to create team '%s': %v", name, err)
	}

	t.Cleanup(func() {
		client.Teams.DeleteByName(ctx, *team.FullyQualifiedName, &ometa.DeleteTeamByNameParams{
			HardDelete: ometa.Bool(true),
		})
	})

	return team
}

func TestCreateTeam(t *testing.T) {
	ctx := context.Background()
	team := createTestTeam(t, ctx, "test_create_team")

	if team.Name != "test_create_team" {
		t.Errorf("expected name 'test_create_team', got '%s'", team.Name)
	}
}

func TestGetTeamByID(t *testing.T) {
	ctx := context.Background()
	team := createTestTeam(t, ctx, "test_get_team_by_id")

	got, err := client.Teams.GetByID(ctx, team.Id.String(), nil)
	if err != nil {
		t.Fatalf("failed to get team by ID: %v", err)
	}

	if got.Id != team.Id {
		t.Errorf("expected ID '%s', got '%s'", team.Id, got.Id)
	}
}

func TestGetTeamByName(t *testing.T) {
	ctx := context.Background()
	team := createTestTeam(t, ctx, "test_get_team_by_name")

	got, err := client.Teams.GetByName(ctx, *team.FullyQualifiedName, nil)
	if err != nil {
		t.Fatalf("failed to get team by name: %v", err)
	}

	if got.Id != team.Id {
		t.Errorf("expected ID '%s', got '%s'", team.Id, got.Id)
	}
}

func TestListTeams(t *testing.T) {
	ctx := context.Background()
	team := createTestTeam(t, ctx, "test_list_team")

	found := false
	for tm, err := range client.Teams.List(ctx, &ometa.ListTeamsParams{
		Limit: ometa.Int32(100),
	}) {
		if err != nil {
			t.Fatalf("error during list: %v", err)
		}
		if tm.Id == team.Id {
			found = true
			break
		}
	}

	if !found {
		t.Errorf("created team not found in list results")
	}
}

func TestDeleteTeam(t *testing.T) {
	ctx := context.Background()

	team, err := client.Teams.Create(ctx, &ometa.CreateTeam{
		Name:     "test_delete_team",
		TeamType: "Group",
	})
	if err != nil {
		t.Fatalf("failed to create team: %v", err)
	}

	err = client.Teams.Delete(ctx, team.Id.String(), &ometa.DeleteTeamParams{
		HardDelete: ometa.Bool(true),
	})
	if err != nil {
		t.Fatalf("failed to delete team: %v", err)
	}

	_, err = client.Teams.GetByID(ctx, team.Id.String(), nil)
	if !ometa.IsNotFound(err) {
		t.Errorf("expected 404 after delete, got: %v", err)
	}
}
