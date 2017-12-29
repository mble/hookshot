package version_test

import (
	"encoding/json"
	"testing"

	"github.com/mble/hookshot/version"
)

func TestNewVersionData(t *testing.T) {
	version.VERSION = "0.1.0"
	version.GITCOMMIT = "2c0aa57"

	data := version.NewVersionData()

	if data.Version != version.VERSION {
		t.Errorf("expected version: %s, got: %s", version.VERSION, data.Version)
	}

	if data.Build != version.GITCOMMIT {
		t.Errorf("expected version: %s, got: %s", version.GITCOMMIT, data.Build)
	}
}

func TestVersionDataSerialisation(t *testing.T) {
	version.VERSION = "0.1.0"
	version.GITCOMMIT = "2c0aa57"
	expectedJSON := "{\"version\":\"0.1.0\",\"build\":\"2c0aa57\"}"

	data := version.NewVersionData()
	json, err := json.Marshal(data)

	if err != nil {
		t.Errorf("marshalling error: %s", err)
	}

	if string(json) != expectedJSON {
		t.Errorf("expected JSON: %s, got: %s", expectedJSON, string(json))
	}
}
