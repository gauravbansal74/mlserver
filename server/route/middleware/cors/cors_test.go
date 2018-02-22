package cors

import (
	"net/http"
	"testing"

	"github.com/gauravbansal74/mlserver/server/handlers"
)

func TestHandler(t *testing.T) {
	handler := Handler(http.HandlerFunc(handlers.AllDataSourcesHandler))
}
