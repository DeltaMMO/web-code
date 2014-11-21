package code

import (
	"fmt"
	"net/http"
	"strings"
)

func init() {
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	base := r.URL.Path
	i := strings.Index(r.URL.Path[1:], "/")
	if i >= 0 {
		base = r.URL.Path[:i+1]
	}

	fmt.Fprintf(w, `<!DOCTYPE html>
<html>
	<head>
		<meta name="go-import" content="code.delta-mmo.com%v git https://github.com/DeltaMMO%v.git">
	</head>
	<body>
		Delta code hosting, see <a href="https://www.github.com/DeltaMMO%v">github.com/DeltaMMO%v</a>
	</body>
</html>`, base, base, base, base)
}
