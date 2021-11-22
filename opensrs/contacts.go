package opensrs

func NewContactSet() *ContactSet {
	cs := &ContactSet{}
	return cs
}

func NewContact() *ContactObject {
	co := &ContactObject{}
	return co
}

func (c *ContactSet) SetOwner(co *ContactObject) {
	c.Owner = *co
}

func (c *ContactSet) SetAdmin(co *ContactObject) {
	c.Admin = *co
}

func (c *ContactSet) SetBilling(co *ContactObject) {
	c.Billing = *co
}

func (c *ContactSet) SetTech(co *ContactObject) {
	c.Tech = *co
}

func (c *ContactSet) SetAll(co *ContactObject) {
	c.SetOwner(co)
	c.SetAdmin(co)
	c.SetBilling(co)
	c.SetTech(co)
}
