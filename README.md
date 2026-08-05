# OpenMetadata Go SDK

The official **Go client for [OpenMetadata](https://github.com/open-metadata/OpenMetadata)** — the open-source data catalog and metadata platform for data discovery, lineage, governance, data quality, and AI context.

Typed access to 50+ entity services (tables, dashboards, pipelines, glossaries, domains, data contracts, test cases, and more) with automatic pagination, retries, and JSON Patch support.

[![Go Reference](https://pkg.go.dev/badge/github.com/open-metadata/openmetadata-sdk/openmetadata-go-client.svg)](https://pkg.go.dev/github.com/open-metadata/openmetadata-sdk/openmetadata-go-client)

> **Looking for the Python SDK?** It ships in the `openmetadata-ingestion` package — see [docs.open-metadata.org/latest/sdk/python](https://docs.open-metadata.org/latest/sdk/python). For AI agents and MCP, see [ai-sdk](https://github.com/open-metadata/ai-sdk). This repository is the **Go** client.

## Install

```bash
go get github.com/open-metadata/openmetadata-sdk/openmetadata-go-client@latest
```

Client releases track OpenMetadata server versions — use the client version matching your server (server 1.13.x → client v1.13.x).

## Quick start

```go
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/open-metadata/openmetadata-sdk/openmetadata-go-client/pkg/ometa"
)

func main() {
	client := ometa.NewClient(
		"http://localhost:8585",
		ometa.WithToken("<JWToken>"),
	)

	ctx := context.Background()

	// Fetch a table by fully qualified name
	table, err := client.Tables.GetByName(ctx, "sample_data.ecommerce_db.shopify.raw_product_catalog", nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(table.Name)

	// List returns an iterator — pagination is handled for you
	for t, err := range client.Tables.List(ctx, nil) {
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(t.Name)
	}
}
```

Get a token from **Settings → Bots → ingestion-bot** in the OpenMetadata UI.

## Every service, one pattern

All 50+ services expose the same methods:

```
List · GetByID · GetByName · Create · CreateOrUpdate · Patch · Delete · DeleteByName · ListVersions · GetVersion · Restore
```

```go
// Create
table, err := client.Tables.Create(ctx, &ometa.CreateTable{
	Name:           "my_table",
	DatabaseSchema: "service.database.schema",
	Columns: []ometa.Column{
		{Name: "id", DataType: ometa.ColumnDataTypeINT},
		{Name: "name", DataType: ometa.ColumnDataTypeSTRING},
	},
})

// Patch (JSON Patch)
table, err = client.Tables.Patch(ctx, "<id>", []ometa.JSONPatchOp{
	{Op: "add", Path: "/description", Value: "Updated description"},
})
```

**Services include:** `Tables`, `Topics`, `Dashboards`, `Pipelines`, `Charts`, `Containers`, `SearchIndexes`, `Metrics`, `StoredProcedures`, `Databases`, `DatabaseSchemas`, `APICollections`, `APIEndpoints`, `Glossaries`, `GlossaryTerms`, `Classifications`, `Tags`, `DataProducts`, `Domains`, `DataContracts`, `Policies`, `TestCases`, `TestSuites`, `TestDefinitions`, `Users`, `Teams`, `Roles`, `Personas`, `Bots`, `IngestionPipelines`, `EventSubscriptions`, `WorkflowDefinitions`, and more.

## Client options

| Option | Description | Default |
|--------|-------------|---------|
| `WithToken(token)` | Authentication token | `""` |
| `WithAPIVersion(v)` | Override API version | `"v1"` |
| `WithRetry(count, wait)` | Retry behavior | 3 retries, 30s wait |
| `WithHTTPClient(c)` | Custom `*http.Client` | Default `http.Client` |

Retries apply to `429 Too Many Requests` and `504 Gateway Timeout`.

## Pointer helpers

Optional fields use pointers. Helpers are provided:

```go
ometa.Str("value")   // *string
ometa.Bool(true)     // *bool
ometa.Int32(42)      // *int32
ometa.Int64(42)      // *int64
ometa.Float64(3.14)  // *float64
ometa.Ptr(myValue)   // *T (generic)
```

## Documentation

- **Full client reference:** [`openmetadata-go-client/README.md`](openmetadata-go-client/README.md)
- **Go package docs:** https://pkg.go.dev/github.com/open-metadata/openmetadata-sdk/openmetadata-go-client
- **OpenMetadata docs:** https://docs.open-metadata.org/latest
- **REST API:** https://docs.open-metadata.org/latest/main-concepts/metadata-standard/apis
- **Main project:** https://github.com/open-metadata/OpenMetadata

## Contributing

Issues and pull requests welcome. Build and test targets are in the [`Makefile`](Makefile); run `make help` to list them.

## License

Apache License 2.0
