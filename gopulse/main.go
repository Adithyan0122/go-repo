package main

import (
	_ "encoding/json"
	_ "flag"
	_ "fmt"
	_ "net/http"
	_ "os"
	_ "strings"
	"time"
	_ "time"
)

//URLInput maps directly to the structure of targets.json

/*Why is it URLs and not urls?
Names starting with an uppercase letter are exported (public).
Names starting with lowercase are unexported (private to the package).
There are no public or private keywords. Capitalization is the access control.
This matters here because the encoding/json package lives outside your main package.
For it to write into your field, the field must be visible to it. If you wrote urls []string,
it would compile fine, but json.Unmarshal would silently skip the field and you'd end up with an empty list and no error.
*/

type URLInput struct {
	URLs []string `json:"urls"`
}

// Config holds the CLI configuration settings.
type Config struct {
	Workers  int
	Timeout  time.Duration
	FilePath string
}

// Result packages both the metrics and any encountered error.

type Result struct {
	URL        string
	StatusCode int
	Duration   time.Duration
	Err        error
}
