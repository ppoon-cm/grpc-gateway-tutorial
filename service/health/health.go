package health

import (
	"context"

	healthcheck "github.com/myuser/myrepo/gen/go"
)

type healthserver struct{
	healthcheck.UnimplementedHealthCheckServer
}

func NewServer() *healthserver {
	return &healthserver{}
}

func (s *healthserver) HealthLive(ctx context.Context, in *healthcheck.HealthCheckRequest) (*healthcheck.HealthCheckReply, error) {
	return &healthcheck.HealthCheckReply{Status: "UP"}, nil
}
