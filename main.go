package main

import (
	"net/http"
	"os"

	flag "github.com/spf13/pflag"

	"github.com/PuerkitoBio/goquery"
	"github.com/n0ncetonic/crt_sheep/pkg/logger"
)

var opts = make(map[string]string)

func init() {
	var debug = flag.BoolP("debug", "d", false, "enable debug output")
	var queryFlag = flag.StringP("query", "q", "", "desired search query [required]")

	flag.Parse()

	if *queryFlag == "" {
		flag.Usage()
		os.Exit(1)
	}

	opts["query"] = *queryFlag

	if !*debug {
		logger.HideDebug()
	}

}

func getIDs(s string) (ids []string) {
	req, err := http.NewRequest("GET", "https://crt.sh/atom/", nil)
	if err != nil {
		logger.Debug(err)
	}

	q := req.URL.Query()
	q.Add("q", s)
	req.URL.RawQuery = q.Encode()

	res, err := http.Get(req.URL.String())
	if err != nil {
		logger.Info(err)
		os.Exit(1)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		logger.Debug("status code error: ", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		logger.Debug(err)
	}

	// body > table:nth-child(8) > tbody > tr > td > table > tbody > tr:nth-child(2) > td:nth-child(1)
	doc.Find("td.outer table tbody tr td:nth-child(1)").Each(func(i int, s *goquery.Selection) {
		logger.Debug(i)
		result := s.Find("a").Text()
		ids = append(ids, result)
		logger.Debug(result)
	})
	return ids
}

func main() {
	query := opts["query"]
	ids := getIDs(query)
	logger.Info(ids)

}
