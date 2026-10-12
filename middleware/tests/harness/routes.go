package harness

import "net/http"

type Route struct {
	Method string
	Path   string
}

const (
	FrodoId       = "00000000-0000-4000-8000-000000000001"
	SamwiseId     = "00000000-0000-4000-8000-000000000002"
	BoromirId     = "00000000-0000-4000-8000-000000000003"
	FellowshipId  = "00000000-0000-4000-8000-000000000004"
	CourtId       = "00000000-0000-4000-8000-000000000005"
	PeerReviewId  = "00000000-0000-4000-8000-000000000006"
	VetoPowerId   = "00000000-0000-4000-8000-000000000007"
	RegentsWordId = "00000000-0000-4000-8000-000000000008"
	ShireId       = "00000000-0000-4000-8000-000000000009"
	MordorId      = "00000000-0000-4000-8000-00000000000a"
)

var (
	CreateUser          = Route{http.MethodPost, "/api/iam/user/create"}
	UpdateFrodo         = Route{http.MethodPost, "/api/iam/user/" + FrodoId + "/update"}
	UpdateSamwise       = Route{http.MethodPost, "/api/iam/user/" + SamwiseId + "/update"}
	UpdateWithoutUuid   = Route{http.MethodPost, "/api/iam/user/frodo/update"}
	RenameFrodo         = Route{http.MethodPost, "/api/iam/user/" + FrodoId + "/rename"}
	FrodosHistory       = Route{http.MethodGet, "/api/iam/user/" + FrodoId + "/history"}
	DeleteBoromir       = Route{http.MethodDelete, "/api/iam/user/" + BoromirId + "/delete"}
	CreateGroup         = Route{http.MethodPost, "/api/iam/group/create"}
	DeleteFellowship    = Route{http.MethodDelete, "/api/iam/group/" + FellowshipId + "/delete"}
	AddUserToFellowship = Route{http.MethodPost, "/api/iam/group/" + FellowshipId + "/add-user"}
	RemoveEomer         = Route{http.MethodPost, "/api/iam/group/" + CourtId + "/remove-user"}
	UpdatePolicy        = Route{http.MethodPost, "/api/iam/policy/" + PeerReviewId + "/update"}
	DeletePolicy        = Route{http.MethodDelete, "/api/iam/policy/" + VetoPowerId + "/delete"}
	AddPermission       = Route{http.MethodPost, "/api/iam/policy/" + RegentsWordId + "/add-permission"}
	MusterTheArmy       = Route{http.MethodPost, "/api/rohan/army/muster"}
	ShireAddsBook       = Route{http.MethodPost, "/api/library/" + ShireId + "/add-book"}
	MordorAddsBook      = Route{http.MethodPost, "/api/library/" + MordorId + "/add-book"}
	MisconfiguredRoute  = Route{http.MethodPost, "/api/iam/group/remove-user-without-id"}
	PublicEvents        = Route{http.MethodGet, "/api/calendar/public-events"}
	LatestNews          = Route{http.MethodGet, "/api/forum/latest-news"}
	SetFrodosPassword   = Route{http.MethodPost, "/api/iam/user/" + FrodoId + "/set-password"}
)
