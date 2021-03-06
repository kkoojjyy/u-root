// Copyright 2018 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package strace

import (
	"golang.org/x/sys/unix"
)

const sizeOfInt32 int = 4

// Address is a byte slice cast as a string that represents the address of a
// network node. Or, in the case of unix endpoints, it may represent a path.
type Address string

type FullAddress struct {
	// Addr is the network address.
	Addr Address

	// Port is the transport port.
	//
	// This may not be used by all endpoint types.
	Port uint16
}

// GetAddress reads an sockaddr struct from the given address and converts it
// to the FullAddress format. It supports AF_UNIX, AF_INET and AF_INET6
// addresses.
func GetAddress(t *Tracer, sfamily int, addr []byte) (FullAddress, error) {
	/*
		// Make sure we have at least 2 bytes for the address family.
		if len(addr) < 2 {
			return FullAddress{}, syserr.ErrInvalidArgument
		}

		family := usermem.ByteOrder.Uint16(addr)
		if family != uint16(sfamily) {
			return FullAddress{}, syserr.ErrAddressFamilyNotSupported
		}

		// Get the rest of the fields based on the address family.
		switch family {
		case linux.AF_UNIX:
			path := addr[2:]
			if len(path) > linux.UnixPathMax {
				return FullAddress{}, syserr.ErrInvalidArgument
			}
			// Drop the terminating NUL (if one exists) and everything after
			// it for filesystem (non-abstract) addresses.
			if len(path) > 0 && path[0] != 0 {
				if n := bytes.IndexByte(path[1:], 0); n >= 0 {
					path = path[:n+1]
				}
			}
			return FullAddress{
				Addr: Address(path),
			}, nil

		case linux.AF_INET:
			var a linux.SockAddrInet
			if len(addr) < sockAddrInetSize {
				return FullAddress{}, syserr.ErrBadAddress
			}
			binary.Unmarshal(addr[:sockAddrInetSize], usermem.ByteOrder, &a)

			out := FullAddress{
				Addr: Address(a.Addr[:]),
				Port: ntohs(a.Port),
			}
			if out.Addr == "\x00\x00\x00\x00" {
				out.Addr = ""
			}
			return out, nil

		case linux.AF_INET6:
			var a linux.SockAddrInet6
			if len(addr) < sockAddrInet6Size {
				return FullAddress{}, syserr.ErrBadAddress
			}
			binary.Unmarshal(addr[:sockAddrInet6Size], usermem.ByteOrder, &a)

			out := FullAddress{
				Addr: Address(a.Addr[:]),
				Port: ntohs(a.Port),
			}
			if isLinkLocal(out.Addr) {
				out.NIC = NICID(a.Scope_id)
			}
			if out.Addr == Address(strings.Repeat("\x00", 16)) {
				out.Addr = ""
			}
			return out, nil

		default:
	*/
	return FullAddress{}, unix.ENOTSUP
	//}
}
