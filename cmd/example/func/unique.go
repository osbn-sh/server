package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/stream", func(w http.ResponseWriter, r *http.Request) {
		url := "https://dl.rozmusic.com/Music/1404/12/19/Reza%20Sadeghi%20-%20Khoda%20Madar%20Vatan.mp3"

		// درخواست به سرور اصلی
		resp, err := http.Get(url)
		if err != nil {
			http.Error(w, "خطا در دریافت موزیک از منبع", http.StatusBadGateway)
			return
		}
		defer resp.Body.Close()

		// تنظیم نوع محتوای اصلی (MIME type)
		w.Header().Set("Content-Type", resp.Header.Get("Content-Type"))
		w.WriteHeader(resp.StatusCode)

		// انتقال مستقیم داده (stream)
		if _, err := io.Copy(w, resp.Body); err != nil {
			log.Println("خطا در ارسال استریم:", err)
		}
	})

	log.Println("Server running on http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
