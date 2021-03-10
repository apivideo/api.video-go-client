package apivideosdk

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

var accountJSONResponse string = `{
	"quota": {
	  "quotaUsed": 6,
	  "quotaRemaining": 54,
	  "quotaTotal": 60
	},
	"features": [
	  "app.dynamic_metadata",
	  "app.event_log",
	  "player.white_label",
	  "stats.player_events",
	  "transcode.mp4_support"
	]
  }`

var accountStruct = Account{
	Quota: &AccountQuota{
		QuotaUsed:      PtrFloat32(6),
		QuotaRemaining: PtrFloat32(54),
		QuotaTotal:     PtrFloat32(60),
	},
	Features: &[]string{
		"app.dynamic_metadata",
		"app.event_log",
		"player.white_label",
		"stats.player_events",
		"transcode.mp4_support",
	},
}

func TestAccount_Get(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/account", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, accountJSONResponse)
	})

	account, err := client.Account.Get()
	if err != nil {
		t.Errorf("Account.Get error: %v", err)
	}

	expected := &accountStruct
	if !reflect.DeepEqual(account, expected) {
		t.Errorf("Account.Get\n got=%#v\nwant=%#v", account, expected)
	}
}
