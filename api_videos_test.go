package apivideosdk

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"testing"
	"time"
)

var videoJSONResponses = []string{`{
	"videoId": "vi4k0jvEUuaTdRAEjQ4Jfagz",
	"playerId": "pl45KFKdlddgk654dspkze",
	"title": "Maths video",
	"description": "An amazing video explaining the string theory",
	"public": true,
	"panoramic": false,
	"mp4Support":true,
	"tags": [
	  "maths",
	  "string theory",
	  "video"
	],
	"metadata": [
	  {
		"key": "Author",
		"value": "John Doe"
	  },
	  {
		"key": "Format",
		"value": "Tutorial"
	  }
	],
	"publishedAt": "2019-07-14T23:36:18.598Z",
	"updatedAt": "2019-07-14T23:49:18.598Z",
	"source": {
	  "uri": "/videos/vi4k0jvEUuaTdRAEjQ4Jfagz/source"
	},
	"assets": {
	  "iframe": "<iframe src='//embed.api.video/vod/vi4k0jvEUuaTdRAEjQ4Jfagz' width='100%' height='100%' frameborder='0' scrolling='no' allowfullscreen=''></iframe>",
	  "player": "https://embed.api.video/vod/vi4k0jvEUuaTdRAEjQ4Jfagz",
	  "hls": "https://cdn.api.video/vod/vi4k0jvEUuaTdRAEjQ4Jfagz/hls/manifest.m3u8",
	  "thumbnail": "https://cdn.api.video/vod/vi4k0jvEUuaTdRAEjQ4Jfagz/thumbnail.jpg",
	  "mp4": "https://cdn.api.video/vod/vi4k0jvEUuaTdRAEjQ4Jfagz/mp4/360/source.mp4"
	}
  }`,
	`{
	"videoId": "vi6HangYsow3vXxwdx3YMlAb",
	"playerId": "pl3zY7qtojdW2EvMIU37707q",
	"title": "Maths video 2",
	"description": "An amazing video explaining the string theory 2",
	"public": true,
	"panoramic": false,
	"mp4Support":false,
	"tags": [
	  "maths",
	  "string theory"
	],
	"metadata": [
	  {
		"key": "Author",
		"value": "John Doe"
	  }
	],
	"publishedAt": "2019-07-16T23:36:18.598Z",
	"updatedAt": "2019-07-16T23:49:18.598Z",
	"source": {
	  "uri": "/videos/vi6HangYsow3vXxwdx3YMlAb/source"
	},
	"assets": {
	  "iframe": "<iframe src='//embed.api.video/vod/vi6HangYsow3vXxwdx3YMlAb' width='100%' height='100%' frameborder='0' scrolling='no' allowfullscreen=''></iframe>",
	  "player": "https://embed.api.video/vod/vi6HangYsow3vXxwdx3YMlAb",
	  "hls": "https://cdn.api.video/vod/vi6HangYsow3vXxwdx3YMlAb/hls/manifest.m3u8",
	  "thumbnail": "https://cdn.api.video/vod/vi6HangYsow3vXxwdx3YMlAb/thumbnail.jpg"			
	}
  }`,
}

func getTime(timStr string) time.Time {
	parsed, _ := time.Parse("2006-01-02T15:04:05.000Z", timStr)
	return parsed
}

