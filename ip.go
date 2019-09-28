package iputils

import (
	"fmt"
	"math/big"
	"net"
)

var (
	Zero    = big.NewInt(0)
	MaxIPv4 = big.NewInt(0).SetBytes([]byte{0xff, 0xff, 0xff, 0xff})
	MaxIPv6 = big.NewInt(0).SetBytes([]byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff})

	RangeError = fmt.Errorf("range error")
)

func IsIPv4(ip net.IP) bool {
	p := ip
	if p.To4() != nil {
		return true
	}
	return false
}

func IsIPv6(ip net.IP) bool {
	p := ip
	if p.To4() != nil {
		return false
	}
	return true
}

func Add(ip net.IP, n uint64) (net.IP, error) {
	return calc(ip, big.NewInt(0).SetUint64(n))
}

func AddBigInt(ip net.IP, n *big.Int) (net.IP, error) {
	return calc(ip, n)
}

func calc(ip net.IP, n *big.Int) (net.IP, error) {
	var ipLen int
	var max *big.Int
	p := ip
	if IsIPv4(p) {
		p = p.To4()
		max = MaxIPv4
		ipLen = 4
	} else {
		p = p.To16()
		max = MaxIPv6
		ipLen = 16
	}
	z := big.NewInt(0)
	ipz := big.NewInt(0).SetBytes([]byte(p))
	z.Add(ipz, n)
	if z.Cmp(Zero) == -1 || z.Cmp(max) == 1 {
		return nil, RangeError
	}
	res := z.Bytes()
	buf := make([]byte, ipLen-len(res))
	buf = append(buf, res...)
	return net.IP(buf), nil

}

func Sub(ip net.IP, n uint64) (net.IP, error) {
	return SubBigInt(ip, big.NewInt(0).SetUint64(n))
}

func SubBigInt(ip net.IP, n *big.Int) (net.IP, error) {
	return calc(ip, big.NewInt(0).Neg(n))
}
