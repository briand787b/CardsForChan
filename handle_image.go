package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"strconv"
)

func HandleImageShow(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	imageIDStr := params.ByName("imageID")
	imageID, err := strconv.Atoi(imageIDStr)
	if err != nil {
		// TODO: make this a bad request, not a 404
		http.NotFound(w, r)
		return
	}
	image, err := globalImageStore.Find(imageID)
	if err != nil {
		panic(err)
	}

	// 404
	if image == nil {
		http.NotFound(w, r)
		return
	}

	RenderTemplate(w, r, "images/show", map[string]interface{}{
		"Image": image,
	})
}
