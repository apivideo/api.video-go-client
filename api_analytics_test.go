package apivideosdk

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

var liveStreamsPlaysResponses = []string{
	`{
		"data" : [ {
				"value" : "li3q7HxhApxRF1c8F8r6VeaI",
				"plays" : 100
			}, {
				"value" : "li3q7HxhApxRF1c8F8r6VeaB",
				"plays" : 10
			}, {
				"value" : "li3q7HxhApxRF1c8F8r6VeaD",
				"plays" : 1
			} 
		]
	}`,
	`{
		"data" : [ {
				"value" : "france",
				"plays" : 100
			}, {
				"value" : "united states",
				"plays" : 10
			}, {
				"value" : "spain",
				"plays" : 1
			} 
		]
	}`,
	`{
	  	"data" : [ {
				"value" : "2023-06-10T10:00:00.000Z",
				"plays" : 100
	  		}, {
				"value" : "2023-06-10T11:00:00.000Z",
				"plays" : 10
	  		}, {
				"value" : "2023-06-10T12:00:00.000Z",
				"plays" : 1
	  		}
		]
	}`,
}

var liveStreamsPlaysStructs = []AnalyticsPlaysResponse{
	{
		Data: []AnalyticsData{
			{Value: "li3q7HxhApxRF1c8F8r6VeaI", Plays: 100},
			{Value: "li3q7HxhApxRF1c8F8r6VeaB", Plays: 10},
			{Value: "li3q7HxhApxRF1c8F8r6VeaD", Plays: 1},
		},
	},
	{
		Data: []AnalyticsData{
			{Value: "france", Plays: 100},
			{Value: "united states", Plays: 10},
			{Value: "spain", Plays: 1},
		},
	},
	{
		Data: []AnalyticsData{
			{Value: "2023-06-10T10:00:00.000Z", Plays: 100},
			{Value: "2023-06-10T11:00:00.000Z", Plays: 10},
			{Value: "2023-06-10T12:00:00.000Z", Plays: 1},
		},
	},
}

func TestAnalytics_GetLiveStreamsPlaysByLiveStreamId(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/analytics/live-streams/plays", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, liveStreamsPlaysResponses[0])
	})

	plays, err := client.Analytics.GetLiveStreamsPlays(AnalyticsApiGetLiveStreamsPlaysRequest{}.From("2023-04-01").Dimension("liveStreamId"))
	if err != nil {
		t.Errorf("Analytics.GetLiveStreamsPlays error: %v", err)
	}

	expected := &liveStreamsPlaysStructs[0]
	if !reflect.DeepEqual(plays, expected) {
		t.Errorf("Analytics.GetLiveStreamsPlays\n got=%#v\nwant=%#v", plays, expected)
	}
}

func TestAnalytics_GetLiveStreamsPlaysByCountry(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/analytics/live-streams/plays", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, liveStreamsPlaysResponses[1])
	})

	plays, err := client.Analytics.GetLiveStreamsPlays(AnalyticsApiGetLiveStreamsPlaysRequest{}.From("2023-04-01").Dimension("country"))
	if err != nil {
		t.Errorf("Analytics.GetLiveStreamsPlays error: %v", err)
	}

	expected := &liveStreamsPlaysStructs[1]
	if !reflect.DeepEqual(plays, expected) {
		t.Errorf("Analytics.GetLiveStreamsPlays\n got=%#v\nwant=%#v", plays, expected)
	}
}

func TestAnalytics_GetLiveStreamsPlaysByEmittedAt(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/analytics/live-streams/plays", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, liveStreamsPlaysResponses[2])
	})

	plays, err := client.Analytics.GetLiveStreamsPlays(AnalyticsApiGetLiveStreamsPlaysRequest{}.From("2023-04-01").Dimension("emittedAt"))
	if err != nil {
		t.Errorf("Analytics.GetLiveStreamsPlays error: %v", err)
	}

	expected := &liveStreamsPlaysStructs[2]
	if !reflect.DeepEqual(plays, expected) {
		t.Errorf("Analytics.GetLiveStreamsPlays\n got=%#v\nwant=%#v", plays, expected)
	}
}

