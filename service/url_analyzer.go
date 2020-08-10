package service

import (
	"encoding/json"
	"errors"
	"log"
	"strconv"

	"github.com/gocolly/colly"
	"github.com/valyala/fasthttp"

	. "go-picture-spider/model"
)

const (
	_pixivAjaxUrl    = "https://www.pixiv.net/touch/ajax/illust/details?illust_id="
	_donmaiUrl       = "https://danbooru.donmai.us/posts/"
	_medibangAjaxUrl = "https://medibang.com/api/picture/"
	_yandeUrl        = "https://yande.re/post/show/"
)

// 读取pixiv图片url
func Pixiv(id string, manga int64) *UrlAnalyzerResult {
	url := _pixivAjaxUrl + id

	c := &fasthttp.Client{}

	// c.Dial = fasthttpproxy.FasthttpSocksDialer("127.0.0.1:1080")

	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()

	defer func() {
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(resp)
	}()

	req.Header.SetMethod("GET")

	req.SetRequestURI(url)

	if err := c.Do(req, resp); err != nil {
		log.Println("请求失败：", err.Error())
		return &UrlAnalyzerResult{
			Code: Failed,
			Err:  errors.New("request error"),
		}
	}

	r := PixivAjaxResponse{}

	if err := json.Unmarshal(resp.Body(), &r); err != nil {
		log.Println("Decode Json String: ", err.Error())
		return &UrlAnalyzerResult{
			Code: Failed,
			Err:  errors.New("decode json error"),
		}
	}

	var originUrl string

	if manga != 0 {
		if pageCount, _ := strconv.ParseInt(r.Body.IllustDetails.PageCount, 10, 32); manga > pageCount-1 {
			return &UrlAnalyzerResult{
				Code:        IllegalParameters,
				Err:         errors.New("no such picture, check manga"),
				OriginalUrl: "",
			}
		}
		originUrl = r.Body.IllustDetails.MangaA[manga].URLBig
	} else {
		originUrl = r.Body.IllustDetails.URLBig
	}

	var isR18 bool

	if r.Body.IllustDetails.Tags[0] == "R-18" {
		isR18 = true
	}

	return &UrlAnalyzerResult{
		Code:        Success,
		Err:         nil,
		OriginalUrl: originUrl,
		IsR18:       isR18,
		Referer:     url,
	}
}

func Donmai(id string) *UrlAnalyzerResult {
	url := _donmaiUrl + id

	c := colly.NewCollector()

	var err error
	/*
		err = c.SetProxy("socks5://127.0.0.1:1080")
		if err != nil {
			fmt.Println(err)
		}
	*/

	originUrl, imgUrl := "", ""

	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.117 Safari/537.36"

	c.OnHTML("a#image-resize-link", func(e *colly.HTMLElement) {
		text := e.Attr("href")

		originUrl = text
	})

	c.OnHTML("img#image", func(e *colly.HTMLElement) {
		text := e.Attr("src")

		imgUrl = text
	})

	var xErr error
	c.OnError(func(response *colly.Response, err error) {
		xErr = err
	})

	err = c.Visit(url)

	if err != nil {
		return &UrlAnalyzerResult{
			Code: Failed,
			Err:  err,
		}
	}

	if xErr != nil {
		return &UrlAnalyzerResult{
			Code: Failed,
			Err:  xErr,
		}
	}

	if originUrl != "" {
		return &UrlAnalyzerResult{
			Code:        Success,
			OriginalUrl: originUrl,
			Referer:     url,
		}
	} else if imgUrl != "" {
		return &UrlAnalyzerResult{
			Code:        Success,
			OriginalUrl: imgUrl,
			Referer:     url,
		}
	} else {
		return &UrlAnalyzerResult{
			Code: NotFound,
			Err:  errors.New("fail get image url"),
		}
	}
}

func Medibang(id string) *UrlAnalyzerResult {
	url := _medibangAjaxUrl + id

	c := &fasthttp.Client{}

	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()

	defer func() {
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(resp)
	}()

	req.Header.SetMethod("GET")

	req.SetRequestURI(url)

	if err := c.Do(req, resp); err != nil {
		log.Println("请求失败：", err.Error())
		return &UrlAnalyzerResult{
			Code: Failed,
			Err:  errors.New("request error"),
		}
	}

	r := MedibangAjaxResponse{}

	if err := json.Unmarshal(resp.Body(), &r); err != nil {
		log.Println("Decode Json String: ", err.Error())
		return &UrlAnalyzerResult{
			Code: Failed,
			Err:  errors.New("decode json error"),
		}
	}

	originUrl := r.PictureDetailBean.ImageURLOriginal

	return &UrlAnalyzerResult{
		Code:        Success,
		OriginalUrl: originUrl,
		Referer:     url,
	}
}

func Yande(id string) *UrlAnalyzerResult {
	url := _yandeUrl + id

	c := colly.NewCollector()

	var err error
	/*
		err = c.SetProxy("socks5://127.0.0.1:1080")
		if err != nil {
			fmt.Println(err)
		}
	*/

	originUrl, higherUrl, imgUrl := "", "", ""

	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.117 Safari/537.36"

	c.OnHTML("a#png", func(e *colly.HTMLElement) {
		text, _ := e.DOM.First().Attr("href")

		originUrl = text
	})

	c.OnHTML("a#highres", func(e *colly.HTMLElement) {
		text, _ := e.DOM.First().Attr("href")

		higherUrl = text
	})

	c.OnHTML("img#image", func(e *colly.HTMLElement) {
		text, _ := e.DOM.First().Attr("src")

		imgUrl = text
	})

	var xErr error
	c.OnError(func(response *colly.Response, err error) {
		xErr = err
	})

	err = c.Visit(url)

	if err != nil {
		return &UrlAnalyzerResult{
			Code: Failed,
			Err:  err,
		}
	}

	if xErr != nil {
		return &UrlAnalyzerResult{
			Code: Failed,
			Err:  xErr,
		}
	}

	if originUrl != "" {
		return &UrlAnalyzerResult{
			Code:        Success,
			OriginalUrl: originUrl,
			Referer:     url,
		}
	} else if higherUrl != "" {
		return &UrlAnalyzerResult{
			Code:        Success,
			OriginalUrl: higherUrl,
			Referer:     url,
		}
	} else if imgUrl != "" {
		return &UrlAnalyzerResult{
			Code:        Success,
			OriginalUrl: imgUrl,
			Referer:     url,
		}
	} else {
		return &UrlAnalyzerResult{
			Code: NotFound,
			Err:  errors.New("fail get image url"),
		}
	}
}
