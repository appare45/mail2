package config

import (
	"fmt"
	"net"
)

type SmtpConfig struct {
	Host string
	Ip   string
	Port int
}

func (c *SmtpConfig) defaultConfig() {
	c.Port = 587
}

func (c *SmtpConfig) validate() error {
	if c.Host == "" && c.Ip == "" {
		return fmt.Errorf("smtp host is required")
	}
	return nil
}

// TODO: MXレコードがない場合はA/CNAMEレコードを検索する
/*
Once an SMTP client lexically identifies a domain to which mail will
  be delivered for processing (as described in sections 3.6 and 3.7), a
  DNS lookup MUST be performed to resolve the domain name [22].  The
  names are expected to be fully-qualified domain names (FQDNs):
  mechanisms for inferring FQDNs from partial names or local aliases
  are outside of this specification and, due to a history of problems,
  are generally discouraged.  The lookup first attempts to locate an MX
  record associated with the name.  If a CNAME record is found instead,
  the resulting name is processed as if it were the initial name.  If
  no MX records are found, but an A RR is found, the A RR is treated as
  if it was associated with an implicit MX RR, with a preference of 0,
  pointing to that host.  If one or more MX RRs are found for a given
  name, SMTP systems MUST NOT utilize any A RRs associated with that
  name unless they are located using the MX RRs; the "implicit MX" rule
  above applies only if there are no MX records present.  If MX records
  are present, but none of them are usable, this situation MUST be
  reported as an error.
 https://datatracker.ietf.org/doc/html/rfc2821#section-5
*/
func lookupIpaddr(domain string) (net.IP, error) {
	mx, err := net.LookupMX(domain)
	if err != nil {
		return nil, err
	}
	var ipaddr net.IP
	for i := 0; i < len(mx); i++ {
		ips, _ := net.LookupIP(mx[i].Host)
		for j := 0; j < len(ips); j++ {
			ipaddr = ips[j]
			break
		}
	}
	if ipaddr == nil {
		return nil, fmt.Errorf("No ip address found for %s", domain)
	}
	return ipaddr, nil
}
func (c *SmtpConfig) Ipaddr() (*net.TCPAddr, error) {
	var ipaddr net.IP
	ipaddr = net.ParseIP(c.Ip)
	if c.Host != "" {
		ipaddr, _ = lookupIpaddr(c.Host)
	}
	if ipaddr == nil {
		return nil, fmt.Errorf("Invalid ip address %s", c.Ip)
	}
	return &net.TCPAddr{
		IP:   ipaddr,
		Port: c.Port,
	}, nil
}