var videoStructs = []Video{
	{
		VideoId:     PtrString("vi4k0jvEUuaTdRAEjQ4Jfagz"),
		Title:       PtrString("Maths video"),
		Description: PtrString("An amazing video explaining the string theory"),
		PublishedAt: PtrString("2019-07-14T23:36:18.598Z"),
		UpdatedAt:   PtrString("2019-07-14T23:49:18.598Z"),
		Tags:        &[]interface{}{"maths", "string theory", "video"},
		Metadata: &[]Metadata{
			{
				Key:   PtrString("Author"),
				Value: PtrString("John Doe"),
			},
			{
				Key:   PtrString("Format"),
				Value: PtrString("Tutorial"),
			},
		},
		Source: &VideoSource{
			Uri: PtrString("/videos/vi4k0jvEUuaTdRAEjQ4Jfagz/source"),
		},
		Assets: &VideoAssets{
			Hls:       PtrString("https://cdn.api.video/vod/vi4k0jvEUuaTdRAEjQ4Jfagz/hls/manifest.m3u8"),
			Iframe:    PtrString("<iframe src='//embed.api.video/vod/vi4k0jvEUuaTdRAEjQ4Jfagz' width='100%' height='100%' frameborder='0' scrolling='no' allowfullscreen=''></iframe>"),
			Thumbnail: PtrString("https://cdn.api.video/vod/vi4k0jvEUuaTdRAEjQ4Jfagz/thumbnail.jpg"),
			Player:    PtrString("https://embed.api.video/vod/vi4k0jvEUuaTdRAEjQ4Jfagz"),
			Mp4:       PtrString("https://cdn.api.video/vod/vi4k0jvEUuaTdRAEjQ4Jfagz/mp4/360/source.mp4"),
		},
		PlayerId:   PtrString("pl45KFKdlddgk654dspkze"),
		Public:     PtrBool(true),
		Panoramic:  PtrBool(false),
		Mp4Support: PtrBool(true),
	},
	{
		VideoId:     PtrString("vi6HangYsow3vXxwdx3YMlAb"),
		Title:       PtrString("Maths video 2"),
		Description: PtrString("An amazing video explaining the string theory 2"),
		PublishedAt: PtrString("2019-07-16T23:36:18.598Z"),
		UpdatedAt:   PtrString("2019-07-16T23:49:18.598Z"),
		Tags:        &[]interface{}{"maths", "string theory"},
		Metadata: &[]Metadata{
			{
				Key:   PtrString("Author"),
				Value: PtrString("John Doe"),
			},
		},
		Source: &VideoSource{
			Uri: PtrString("/videos/vi6HangYsow3vXxwdx3YMlAb/source"),
		},
		Assets: &VideoAssets{
			Hls:       PtrString("https://cdn.api.video/vod/vi6HangYsow3vXxwdx3YMlAb/hls/manifest.m3u8"),
			Iframe:    PtrString("<iframe src='//embed.api.video/vod/vi6HangYsow3vXxwdx3YMlAb' width='100%' height='100%' frameborder='0' scrolling='no' allowfullscreen=''></iframe>"),
			Thumbnail: PtrString("https://cdn.api.video/vod/vi6HangYsow3vXxwdx3YMlAb/thumbnail.jpg"),
			Player:    PtrString("https://embed.api.video/vod/vi6HangYsow3vXxwdx3YMlAb"),
		},
		PlayerId:   PtrString("pl3zY7qtojdW2EvMIU37707q"),
		Public:     PtrBool(true),
		Panoramic:  PtrBool(false),
		Mp4Support: PtrBool(false),
	},
}

var videoRequestJSON = `{
	"title": "Maths video",
	"description": "An amazing video explaining the string theory",
	"public": true,
	"mp4Support":true,
	"playerId": "pl45KFKdlddgk654dspkze",
	"tags": [
	  "maths",
	  "string theory",
	  "video"
	],
	"metadata": [
	  {
		"key": "Author",
		"value": "John Doe"
	  },
	  {
		"key": "Format",
		"value": "Tutorial"
	  }
	]
  }`

var videoCreatePayload = VideoCreatePayload{
	Title:       "Maths video",
	Description: PtrString("An amazing video explaining the string theory"),
	Public:      PtrBool(true),
	Panoramic:   PtrBool(false),
	PlayerId:    PtrString("pl45KFKdlddgk654dspkze"),
	Tags:        &[]string{"maths", "string theory", "video"},
	Metadata: &[]Metadata{
		{
			Key:   PtrString("Author"),
			Value: PtrString("John Doe"),
		},
		{
			Key:   PtrString("Format"),
			Value: PtrString("Tutorial"),
		},
	},
	Mp4Support: PtrBool(true),
}

var videoUpdateStruct = VideoUpdatePayload{
	Title:       PtrString("Maths video"),
	Description: PtrString("An amazing video explaining the string theory"),
	Public:      PtrBool(true),
	Panoramic:   PtrBool(false),
	PlayerId:    PtrString("pl45KFKdlddgk654dspkze"),
	Tags:        &[]string{"maths", "string theory", "video"},
	Metadata: &[]Metadata{
		{
			Key:   PtrString("Author"),
			Value: PtrString("John Doe"),
		},
		{
			Key:   PtrString("Format"),
			Value: PtrString("Tutorial"),
		},
	},
	Mp4Support: PtrBool(true),
}

