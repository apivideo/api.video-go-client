package apivideosdk

import (
	"os"
	"testing"
)

func IntegrationUploadTests(t *testing.T) {
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

	// delete the video
	err = client.Videos.Delete(video.VideoId)

	if err != nil {
		t.Errorf("Videos.Delete error: %v", err)
	}
}

func IntegrationUploadWithUploadTokenTests(t *testing.T) {
	client := ClientBuilder(os.Getenv("API_KEY")).Build()

	// create a video
	payload := TokenCreationPayload{}
	token, err := client.UploadTokens.CreateToken(payload)

	if err != nil {
		t.Errorf("UploadTokens.CreateToken error: %v", err)
	}

	// upload a video file
	videoFile, err := os.Open("sample.mp4")
	video, err := client.Videos.UploadWithUploadTokenFile(*token.Token, videoFile)

	if err != nil {
		t.Errorf("Videos.Delete error: %v", err)
	}

	// delete the video
	err = client.Videos.Delete(video.VideoId)

	if err != nil {
		t.Errorf("Videos.Delete error: %v", err)
	}

	// delete the token
	err = client.UploadTokens.DeleteToken(*token.Token)

	if err != nil {
		t.Errorf("UploadTokens.DeleteToken error: %v", err)
	}
}

func IntegrationUploadWithUploadTokenWithVideoIdTests(t *testing.T) {
	client := ClientBuilder(os.Getenv("API_KEY")).Build()

	// create an upload token
	tokenPayload := TokenCreationPayload{}
	token, err := client.UploadTokens.CreateToken(tokenPayload)

	if err != nil {
		t.Errorf("UploadTokens.CreateToken error: %v", err)
	}

	// create a video
	videoPayload := VideoCreationPayload{
		Title: "go client test",
	}
	video, err := client.Videos.Create(videoPayload)

	if err != nil {
		t.Errorf("Videos.Create error: %v", err)
	}

	// upload a video file
	videoFile, err := os.Open("sample.mp4")
	_, err = client.Videos.UploadWithUploadTokenFileWithVideoId(&video.VideoId, *token.Token, videoFile)

	if err != nil {
		t.Errorf("Videos.Delete error: %v", err)
	}

	// delete the video
	err = client.Videos.Delete(video.VideoId)

	if err != nil {
		t.Errorf("Videos.Delete error: %v", err)
	}

	// delete the token
	err = client.UploadTokens.DeleteToken(*token.Token)

	if err != nil {
		t.Errorf("UploadTokens.DeleteToken error: %v", err)
	}
}
