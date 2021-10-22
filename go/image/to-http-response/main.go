package main

import (
	"encoding/hex"
	"encoding/json"
	"image"
	"image/color"
	"image/png"
	"net/http"
)

type ImageInfo struct {
	PosX int `json:"posX"`
	PosY int `json:"posY"`
	Width int `json:"width"`
	Height int `json:"height"`
	HeartColor string `json:"heartColor"`
	BackgroundColor string `json:"backgroundColor"`
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

// http request
// POST /image
/**
{
    "posX": 0,
    "posY": 0,
    "width": 500,
    "height": 500,
    "heartColor": "FFFF0000",
    "backgroundColor": "FFFFF000"
}
*/

func main() {

	// (json data) -> (특정 타입 데이터)
	// json.Unmarshal
	// json.NewDecoder

	// (특정 타입 데이터) -> (json data)
	// json.Marshal
	// json.NewEncoder

	//RESTful API

	http.HandleFunc("/image", func(writer http.ResponseWriter, request *http.Request) {
		var from ImageInfo
		_ = json.NewDecoder(request.Body).Decode(&from)
		_ = request.Body.Close()


		width := from.Width
		height := from.Height
		hWidth := width / 2
		hHeight := from.Height / 2

		heartColor, _ := hex.DecodeString(from.HeartColor)
		backgroundColor, _ := hex.DecodeString(from.BackgroundColor)
		// A R G B
		img := image.NewRGBA(image.Rect(0, 0, width, height))
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				dx := x - hWidth - from.PosX
				dy := hHeight - y - from.PosY
				if (dx*dx)-(abs(dx)*dy)+(dy*dy) <= 5000 {
					img.SetRGBA(x, y, color.RGBA{
						R: heartColor[1],
						G: heartColor[2],
						B: heartColor[3],
						A: heartColor[0],
					})
				} else {
					img.SetRGBA(x, y, color.RGBA{
						R: backgroundColor[1],
						G: backgroundColor[2],
						B: backgroundColor[3],
						A: backgroundColor[0],
					})
				}
			}
		}

		writer.Header().Set("Content-type", "image/png")
		png.Encode(writer, img)
	})
	http.ListenAndServe(":8000", nil)
}
