package main

import (
	"bytes"
	_ "embed"
	"html/template"
	"io"
	"net/http"
	"sync"
)

//go:embed index.html
var indexTpl string

func (s *server) handleHome() http.HandlerFunc {
	var (
		init   sync.Once
		tpl    *template.Template
		tplerr error
	)
	return func(w http.ResponseWriter, r *http.Request) {
		init.Do(func() {
			tpl, tplerr = template.New("").Parse(indexTpl)
		})
		if tplerr != nil {
			http.Error(w, tplerr.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)

		var buf bytes.Buffer
		tplerr = tpl.Execute(&buf, "index")
		if tplerr != nil {
			http.Error(w, tplerr.Error(), http.StatusInternalServerError)
			return
		}

		io.Copy(w, &buf)
	}
}
