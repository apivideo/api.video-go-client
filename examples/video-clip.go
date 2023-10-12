/ First install the go client with "go get github.com/apivideo/api.video-go-client"
// Documentation: https://github.com/apivideo/api.video-go-client/blob/main/docs/VideosApi.md#create

package main

import (
    "context"
    "fmt"
    "os"
    apivideosdk "github.com/apivideo/api.video-go-client"
)
// upload a video with a clip
func main() {
    client := apivideosdk.ClientBuilder("YOUR_API_KEY").Build()
    // if you rather like to use the sandbox environment:
    // client := apivideosdk.SandboxClientBuilder("YOUR_SANDBOX_API_KEY").Build()
    video clip := apivideosdk.VideoClip{}
		videoClip.SetStartTimecode("00:00:20")
		videoClip.SetEndTimecode("00:00:22")
		videoCreationPayload := apivideosdk.VideoCreationPayload{}
		videoCreationPayload.SetTitle("my title")
		videoCreationPayload.SetClip(videoClip)
    res, err := client.Videos.Create(videoCreationPayload)
    videoId := res.videoId
	  uploadRes, err := apiVideoClient.Videos.UploadFile(videoId, videoFile)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Videos.Create``: %v\
", err)
    }
    // response from `Create`: Video
    fmt.Fprintf(os.Stdout, "Response from `Videos.Create`: %v\
", res)
}
