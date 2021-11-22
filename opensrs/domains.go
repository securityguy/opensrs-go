package opensrs

import (
	"fmt"
	"strconv"
)

// DomainsService handles communication with the domain related methods of the OpenSRS API.
type DomainsService struct {
	client *Client
}

// GetDomain fetches a domain.
func (s *DomainsService) GetDomain(domainIdentifier string, domainType string, limit int) (*OpsResponse, error) {
	opsResponse := OpsResponse{}
	opsRequestAttributes := OpsRequestAttributes{Domain: domainIdentifier, Type: domainType}
	if limit > 0 {
		opsRequestAttributes.Limit = strconv.Itoa(limit)
	}

	resp, err := s.client.post("GET", "DOMAIN", opsRequestAttributes, &opsResponse)
	if err != nil {
		return nil, err
	}
	_ = resp
	return &opsResponse, nil
}

// UpdateDomainNameservers changes domain servers on a domain.
func (s *DomainsService) UpdateDomainNameservers(domainIdentifier string, newDs []string) (*OpsResponse, error) {
	opsResponse := OpsResponse{}
	opsRequestAttributes := OpsRequestAttributes{Domain: domainIdentifier, AssignNs: newDs, OpType: "assign"}

	resp, err := s.client.post("ADVANCED_UPDATE_NAMESERVERS", "DOMAIN", opsRequestAttributes, &opsResponse)
	if err != nil {
		return nil, err
	}
	_ = resp
	return &opsResponse, nil
}

// NewRegistration returns a RegistrationRequest structure
func (s *DomainsService) NewRegistration() *RegistrationRequest {
	// Use most common defaults
	reg := &RegistrationRequest{
		dsvc:              s,
		Domain:            "",
		CustomNameservers: false,
		CustomTechContact: false,
		LockDomain:        true,
		WhoisPrivacy:      false,
		Period:            1,
		RegType:           "new",
		Handle:            "process",
		Username:          "",
		Password:          "",
	}
	return reg
}

// Register a domain.
func (r *RegistrationRequest) Register() (*OpsResponse, error) {
	opsResponse := OpsResponse{}
	opsRequestAttributes := OpsRequestAttributes{
		Domain:            r.Domain,
		ContactSet:        *r.ContactSet,
		CustomNameservers: bool01(r.CustomNameservers),
		CustomTechContact: bool01(r.CustomTechContact),
		FLockDomain:       bool01(r.LockDomain),
		FWhoisPrivacy:     bool01(r.WhoisPrivacy),
		Period:            fmt.Sprintf("%d", r.Period),
		RegUsername:       r.Username,
		RegPassword:       r.Password,
		RegType:           r.RegType,
		Handle:            r.Handle,
	}

	resp, err := r.dsvc.client.post("SW_REGISTER", "DOMAIN", opsRequestAttributes, &opsResponse)
	if err != nil {
		return nil, err
	}
	_ = resp
	return &opsResponse, nil
}

// Convert boolean to a string containing "0" or "1"
func bool01(b bool) string {
	if b {
		return "1"
	}
	return "0"
}
