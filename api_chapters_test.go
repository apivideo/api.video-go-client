package apivideosdk

import (
	"fmt"
	"net/http"
	"os"
	"reflect"
	"testing"
)

var chapterJSONResponses = []string{
	`{
		"uri": "/videos/vi2ZEQZrOQckdYZ3X5sjPse8/chapters/en",
		"src": "https://cdn.api.video/vod/vi2ZEQZrOQckdYZ3X5sjPse8/chapters/en.vtt",
		"language": "en"
	}`,
	`{
		"uri": "/videos/vi2ZEQZrOQckdYZ3X5sjPse8/chapters/fr",
		"src": "https://cdn.api.video/vod/vi2ZEQZrOQckdYZ3X5sjPse8/chapters/fr.vtt",
		"language": "fr"
	}`,
}

var chapterStructs = []Chapter{
	{
		Uri:      PtrString("/videos/vi2ZEQZrOQckdYZ3X5sjPse8/chapters/en"),
		Src:      PtrString("https://cdn.api.video/vod/vi2ZEQZrOQckdYZ3X5sjPse8/chapters/en.vtt"),
		Language: PtrString("en"),
	},
	{
		Uri:      PtrString("/videos/vi2ZEQZrOQckdYZ3X5sjPse8/chapters/fr"),
		Src:      PtrString("https://cdn.api.video/vod/vi2ZEQZrOQckdYZ3X5sjPse8/chapters/fr.vtt"),
		Language: PtrString("fr"),
	},
}

func TestChapters_Get(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/videos/vi2ZEQZrOQckdYZ3X5sjPse8/chapters/en", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, chapterJSONResponses[0])
	})

	chapter, err := client.Chapters.Get("vi2ZEQZrOQckdYZ3X5sjPse8", "en")
	if err != nil {
		t.Errorf("Chapters.Get error: %v", err)
	}

	expected := &chapterStructs[0]
	if !reflect.DeepEqual(chapter, expected) {
		t.Errorf("Chapters.Get\n got=%#v\nwant=%#v", chapter, expected)
	}
}

func TestChapters_List(t *testing.T) {
	setup()
	defer teardown()
	JSONResp := fmt.Sprintf(
		`{"data":[%s,%s], "pagination":%s}`,
		chapterJSONResponses[0],
		chapterJSONResponses[1],
		paginationJSON)

	mux.HandleFunc("/videos/vi2ZEQZrOQckdYZ3X5sjPse8/chapters", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, JSONResp)
	})

	chapters, err := client.Chapters.List("vi2ZEQZrOQckdYZ3X5sjPse8", ChaptersApiListRequest{})
	if err != nil {
		t.Errorf("Chapters.List error: %v", err)
	}

	expected := &ChaptersListResponse{
		Data:       &chapterStructs,
		Pagination: &paginationStruct,
	}
	if !reflect.DeepEqual(chapters, expected) {
		t.Errorf("Chapters.List\n got=%#v\nwant=%#v", chapters, expected)
	}
}

func TestChapters_Upload(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/videos/vi2ZEQZrOQckdYZ3X5sjPse8/chapters/en", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, chapterJSONResponses[0])
	})

	file := createTempFile("test.chapter", 1024*1024)
	defer os.Remove(file.Name())

	chapter, err := client.Chapters.UploadFile("vi2ZEQZrOQckdYZ3X5sjPse8", "en", file)
	if err != nil {
		t.Errorf("Chapters.Upload error: %v", err)
	}

	expected := &chapterStructs[0]
	if !reflect.DeepEqual(chapter, expected) {
		t.Errorf("Chapters.Upload\n got=%#v\nwant=%#v", chapter, expected)
	}
}

func TestChapters_Delete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/videos/vi2ZEQZrOQckdYZ3X5sjPse8/chapters/en", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	err := client.Chapters.Delete("vi2ZEQZrOQckdYZ3X5sjPse8", "en")
	if err != nil {
		t.Errorf("Chapters.Delete error: %v", err)
	}
}
