package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRouter(t *testing.T) {
	r := getRouter()
	mockServer := httptest.NewServer(r)
	resp, err := http.Get(mockServer.URL + "/api/v1/products")

	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Status should be ok, got %d", resp.StatusCode)
	}

	defer resp.Body.Close()
}

func TestRouterForNonExistentRoute(t *testing.T) {
	r := getRouter()
	mockServer := httptest.NewServer(r)
	resp, err := http.Post(mockServer.URL+"/api/v1/products", "", nil)

	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("Status should be 405, got %d", resp.StatusCode)
	}

	defer resp.Body.Close()
}

func TestJsonProductsDataRouterContentType(t *testing.T) {
	r := getRouter()
	mockServer := httptest.NewServer(r)

	resp, err := http.Get(mockServer.URL + "/api/v1/products")
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Status should be 200, got %d", resp.StatusCode)
	}

	contentType := resp.Header.Get("Content-Type")
	expectedContentType := "application/json"

	if expectedContentType != contentType {
		t.Errorf("Wrong content type, expected %s, got %s",
			expectedContentType, contentType)
	}

	defer resp.Body.Close()
}

func TestJsonProductDataRouterContentType(t *testing.T) {
	r := getRouter()
	mockServer := httptest.NewServer(r)

	resp, err := http.Get(mockServer.URL + "/api/v1/product/" + "0000001")
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Status should be 200, got %d", resp.StatusCode)
	}

	contentType := resp.Header.Get("Content-Type")
	expectedContentType := "application/json"

	if expectedContentType != contentType {
		t.Errorf("Wrong content type, expected %s, got %s",
			expectedContentType, contentType)
	}

	defer resp.Body.Close()
}

func TestJsonProductDataRouterContentDataNotEmpty(t *testing.T) {
	r := getRouter()
	mockServer := httptest.NewServer(r)

	resp, err := http.Get(mockServer.URL + "/api/v1/product/0000001")
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Status should be 200, got %d", resp.StatusCode)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("Error when trying get body.")
	}
	bodyString := string(bodyBytes)
	if bodyString == "" {
		t.Errorf("Status should be not empty %s", bodyString)
	}

	defer resp.Body.Close()
}
