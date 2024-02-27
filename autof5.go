package autof5

import (
	"fmt"
	"net/http"
	"strings"
)

const js = `
<script>
let _wait = () => fetch("/_autoF5_wait", { mode: 'no-cors' }).catch(function(err) {
	let refresh = () => {
		fetch("/_autoF5").then(() => {
			location.reload();
			_wait();
		}).catch(() => {
			setTimeout(() => {
				refresh();
			}, 500);
		});
	};
	refresh();            
});
_wait();
</script>
`

type responseRecorder struct {
	w      http.ResponseWriter
	status int
	body   string
}

func (r *responseRecorder) Header() http.Header {
	return r.w.Header()
}

func (r *responseRecorder) Write(b []byte) (int, error) {
	r.body += string(b)
	return len(b), nil
}

func (r *responseRecorder) WriteHeader(status int) {
	r.status = status
	r.w.WriteHeader(status)
}

func (r *responseRecorder) Flush() {
	_, _ = r.w.Write([]byte(r.body))
}

func AutoF5(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/_autoF5_wait" {
			<-r.Context().Done()
			return
		}
		if r.URL.Path == "/_autoF5" {
			fmt.Fprintf(w, "ok")
			return
		}

		rec := &responseRecorder{w: w}
		next.ServeHTTP(rec, r)
		rec.body = strings.Replace(rec.body, "</body>", js+"</body>", 1)
		rec.Flush()
	})
}
