package ometa_test

import (
	"fmt"
	"testing"

	"github.com/open-metadata/openmetadata-sdk/openmetadata-go-client/pkg/ometa"
	"github.com/open-metadata/openmetadata-sdk/openmetadata-go-client/pkg/ometa/testdata"
	"github.com/stretchr/testify/assert"
)

func TestListTables(t *testing.T) {
	restFixture := testdata.RestFixture()
	path := "tables"
	body, err := restFixture.Get(path, nil, nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	data, ok := body["data"].([]interface{})
	assert.Equal(t, ok, true, "data key should exist")
	assert.NotNil(t, data, "body[data] should not be nil")
	assert.Equal(t, len(data), 10, "lenght of body[data] should be 10")

	tables, err := ometa.DeserializeSlice(data, &ometa.Table{})
	if err != nil {
		t.Fatal(err)
	}

	tableSlice, ok := tables.([]ometa.Table)
	assert.True(t, ok, "table should be of type []ometa.Table")
	assert.Equal(t, len(tableSlice), 10, "lenght of table should be 10")
}

func TestListTablesWithQueryParams(t *testing.T) {
	restFixture := testdata.RestFixture()
	path := "tables"
	queryParams := map[string]string{
		"limit":  "2",
		"fields": "tableConstraints",
	}
	body, err := restFixture.Get(path, nil, nil, queryParams)
	if err != nil {
		t.Fatal(err)
	}
	data, ok := body["data"].([]interface{})
	assert.True(t, ok, "data key should exist")
	assert.NotNil(t, data, "body[data] should not be nil")
	assert.Equal(t, len(data), 2, "lenght of body[data] should be 2")

	tableOne := data[0].(map[string]interface{})
	_, ok = tableOne["tableConstraints"]
	assert.True(t, ok, "tableConstraints key should exist")
}

func TestGetTableByName(t *testing.T) {
	restFixture := testdata.RestFixture()
	path := `tables/name/sample_data.ecommerce_db.shopify.%22dim.shop%22`
	body, err := restFixture.Get(path, nil, nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	name, ok := body["name"].(string)
	assert.True(t, ok, "`name` key should exist")
	assert.NotNil(t, body, "body should not be nil")
	assert.Equal(t, name, "dim.shop", "name should be `dim.shop`")
	table, err := ometa.Deserialize(body, &ometa.Table{})
	if err != nil {
		t.Fatal(err)
	}

	table, ok = table.(*ometa.Table)
	if !ok {
		t.Fatalf("Expected type *ometa.Table, got %T", table)
	}
	assert.True(t, ok, "table should be of type ometa.Table")
}

func TestCreateTable(t *testing.T) {
	restFixture := testdata.RestFixture()
	path := "tables"
	body, err := restFixture.Post(path, testdata.Data, nil, nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, body["name"], "goClientTestTable", "name should be `goClientTestTable`")
	deleteTable(body["id"].(string))
}

func TestUpdateTablePut(t *testing.T) {
	table := createTable()
	restFixture := testdata.RestFixture()
	path := "tables"
	data := map[string]interface{}{
		"name":           "goClientTestTable",
		"databaseSchema": "sample_data.ecommerce_db.shopify",
		"columns": []map[string]interface{}{
			{
				"name":     "columnOne",
				"dataType": "NUMBER",
			},
			{
				"name":     "columnTwo",
				"dataType": "NUMBER",
			},
			{
				"name":     "columnThree",
				"dataType": "NUMBER",
			},
		},
	}
	body, err := restFixture.Put(path, data, nil, nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	columns := body["columns"].([]interface{})
	assert.Equal(t, len(columns), 3, "lenght of columns should be 3")
	deleteTable(table["id"].(string))
}

func TestUpdateTablePatch(t *testing.T) {
	table := createTable()
	id := table["id"].(string)

	restFixture := testdata.RestFixture()

	patchPath := fmt.Sprintf("tables/%s", id)
	data := []map[string]interface{}{
		{
			"op":    "add",
			"path":  "/description",
			"value": "This is a test table",
		},
	}

	body, err := restFixture.Patch(patchPath, data, nil, nil, nil)
	if err != nil {
		t.Fatal(err)
	}

	description, ok := body["description"].(string)
	assert.True(t, ok, "`description` key should exist")
	assert.Equal(t, description, "This is a test table", "description should be `This is a test table`")
	deleteTable(id)
}

func TestDeleteTable(t *testing.T) {
	table := createTable()
	id := table["id"].(string)

	restFixture := testdata.RestFixture()
	path := fmt.Sprintf("tables/%s", id)
	queryParams := map[string]string{
		"hardDelete": "true",
		"recursive":  "true",
	}
	restFixture.Delete(path, nil, nil, queryParams)
	path = fmt.Sprintf("tables/%s", id)
	body, _ := restFixture.Get(path, nil, nil, nil)
	assert.Equal(t, body["code"].(float64), 404.0, "statusCode should be 404")
}

func TestAddLineage(t *testing.T) {
	restFixture := testdata.RestFixture()
	path := "tables"
	body, err := restFixture.Get(path, nil, nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	data, ok := body["data"].([]interface{})
	if !ok {
		t.Fatal("data key should exist")
	}
	tables, err := ometa.DeserializeSlice(data, &ometa.Table{})
	if err != nil {
		t.Fatal(err)
	}

	tableSlice, ok := tables.([]ometa.Table)
	if !ok {
		t.Fatal("tableSlice should be of type []ometa.Table")
	}
	fromTable := tableSlice[0]
	toTable := tableSlice[1]

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

	path = "lineage"
	payload, err := ometa.Serialize(addLineage)
	if err != nil {
		t.Fatal(err)
	}
	_, err = restFixture.Put(path, payload, nil, nil, nil)
	if err != nil {
		t.Fatal(err)
	}
}

func createTable() map[string]any {
	restFixture := testdata.RestFixture()
	path := "tables"
	body, _ := restFixture.Post(path, testdata.Data, nil, nil, nil)
	return body
}

func deleteTable(id string) map[string]any {
	restFixture := testdata.RestFixture()
	path := fmt.Sprintf("tables/%s", id)
	queryParams := map[string]string{
		"hardDelete": "true",
		"recursive":  "true",
	}
	body, _ := restFixture.Delete(path, nil, nil, queryParams)
	return body
}
