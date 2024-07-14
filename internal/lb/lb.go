package lb

import (
	"github.com/startcodextech/managerlb/internal/config"
	"github.com/startcodextech/managerlb/pkg/nginx"
	"github.com/startcodextech/managerlb/pkg/system"
	"sync"
)

type LoadBalancer struct {
	config *config.Config
	nginx  *nginx.Nginx
	apps   []config.LbApp
	module *config.Module
	mu     sync.Mutex
}

func New(os system.OSFamily, cfg *config.Config) (*LoadBalancer, error) {
	nginxLB, err := nginx.New(os)
	if err != nil {
		return nil, err
	}

	nginxVersion, err := nginxLB.GetVersion()
	if err != nil {
		return nil, err
	}

	module, err := cfg.RegistreModule(&config.Module{
		Name:          "Load Balancer",
		DefaultSystem: true,
		Installed:     false,
		Enabled:       false,
		Started:       false,
		Tool:          "nginx",
		ToolVersion:   nginxVersion,
	})
	if err != nil {
		panic(err)
	}

	lb := &LoadBalancer{
		config: cfg,
		nginx:  nginxLB,
		module: module,
		apps:   make([]config.LbApp, 0),
	}

	return lb, nil
}

func (lb *LoadBalancer) Name() string {
	return lb.module.Name
}

func (lb *LoadBalancer) IsEnabled() bool {
	return lb.module.Enabled
}

func (lb *LoadBalancer) IsInstalled() bool {
	return lb.module.Installed
}

func (lb *LoadBalancer) IsDefaultSystem() bool {
	return lb.module.DefaultSystem
}

func (lb *LoadBalancer) Install() error {
	lb.mu.Lock()
	defer lb.mu.Unlock()

	err := lb.nginx.Install()
	if err != nil {
		return err
	}

	lb.module.Enabled = true
	lb.module.Installed = true
	err = lb.config.UpdateModule(lb.module)
	if err != nil {
		return err
	}

	return nil
}

func (lb *LoadBalancer) Start() error {
	err := lb.nginx.Enable()
	if err != nil {
		return err
	}

	err = lb.nginx.Start()
	if err != nil {
		return err
	}

	lb.module.Started = true
	err = lb.config.UpdateModule(lb.module)
	if err != nil {
		return err
	}

	return nil
}

func (lb *LoadBalancer) Stop() error {
	err := lb.nginx.Disable()
	if err != nil {
		return err
	}

	err = lb.nginx.Stop()
	if err != nil {
		return err
	}

	lb.module.Started = false
	err = lb.config.UpdateModule(lb.module)
	if err != nil {
		return err
	}

	return nil
}

func (lb *LoadBalancer) Uninstall() error {
	err := lb.nginx.Uninstall()
	if err != nil {
		return err
	}

	lb.module.Enabled = false
	lb.module.Started = false
	lb.module.Installed = false
	err = lb.config.UpdateModule(lb.module)
	if err != nil {
		return err
	}
	return nil
}

func (lb *LoadBalancer) Reload() error {
	err := lb.nginx.Reload()
	if err != nil {
		lb.module.Started = false
		lb.config.UpdateModule(lb.module)
		return err
	}
	return err
}