var videoStatusJSONResponse = `{
	"ingest": {
	  "status": "uploaded",
	  "filesize": 273579401,
	  "receivedBytes": [
		{
		  "to": 134217727,
		  "from": 0,
		  "total": 273579401
		},
		{
		  "to": 268435455,
		  "from": 134217728,
		  "total": 273579401
		}
	  ]
	},
	"encoding": {
	  "playable": true,
	  "qualities": [
		{
		  "quality": "360p",
		  "status": "encoded"
		},
		{
		  "quality": "480p",
		  "status": "encoded"
		},
		{
		  "quality": "720p",
		  "status": "encoded"
		},
		{
		  "quality": "1080p",
		  "status": "encoding"
		},
		{
		  "quality": "2160p",
		  "status": "waiting"
		}
	  ],
	  "metadata": {
		"width": 424,
		"height": 240,
		"bitrate": 411,
		"duration": 4176,
		"framerate": 24,
		"samplerate": 48000,
		"videoCodec": "h264",
		"audioCodec": "aac",
		"aspectRatio": "16/9"
	  }
	}
  }`

var videoStatusStruct = Videostatus{
	Ingest: &VideostatusIngest{
		Status:   PtrString("uploaded"),
		Filesize: PtrInt32(273579401),
		ReceivedBytes: &[]BytesRange{
			{
				To:    PtrInt32(134217727),
				From:  PtrInt32(0),
				Total: PtrInt32(273579401),
			},
			{
				To:    PtrInt32(268435455),
				From:  PtrInt32(134217728),
				Total: PtrInt32(273579401),
			},
		},
	},
	Encoding: &VideostatusEncoding{
		Playable: PtrBool(true),
		Qualities: &[]Quality{
			{
				Quality: PtrString("360p"),
				Status:  PtrString("encoded"),
			},
			{
				Quality: PtrString("480p"),
				Status:  PtrString("encoded"),
			},
			{
				Quality: PtrString("720p"),
				Status:  PtrString("encoded"),
			},
			{
				Quality: PtrString("1080p"),
				Status:  PtrString("encoding"),
			},
			{
				Quality: PtrString("2160p"),
				Status:  PtrString("waiting"),
			},
		},
		Metadata: &VideostatusEncodingMetadata{
			Width:       PtrInt32(424),
			Height:      PtrInt32(240),
			Bitrate:     PtrFloat32(411),
			Duration:    PtrInt32(4176),
			Framerate:   PtrInt32(24),
			Samplerate:  PtrInt32(48000),
			VideoCodec:  PtrString("h264"),
			AudioCodec:  PtrString("aac"),
			AspectRatio: PtrString("16/9"),
		},
	},
}

func TestVideos_Get(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/videos/vi4k0jvEUuaTdRAEjQ4Jfagz", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, videoJSONResponses[0])
	})

	video, err := client.Videos.Get("vi4k0jvEUuaTdRAEjQ4Jfagz")
	if err != nil {
		t.Errorf("Videos.Get error: %v", err)
	}

	expected := &videoStructs[0]
	if !reflect.DeepEqual(video, expected) {
		t.Errorf("Videos.Get\n got=%#v\nwant=%#v", video, expected)
	}
}

func TestVideos_List(t *testing.T) {
	setup()
	defer teardown()
	JSONResp := fmt.Sprintf(
		`{"data":[%s,%s], "pagination":%s}`,
		videoJSONResponses[0],
		videoJSONResponses[1],
		paginationJSON)

	mux.HandleFunc("/videos", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		expectedQuery := url.Values{
			"currentPage":    []string{"1"},
			"pageSize":       []string{"25"},
			"sortBy":         []string{"publishedAt"},
			"sortOrder":      []string{"desc"},
			"tags":           []string{"tag1", "tag2"},
			"metadata[key]":  []string{"value"},
			"metadata[key2]": []string{"value2"},
		}
		fmt.Println(r.URL.Query())
		if !reflect.DeepEqual(r.URL.Query(), expectedQuery) {
			t.Errorf("Request querystring\n got=%#v\nwant=%#v", r.URL.Query(), expectedQuery)
		}
		fmt.Fprint(w, JSONResp)
	})

	videos, err := client.Videos.List(new(VideosApiListRequest).
		CurrentPage(1).
		PageSize(25).
		SortBy("publishedAt").
		SortOrder("desc").
		Tags([]string{"tag1", "tag2"}).
		Metadata(map[string]string{"key": "value", "key2": "value2"}))

	if err != nil {
		t.Errorf("Videos.List error: %v", err)
	}

	expected := &VideosListResponse{
		Data:       videoStructs,
		Pagination: paginationStruct,
	}
	if !reflect.DeepEqual(videos, expected) {
		t.Errorf("Videos.List\n got=%#v\nwant=%#v", videos, expected)
	}
}

