package integration

import (
	"os"
	"testing"

	"github.com/open-metadata/openmetadata-sdk/openmetadata-go-client/pkg/ometa"
)

var client *ometa.Client

const (
	testSchema           = "sample_data.ecommerce_db.shopify"
	testDatabase         = "sample_data.ecommerce_db"
	testDatabaseService  = "sample_data"
	testDashboardService = "sample_superset"
	testPipelineService  = "sample_airflow"
	testMessagingService = "sample_kafka"
)

func TestMain(m *testing.M) {
	client = ometa.NewClient(
		"http://localhost:8585",
		ometa.WithToken(os.Getenv("OM_TEST_TOKEN")),
	)
	os.Exit(m.Run())
}
