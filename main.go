package main

import (
	"io/ioutil"
	"net/http"
	"os"
	"regexp"

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

// From http://www.golangprograms.com/remove-duplicate-values-from-slice.html
func unique(stringSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range stringSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func grabSAN(ids []string) (sans []string) {
	for _, id := range ids {
		logger.Debug("parsing ID ", id, "..")
		req, err := http.NewRequest("GET", "https://crt.sh/", nil)
		if err != nil {
			logger.Debug(err)
		}
		q := req.URL.Query()
		q.Add("id", id)
		q.Add("output", "json")
		req.URL.RawQuery = q.Encode()

		res, err := http.Get(req.URL.String())
		if err != nil {
			logger.Debug(err)
		}
		defer res.Body.Close()
		if res.StatusCode != 200 {
			logger.Debug("status code error: ", res.StatusCode, res.Status)
		}

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			logger.Debug(err)
			return
		}

		src := string(body)

		r, _ := regexp.Compile(`;DNS:([a-z\.]+)<BR>`)
		matches := r.FindAllStringSubmatch(src, -1)
		for _, match := range matches {
			sans = append(sans, match[1])
		}

	}

	return unique(sans)
}

func getIDs(s string) (ids []string) {
	logger.Debug("grabbing IDs for ", s, "..")
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

	doc.Find("td.outer table tbody tr td:nth-child(1)").Each(func(i int, s *goquery.Selection) {
		result := s.Find("a").Text()
		ids = append(ids, result)
	})
	return ids
}

func main() {
	query := opts["query"]
	ids := getIDs(query)
	san := grabSAN(ids)
	logger.Info(san)

}
