package integration

import (
	"context"
	"testing"

	"github.com/open-metadata/openmetadata-sdk/openmetadata-go-client/pkg/ometa"
)

func createTestDashboard(t *testing.T, ctx context.Context, name string) *ometa.Dashboard {
	t.Helper()

	dashboard, err := client.Dashboards.Create(ctx, &ometa.CreateDashboard{
		Name:    name,
		Service: testDashboardService,
	})
	if err != nil {
		t.Fatalf("failed to create dashboard '%s': %v", name, err)
	}

	t.Cleanup(func() {
		client.Dashboards.DeleteByName(ctx, *dashboard.FullyQualifiedName, &ometa.DeleteDashboardByFQNParams{
			HardDelete: ometa.Bool(true),
		})
	})

	return dashboard
}

func TestCreateDashboard(t *testing.T) {
	ctx := context.Background()
	dashboard := createTestDashboard(t, ctx, "test_create_dashboard")

	if dashboard.Name != "test_create_dashboard" {
		t.Errorf("expected name 'test_create_dashboard', got '%s'", dashboard.Name)
	}
}

func TestGetDashboardByID(t *testing.T) {
	ctx := context.Background()
	dashboard := createTestDashboard(t, ctx, "test_get_dashboard_by_id")

	got, err := client.Dashboards.GetByID(ctx, dashboard.Id.String(), nil)
	if err != nil {
		t.Fatalf("failed to get dashboard by ID: %v", err)
	}

	if got.Id != dashboard.Id {
		t.Errorf("expected ID '%s', got '%s'", dashboard.Id, got.Id)
	}
}

func TestGetDashboardByName(t *testing.T) {
	ctx := context.Background()
	dashboard := createTestDashboard(t, ctx, "test_get_dashboard_by_name")

	got, err := client.Dashboards.GetByName(ctx, *dashboard.FullyQualifiedName, nil)
	if err != nil {
		t.Fatalf("failed to get dashboard by name: %v", err)
	}

	if got.Id != dashboard.Id {
		t.Errorf("expected ID '%s', got '%s'", dashboard.Id, got.Id)
	}
}

func TestDeleteDashboard(t *testing.T) {
	ctx := context.Background()

	dashboard, err := client.Dashboards.Create(ctx, &ometa.CreateDashboard{
		Name:    "test_delete_dashboard",
		Service: testDashboardService,
	})
	if err != nil {
		t.Fatalf("failed to create dashboard: %v", err)
	}

	err = client.Dashboards.Delete(ctx, dashboard.Id.String(), &ometa.DeleteDashboardParams{
		HardDelete: ometa.Bool(true),
	})
	if err != nil {
		t.Fatalf("failed to delete dashboard: %v", err)
	}

	_, err = client.Dashboards.GetByID(ctx, dashboard.Id.String(), nil)
	if !ometa.IsNotFound(err) {
		t.Errorf("expected 404 after delete, got: %v", err)
	}
}
