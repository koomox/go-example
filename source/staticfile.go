package main

import (
	"fmt"
	"net/http"
	"github.com/koomox/static"
	"strings"
)

func main() {
	libs := append(static.JQuery, static.Bootstrap...)
	libs = append(libs, static.FontAwesome...)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/", "/index.html":
			for _, f := range files {
				if strings.EqualFold("/index.html", f.Name) {
					w.Header().Set("Content-Type", f.ContentType)
					w.WriteHeader(http.StatusOK)
					b, err := static.DeCompress(f.Content)
					if err != nil {
						fmt.Fprintf(w, "%v Error Not found!", r.URL.Path)
						return
					}
					fmt.Fprintf(w, "%v", string(b))
					return
				}
			}
			w.Header().Set("Content-Type", "text/plain; charset=utf-8")
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "%v Error Not found!", r.URL.Path)
		default:
			if strings.HasPrefix(r.URL.Path, "/ajax/libs") {
				for _, f := range libs {
					if strings.EqualFold(r.URL.Path, f.Name) {
						w.Header().Set("Content-Type", f.ContentType)
						w.WriteHeader(http.StatusOK)
						b, err := static.DeCompress(f.Content)
						if err != nil {
							fmt.Fprintf(w, "%v Error Not found!", r.URL.Path)
							return
						}
						fmt.Fprintf(w, "%v", string(b))
						return
					}
				}
			}
			
			w.Header().Set("Content-Type", "text/plain; charset=utf-8")
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "%v Error Not found!", r.URL.Path)
		}
	})

	http.ListenAndServe(":3000", nil)
}

var files = []static.File{
	static.File{
		Name: "/index.html",
		ContentType: "text/html; charset=utf-8",
		Content: []byte("eJyslN9r2zAQx98L/R80ddAEaivdChudbBjtxh4G20ML2+NZvtSX6kcqXZxmpf/7sJ2mSbsxCvOLdKevPnfcnaxfnX87u/j5/ZNo2Nlyf093q7DgrwqJXvYehLrc3xNCCO2QQZgGYkIu5OXF5+y93Dnz4LCQLeFyHiJLYYJn9FzIJdXcFDW2ZDDrjSNBnpjAZsmAxeI4n+yyGuZ5hjcLagv5I7v8mJ0FNwemyuIWmLDA+go3V5nYYnmBibUa9usDS/5aRLSFTLyymBpElqKJOC2kghncKktVUtPgOYMlpuBQneTv8okyadedO/K5SUm+EF2FwIkjzNVJfpIf99yN7xk0mUhzFimaHcjsZoFxpd72hMHIeqMHzJIstRqu/gP0NJvZ02RewFonNSz5TTShxj8xtFpPk65CvXqAemgFRILMQtUVsYoItYkLVz0Uo5cFK4yFlP4m6BvxXJIRo5OlhnVDDmSpN7IpiClkTXDYrdNlly2VX4JDraDUytJ/CVCFcL0d4CtVEeLqhTEEGKYW5VAss4ixfwFz6Mb/HBiesLQKdmNq5aF9tGpqBdWFHHrVlfFRWNOWcKf/w/d6VAezcOh5nHcZrkbThTdMwY/Gd9vCTnp4MIQ4HK/nYnTXP/9T8WYyOWqQrhoe9oy3fCpkol8o6oDJH7JwwIxR3o8/bHPvx4/Jbs3Weqa06n9mvwMAAP//bOF4Rw"),
	},
}