func TestVideos_ListUpdatedAt(t *testing.T) {
	setup()
	defer teardown()
	JSONResp := fmt.Sprintf(
		`{"data":[%s,%s], "pagination":%s}`,
		videoJSONResponses[0],
		videoJSONResponses[1],
		paginationJSON)

	mux.HandleFunc("/videos", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		expectedQuery := url.Values{
			"currentPage":    []string{"1"},
			"pageSize":       []string{"25"},
			"sortBy":         []string{"updatedAt"},
			"sortOrder":      []string{"desc"},
			"tags":           []string{"tag1", "tag2"},
			"metadata[key]":  []string{"value"},
			"metadata[key2]": []string{"value2"},
		}
		if !reflect.DeepEqual(r.URL.Query(), expectedQuery) {
			t.Errorf("Request querystring\n got=%#v\nwant=%#v", r.URL.Query(), expectedQuery)
		}
		fmt.Fprint(w, JSONResp)
	})

	videos, err := client.Videos.List(new(VideosApiListRequest).
		CurrentPage(1).
		PageSize(25).
		SortBy("updatedAt").
		SortOrder("desc").
		Tags([]string{"tag1", "tag2"}).
		Metadata(map[string]string{"key": "value", "key2": "value2"}))

	if err != nil {
		t.Errorf("Videos.List error: %v", err)
	}

	expected := &VideosListResponse{
		Data:       videoStructs,
		Pagination: paginationStruct,
	}
	if !reflect.DeepEqual(videos, expected) {
		t.Errorf("Videos.List\n got=%#v\nwant=%#v", videos, expected)
	}
}

func TestVideos_Create(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/videos", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		expectedBody := map[string]interface{}{
			"title":       "Maths video",
			"description": "An amazing video explaining the string theory",
			"public":      true,
			"mp4Support":  true,
			"panoramic":   false,
			"playerId":    "pl45KFKdlddgk654dspkze",
			"tags":        []interface{}{"maths", "string theory", "video"},
			"metadata": []interface{}{
				map[string]interface{}{"key": "Author", "value": "John Doe"},
				map[string]interface{}{"key": "Format", "value": "Tutorial"},
			},
		}
		var v map[string]interface{}
		err := json.NewDecoder(r.Body).Decode(&v)
		if err != nil {
			t.Fatalf("decode json: %v", err)
		}

		if !reflect.DeepEqual(v, expectedBody) {
			t.Errorf("Request body\n got=%#v\nwant=%#v", v, expectedBody)
		}
		fmt.Fprint(w, videoJSONResponses[0])
	})

	video, err := client.Videos.Create(videoCreatePayload)
	if err != nil {
		t.Errorf("Videos.Create error: %v", err)
	}

	expected := &videoStructs[0]
	if !reflect.DeepEqual(video, expected) {
		t.Errorf("Videos.Create\n got=%#v\nwant=%#v", video, expected)
	}
}

func TestVideos_Update(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/videos/vi4k0jvEUuaTdRAEjQ4Jfagz", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPatch)
		expectedBody := map[string]interface{}{
			"title":       "Maths video",
			"description": "An amazing video explaining the string theory",
			"public":      true,
			"mp4Support":  true,
			"panoramic":   false,
			"playerId":    "pl45KFKdlddgk654dspkze",
			"tags":        []interface{}{"maths", "string theory", "video"},
			"metadata": []interface{}{
				map[string]interface{}{"key": "Author", "value": "John Doe"},
				map[string]interface{}{"key": "Format", "value": "Tutorial"},
			},
		}
		var v map[string]interface{}
		err := json.NewDecoder(r.Body).Decode(&v)
		if err != nil {
			t.Fatalf("decode json: %v", err)
		}

		if !reflect.DeepEqual(v, expectedBody) {
			t.Errorf("Request body\n got=%#v\nwant=%#v", v, expectedBody)
		}
		fmt.Fprint(w, videoJSONResponses[0])
	})

	video, err := client.Videos.Update("vi4k0jvEUuaTdRAEjQ4Jfagz", videoUpdateStruct)
	if err != nil {
		t.Errorf("Videos.Update error: %v", err)
	}

	expected := &videoStructs[0]
	if !reflect.DeepEqual(video, expected) {
		t.Errorf("Videos.Update\n got=%#v\nwant=%#v", video, expected)
	}
}

func TestVideos_Delete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/videos/vi4k0jvEUuaTdRAEjQ4Jfagz", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	err := client.Videos.Delete("vi4k0jvEUuaTdRAEjQ4Jfagz")
	if err != nil {
		t.Errorf("Videos.Delete error: %v", err)
	}
}

