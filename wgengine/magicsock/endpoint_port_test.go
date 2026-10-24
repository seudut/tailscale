// Copyright (c) Tailscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

package magicsock

import (
	"net/netip"
	"testing"
)

func TestEndpointPortForAddr(t *testing.T) {
	tests := []struct {
		name   string
		ip     netip.Addr
		port4  uint16
		port6  uint16
		want   uint16
		wantOK bool
	}{
		{
			name:   "ipv4 uses pconn4 port",
			ip:     netip.MustParseAddr("192.0.2.1"),
			port4:  41641,
			port6:  53001,
			want:   41641,
			wantOK: true,
		},
		{
			name:   "ipv4-mapped uses pconn4 port",
			ip:     netip.MustParseAddr("::ffff:192.0.2.1"),
			port4:  41641,
			port6:  53001,
			want:   41641,
			wantOK: true,
		},
		{
			name:   "ipv6 uses pconn6 port",
			ip:     netip.MustParseAddr("2001:db8::1"),
			port4:  41641,
			port6:  53001,
			want:   53001,
			wantOK: true,
		},
		{
			name:   "ipv6 without pconn6 is skipped",
			ip:     netip.MustParseAddr("2001:db8::1"),
			port4:  41641,
			port6:  0,
			want:   0,
			wantOK: false,
		},
		{
			name:   "ipv4 without pconn4 is skipped",
			ip:     netip.MustParseAddr("192.0.2.1"),
			port4:  0,
			port6:  53001,
			want:   0,
			wantOK: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := endpointPortForAddr(tt.ip, tt.port4, tt.port6)
			if got != tt.want || ok != tt.wantOK {
				t.Fatalf("endpointPortForAddr(%v, %d, %d) = (%d, %v), want (%d, %v)",
					tt.ip, tt.port4, tt.port6, got, ok, tt.want, tt.wantOK)
			}
		})
	}
}
