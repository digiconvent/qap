package harness

import (
	"encoding/json"
	"io"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"

	qap "github.com/digiconvent/qap/middleware"
)

type Member struct {
	realm *Realm
	name  string
}

func (m Member) Request(at Route, headers map[string]string, body string) Answer {
	m.realm.t.Helper()
	req := httptest.NewRequest(at.Method, at.Path, strings.NewReader(body))
	if m.name != "" {
		req.Header.Set("X-User", m.name)
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}
	resp, err := m.realm.app.Test(req)
	if err != nil {
		m.realm.t.Fatal(err)
	}
	raw, _ := io.ReadAll(resp.Body)
	a := Answer{Status: resp.StatusCode, Raw: string(raw)}
	json.Unmarshal(raw, &a.Body)
	json.Unmarshal(raw, &a.List)
	return a
}

func (m Member) Intent(at Route, intent string, headers map[string]string, body string) Answer {
	all := map[string]string{"X-QAP-Intent": intent}
	for key, value := range headers {
		all[key] = value
	}
	return m.Request(at, all, body)
}

func (m Member) Discovers(at Route) Answer {
	return m.Intent(at, "discover", nil, "")
}

func (m Member) Proposes(at Route, policy qap.Policy, body string) Answer {
	return m.Intent(at, "submit", map[string]string{"X-QAP-Policy-ID": policy.Id}, body)
}

func (m Member) ProposesWithoutPolicy(at Route) Answer {
	return m.Intent(at, "submit", nil, "")
}

func (m Member) Lists(at Route) Answer {
	return m.Intent(at, "list", nil, "")
}

func (m Member) Polls(at Route, id string) Answer {
	return m.Intent(at, "poll", map[string]string{"X-QAP-ID": id}, "")
}

func (m Member) Votes(at Route, id string, vote int) Answer {
	return m.Intent(at, "vote", map[string]string{"X-QAP-ID": id, "X-QAP-Vote": strconv.Itoa(vote)}, "")
}

func (m Member) VotesFavour(at Route, id string) Answer {
	return m.Votes(at, id, 1)
}

func (m Member) VotesAgainst(at Route, id string) Answer {
	return m.Votes(at, id, -1)
}

func (m Member) Withholds(at Route, id string) Answer {
	return m.Votes(at, id, 0)
}

func (m Member) Comments(at Route, id string, text string) Answer {
	return m.Intent(at, "comment", map[string]string{"X-QAP-ID": id, "X-QAP-Comment": url.PathEscape(text)}, "")
}

func (m Member) Withdraws(at Route, id string) Answer {
	return m.Intent(at, "delete-submission", map[string]string{"X-QAP-ID": id}, "")
}
