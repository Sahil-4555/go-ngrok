package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	imageDir := "./images" 
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		urlPath := r.URL.Path
		fmt.Println(urlPath)

		if !strings.HasPrefix(urlPath, "/images/") {
			http.NotFound(w, r)
			return
		}

		imagePath := filepath.Join(imageDir, urlPath[8:])
		fmt.Println(imagePath)

		_, err := os.Stat(imagePath)
		if err != nil {
			http.NotFound(w, r)
			return
		}

		ext := filepath.Ext(imagePath)
		contentType := getContentType(ext)
		w.Header().Set("Content-Type", contentType)

		http.ServeFile(w, r, imagePath)
	})

	port := 8080 
	addr := fmt.Sprintf(":%d", port)

	fmt.Printf("Server started on http://localhost%s\n", addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		panic(err)
	}
}

func getContentType(extension string) string {
	switch extension {
	case ".jpg", ".jpeg":
		return "image/jpeg"
	case ".png":
		return "image/png"
	default:
		return "application/octet-stream"
	}
}
