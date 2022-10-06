package main

import (
	"bytes"
	"encoding/json"
	"golang.design/x/clipboard"
	"io"
	"log"
	"net/http"
	"time"
)

const uploadUrl = "https://api.imgur.com/3/image?client_id=546c25a59c58ad7"

type ImgurUploadImageResponse struct {
	Data struct {
		Id          string        `json:"id"`
		Title       interface{}   `json:"title"`
		Description interface{}   `json:"description"`
		Datetime    int           `json:"datetime"`
		Type        string        `json:"type"`
		Animated    bool          `json:"animated"`
		Width       int           `json:"width"`
		Height      int           `json:"height"`
		Size        int           `json:"size"`
		Views       int           `json:"views"`
		Bandwidth   int           `json:"bandwidth"`
		Vote        interface{}   `json:"vote"`
		Favorite    bool          `json:"favorite"`
		Nsfw        interface{}   `json:"nsfw"`
		Section     interface{}   `json:"section"`
		AccountUrl  interface{}   `json:"account_url"`
		AccountId   int           `json:"account_id"`
		IsAd        bool          `json:"is_ad"`
		InMostViral bool          `json:"in_most_viral"`
		HasSound    bool          `json:"has_sound"`
		Tags        []interface{} `json:"tags"`
		AdType      int           `json:"ad_type"`
		AdUrl       string        `json:"ad_url"`
		Edited      string        `json:"edited"`
		InGallery   bool          `json:"in_gallery"`
		Deletehash  string        `json:"deletehash"`
		Name        string        `json:"name"`
		Link        string        `json:"link"`
	} `json:"data"`
	Success bool `json:"success"`
	Status  int  `json:"status"`
}

func main() {
	for {
		image := clipboard.Read(clipboard.FmtImage)

		if len(image) > 0 {
			link := uploadImageToImgur(&image)
			clipboard.Write(clipboard.FmtText, []byte(link))

			log.Println(link)
		}

		time.Sleep(500 * time.Millisecond)
	}
}

func uploadImageToImgur(image *[]byte) string {
	postData := bytes.NewReader(*image)

	response, err := http.Post(uploadUrl, "", postData)
	if err != nil {
		log.Println("Error upload")
		return ""
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println("Error read response body")
		return ""
	}

	var encodedBody ImgurUploadImageResponse

	err = json.Unmarshal(body, &encodedBody)
	if err != nil {
		log.Println("Error unmarshall")
		return ""
	}

	return encodedBody.Data.Link
}
