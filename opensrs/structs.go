package opensrs

type OpsRequest struct {
	Action     string               `json:"action"`
	Object     string               `json:"object"`
	Protocol   string               `json:"protocol"`
	Attributes OpsRequestAttributes `json:"attributes"`
}

type OpsRequestAttributes struct {
	Domain            string         `json:"domain"`
	Limit             string         `json:"limit,omitempty"`
	Type              string         `json:"type,omitempty"`
	Data              string         `json:"data,omitempty"`
	AffectDomains     string         `json:"affect_domains,omitempty"`
	NameserverList    NameserverList `json:"nameserver_list,omitempty"`
	OpType            string         `json:"op_type,omitempty"`
	AssignNs          []string       `json:"assign_ns,omitempty"`
	ContactSet        ContactSet     `json:"contact_set,omitempty"`
	CustomNameservers string         `json:"custom_nameservers,omitempty"`
	CustomTechContact string         `json:"custom_tech_contact,omitempty"`
	FLockDomain       string         `json:"f_lock_domain,omitempty"`
	FWhoisPrivacy     string         `json:"f_whois_privacy,omitempty"`
	Period            string         `json:"period,omitempty"`
	RegUsername       string         `json:"reg_username,omitempty"`
	RegPassword       string         `json:"reg_password,omitempty"`
	RegType           string         `json:"reg_type,omitempty"`
	Handle            string         `json:"handle,omitempty"`
}

type OpsResponse struct {
	Action       string                `json:"action,omitempty"`
	Object       string                `json:"object,omitempty"`
	Protocol     string                `json:"protocol,omitempty"`
	IsSuccess    string                `json:"is_success,omitempty"`
	ResponseCode string                `json:"response_code,omitempty"`
	ResponseText string                `json:"response_text,omitempty"`
	Attributes   OpsResponseAttributes `json:"attributes,omitempty"`
}

type OpsResponseAttributes struct {
	AffiliateId        string         `json:"affiliate_id,omitempty"`
	AutoRenew          string         `json:"auto_renew,omitempty"`
	ExpireDate         string         `json:"expiredate,omitempty"`
	GdprConsentStatus  string         `json:"gdpr_consent_status,omitempty"`
	LetExpire          string         `json:"let_expire,omitempty"`
	RegistryCreateDate string         `json:"registry_createdate,omitempty"`
	RegistryExpireDate string         `json:"registry_expiredate,omitempty"`
	RegistryUpdateDate string         `json:"registry_updatedate,omitempty"`
	SponsoringRsp      string         `json:"sponsoring_rsp,omitempty"`
	NameserverList     NameserverList `json:"nameserver_list,omitempty"`
	Contacts           *ContactSet    `json:"contact_set,omitempty"`
	Type               string         `json:"type,omitempty"`
	LockState          string         `json:"lock_state,omitempty"`
	TLDData            interface{}    `json:"tld_data,omitempty"`
	AdminEmail         string         `json:"admin_email,omitempty"`
	ID                 string         `json:"id,omitempty"`
	RegistrationText   string         `json:"registration_text,omitempty"`
	RegistrationCode   string         `json:"registration_code,omitempty"`
	DomainID           string         `json:"domain_id,omitempty"`
}

type NameserverList []struct {
	Name      string `json:"name"`
	IpAddress string `json:"ipaddress,omitempty"`
	Ipv6      string `json:"ipv6,omitempty"`
	SortOrder string `json:"sortorder,omitempty"`
}

type ContactSet struct {
	Admin   ContactObject `json:"admin,omitempty"`
	Billing ContactObject `json:"billing,omitempty"`
	Owner   ContactObject `json:"owner,omitempty"`
	Tech    ContactObject `json:"tech,omitempty"`
}

type ContactObject struct {
	Address1          string `json:"address1,omitempty"`
	Address2          string `json:"address2,omitempty"`
	Address3          string `json:"address3,omitempty"`
	City              string `json:"city,omitempty"`
	Country           string `json:"country,omitempty"`
	Email             string `json:"email,omitempty"`
	Fax               string `json:"fax,omitempty"`
	FirstName         string `json:"first_name,omitempty"`
	GdprConsentStatus string `json:"gdpr_consent_status,omitempty"`
	LastName          string `json:"last_name,omitempty"`
	OrgName           string `json:"org_name,omitempty"`
	Phone             string `json:"phone,omitempty"`
	PostalCode        string `json:"postal_code,omitempty"`
	State             string `json:"state,omitempty"`
	Status            string `json:"status,omitempty"`
}

type RegistrationRequest struct {
	dsvc              *DomainsService
	Domain            string      `json:"domain,omitempty"`
	ContactSet        *ContactSet `json:"contact_set,omitempty"`
	Username          string      `json:"reg_username,omitempty"`
	Password          string      `json:"reg_password,omitempty"`
	CustomNameservers bool        `json:"custom_nameservers,omitempty"`
	CustomTechContact bool        `json:"custom_tech_contact,omitempty"`
	LockDomain        bool        `json:"f_lock_domain,omitempty"`
	WhoisPrivacy      bool        `json:"f_whois_privacy,omitempty"`
	Period            int         `json:"period,omitempty"`
	RegType           string      `json:"reg_type,omitempty"`
	Handle            string      `json:"handle,omitempty"`
}
