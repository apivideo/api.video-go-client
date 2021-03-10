package apivideosdk

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"
)

var (
	mux *http.ServeMux

	client *Client

	server *httptest.Server
)

var paginationJSON = `{
    "currentPage": 1,
    "pageSize": 25,
    "pagesTotal": 1,
    "itemsTotal": 11,
    "currentPageItems": 11,
    "links": [
      {
        "rel": "self",
        "uri": "https://ws.api.video"
      },
      {
        "rel": "first",
        "uri": "https://ws.api.video"
      },
      {
        "rel": "last",
        "uri": "https://ws.api.video"
      }
    ]
  }`

var paginationStruct = Pagination{
	CurrentPage:      PtrInt32(1),
	PageSize:         PtrInt32(25),
	PagesTotal:       PtrInt32(1),
	ItemsTotal:       PtrInt32(11),
	CurrentPageItems: PtrInt32(11),
	Links: []PaginationLink{
		{
			Rel: PtrString("self"),
			Uri: PtrString("https://ws.api.video"),
		},
		{
			Rel: PtrString("first"),
			Uri: PtrString("https://ws.api.video"),
		},
		{
			Rel: PtrString("last"),
			Uri: PtrString("https://ws.api.video"),
		},
	},
}

func setup() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	mux.HandleFunc("/auth/api-key", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{
			"token_type": "Bearer",
			"expires_in": 3600,
			"access_token": "fakeToken",
			"refresh_token": "fakeToken"
		  }`)
	})

	client = ClientBuilder("apiKey").Build()
	url, _ := url.Parse(server.URL)
	client.BaseURL = url
}

func teardown() {
	server.Close()
}

func testMethod(t *testing.T, r *http.Request, expected string) {
	if expected != r.Method {
		t.Errorf("Request method = %v, expected %v", r.Method, expected)
	}
}

func createTempFile(filename string, size int64) *os.File {
	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}

	if err := f.Truncate(size); err != nil {
		log.Fatal(err)
	}
	return f
}
