package main

import (
	"net/http"
	"os"

	flag "github.com/spf13/pflag"

	"astuart.co/goq"
	"github.com/n0ncetonic/crt_sheep/pkg/logger"
)

// Structured representation for github file name table
type example struct {
	ID    string   `goquery:"table:nth-child(8) > tbody > tr > td > table > tbody > tr:nth-child(2) > td > a"`
	Files []string `goquery:"table.files tbody tr.js-navigation-item td.content,text"`
}

func initFlags() (opts map[string]string) {
	o := make(map[string]string)

	queryFlag := flag.StringP("query", "q", "", "desired search query [required]")

	flag.Parse()

	if *queryFlag == "" {
		flag.Usage()
		os.Exit(1)
	}

	o["query"] = *queryFlag
	return opts
}

func init() {
	var debug = flag.BoolP("debug", "d", false, "enable debug output")

	if !*debug {
		logger.HideDebug()
	}
}

func main() {
	// options := initFlags()
	_ = initFlags()

	res, err := http.Get("https://crt.sh/atom/?q=Dropbox.com.")
	if err != nil {
		logger.Info(err)
		os.Exit(1)
	}
	defer res.Body.Close()

	var ex example

	err = goq.NewDecoder(res.Body).Decode(&ex)
	if err != nil {
		logger.Info(err)
		os.Exit(1)
	}

	logger.Info(ex.ID)
}
