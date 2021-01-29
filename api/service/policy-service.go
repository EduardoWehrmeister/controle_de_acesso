package service

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/persist"

	"controle_de_acesso.com/api/repository"
)

type IPolicyService interface {
	AddPolicy(role string, domain string, resource string, action string) bool
}

type policyService struct{}

var (
	cRepo           repository.ICasbinRepository
	cMongoDbAdapter persist.BatchAdapter
	enf             *casbin.Enforcer
	serviceError    error
)

func NewPolicyService(casbinRepository repository.ICasbinRepository) IPolicyService {
	cRepo = casbinRepository
	cMongoDbAdapter = casbinRepository.GetTheAdapter()
	enf, serviceError = casbin.NewEnforcer("./configuracao/rbac_with_domains_model.conf", cMongoDbAdapter)
	if serviceError != nil {
		panic(serviceError)
	}
	return &policyService{}
}

func (*policyService) AddPolicy(role string, domain string, resource string, action string) bool {
	result, errs := enf.AddPolicy(role, domain, resource, action)
	if errs != nil {
		panic(errs)
	}

	return result
}
