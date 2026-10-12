package realm

import (
	"encoding/json"

	"github.com/digiconvent/qap/demo/server/story"
	qap "github.com/digiconvent/qap/middleware"
)

type payload map[string]any

type action func(r *Realm, request qap.Request, p payload) qap.Result

func (r *Realm) handler(act action) qap.Handler {
	return qap.Handler{Execute: func(request qap.Request) qap.Result {
		var p payload
		json.Unmarshal([]byte(request.Body), &p)
		r.mutex.Lock()
		defer r.mutex.Unlock()
		r.executed = append(r.executed, execution{request.Method, request.Url, p, request.User})
		return act(r, request, p)
	}}
}

func (r *Realm) EndpointFor(permission string) qap.Handler {
	switch permission {
	case "rohan.court.grant-authority":
		return r.handler(grantsAuthority)
	case "shire.red-book.read":
		return r.handler(readsTheRedBook)
	}
	return r.handler(carriesOut)
}

func carriesOut(r *Realm, request qap.Request, p payload) qap.Result {
	return qap.Result{Status: 204, Headers: map[string]string{}}
}

func grantsAuthority(r *Realm, request qap.Request, p payload) qap.Result {
	if to, ok := p["to"].(string); ok {
		r.grantAuthority(to)
	}
	return carriesOut(r, request, p)
}

func readsTheRedBook(r *Realm, request qap.Request, p payload) qap.Result {
	return qap.Result{Status: 200, Headers: map[string]string{"content-type": "application/json"}, Body: story.RedBook}
}
