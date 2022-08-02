package dnsmasq

type DNSMasq struct {
	Path string
}

type Interface struct {
	Address    string
	Gateway    string
	DHCPServer string
}

func (d *DNSMasq) Run() error {
	return nil
}