func TestVideos_Upload(t *testing.T) {
	setup()
	defer teardown()
	var count int
	mux.HandleFunc("/videos/vi4k0jvEUuaTdRAEjQ4Jfagz/source", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		if r.Header.Get("Content-Range") != "" {
			t.Errorf("Videos.Upload Content-Range header shouldn't be set in non chunked uploads")
		}
		fmt.Fprint(w, videoJSONResponses[0])
		count++
	})

	file := createTempFile("test.video", 8*1024*1024)
	defer os.Remove(file.Name())

	video, err := client.Videos.UploadFile("vi4k0jvEUuaTdRAEjQ4Jfagz", file)
	if err != nil {
		t.Errorf("Videos.Upload error: %v", err)
	}

	if count != 1 {
		t.Errorf("Videos.Upload endpoint should be called 1 time on non chunked uploads")
	}

	expected := &videoStructs[0]
	if !reflect.DeepEqual(video, expected) {
		t.Errorf("Videos.Upload\n got=%#v\nwant=%#v", video, expected)
	}
}

func TestVideos_ChunkedUpload(t *testing.T) {
	setup()
	defer teardown()
	var count int64
	headers := []string{
		"bytes 0-2097151/8388608",
		"bytes 2097152-4194303/8388608",
		"bytes 4194304-6291455/8388608",
		"bytes 6291456-8388607/8388608",
	}
	mux.HandleFunc("/videos/vi4k0jvEUuaTdRAEjQ4Jfagz/source", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		if r.Header.Get("Content-Range") == "" {
			t.Errorf("Videos.ChunkedUpload Content-Range header should be set in chunked uploads")
		}
		if r.Header.Get("Content-Range") != headers[count] {
			t.Errorf("Videos.ChunkedUpload Bad Content-Range got=%#v\nwant=%#v", r.Header.Get("Content-Range"), headers[count])
		}
		fmt.Fprint(w, videoJSONResponses[0])
		count++
	})

	filesize := int64(8 * 1024 * 1024)
	chunksize := int64(2 * 1024 * 1024)
	nbRequests := (filesize / chunksize)
	file := createTempFile("test.video", filesize)
	defer os.Remove(file.Name())

	client.ChunkSize(chunksize)
	video, err := client.Videos.UploadFile("vi4k0jvEUuaTdRAEjQ4Jfagz", file)
	if err != nil {
		t.Errorf("Videos.ChunkedUpload error: %v", err)
	}

	if count != nbRequests {
		t.Errorf("Videos.ChunkedUpload endpoint should be called %d times, got %d", nbRequests, count)
	}

	expected := &videoStructs[0]
	if !reflect.DeepEqual(video, expected) {
		t.Errorf("Videos.ChunkedUpload\n got=%#v\nwant=%#v", video, expected)
	}
}

func TestVideos_Status(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/videos/vi4k0jvEUuaTdRAEjQ4Jfagz/status", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, videoStatusJSONResponse)
	})

	status, err := client.Videos.GetStatus("vi4k0jvEUuaTdRAEjQ4Jfagz")
	if err != nil {
		t.Errorf("Videos.Status error: %v", err)
	}

	expected := &videoStatusStruct
	if !reflect.DeepEqual(status, expected) {
		t.Errorf("Videos.Status\n got=%#v\nwant=%#v", status, expected)
	}

}

func TestVideos_PickThumbnail(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/videos/vi4k0jvEUuaTdRAEjQ4Jfagz/thumbnail", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPatch)
		fmt.Fprint(w, videoJSONResponses[0])
	})

	video, err := client.Videos.PickThumbnail("vi4k0jvEUuaTdRAEjQ4Jfagz", VideoThumbnailPickPayload{Timecode: "00:00:01:02"})
	if err != nil {
		t.Errorf("Videos.PickThumbnail error: %v", err)
	}

	expected := &videoStructs[0]
	if !reflect.DeepEqual(video, expected) {
		t.Errorf("Videos.PickThumbnail\n got=%#v\nwant=%#v", video, expected)
	}

}

func TestVideos_UploadThumbnail(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/videos/vi4k0jvEUuaTdRAEjQ4Jfagz/thumbnail", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, videoJSONResponses[0])
	})

	file := createTempFile("test.thumbnail", 1024*1024)
	defer os.Remove(file.Name())

	video, err := client.Videos.UploadThumbnailFile("vi4k0jvEUuaTdRAEjQ4Jfagz", file)
	if err != nil {
		t.Errorf("Videos.UploadThumbnail error: %v", err)
	}

	expected := &videoStructs[0]
	if !reflect.DeepEqual(video, expected) {
		t.Errorf("Videos.UploadThumbnail\n got=%#v\nwant=%#v", video, expected)
	}
}
