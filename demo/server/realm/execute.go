package realm

import (
	"github.com/digiconvent/qap/demo/server/story"
	qap "github.com/digiconvent/qap/middleware"
)

type payload map[string]any

type action func(r *Realm, request qap.Request[payload]) qap.Result

type endpoint struct {
	realm  *Realm
	action action
}

func (e endpoint) Check(p payload) map[string][]error {
	return nil
}

func (e endpoint) Execute(request qap.Request[payload]) qap.Result {
	e.realm.mutex.Lock()
	defer e.realm.mutex.Unlock()
	e.realm.executed = append(e.realm.executed, execution{request.Method, request.Url, request.Payload, request.User})
	return e.action(e.realm, request)
}

func (r *Realm) EndpointFor(permission string) qap.Handler {
	switch permission {
	case "rohan.court.grant-authority":
		return qap.Handle(endpoint{r, grantsAuthority})
	case "shire.red-book.read":
		return qap.Handle(endpoint{r, readsTheRedBook})
	}
	return qap.Handle(endpoint{r, carriesOut})
}

func carriesOut(r *Realm, request qap.Request[payload]) qap.Result {
	return qap.Result{Status: 204, Headers: map[string]string{}}
}

func grantsAuthority(r *Realm, request qap.Request[payload]) qap.Result {
	if to, ok := request.Payload["to"].(string); ok {
		r.grantAuthority(to)
	}
	return carriesOut(r, request)
}

func readsTheRedBook(r *Realm, request qap.Request[payload]) qap.Result {
	return qap.Result{Status: 200, Headers: map[string]string{"content-type": "application/json"}, Body: story.RedBook}
}
