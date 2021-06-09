package apivideosdk

import (
	"os"
	"testing"
)

func IntegrationTests(t *testing.T) {
	client := ClientBuilder(os.Getenv("API_KEY")).Build()

	// create a video
	payload := VideoCreationPayload{
		Title: "go client test",
	}
	video, err := client.Videos.Create(payload)

	if err != nil {
		t.Errorf("Videos.Create error: %v", err)
	}

	// upload a video file
	videoFile, err := os.Open("sample.mp4")
	_, err = client.Videos.UploadFile(video.VideoId, videoFile)

	if err != nil {
		t.Errorf("Videos.Delete error: %v", err)
	}

	// list video sessions
	metadata := map[string]string{"key": "value", "key2": "value2"}
	listVideoSessionPayload := RawStatisticsApiListVideoSessionsRequest{
		metadata: &metadata,
	}
	_, err = client.RawStatistics.ListVideoSessions(video.VideoId, listVideoSessionPayload)

	if err != nil {
		t.Errorf("RawStatistics.ListVideoSessions error: %v", err)
	}

	// delete the video
	err = client.Videos.Delete(video.VideoId)

	if err != nil {
		t.Errorf("Videos.Delete error: %v", err)
	}
}
