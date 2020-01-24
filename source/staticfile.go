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
		Content: []byte("eJyck89u1DAQxu+V+g7GXHFMRQ8c7EioBXFAgsNWguPYGZpZ/Ce1Z1P27dF601U3CCE4JZ9n5peZL2Pz4vbzzebbl/di5Bj6ywtzeIoA6d5KTLKdIAz95YUQQpiIDMKPUCqylXebD+qtPIsliGjlTPg45cJS+JwYE1v5SAOPdsCZPKomXglKxARBVQ8B7VX3+pw1Mk8KH3Y0W/lV3b1TNzlOwOQCPgMTWhzu8VTKxAH7DVY2+vi+BAKlH6JgsLLyPmAdEVmKseB3KzVs4acO5Kp2OXPlApO+7q67K+3rs7MuUup8raevVV9oYlGLP4NsH3ZY9vpNIxyFaqIBtlX2Rh9L/wJad7NdN7NiGb38LuPysH+CJ5gFFAIVwB0McAVh8GUX3dMgLS0H4QPU+qeEZuLvKYoYo+wNLGa+lP3HHNFo6I0O9F+AT+QKlP0/MgR4phnlcVi/K6WtyASH/bgFhhXL6BxO0ugEc7Nwsc7odil+BQAA//8f5vWZ"),
	},
}