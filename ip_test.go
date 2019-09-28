package iputils_test

import (
	"math"
	"math/big"
	"net"
	"testing"

	"github.com/mimuret/iputils"
	"github.com/stretchr/testify/assert"
)

func TestIsIPv4(t *testing.T) {
	ip := net.ParseIP("127.0.0.1")
	assert.True(t, iputils.IsIPv4(ip))
	ip = net.ParseIP("2001:6db::1")
	assert.False(t, iputils.IsIPv4(ip))
}

func TestIsIPv6(t *testing.T) {
	ip := net.ParseIP("127.0.0.1")
	assert.False(t, iputils.IsIPv6(ip))
	ip = net.ParseIP("2001:6db::1")
	assert.True(t, iputils.IsIPv6(ip))
}

type testCaseTestAdd struct {
	X string
	Y uint64
	B *big.Int
	Z string
	E bool
}

func TestAdd(t *testing.T) {
	testcases := []testCaseTestAdd{
		testCaseTestAdd{
			X: "0.0.0.0",
			Y: 1,
			Z: "0.0.0.1",
		},
		testCaseTestAdd{
			X: "0.0.0.0",
			Y: uint64(math.Pow(2, 8)) - 1,
			Z: "0.0.0.255",
		},
		testCaseTestAdd{
			X: "0.0.0.0",
			Y: uint64(math.Pow(2, 8)),
			Z: "0.0.1.0",
		},
		testCaseTestAdd{
			X: "0.0.0.0",
			Y: uint64(math.Pow(2, 16) - 1),
			Z: "0.0.255.255",
		},
		testCaseTestAdd{
			X: "0.0.0.0",
			Y: uint64(math.Pow(2, 16)),
			Z: "0.1.0.0",
		},
		testCaseTestAdd{
			X: "0.0.0.0",
			Y: uint64(math.Pow(2, 24) - 1),
			Z: "0.255.255.255",
		},
		testCaseTestAdd{
			X: "0.0.0.0",
			Y: uint64(math.Pow(2, 24)),
			Z: "1.0.0.0",
		},
		testCaseTestAdd{
			X: "0.0.0.0",
			Y: uint64(math.Pow(2, 32) - 1),
			Z: "255.255.255.255",
		},
		testCaseTestAdd{
			X: "0.0.0.0",
			Y: uint64(math.Pow(2, 32)),
			Z: "<nil>",
			E: true,
		},
		testCaseTestAdd{
			X: "::0",
			Y: 1,
			Z: "::1",
		},
		testCaseTestAdd{
			X: "::0",
			Y: uint64(math.Pow(2, 16) - 1),
			Z: "::FFFF",
		},
		testCaseTestAdd{
			X: "::0",
			Y: uint64(math.Pow(2, 16)),
			Z: "::1:0",
		},
		testCaseTestAdd{
			X: "::0",
			Y: uint64(math.Pow(2, 32) - 1),
			Z: "::FFFF:FFFF",
		},
		testCaseTestAdd{
			X: "::0",
			Y: uint64(math.Pow(2, 32)),
			Z: "::1:0:0",
		},
		testCaseTestAdd{
			X: "::0",
			Y: uint64(math.Pow(2, 48) - 1),
			Z: "::FFFF:FFFF:FFFF",
		},
		testCaseTestAdd{
			X: "::0",
			Y: uint64(math.Pow(2, 48)),
			Z: "::1:0:0:0",
		},
		testCaseTestAdd{
			X: "FFFF:FFFF:FFFF:FFFF:FFFF:FFFF:FFFF:FFFF",
			Y: 1,
			Z: "<nil>",
			E: true,
		},
	}

	for _, testcase := range testcases {
		ip := net.ParseIP(testcase.X)
		p, err := iputils.Add(ip, testcase.Y)
		t.Logf("%s + %d = %s, result=%s len=%d", ip.String(), testcase.Y, testcase.Z, p.String(), len(ip))
		if testcase.E {
			assert.Error(t, err)
			assert.Nil(t, p)
		} else {
			assert.True(t, p.Equal(net.ParseIP(testcase.Z)))
			assert.NoError(t, err)
		}
	}

}

