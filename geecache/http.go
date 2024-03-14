package geecache

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

// default path e.g. http://www.xxx.com/_geecache/
const defaultBasePath = "/_geecache/"

type HTTPPool struct {
	self     string
	basePath string
}

func NewHTTPPool(self string) *HTTPPool {
	return &HTTPPool{
		self:     self,
		basePath: defaultBasePath,
	}
}

// print log info
func (p *HTTPPool) Log(format string, v ...any) {
	log.Printf("[Server %s] %s", p.self, fmt.Sprintf(format, v...))
}

func (p *HTTPPool) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !strings.HasPrefix(r.URL.Path, p.basePath) { // not same with the defaultpath
		panic("HTTPPool serving unexpected path: " + r.URL.Path)
	}
	// print log info contains method and path
	p.Log("%s %s", r.Method, r.URL.Path)
	// first use slice to split and the use / to split
	parts := strings.SplitN(r.URL.Path[len(p.basePath):], "/", 2)
	// <groupname>/<key> required
	if len(parts) != 2 {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	// /<basepath>/<groupname>/<key> required
	groupName := parts[0]
	key := parts[1]

	group := GetGroup(groupName)
	if group == nil {
		http.Error(w, "no such group"+groupName, http.StatusNotFound)
		return
	}

	view, err := group.Get(key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(view.ByteSlice())
}
