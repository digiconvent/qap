package main

import (
	"sync"

	"github.com/digiconvent/qap/demo/server/realm"
	"github.com/gofiber/fiber/v3"
)

type server struct {
	mutex sync.Mutex
	realm *realm.Realm
}

func (s *server) current() *realm.Realm {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.realm
}

func (s *server) requires(permission string) fiber.Handler {
	return func(ctx fiber.Ctx) error {
		current := s.current()
		return current.Guard.RequiresPermission(permission).Run(current.EndpointFor(permission))(ctx)
	}
}

func (s *server) reset(ctx fiber.Ctx) error {
	s.mutex.Lock()
	s.realm = realm.New()
	s.mutex.Unlock()
	return ctx.SendStatus(fiber.StatusNoContent)
}

func (s *server) record(ctx fiber.Ctx) error {
	return ctx.JSON(s.current().Recorded())
}