var videosPlaysResponses = []string{
	`{
		"data" : [ {
				"value" : "vi3q7HxhApxRF1c8F8r6VeaI",
				"plays" : 100
			}, {
				"value" : "vi3q7HxhApxRF1c8F8r6VeaF",
				"plays" : 10
			}, {
				"value" : "vi3q7HxhApxRF1c8F8r6VeaH",
				"plays" : 1
			} 
		]
	}`,
	`{
		"data" : [ {
				"value" : "france",
				"plays" : 100
			}, {
				"value" : "united states",
				"plays" : 10
			}, {
				"value" : "spain",
				"plays" : 1
			} 
		]
	}`,
	`{
	  	"data" : [ {
				"value" : "2023-06-10T10:00:00.000Z",
				"plays" : 100
	  		}, {
				"value" : "2023-06-10T11:00:00.000Z",
				"plays" : 10
	  		}, {
				"value" : "2023-06-10T12:00:00.000Z",
				"plays" : 1
	  		}
		]
	}`,
}

var videosPlaysStructs = []AnalyticsPlaysResponse{
	{
		Data: []AnalyticsData{
			{Value: "vi3q7HxhApxRF1c8F8r6VeaI", Plays: 100},
			{Value: "vi3q7HxhApxRF1c8F8r6VeaF", Plays: 10},
			{Value: "vi3q7HxhApxRF1c8F8r6VeaH", Plays: 1},
		},
	},
	{
		Data: []AnalyticsData{
			{Value: "france", Plays: 100},
			{Value: "united states", Plays: 10},
			{Value: "spain", Plays: 1},
		},
	},
	{
		Data: []AnalyticsData{
			{Value: "2023-06-10T10:00:00.000Z", Plays: 100},
			{Value: "2023-06-10T11:00:00.000Z", Plays: 10},
			{Value: "2023-06-10T12:00:00.000Z", Plays: 1},
		},
	},
}

func TestAnalytics_GetVideosPlaysByVideoId(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/analytics/videos/plays", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, videosPlaysResponses[0])
	})

	plays, err := client.Analytics.GetVideosPlays(AnalyticsApiGetVideosPlaysRequest{}.From("2023-04-01").Dimension("videoId"))
	if err != nil {
		t.Errorf("Analytics.GetVideosPlays error: %v", err)
	}

	expected := &videosPlaysStructs[0]
	if !reflect.DeepEqual(plays, expected) {
		t.Errorf("Analytics.GetVideosPlays\n got=%#v\nwant=%#v", plays, expected)
	}
}

func TestAnalytics_GetVideosPlaysByCountry(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/analytics/videos/plays", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, videosPlaysResponses[1])
	})

	plays, err := client.Analytics.GetVideosPlays(AnalyticsApiGetVideosPlaysRequest{}.From("2023-04-01").Dimension("country"))
	if err != nil {
		t.Errorf("Analytics.GetVideosPlays error: %v", err)
	}

	expected := &videosPlaysStructs[1]
	if !reflect.DeepEqual(plays, expected) {
		t.Errorf("Analytics.GetVideosPlays\n got=%#v\nwant=%#v", plays, expected)
	}
}

func TestAnalytics_GetVideosPlaysByEmittedAt(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/analytics/videos/plays", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, videosPlaysResponses[2])
	})

	plays, err := client.Analytics.GetVideosPlays(AnalyticsApiGetVideosPlaysRequest{}.From("2023-04-01").Dimension("emittedAt"))
	if err != nil {
		t.Errorf("Analytics.GetVideosPlays error: %v", err)
	}

	expected := &videosPlaysStructs[2]
	if !reflect.DeepEqual(plays, expected) {
		t.Errorf("Analytics.GetVideosPlays\n got=%#v\nwant=%#v", plays, expected)
	}
}
