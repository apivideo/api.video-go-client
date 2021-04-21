package apivideosdk

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"testing"
)

var playerJSONResponses = []string{
	`{
		"playerId": "pt3Lony8J6NozV71Yxn8KVFn",
		"assets": {
			"logo": "https://cdn.api.video/player/pl3Lony8J6NozV71Yxn8KVFn/logo.png",
			"link": "https://api.video"
		},
		"text": "rgba(255, 255, 255, 0.95)",
		"link": "rgba(255, 0, 0, 0.95)",
		"linkHover": "rgba(255, 255, 255, 0.75)",
		"trackPlayed": "rgba(255, 255, 255, 0.95)",
		"trackUnplayed": "rgba(255, 255, 255, 0.1)",
		"trackBackground": "rgba(0, 0, 0, 0)",
		"backgroundTop": "rgba(72, 4, 45, 1)",
		"backgroundBottom": "rgba(94, 95, 89, 1)",
		"backgroundText": "rgba(255, 255, 255, 0.95)",
		"enableApi": false,
		"enableControls": true,
		"forceAutoplay": false,
		"hideTitle": false,
		"forceLoop": false
	}`,
	`{
		"playerId": "pl3zY7qtojdW2EvMIU37707q",
		"assets": {
			"logo": "https://cdn.api.video/player/pl3zY7qtojdW2EvMIU37707q/logo.png",
			"link": "https://api.video"
		},
		"text": "rgba(255, 255, 255, 0.95)",
		"link": "rgba(255, 0, 0, 0.95)",
		"linkHover": "rgba(255, 255, 255, 0.75)",
		"trackPlayed": "rgba(255, 255, 255, 0.95)",
		"trackUnplayed": "rgba(255, 255, 255, 0.1)",
		"trackBackground": "rgba(0, 0, 0, 0)",
		"backgroundTop": "rgba(72, 4, 45, 1)",
		"backgroundBottom": "rgba(94, 95, 89, 1)",
		"backgroundText": "rgba(255, 255, 255, 0.95)",
		"enableApi": false,
		"enableControls": true,
		"forceAutoplay": false,
		"hideTitle": false,
		"forceLoop": false
	}`,
}

var playerStructs = []Player{
	{
		PlayerId:         PtrString("pt3Lony8J6NozV71Yxn8KVFn"),
		Text:             PtrString("rgba(255, 255, 255, 0.95)"),
		Link:             PtrString("rgba(255, 0, 0, 0.95)"),
		LinkHover:        PtrString("rgba(255, 255, 255, 0.75)"),
		TrackPlayed:      PtrString("rgba(255, 255, 255, 0.95)"),
		TrackUnplayed:    PtrString("rgba(255, 255, 255, 0.1)"),
		TrackBackground:  PtrString("rgba(0, 0, 0, 0)"),
		BackgroundTop:    PtrString("rgba(72, 4, 45, 1)"),
		BackgroundBottom: PtrString("rgba(94, 95, 89, 1)"),
		BackgroundText:   PtrString("rgba(255, 255, 255, 0.95)"),
		EnableApi:        PtrBool(false),
		EnableControls:   PtrBool(true),
		ForceAutoplay:    PtrBool(false),
		HideTitle:        PtrBool(false),
		ForceLoop:        PtrBool(false),
		Assets: &PlayerAllOfAssets{
			Logo: PtrString("https://cdn.api.video/player/pl3Lony8J6NozV71Yxn8KVFn/logo.png"),
			Link: PtrString("https://api.video"),
		},
	},
	{
		PlayerId:         PtrString("pl3zY7qtojdW2EvMIU37707q"),
		Text:             PtrString("rgba(255, 255, 255, 0.95)"),
		Link:             PtrString("rgba(255, 0, 0, 0.95)"),
		LinkHover:        PtrString("rgba(255, 255, 255, 0.75)"),
		TrackPlayed:      PtrString("rgba(255, 255, 255, 0.95)"),
		TrackUnplayed:    PtrString("rgba(255, 255, 255, 0.1)"),
		TrackBackground:  PtrString("rgba(0, 0, 0, 0)"),
		BackgroundTop:    PtrString("rgba(72, 4, 45, 1)"),
		BackgroundBottom: PtrString("rgba(94, 95, 89, 1)"),
		BackgroundText:   PtrString("rgba(255, 255, 255, 0.95)"),
		EnableApi:        PtrBool(false),
		EnableControls:   PtrBool(true),
		ForceAutoplay:    PtrBool(false),
		HideTitle:        PtrBool(false),
		ForceLoop:        PtrBool(false),
		Assets: &PlayerAllOfAssets{
			Logo: PtrString("https://cdn.api.video/player/pl3zY7qtojdW2EvMIU37707q/logo.png"),
			Link: PtrString("https://api.video"),
		},
	},
}

