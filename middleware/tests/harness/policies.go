package harness

import qap "github.com/digiconvent/qap/middleware"

func number(n int) *int {
	return &n
}

var (
	PeerReview     = qap.Policy{Id: "peer review", Name: "peer review", Groups: []string{Fellowship}, FavourThreshold: 2}
	SeniorApproval = qap.Policy{Id: "senior approval", Name: "senior approval", Groups: []string{WhiteCouncil}, FavourThreshold: 3}
	VetoPower      = qap.Policy{Id: "veto power", Name: "veto power", Groups: []string{WhiteCouncil}, FavourThreshold: 2, AgainstThreshold: number(1), CommentsInAudit: true}
	FastTrack      = qap.Policy{Id: "fast track", Name: "fast track", Groups: []string{Fellowship}, FavourThreshold: 1}
	KingsCouncil   = qap.Policy{Id: "the king's council", Name: "the king's council", Groups: []string{CourtOfEdoras}, FavourThreshold: 3}
	RegentsWord    = qap.Policy{Id: "the regent's word", Name: "the regent's word", Groups: []string{CourtOfEdoras}, FavourThreshold: 3, DeadlineSeconds: number(60)}
	FullFellowship = qap.Policy{Id: "full fellowship", Name: "full fellowship", Groups: []string{Fellowship}, FavourThreshold: 2, Quorum: number(3)}
	SecretBallot   = qap.Policy{Id: "secret ballot", Name: "secret ballot", Groups: []string{WhiteCouncil}, FavourThreshold: 2}

	HalfTheCourt = qap.Policy{Id: "half the court", Name: "half the court", Groups: []string{CourtOfEdoras}, FavourThreshold: -50, Quorum: number(3), WithholdPermitted: true}
	ShireLibrary = qap.Policy{Id: "shire library", Name: "shire library", Groups: []string{Fellowship}, FavourThreshold: 1}
)

func policies() map[string][]qap.Policy {
	return map[string][]qap.Policy{
		"iam.user.create":                  {PeerReview, FastTrack},
		"iam.user.update":                  {PeerReview},
		"iam.user.history":                 {PeerReview},
		"iam.user.set-password":            {PeerReview},
		"iam.user.delete":                  {SeniorApproval},
		"iam.group.create":                 {PeerReview},
		"iam.group.delete":                 {SeniorApproval},
		"iam.group.add-user":               {PeerReview},
		"iam.group.remove-user":            {PeerReview, KingsCouncil, RegentsWord},
		"iam.policy.update":                {PeerReview, FullFellowship},
		"iam.policy.delete":                {SeniorApproval, SecretBallot},
		"iam.policy.add-permission":        {VetoPower},
		"rohan.army.muster":                {HalfTheCourt},
		"library." + ShireId + ".add-book": {ShireLibrary},
	}
}
