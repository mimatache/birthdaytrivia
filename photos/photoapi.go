package photos

import (
	"bytes"
	"embed"
	"image"
	"image/jpeg"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/nfnt/resize"
)

//go:embed samples/*
var imageList embed.FS

func Register(router *mux.Router) {
	router.HandleFunc("/photo", handleImageRequest).Methods(http.MethodGet)
}

func handleImageRequest(rw http.ResponseWriter, r *http.Request) {
	fname := r.URL.Query().Get("name")
	widthStr := r.URL.Query().Get("width")
	heightStr := r.URL.Query().Get("height")

	width, err := strconv.Atoi(widthStr)
	if err != nil || width == 0 {
		width = 1920
	}

	height, err := strconv.Atoi(heightStr)
	if err != nil || height == 0 {
		height = 1080
	}

	data, err := getPhoto(fname, width/2, height/2)
	if err != nil {
		_, _ = rw.Write([]byte("error occurred" + err.Error()))
		return
	}

	_ = rw.Header().Write(bytes.NewBufferString("Content-Type: image/png"))
	_ = rw.Header().Write(bytes.NewBufferString("Content-Length: " + strconv.Itoa(len(data))))
	_, _ = rw.Write(data)
}

func getPhoto(name string, maxWidth, maxHeight int) ([]byte, error) {
	data, err := imageList.ReadFile("samples/" + name)
	if err != nil {
		return nil, err
	}
	im, _, err := image.DecodeConfig(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	overWidth := im.Width / maxWidth
	overHight := im.Height / maxHeight

	if overWidth > 1 || overHight > 1 {
		toReduce := max(overWidth, overHight)

		img, _, err := image.Decode(bytes.NewReader(data))
		if err != nil {
			return nil, err
		}

		newImage := resize.Resize(uint(im.Width/toReduce), uint(im.Height/toReduce), img, resize.Lanczos3)
		buf := new(bytes.Buffer)
		err = jpeg.Encode(buf, newImage, nil)
		if err != nil {
			return nil, err
		}
		data = buf.Bytes()
	}

	return data, nil
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