var playerCreationPayload = PlayerCreationPayload{
	Text:             PtrString("rgba(255, 255, 255, 0.95)"),
	Link:             PtrString("rgba(255, 0, 0, 0.95)"),
	LinkHover:        PtrString("rgba(255, 255, 255, 0.75)"),
	TrackPlayed:      PtrString("rgba(255, 255, 255, 0.95)"),
	TrackUnplayed:    PtrString("rgba(255, 255, 255, 0.1)"),
	TrackBackground:  PtrString("rgba(0, 0, 0, 0)"),
	BackgroundTop:    PtrString("rgba(72, 4, 45, 1)"),
	BackgroundBottom: PtrString("rgba(94, 95, 89, 1)"),
	BackgroundText:   PtrString("rgba(255, 255, 255, 0.95)"),
	EnableApi:        PtrBool(false),
	EnableControls:   PtrBool(true),
	ForceAutoplay:    PtrBool(false),
	HideTitle:        PtrBool(false),
	ForceLoop:        PtrBool(false),
}

var playerUpdatePayload = PlayerUpdatePayload{
	Text:             PtrString("rgba(255, 255, 255, 0.95)"),
	Link:             PtrString("rgba(255, 0, 0, 0.95)"),
	LinkHover:        PtrString("rgba(255, 255, 255, 0.75)"),
	TrackPlayed:      PtrString("rgba(255, 255, 255, 0.95)"),
	TrackUnplayed:    PtrString("rgba(255, 255, 255, 0.1)"),
	TrackBackground:  PtrString("rgba(0, 0, 0, 0)"),
	BackgroundTop:    PtrString("rgba(72, 4, 45, 1)"),
	BackgroundBottom: PtrString("rgba(94, 95, 89, 1)"),
	BackgroundText:   PtrString("rgba(255, 255, 255, 0.95)"),
	EnableApi:        PtrBool(false),
	EnableControls:   PtrBool(true),
	ForceAutoplay:    PtrBool(false),
	HideTitle:        PtrBool(false),
	ForceLoop:        PtrBool(false),
}

func TestPlayers_Get(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/players/pt3Lony8J6NozV71Yxn8KVFn", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, playerJSONResponses[0])
	})

	player, err := client.PlayerThemes.Get("pt3Lony8J6NozV71Yxn8KVFn")
	if err != nil {
		t.Errorf("Players.Get error: %v", err)
	}

	expected := &playerStructs[0]
	if !reflect.DeepEqual(player, expected) {
		t.Errorf("Players.Get\n got=%#v\nwant=%#v", player, expected)
	}
}

func TestPlayers_List(t *testing.T) {
	setup()
	defer teardown()
	JSONResp := fmt.Sprintf(
		`{"data":[%s,%s], "pagination":%s}`,
		playerJSONResponses[0],
		playerJSONResponses[1],
		paginationJSON)

	mux.HandleFunc("/players", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		expectedQuery := url.Values{
			"currentPage": []string{"1"},
			"pageSize":    []string{"25"},
		}
		if !reflect.DeepEqual(r.URL.Query(), expectedQuery) {
			t.Errorf("Request querystring\n got=%#v\nwant=%#v", r.URL.Query(), expectedQuery)
		}
		fmt.Fprint(w, JSONResp)
	})

	opts := new(PlayerThemesApiListRequest).CurrentPage(1).PageSize(25)

	players, err := client.PlayerThemes.List(opts)
	if err != nil {
		t.Errorf("Players.List error: %v", err)
	}

	expected := PlayersListResponse{
		Data:       &playerStructs,
		Pagination: &paginationStruct,
	}
	if !reflect.DeepEqual(*players, expected) {
		t.Errorf("Players.List\n got=%#v\nwant=%#v", *players, expected)
	}
}

