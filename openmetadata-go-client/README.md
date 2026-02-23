# OpenMetadata Go Client

## Installation

```bash
go get github.com/open-metadata/openmetadata-sdk/openmetadata-go-client@latest
```

## Quick Start

### Creating a Client

```go
import "github.com/open-metadata/openmetadata-sdk/openmetadata-go-client/pkg/ometa"

client := ometa.NewClient(
    "http://localhost:8585",
    ometa.WithToken("<JWToken>"),
)
```

### Client Options

| Option | Description | Default |
|--------|-------------|---------|
| `WithToken(token)` | Set authentication token | `""` |
| `WithAPIVersion(v)` | Override API version | `"v1"` |
| `WithRetry(count, wait)` | Configure retry behavior | `3` retries, `30s` wait |
| `WithHTTPClient(c)` | Use a custom `*http.Client` | Default `http.Client` |

Retryable status codes: `429` (Too Many Requests), `504` (Gateway Timeout).

## Usage

All entity types (Tables, Databases, Users, Glossaries, etc.) follow the same service pattern accessible from the client.

### List Entities

`List` returns an iterator (`iter.Seq2[T, error]`) for automatic pagination.

```go
for table, err := range client.Tables.List(ctx, nil) {
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(table.Name)
}
```

### Get by ID

```go
table, err := client.Tables.GetByID(ctx, "<id>", nil)
```

### Get by Fully Qualified Name

```go
table, err := client.Tables.GetByName(ctx, "service.database.schema.table", nil)
```

### Create

```go
table, err := client.Tables.Create(ctx, &ometa.CreateTable{
    Name:           "my_table",
    DatabaseSchema: "service.database.schema",
    Columns: []ometa.Column{
        {Name: "id", DataType: ometa.ColumnDataTypeINT},
        {Name: "name", DataType: ometa.ColumnDataTypeSTRING},
    },
})
```

### Create or Update

```go
table, err := client.Tables.CreateOrUpdate(ctx, &ometa.CreateTable{
    Name:           "my_table",
    DatabaseSchema: "service.database.schema",
    Columns: []ometa.Column{
        {Name: "id", DataType: ometa.ColumnDataTypeINT},
    },
})
```

### Patch (JSON Patch)

```go
table, err := client.Tables.Patch(ctx, "<id>", []ometa.JSONPatchOp{
    {Op: "add", Path: "/description", Value: "Updated description"},
})
```

### Delete

```go
err := client.Tables.Delete(ctx, "<id>", nil)
```

### Delete by Name

```go
err := client.Tables.DeleteByName(ctx, "service.database.schema.table", nil)
```

### Entity Versions

```go
history, err := client.Tables.ListVersions(ctx, "<id>")

table, err := client.Tables.GetVersion(ctx, "<id>", "0.2")
```

### Restore a Soft-Deleted Entity

```go
table, err := client.Tables.Restore(ctx, &ometa.RestoreEntity{
    Id: ometa.Str("<id>"),
})
```

## Pointer Helpers

Many optional fields in request/response structs use pointers. Use the provided helpers:

```go
ometa.Str("value")     // *string
ometa.Bool(true)        // *bool
ometa.Int32(42)         // *int32
ometa.Int64(42)         // *int64
ometa.Float64(3.14)     // *float64
ometa.Ptr(myValue)      // *T (generic)
```

## Available Services

The client exposes 50+ typed services including:

- **Data Assets**: `Tables`, `Topics`, `Dashboards`, `Pipelines`, `Charts`, `Containers`, `SearchIndexes`, `Metrics`, `StoredProcedures`
- **Data Organization**: `Databases`, `DatabaseSchemas`, `APICollections`, `APIEndpoints`
- **Governance**: `Glossaries`, `GlossaryTerms`, `Classifications`, `Tags`, `DataProducts`, `Domains`, `DataContracts`, `Policies`
- **Quality**: `TestCases`, `TestSuites`, `TestDefinitions`
- **Services**: `DatabaseServices`, `DashboardServices`, `PipelineServices`, `MessagingServices`, `SearchServices`, `StorageServices`, `LLMServices`
- **People**: `Users`, `Teams`, `Roles`, `Personas`, `Bots`
- **Other**: `IngestionPipelines`, `EventSubscriptions`, `WorkflowDefinitions`, `Documents`

All services share the same method pattern: `List`, `GetByID`, `GetByName`, `Create`, `CreateOrUpdate`, `Patch`, `Delete`, `DeleteByName`, `ListVersions`, `GetVersion`, `Restore`.
