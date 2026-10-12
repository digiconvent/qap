package story

import qap "github.com/digiconvent/qap/middleware"

var (
	mastersWord      = qap.Policy{Id: "the master's word", Name: "the master's word", Groups: []string{masterOfBagEnd}, FavourThreshold: 1}
	councilsConsent  = qap.Policy{Id: "the council's consent", Name: "the council's consent", Groups: []string{council}, FavourThreshold: 4, AgainstThreshold: number(2), CommentsInAudit: true}
	halfTheCouncil   = qap.Policy{Id: "half the council", Name: "half the council", Groups: []string{council}, FavourThreshold: -50, WithholdPermitted: true}
	fellowshipAgrees = qap.Policy{Id: "the fellowship agrees", Name: "the fellowship agrees", Groups: []string{fellowship}, FavourThreshold: 3}
	kingsCouncil     = qap.Policy{Id: "the king's council", Name: "the king's council", Groups: []string{court}, FavourThreshold: 2, AgainstThreshold: number(3), CommentsInAudit: true}
	kingsWord        = qap.Policy{Id: "the king's word", Name: "the king's word", Groups: []string{KingsAuthority}, FavourThreshold: 1}
	entmoot          = qap.Policy{Id: "the entmoot", Name: "the entmoot", Groups: []string{ents}, FavourThreshold: 2, Quorum: number(4)}
	bearersChoice    = qap.Policy{Id: "the bearer's choice", Name: "the bearer's choice", Groups: []string{ringBearers}, FavourThreshold: 2}
	parleyAtOrthanc  = qap.Policy{Id: "the parley", Name: "the parley", Groups: []string{parley}, FavourThreshold: 3, DeadlineSeconds: number(45), CommentsInAudit: true}
	keepersOfTheBook = qap.Policy{Id: "the keepers of the book", Name: "the keepers of the book", Groups: []string{bagEnd}, FavourThreshold: 2}
)

var policies = map[string][]qap.Policy{
	"shire.bag-end.bequeath":      {mastersWord},
	"council.ring.decide":         {councilsConsent},
	"council.fellowship.form":     {halfTheCouncil},
	"fellowship.ring.bearer":      {bearersChoice},
	"isengard.peace.offer":        {parleyAtOrthanc},
	"fellowship.journey.plan":     {fellowshipAgrees},
	"rohan.court.grant-authority": {kingsCouncil},
	"rohan.court.banish":          {kingsWord},
	"fangorn.isengard.march":      {entmoot},
	"shire.red-book.read":         {keepersOfTheBook},
}
