package harness

import (
	"errors"
	"maps"

	qap "github.com/digiconvent/qap/middleware"
)

type Payload map[string]any

type endpoint struct {
	realm   *Realm
	records bool
	updates bool
	reply   qap.Result
}

type readingEndpoint struct {
	endpoint
}

func (e endpoint) Check(payload Payload) map[string][]error {
	e.realm.mutex.Lock()
	defer e.realm.mutex.Unlock()
	if e.realm.refusing {
		return map[string][]error{"body": {errors.New("the host refuses this request")}}
	}
	return nil
}

func (e endpoint) Execute(request qap.Request[Payload]) qap.Result {
	e.realm.mutex.Lock()
	defer e.realm.mutex.Unlock()
	if e.records {
		e.realm.executed = append(e.realm.executed, request)
	}
	if person, known := e.realm.people[request.Params["id"]]; e.updates && known {
		for field, value := range request.Payload {
			person[field] = value
		}
	}
	return e.reply
}

func (e readingEndpoint) Read(id string) any {
	e.realm.mutex.Lock()
	defer e.realm.mutex.Unlock()
	person, known := e.realm.people[id]
	if !known {
		return nil
	}
	return maps.Clone(person)
}

func (r *Realm) RecordingEndpoint() qap.Handler {
	return r.carriedOut()
}

func (r *Realm) reached() qap.Handler {
	return qap.Handle(endpoint{realm: r, reply: qap.Result{Status: 200, Body: "reached"}})
}

func (r *Realm) carriedOut() qap.Handler {
	return qap.Handle(endpoint{realm: r, records: true, reply: qap.Result{Status: 204, Headers: map[string]string{}}})
}

func (r *Realm) updated() qap.Handler {
	return qap.Handle(readingEndpoint{endpoint{realm: r, records: true, updates: true, reply: qap.Result{Status: 204, Headers: map[string]string{}}}})
}

func (r *Realm) history() qap.Handler {
	body := map[string]any{"user": "frodo", "history": []string{"Bag End", "Rivendell"}}
	return qap.Handle(endpoint{realm: r, records: true, reply: qap.Result{Status: 200, Headers: map[string]string{"content-type": "application/json"}, Body: body}})
}
