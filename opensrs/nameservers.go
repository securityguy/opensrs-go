package opensrs

func (n NameserverList) ToString() []string {
	domains := make([]string, len(n))
	for i, ns := range n {
		domains[i] = ns.Name
	}
	return domains
}
