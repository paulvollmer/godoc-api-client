package godoc

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSearch(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		if req.URL.String() != "/search?q=foo" {
			t.Error("Request URL not equal")
		}
		// Send response to be tested
		rw.Write([]byte(`{
			"results": [
			  {
					"name": "test",
					"path": "testing",
					"import_count": 1,
					"synopsis": "Package testing",
					"score": 0.99
				}
			]
		}`))
	}))
	// Close the server when test finishes
	defer server.Close()

	// Use Client & URL from our local test server
	api := API{server.Client(), server.URL}
	results, res, err := api.Search("foo")
	if err != nil {
		t.Error(err)
	}
	if res.StatusCode != 200 {
		t.Error("StatusCode should be 200")
	}
	if len(results.Results) != 1 {
		t.Error("Results should be length of 1")
	}
	if results.Results[0].Name != "test" {
		t.Error("Result Name should be length of test")
	}
}

func TestPackages(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if req.URL.String() != "/packages" {
			t.Error("Request URL not equal")
		}
		rw.Write([]byte(`{
			"results": [
			  {
					"name": "test",
					"path": "testing",
					"import_count": 1,
					"synopsis": "Package testing",
					"score": 0.99
				}
			]
		}`))
	}))
	defer server.Close()

	api := API{server.Client(), server.URL}
	results, res, err := api.Packages()
	if err != nil {
		t.Error(err)
	}
	if res.StatusCode != 200 {
		t.Error("StatusCode should be 200")
	}
	if len(results.Results) != 1 {
		t.Error("Results should be length of 1")
	}
	if results.Results[0].Name != "test" {
		t.Error("Result Name should be length of test")
	}
}

func TestImporters(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if req.URL.String() != "/importers/foo" {
			t.Error("Request URL not equal")
		}
		rw.Write([]byte(`{
			"results": [
				{
					"path": "testing",
					"import_count": 1,
					"synopsis": "Package testing"
				}
			]
		}`))
	}))
	defer server.Close()

	api := API{server.Client(), server.URL}
	results, res, err := api.Importers("foo")
	if err != nil {
		t.Error(err)
	}
	if res.StatusCode != 200 {
		t.Error("StatusCode should be 200")
	}
	if len(results.Results) != 1 {
		t.Error("Results should be length of 1")
	}
	if results.Results[0].Path != "testing" {
		t.Error("Result Path should be length of test")
	}
}

func TestImporter(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if req.URL.String() != "/imports/foo" {
			t.Error("Request URL not equal")
		}
		rw.Write([]byte(`{
			"imports": [
				{
					"path": "bytes",
					"import_count": 0,
					"synopsis": "Package bytes implements functions for the manipulation of byte slices."
				}
			],
			"testImports": [
				{
					"path": "testing",
					"import_count": 0,
					"synopsis": "Package testing provides support for automated testing of Go packages."
				}
			]
		}`))
	}))
	defer server.Close()

	api := API{server.Client(), server.URL}
	results, res, err := api.Imports("foo")
	if err != nil {
		t.Error(err)
	}
	if res.StatusCode != 200 {
		t.Error("StatusCode should be 200")
	}
	if len(results.Imports) != 1 {
		t.Error("Results should be length of 1")
	}
	if results.Imports[0].Path != "bytes" {
		t.Error("Result Path should be length of test")
	}
	if len(results.TestImports) != 1 {
		t.Error("Results should be length of 1")
	}
	if results.TestImports[0].Path != "testing" {
		t.Error("Result Path should be length of test")
	}
}
