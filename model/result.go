package model

const (
	Success           = 200
	IllegalParameters = 400
	Forbidden         = 403
	NotFound          = 404
	Failed            = 500
)

type UrlAnalyzerResult struct {
	Code        int32
	Err         error
	OriginalUrl string
	IsR18       bool
	Referer     string
}
