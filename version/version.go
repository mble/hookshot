package version

import "net/http"

// VERSION indicates the version of the binary that is running
var VERSION string

// GITCOMMIT indicates the short SHA1 of the commit the binary was built from
var GITCOMMIT string

// Data is a struct to represent the package-level version information
type Data struct {
	Version string `json:"version"`
	Build   string `json:"build"`
}

// NewVersionData returns a Data reference to pass around when version information is needed
func NewVersionData() *Data {
	data := new(Data)
	data.Version = VERSION
	data.Build = GITCOMMIT
	return data
}

func (d *Data) Render(w http.ResponseWriter, req *http.Request) error {
	return nil
}