func TestPlayers_Create(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/players", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		expectedBody := map[string]interface{}{
			"text":             "rgba(255, 255, 255, 0.95)",
			"link":             "rgba(255, 0, 0, 0.95)",
			"linkHover":        "rgba(255, 255, 255, 0.75)",
			"trackPlayed":      "rgba(255, 255, 255, 0.95)",
			"trackUnplayed":    "rgba(255, 255, 255, 0.1)",
			"trackBackground":  "rgba(0, 0, 0, 0)",
			"backgroundTop":    "rgba(72, 4, 45, 1)",
			"backgroundBottom": "rgba(94, 95, 89, 1)",
			"backgroundText":   "rgba(255, 255, 255, 0.95)",
			"enableApi":        false,
			"enableControls":   true,
			"forceAutoplay":    false,
			"hideTitle":        false,
			"forceLoop":        false,
		}
		var v map[string]interface{}
		err := json.NewDecoder(r.Body).Decode(&v)
		if err != nil {
			t.Fatalf("decode json: %v", err)
		}
		if !reflect.DeepEqual(v, expectedBody) {
			t.Errorf("Request body\ngot=%#v\nwant=%#v", v, expectedBody)
		}
		fmt.Fprint(w, playerJSONResponses[0])
	})

	player, err := client.PlayerThemes.Create(playerCreationPayload)
	if err != nil {
		t.Errorf("Players.Create error: %v", err)
	}

	expected := &playerStructs[0]
	if !reflect.DeepEqual(player, expected) {
		t.Errorf("Players.Create\n got=%#v\nwant=%#v", player, expected)
	}
}

func TestPlayers_Update(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/players/pt3Lony8J6NozV71Yxn8KVFn", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPatch)
		expectedBody := map[string]interface{}{
			"text":             "rgba(255, 255, 255, 0.95)",
			"link":             "rgba(255, 0, 0, 0.95)",
			"linkHover":        "rgba(255, 255, 255, 0.75)",
			"trackPlayed":      "rgba(255, 255, 255, 0.95)",
			"trackUnplayed":    "rgba(255, 255, 255, 0.1)",
			"trackBackground":  "rgba(0, 0, 0, 0)",
			"backgroundTop":    "rgba(72, 4, 45, 1)",
			"backgroundBottom": "rgba(94, 95, 89, 1)",
			"backgroundText":   "rgba(255, 255, 255, 0.95)",
			"enableApi":        false,
			"enableControls":   true,
			"forceAutoplay":    false,
			"hideTitle":        false,
			"forceLoop":        false,
		}
		var v map[string]interface{}
		err := json.NewDecoder(r.Body).Decode(&v)
		if err != nil {
			t.Fatalf("decode json: %v", err)
		}
		fmt.Println(v)
		if !reflect.DeepEqual(v, expectedBody) {
			t.Errorf("Request body\n got=%#v\n want=%#v", v, expectedBody)
		}
		fmt.Fprint(w, playerJSONResponses[0])
	})

	player, err := client.PlayerThemes.Update("pt3Lony8J6NozV71Yxn8KVFn", playerUpdatePayload)
	if err != nil {
		t.Errorf("Players.Update error: %v", err)
	}

	expected := &playerStructs[0]
	if !reflect.DeepEqual(player, expected) {
		t.Errorf("Players.Update\n got=%#v\nwant=%#v", player, expected)
	}
}

func TestPlayers_Delete(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/players/pt3Lony8J6NozV71Yxn8KVFn", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	err := client.PlayerThemes.Delete("pt3Lony8J6NozV71Yxn8KVFn")
	if err != nil {
		t.Errorf("Players.Delete error: %v", err)
	}
}

func TestPlayers_UploadLogo(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/players/pt3Lony8J6NozV71Yxn8KVFn/logo", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, playerJSONResponses[0])
	})

	file := createTempFile("test.logo", 1024*1024)
	defer os.Remove(file.Name())

	player, err := client.PlayerThemes.UploadLogoFile("pt3Lony8J6NozV71Yxn8KVFn", file, "https://api.video")
	if err != nil {
		t.Errorf("Captions.Upload error: %v", err)
	}

	expected := &playerStructs[0]
	if !reflect.DeepEqual(player, expected) {
		t.Errorf("Captions.Upload\n got=%#v\nwant=%#v", player, expected)
	}
}

func TestPlayers_DeleteLogo(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/players/pt3Lony8J6NozV71Yxn8KVFn/logo", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	err := client.PlayerThemes.DeleteLogo("pt3Lony8J6NozV71Yxn8KVFn")

	if err != nil {
		t.Errorf("Players.DeleteLogo error: %v", err)
	}
}
