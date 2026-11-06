package pokeapi

import (
	"testing"
	"strings"
)

func TestDecodePokePage(t *testing.T) {
	jsonBody := `{
		"count": 2, 
		"next": null, 
		"prrevious": null, 
		"results": [{"name": "foo", "url": "x"}]
	}`
	page, err := decodePokePage(strings.NewReader(jsonBody))

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(page.Results) != 1 {
		t.Fatalf("expected 1 result, got %d", len(page.Results))
	}
}

