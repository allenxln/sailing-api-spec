package wireformat

import (
	"encoding/json"
	"testing"
)

type idResp struct {
	ID int64 `json:"id,string"`
}

type optIDResp struct {
	ID *int64 `json:"id,string,omitempty"`
}

func TestInt64MarshalsToJSONString(t *testing.T) {
	b, err := json.Marshal(idResp{ID: 123})
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	if got := string(b); got != `{"id":"123"}` {
		t.Fatalf("got %s, want {\"id\":\"123\"}", got)
	}
}

func TestJSONStringUnmarshalsToInt64(t *testing.T) {
	var r idResp
	if err := json.Unmarshal([]byte(`{"id":"123"}`), &r); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if r.ID != 123 {
		t.Fatalf("got %d, want 123", r.ID)
	}
}

func TestZeroInt64MarshalsToQuotedZero(t *testing.T) {
	b, _ := json.Marshal(idResp{ID: 0})
	if got := string(b); got != `{"id":"0"}` {
		t.Fatalf("got %s, want {\"id\":\"0\"}", got)
	}
}

func TestNilPointerWithOmitemptyOmitsField(t *testing.T) {
	b, _ := json.Marshal(optIDResp{ID: nil})
	if got := string(b); got != `{}` {
		t.Fatalf("got %s, want {}", got)
	}
}