func TestAddBigInt(t *testing.T) {
	testcases := []testCaseTestAdd{
		testCaseTestAdd{
			X: "::0",
			B: big.NewInt(0).SetBytes([]byte{0x01}),
			Z: "::1",
		},
		testCaseTestAdd{
			X: "::0",
			B: big.NewInt(0).SetBytes([]byte{0xff, 0xff}),
			Z: "::FFFF",
		},
		testCaseTestAdd{
			X: "::0",
			B: big.NewInt(0).SetBytes([]byte{0x01, 0x00, 0x00}),
			Z: "::1:0",
		},
		testCaseTestAdd{
			X: "::0",
			B: big.NewInt(0).SetBytes([]byte{0xff, 0xff, 0xff, 0xff}),
			Z: "::FFFF:FFFF",
		},
		testCaseTestAdd{
			X: "::0",
			B: big.NewInt(0).SetBytes([]byte{0x01, 0x00, 0x00, 0x00, 0x00}),
			Z: "::1:0:0",
		},
		testCaseTestAdd{
			X: "::0",
			B: big.NewInt(0).SetBytes([]byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff}),
			Z: "::FFFF:FFFF:FFFF",
		},
		testCaseTestAdd{
			X: "::0",
			B: big.NewInt(0).SetBytes([]byte{0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}),
			Z: "::1:0:0:0",
		},
		testCaseTestAdd{
			X: "::0",
			B: big.NewInt(0).SetBytes([]byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}),
			Z: "::FFFF:FFFF:FFFF:FFFF",
		},
		testCaseTestAdd{
			X: "::0",
			B: big.NewInt(0).SetBytes([]byte{0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}),
			Z: "0:0:0:1::",
		},
		testCaseTestAdd{
			X: "::0",
			B: big.NewInt(0).SetBytes([]byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}),
			Z: "0:0:0:FFFF:FFFF:FFFF:FFFF:FFFF",
		},
		testCaseTestAdd{
			X: "::0",
			B: big.NewInt(0).SetBytes([]byte{0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}),
			Z: "0:0:1::",
		},
		testCaseTestAdd{
			X: "::0",
			B: big.NewInt(0).SetBytes([]byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}),
			Z: "0:0:FFFF:FFFF:FFFF:FFFF:FFFF:FFFF",
		},
		testCaseTestAdd{
			X: "::0",
			B: big.NewInt(0).SetBytes([]byte{0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}),
			Z: "0:1::",
		},
		testCaseTestAdd{
			X: "::0",
			B: big.NewInt(0).SetBytes([]byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}),
			Z: "0:FFFF:FFFF:FFFF:FFFF:FFFF:FFFF:FFFF",
		},
		testCaseTestAdd{
			X: "::0",
			B: big.NewInt(0).SetBytes([]byte{0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}),
			Z: "1::",
		},
		testCaseTestAdd{
			X: "::0",
			B: big.NewInt(0).SetBytes([]byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}),
			Z: "FFFF:FFFF:FFFF:FFFF:FFFF:FFFF:FFFF:FFFF",
		},
		testCaseTestAdd{
			X: "::0",
			B: big.NewInt(0).SetBytes([]byte{0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}),
			Z: "<nil>",
			E: true,
		},
	}

	for _, testcase := range testcases {
		ip := net.ParseIP(testcase.X)
		p, err := iputils.AddBigInt(ip, testcase.B)
		t.Logf("%s + %d = %s, result=%s len=%d", ip.String(), testcase.Y, testcase.B.String(), p.String(), len(ip))
		if testcase.E {
			assert.Error(t, err)
			assert.Nil(t, p)
		} else {
			assert.True(t, p.Equal(net.ParseIP(testcase.Z)))
			assert.NoError(t, err)
		}
	}

}
