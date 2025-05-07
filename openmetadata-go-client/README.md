# OpenMetadata Go Client
## Quick Start
### Instantiating a Client
```go
NewRestConfig(
    "http://localhost:8585",
    "",
    0,
    0,
    nil,
    "<JWToken>",
)

rest := ometa.NewRest(restConfig)
```

`NewRestConfig` has the following signature
```go
func NewRestConfig(
	baseURL string,
	apiVersion string, // optional - default: v1 - use "" to set default
	retry int8, // optional - default: 3 - use 0 to set default
	retryWait int16, // optional - default: 30 - use 0 to set default
	retryCodes []int, // optional - default: [429, 504] - use nil to set default
	accessToken string)
```

### Making a request
`Rest` struct has the following interface
```go
type Rest struct {
	Get()
    Post()
    Put()
    Patch()
    Delete()
}
```

Each method has the following signature

```go
// Get Method to make GET calls
func (rest Rest) Get(
	path string,
	header http.Header,
	extraHeader map[string][]string,
	queryParams map[string]string) (map[string]any, error)

body, err := restFixture.Get(path, nil, nil, nil)

// Post Method to make POST calls
func (rest Rest) Post(
	path string,
	data map[string]interface{},
	header http.Header,
	extraHeader map[string][]string,
	queryParams map[string]string) (map[string]any, error)

restFixture.Post(path, testdata.Data, nil, nil, nil)

// Put Method to make PUT calls
func (rest Rest) Put(
	path string,
	data map[string]interface{},
	header http.Header,
	extraHeader map[string][]string,
	queryParams map[string]string) (map[string]any, error)

restFixture.Put(path, data, nil, nil, nil)

// Patch Method to make PATCH calls
func (rest Rest) Patch(
	path string,
	data []map[string]interface{},
	header http.Header,
	extraHeader map[string][]string,
	queryParams map[string]string) (map[string]any, error)

restFixture.Patch(patchPath, data, nil, nil, nil)

// Delete Method to make DELETE calls
func (rest Rest) Delete(
	path string,
	header http.Header,
	extraHeader map[string][]string,
	queryParams map[string]string) (map[string]any, error)

restFixture.Delete(path, nil, nil, queryParams)
```

### SerDeser
response and request can respectively be deserialized and serialized into Go struct

```go
    body, err := restFixture.Get(path, nil, nil, nil)
    if err != nil {
		t.Fatal(err)
	}
	data, ok := body["data"].([]interface{})
	if !ok {
		t.Fatal("data key should exist")
	}
	tables, err := ometa.DeserializeSlice(data, &ometa.Table{})
    ...
    tableSlice, ok := tables.([]ometa.Table)

	addLineage := ometa.AddLineage{
		Edge: ometa.EntitiesEdge{
			FromEntity: ometa.EntityReference{
				Id:   fromTable.Id,
				Type: "table",
			},
			ToEntity: ometa.EntityReference{
				Id:   toTable.Id,
				Type: "table",
			},
		},
	}
    payload, err := ometa.Serialize(addLineage)
	if err != nil {
		t.Fatal(err)
	}
	_, err = restFixture.Put(path, payload, nil, nil, nil)
```

Note that SerDeser have a variation to perform the action on a single or a list of object:
- `Deserialize`
- `Deserialize`
- `Serialize`
- `SerializeSlice`
