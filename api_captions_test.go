package apivideosdk

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"reflect"
	"testing"
)

var captionJSONResponses = []string{
	`{
		"uri": "/videos/vi2ZEQZrOQckdYZ3X5sjPse8/captions/en",
		"src": "https://cdn.api.video/vod/vi2ZEQZrOQckdYZ3X5sjPse8/captions/en.vtt",
		"srclang": "en",
		"default": false
	}`,
	`{
		"uri": "/videos/vi2ZEQZrOQckdYZ3X5sjPse8/captions/fr",
		"src": "https://cdn.api.video/vod/vi2ZEQZrOQckdYZ3X5sjPse8/captions/fr.vtt",
		"srclang": "fr",
		"default": false
	}`,
}

var captionStructs = []Caption{
	{
		Uri:     PtrString("/videos/vi2ZEQZrOQckdYZ3X5sjPse8/captions/en"),
		Src:     PtrString("https://cdn.api.video/vod/vi2ZEQZrOQckdYZ3X5sjPse8/captions/en.vtt"),
		Srclang: PtrString("en"),
		Default: PtrBool(false),
	},
	{
		Uri:     PtrString("/videos/vi2ZEQZrOQckdYZ3X5sjPse8/captions/fr"),
		Src:     PtrString("https://cdn.api.video/vod/vi2ZEQZrOQckdYZ3X5sjPse8/captions/fr.vtt"),
		Srclang: PtrString("fr"),
		Default: PtrBool(false),
	},
}

func TestCaptions_Get(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/videos/vi2ZEQZrOQckdYZ3X5sjPse8/captions/en", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, captionJSONResponses[0])
	})

	caption, err := client.Captions.Get("vi2ZEQZrOQckdYZ3X5sjPse8", "en")
	if err != nil {
		t.Errorf("Captions.Get error: %v", err)
	}

	expected := &captionStructs[0]
	if !reflect.DeepEqual(caption, expected) {
		t.Errorf("Captions.Get\n got=%#v\nwant=%#v", caption, expected)
	}
}

func TestCaptions_List(t *testing.T) {
	setup()
	defer teardown()
	JSONResp := fmt.Sprintf(
		`{"data":[%s,%s], "pagination":%s}`,
		captionJSONResponses[0],
		captionJSONResponses[1],
		paginationJSON)

	mux.HandleFunc("/videos/vi2ZEQZrOQckdYZ3X5sjPse8/captions", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, JSONResp)
	})

	captions, err := client.Captions.List("vi2ZEQZrOQckdYZ3X5sjPse8", CaptionsApiListRequest{})
	if err != nil {
		t.Errorf("Captions.List error: %v", err)
	}

	expected := &CaptionsListResponse{
		Data:       &captionStructs,
		Pagination: &paginationStruct,
	}
	if !reflect.DeepEqual(captions, expected) {
		t.Errorf("Captions.List\n got=%#v\nwant=%#v", captions, expected)
	}
}

func TestCaptions_Upload(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/videos/vi2ZEQZrOQckdYZ3X5sjPse8/captions/en", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, captionJSONResponses[0])
	})

	file := createTempFile("test.caption", 1024*1024)

	defer os.Remove(file.Name())

	caption, err := client.Captions.UploadFile("vi2ZEQZrOQckdYZ3X5sjPse8", "en", file)
	if err != nil {
		t.Errorf("Captions.Upload error: %v", err)
	}

	expected := &captionStructs[0]
	if !reflect.DeepEqual(caption, expected) {
		t.Errorf("Captions.Upload\n got=%#v\nwant=%#v", caption, expected)
	}
}

func TestCaptions_Update(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/videos/vi2ZEQZrOQckdYZ3X5sjPse8/captions/en", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPatch)
		expectedBody := map[string]interface{}{
			"default": true,
		}
		var v map[string]interface{}
		err := json.NewDecoder(r.Body).Decode(&v)
		if err != nil {
			t.Fatalf("decode json: %v", err)
		}

		if !reflect.DeepEqual(v, expectedBody) {
			t.Errorf("Request body\n got=%#v\nwant=%#v", v, expectedBody)
		}
		fmt.Fprint(w, captionJSONResponses[0])
	})

	caption, err := client.Captions.Update("vi2ZEQZrOQckdYZ3X5sjPse8", "en", CaptionsUpdatePayload{
		Default: PtrBool(true),
	})
	if err != nil {
		t.Errorf("Captions.Upload error: %v", err)
	}

	expected := &captionStructs[0]
	if !reflect.DeepEqual(caption, expected) {
		t.Errorf("Captions.Upload\n got=%#v\nwant=%#v", caption, expected)
	}
}

func TestCaptions_Delete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/videos/vi2ZEQZrOQckdYZ3X5sjPse8/captions/en", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	err := client.Captions.Delete("vi2ZEQZrOQckdYZ3X5sjPse8", "en")
	if err != nil {
		t.Errorf("Captions.Delete error: %v", err)
	}
}
