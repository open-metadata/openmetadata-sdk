package integration

import (
	"context"
	"testing"

	"github.com/open-metadata/openmetadata-sdk/openmetadata-go-client/pkg/ometa"
)

func createTestTopic(t *testing.T, ctx context.Context, name string) *ometa.Topic {
	t.Helper()

	topic, err := client.Topics.Create(ctx, &ometa.CreateTopic{
		Name:       name,
		Service:    testMessagingService,
		Partitions: 1,
	})
	if err != nil {
		t.Fatalf("failed to create topic '%s': %v", name, err)
	}

	t.Cleanup(func() {
		client.Topics.DeleteByName(ctx, *topic.FullyQualifiedName, &ometa.DeleteTopicByFQNParams{
			HardDelete: ometa.Bool(true),
		})
	})

	return topic
}

func TestCreateTopic(t *testing.T) {
	ctx := context.Background()
	topic := createTestTopic(t, ctx, "test_create_topic")

	if topic.Name != "test_create_topic" {
		t.Errorf("expected name 'test_create_topic', got '%s'", topic.Name)
	}
}

func TestGetTopicByID(t *testing.T) {
	ctx := context.Background()
	topic := createTestTopic(t, ctx, "test_get_topic_by_id")

	got, err := client.Topics.GetByID(ctx, topic.Id.String(), nil)
	if err != nil {
		t.Fatalf("failed to get topic by ID: %v", err)
	}

	if got.Id != topic.Id {
		t.Errorf("expected ID '%s', got '%s'", topic.Id, got.Id)
	}
}

func TestGetTopicByName(t *testing.T) {
	ctx := context.Background()
	topic := createTestTopic(t, ctx, "test_get_topic_by_name")

	got, err := client.Topics.GetByName(ctx, *topic.FullyQualifiedName, nil)
	if err != nil {
		t.Fatalf("failed to get topic by name: %v", err)
	}

	if got.Id != topic.Id {
		t.Errorf("expected ID '%s', got '%s'", topic.Id, got.Id)
	}
}

func TestDeleteTopic(t *testing.T) {
	ctx := context.Background()

	topic, err := client.Topics.Create(ctx, &ometa.CreateTopic{
		Name:       "test_delete_topic",
		Service:    testMessagingService,
		Partitions: 1,
	})
	if err != nil {
		t.Fatalf("failed to create topic: %v", err)
	}

	err = client.Topics.Delete(ctx, topic.Id.String(), &ometa.DeleteTopicParams{
		HardDelete: ometa.Bool(true),
	})
	if err != nil {
		t.Fatalf("failed to delete topic: %v", err)
	}

	_, err = client.Topics.GetByID(ctx, topic.Id.String(), nil)
	if !ometa.IsNotFound(err) {
		t.Errorf("expected 404 after delete, got: %v", err)
	}
}
