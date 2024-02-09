package fetch

import (
	"encoding/json"
	"errors"
	"sokwva/acfun/dougaInfo/structs"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const (
	VideoApiUrl     = "https://www.acfun.cn/v/ac"
	BangumiApiUrl   = "https://www.acfun.cn/bangumi/aa"
	ArticleApiUrl   = "https://www.acfun.cn/a/ac"
	UserDougaApiUrl = "https://www.acfun.cn/u/"

	VideoApiParam       = "?quickViewId=videoInfo_new&reqID=3&ajaxpipe=1&t=1641532997358"
	ArticleApiParam     = "?quickViewId=videoInfo_new&reqID=3&ajaxpipe=1&t=1641532997358"
	BangumiApiParam     = "?quickViewId=videoInfo&reqID=2&ajaxpipe=1&t=1642342576952"
	UserVideoApiParam   = "?quickViewId=ac-space-video-list&reqID=1&ajaxpipe=1&type=video&order=newest&t=1641623740970"
	UserArticleApiParam = "?quickViewId=ac-space-article-list&reqID=1&ajaxpipe=1&type=article&order=newest&t=1641623740970"
	UserAlbumApiParam   = "?quickViewId=ac-space-album-list&reqID=1&ajaxpipe=1&type=album&order=newest&t=1641623740970"
)

type apiResp struct {
	HTML    string   `json:"html"`
	Style   []string `json:"styles"`
	Scripts []string `json:"scripts"`
	Js      []string `json:"js"`
	Css     []string `json:"css"`
}

func rawToHTML(raw string) (string, error) {
	raw = strings.Replace(raw, "/*<!-- fetch-stream -->*/", "", -1)
	x := &apiResp{}
	err := json.Unmarshal([]byte(raw), x)
	if err != nil {
		return "", err
	}
	return x.HTML, nil
}

func rawHTMLToJson(raw string, featureStr string) ([]byte, error) {
	document, err := goquery.NewDocumentFromReader(strings.NewReader(raw))
	if err != nil {
		return []byte{}, err
	}
	var targetContent string
	document.Find("script.videoInfo").Each(func(i int, s *goquery.Selection) {
		targetContent = s.Text()
	})
	targetContent = strings.Replace(targetContent, featureStr, "", -1)
	return []byte(targetContent), nil
}

func preprocess(raw []byte, target interface{}, dougaType string) error {
	htmlResult, err := rawToHTML(string(raw))
	if err != nil {
		return err
	}
	var jsonResult []byte
	switch dougaType {
	case "video":
		jsonResult, err = rawHTMLToJson(htmlResult, "\n        window.pageInfo = window.videoInfo =")
	case "bangumi":
		jsonResult, err = rawHTMLToJson(htmlResult, "\n        window.pageInfo = window.bangumiData =")
	}
	if err != nil {
		return err
	}
	err = json.Unmarshal(jsonResult, target)
	if err != nil {
		return err
	}
	return nil
}

func GetVideoInfo(acid string) (*structs.DougaInfo, error) {
	result, err := Do(VideoApiUrl + acid + VideoApiParam)
	if err != nil {
		return nil, err
	}
	resultStruct := &structs.DougaInfo{}
	err = preprocess(result, resultStruct, "video")
	return resultStruct, err
}

// Deprecated: no such api interface.
func GetArticleInfo(acid string) (*structs.DougaInfo, error) {
	return nil, errors.New("imple me")
}

func GetBangumiInfo(aaid string) (*structs.BangumiInfo, error) {
	result, err := Do(BangumiApiUrl + aaid + BangumiApiParam)
	if err != nil {
		return nil, err
	}
	resultStruct := &structs.BangumiInfo{}
	err = preprocess(result, resultStruct, "bangumi")
	return resultStruct, err
}

func GetUserVideoDouga(userId string, page uint, pageSize uint) (*[]structs.DougaInfo, error) {
	raw, err := Do(UserDougaApiUrl + userId + UserVideoApiParam + "&page=" + strconv.Itoa(int(page)) + "&pageSize=" + strconv.Itoa(int(pageSize)))
	if err != nil {
		return &[]structs.DougaInfo{}, err
	}
	htmlResult, err := rawToHTML(string(raw))
	if err != nil {
		return &[]structs.DougaInfo{}, err
	}
	document, err := goquery.NewDocumentFromReader(strings.NewReader(htmlResult))
	if err != nil {
		return &[]structs.DougaInfo{}, err
	}
	videoList := []structs.DougaInfo{}
	document.Find("a.ac-space-video").Each(func(i int, s *goquery.Selection) {
		acid, _ := s.Attr("href")
		acid = strings.Replace(acid, "/v/ac", "", 1)
		result, err := GetVideoInfo(acid)
		if err != nil {
			return
		}
		videoList = append(videoList, *result)
	})
	return &videoList, nil
}
