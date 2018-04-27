package app

import (
	"fmt"
	"net/http"
	"time"
)

func pageMaker(w http.ResponseWriter, req *http.Request) {
	loc, _ := time.LoadLocation("Asia/Tokyo")
	time.Local = loc
	t := time.Now()
	fmt.Fprintf(w, "%s <br />\n", t.Format("2006-01-02 15:04:05 MST"))
}
