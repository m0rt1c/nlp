package helpers

import (
	"testing"
)

func TestParseNetLog(t *testing.T) {
	netlog, err := ParseNetLog("./netlog.json")
	if err != nil {
		t.Errorf("got %s, want nil", err)
		return
	}

	got := len(netlog.Events)
	if got != 493 {
		t.Errorf("got %v, want 493", got)
	}

	ids := netlog.Events[9].FindDependenciesIDs()
	if ids[0] != 7 && ids[1] != 10 {
		t.Errorf("got %v, want [7 19]", ids)
	}

	t1 := netlog.Events[64].Type
	if t1 != "SOCKET" {
		t.Errorf("got %v, want SOCKET", t1)
	}

	t2 := netlog.Events[127].Type
	if t2 != "DISK_CACHE_ENTRY" {
		t.Errorf("got %v, want DISK_CACHE_ENTRY", t2)
	}
}

func TestFindRedirections(t *testing.T) {
	netlog, err := ParseNetLog("./netlog.json")
	if err != nil {
		t.Errorf("got %s, want nil", err)
		return
	}

	rs := netlog.FindRedirections()
	want := "https://consent.google.com/intro/?continue=https://www.google.com/&origin=https://www.google.com&if=1&gl=IT&hl=it&pc=s"
	got := rs[0].To
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestFindURLRequests(t *testing.T) {
	netlog, err := ParseNetLog("./netlog.json")
	if err != nil {
		t.Errorf("got %s, want nil", err)
		return
	}

	rs := netlog.FindURLRequests()
	want := 102
	got := len(rs)
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestFindDNSQueries(t *testing.T) {
	netlog, err := ParseNetLog("./netlog.json")
	if err != nil {
		t.Errorf("got %s, want nil", err)
		return
	}

	rs := netlog.FindDNSQueries()
	want := 22
	got := len(rs)
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestFindOpenedSocket(t *testing.T) {
	netlog, err := ParseNetLog("./netlog.json")
	if err != nil {
		t.Errorf("got %s, want nil", err)
		return
	}

	res := netlog.FindOpenedSocket()
	if len(res) != 20 {
		t.Errorf("got %d want 20", len(res))
	}
}

func TestFindSources(t *testing.T) {
	netlog, err := ParseNetLog("./netlog.json")
	if err != nil {
		t.Errorf("got %s, want nil", err)
		return
	}

	res := netlog.FindSources()
	if len(res) != 81 {
		t.Errorf("got %d want 81", len(res))
	}
}
