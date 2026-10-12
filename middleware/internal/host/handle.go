package host

import (
	"github.com/digiconvent/qap/middleware/internal/concepts"
	"github.com/digiconvent/qap/middleware/internal/layer4"
)

func Handle[Payload any](endpoint concepts.TypedEndpoint[Payload]) concepts.Handler {
	handler := concepts.Handler{
		Check: func(body string) map[string][]error {
			payload, problems := layer4.DecodePayload[Payload](body)
			if problems != nil {
				return problems
			}
			return endpoint.Check(payload)
		},
		Execute: func(request concepts.Request) concepts.Result {
			payload, _ := layer4.DecodePayload[Payload](request.Body)
			return endpoint.Execute(concepts.TypedRequest[Payload]{
				Payload:    payload,
				Method:     request.Method,
				Url:        request.Url,
				Route:      request.Route,
				Params:     request.Params,
				Headers:    request.Headers,
				User:       request.User,
				Submission: request.Submission,
			})
		},
	}
	if reader, reads := endpoint.(concepts.Reader); reads {
		handler.Read = reader.Read
	}
	return handler
}
