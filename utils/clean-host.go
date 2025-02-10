package utils

import (
	"net"
	"net/url"
)

func CleanHost(host *string) {
	if u, err := url.Parse(*host); err == nil && u.Host != "" {
		hostname, _, err := net.SplitHostPort(u.Host)
		if err == nil {
			*host = hostname
		} else {
			*host = u.Host
		}
	}
}