package defaultcache

import "testing"

func TestCache(t *testing.T) {
	DefaultCache().Set("key", "wsw")
	get, err := DefaultCache().Get("key")
	if err != nil {
		t.Error("err", err)
		return
	}

	t.Log(get)

	DefaultCache().Delete("key")

	DefaultCache().Set("key", "wsw")
	get, err = DefaultCache().Get("key")
	if err != nil {
		t.Error("err", err)
		return
	}

	t.Log(get)
}
