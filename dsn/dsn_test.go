package dsn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDSN(t *testing.T) {
	data := []struct {
		Input    string
		Expected DSN
	}{
		{
			Input: "host0",
			Expected: DSN{
				HostName: "host0",
			},
		},
		{
			Input: "host10:222",
			Expected: DSN{
				HostName: "host10",
				Port:     "222",
			},
		},
		{
			"host1/db1",
			DSN{
				HostName: "host1",
				Database: "db1",
			},
		},
		{
			"host2:1234/db2",
			DSN{
				HostName: "host2",
				Port:     "1234",
				Database: "db2",
			},
		},
		{
			Input: "user3@host3/db3",
			Expected: DSN{
				UserName: "user3",
				HostName: "host3",
				Database: "db3",
			},
		},
		{
			Input: "user4:passw4@host4/db4",
			Expected: DSN{
				UserName: "user4",
				Password: "passw4",
				HostName: "host4",
				Database: "db4",
			},
		},
		{
			Input: "user5:passw5@host5:555/db5?extra=foo",
			Expected: DSN{
				UserName: "user5",
				Password: "passw5",
				HostName: "host5",
				Port:     "555",
				Database: "db5",
				Extra:    "extra=foo",
			},
		},
	}
	for _, testCase := range data {
		assert.Equal(t, testCase.Expected, Parse(testCase.Input))
	}
}

func TestGetExtra(t *testing.T) {
	input := "user:pass@host/db?key=value&param=foo&query=bar"
	dsn := Parse(input)
	assert.Equal(t, "value", dsn.GetExtra("key"))
	assert.Equal(t, "foo", dsn.GetExtra("param"))
	assert.Equal(t, "bar", dsn.GetExtra("query"))
	assert.Equal(t, "", dsn.GetExtra(""))
	assert.Equal(t, "", dsn.GetExtra("nothing"))
}
