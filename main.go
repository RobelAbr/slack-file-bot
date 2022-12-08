package main

import (
	"fmt"
	"github.com/slack-go/slack"
	"os"
)

func main() {

	os.Setenv("SLACK_BOT_TOKEN", "xoxb-4471312034385-4499035308353-GPvhlIJKkFgsBDmXF10SVw4G")
	os.Setenv("CHANNEL_ID", "C04D23UMN9M")
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channelArr := []string{os.Getenv("CHANNEL_ID")}
	fileArr := []string{"test-for-slack-file"}

	for i := 0; i < len(fileArr); i++ {
		params := slack.FileUploadParameters{
			Channels: channelArr,
			File:     fileArr[i],
		}
		file, err := api.UploadFile(params)
		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}
		fmt.Printf("%s, URL:%s\n", file.Name, file.URL)
	}
}
