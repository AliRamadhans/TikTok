package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func main() {

	url := "https://tiktok-scraper2.p.rapidapi.com/video/no_watermark?video_url=https%3A%2F%2Fwww.tiktok.com%2F%40tiktok%2Fvideo%2F7182229845932363051"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", "0c844eb191mshb938af2853ef196p1fe4bajsn9e79455706d9")
	req.Header.Add("X-RapidAPI-Host", "tiktok-scraper2.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}
