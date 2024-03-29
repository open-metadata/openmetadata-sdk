package testdata

import (
	"github.com/open-metadata/openmetadata-sdk/openmetadata-go-client/pkg/ometa"
)

var Data = map[string]interface{}{
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
	},
}

func RestFixture() *ometa.Rest {
	restConfig := ometa.NewRestConfig(
		"http://localhost:8585",
		"",
		0,
		0,
		nil,
		"eyJraWQiOiJHYjM4OWEtOWY3Ni1nZGpzLWE5MmotMDI0MmJrOTQzNTYiLCJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJzdWIiOiJhZG1pbiIsImlzQm90IjpmYWxzZSwiaXNzIjoib3Blbi1tZXRhZGF0YS5vcmciLCJpYXQiOjE2NjM5Mzg0NjIsImVtYWlsIjoiYWRtaW5Ab3Blbm1ldGFkYXRhLm9yZyJ9.tS8um_5DKu7HgzGBzS1VTA5uUjKWOCU0B_j08WXBiEC0mr0zNREkqVfwFDD-d24HlNEbrqioLsBuFRiwIWKc1m_ZlVQbG7P36RUxhuv2vbSp80FKyNM-Tj93FDzq91jsyNmsQhyNv_fNr3TXfzzSPjHt8Go0FMMP66weoKMgW2PbXlhVKwEuXUHyakLLzewm9UMeQaEiRzhiTMU3UkLXcKbYEJJvfNFcLwSl9W8JCO_l0Yj3ud-qt_nQYEZwqW6u5nfdQllN133iikV4fM5QZsMCnm8Rq1mvLR0y9bmJiD7fwM1tmJ791TUWqmKaTnP49U493VanKpUAfzIiOiIbhg",
	)

	rest := ometa.NewRest(restConfig)
	return rest
}
