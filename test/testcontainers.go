//go:build !exclude_testcontainers

package test

import (
	"context"
	"fmt"

	"github.com/docker/go-connections/nat"
	"github.com/gemurdock/KeyFinder-GoLang/config"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

type Container interface {
	Create(ctx context.Context) error
	Destroy() error
}

type PostgresContainer struct {
	appConfig *config.Config
	ctx       context.Context
	postgresC testcontainers.Container
	host      string
	port      string
}

func (p *PostgresContainer) GetConnInfo() (string, string) {
	return p.host, p.port
}

func (p *PostgresContainer) GetConfig() *config.Config {
	return p.appConfig
}

func (p *PostgresContainer) Create() error {
	p.ctx = context.Background()
	config := config.GetConfigInstance(false)
	config.LoadTestingValues()
	p.appConfig = config

	if p.appConfig == nil {
		panic("appConfig is nil")
	}

	req := testcontainers.ContainerRequest{
		Image:        "postgres",
		ExposedPorts: []string{p.appConfig.DBPort + "/tcp"},
		Env: map[string]string{
			"POSTGRES_HOST":     p.appConfig.DBHost,
			"POSTGRES_PORT":     p.appConfig.DBPort,
			"POSTGRES_USER":     p.appConfig.DBUser,
			"POSTGRES_PASSWORD": p.appConfig.DBPassword,
			"POSTGRES_DB":       p.appConfig.DBName,
		},
		WaitingFor: wait.ForLog("database system is ready to accept connections"),
	}

	postgresC, err := testcontainers.GenericContainer(p.ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	p.postgresC = postgresC
	if err != nil {
		return err
	}

	p.host, p.port, err = getContainerMapping(p.ctx, p.postgresC, p.appConfig.DBPort)
	if err != nil {
		return err
	}
	p.appConfig.DBHost, p.appConfig.DBPort = p.host, p.port // update config with test container values

	return nil
}

func (p *PostgresContainer) Destroy() error {
	if p.postgresC == nil {
		return fmt.Errorf("postgres container is nil")
	}

	err := p.postgresC.Terminate(p.ctx)
	if err != nil {
		return err
	}
	return nil
}

func getContainerMapping(ctx context.Context, container testcontainers.Container, port string) (string, string, error) {
	containerHost, err := container.Host(ctx)
	if err != nil {
		return "", "", err
	}
	containerPort, err := container.MappedPort(ctx, nat.Port(port))
	if err != nil {
		return "", "", err
	}
	return containerHost, containerPort.Port(), nil
}
