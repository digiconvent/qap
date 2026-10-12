package host

import (
	"time"

	"github.com/digiconvent/qap/middleware/internal/concepts"
	"github.com/digiconvent/qap/middleware/internal/layer1"
	"github.com/digiconvent/qap/middleware/internal/layer2"
	"github.com/digiconvent/qap/middleware/internal/layer3"
	"github.com/digiconvent/qap/middleware/internal/layer4"
	"github.com/gofiber/fiber/v3"
)

type Host struct {
	CurrentUser    func(ctx fiber.Ctx) string
	PolicyLookup   func(permission string) []concepts.Policy
	IsMemberOf     func(userId string, groupId string) bool
	CountMembersOf func(groupIds []string) int
	AuditTrail     concepts.AuditTrail
}

type Clock struct {
	Now       func() time.Time
	AfterFunc func(delay time.Duration, callback func()) *time.Timer
	Ticks     <-chan time.Time
}

func Assemble(supplied Host, clock Clock) layer1.Guard {
	store := layer3.NewStore()
	layer3.StartCleanup(store, clock.Ticks, clock.Now)
	eligible := layer3.Eligibility(supplied.IsMemberOf)
	voterCount := layer3.VoterCount(supplied.CountMembersOf)
	execute := layer3.NewExecution(layer3.ExecutionDependencies{Store: store, Executor: layer3.CarryOut, AuditTrail: supplied.AuditTrail, Now: clock.Now})
	reject := layer3.NewRejection(layer3.RejectionDependencies{Store: store, AuditTrail: supplied.AuditTrail, Now: clock.Now})
	deadlines := layer3.NewDeadlineTimer(layer3.DeadlineTimerDependencies{
		AfterFunc: clock.AfterFunc,
		Get:       store.Get,
		Evaluate:  layer3.EvaluatePoll,
		Execute:   execute,
		Reject:    reject,
		Now:       clock.Now,
	})

	intents := layer2.New(layer2.Dependencies{
		PolicyLookup:     supplied.PolicyLookup,
		Eligible:         eligible,
		VoterCount:       voterCount,
		Store:            store,
		Deadlines:        deadlines,
		Now:              clock.Now,
		NewId:            layer4.NewSubmissionId,
		CopyRequest:      layer4.CopyRequest,
		Problems:         layer4.Problems,
		View:             layer4.SubmissionView,
		ReadVote:         layer4.ReadVote,
		WithholdAllowed:  layer4.WithholdAllowed,
		IsAuthor:         layer4.IsAuthor,
		Evaluate:         layer3.EvaluatePoll,
		Execute:          execute,
		Reject:           reject,
		RecordWithdrawal: layer3.NewWithdrawal(layer3.WithdrawalDependencies{AuditTrail: supplied.AuditTrail, Now: clock.Now}),
		Version:          layer4.Version,
	})

	return layer1.New(layer1.Dependencies{CurrentUser: supplied.CurrentUser, Intents: intents.Handlers()})
}
