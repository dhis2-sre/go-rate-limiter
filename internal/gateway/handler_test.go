package gateway

func getBackends() []Backend {
	return []Backend{
		{Name: "backend0", Url: "http://backend0:8080"},
		{Name: "backend1", Url: "http://backend1:8080"},
	}
}

const defaultBackend = "backend0"
const defaultRequestUrl = "http://url:80"

const publicKey = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAtYrBsSkVGXZKQL13lbmd
xFCQcvi6KIssjz3KOHIko/Da6sxE2w67OL84t98wCYbmIuq6xTK6qpEqEs1LaqQS
DnCs2VNDTLk4D1J42R63OpJQfOfebzhTJLx6KldyK2FRGXWILY7AzcoqyuLk433s
lHk6/yFDYgBA4COofeXZvXtUazuzpBWTZCxpEh341ob6XQ5juLYrqr/80XLYzXiu
N1iz24ulxSnD0GV4cRfHEnnzN3oYFzoYTcTQB6dffNAs/ADHNA9IemyLbT0ugvbf
L5MOEBOftYLRwmGFWrXf5s9jccku0FPid2wtZEwsv5Sa+Yvr36KHtrr+PSFksOB1
0QIDAQAB
-----END PUBLIC KEY-----`

const validAccessToken = `eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjQ3OTExNTQ0MTUsImlhdCI6MTYzNzU1NDQxNX0.PtQp6_k5bQ9KE9uk520i4emVnUmxFD8DxyeZsfzgT6CY2oMyXEm7zlIA-4_xz2Q7CrSeqnWxpy0coK9MN0EPE2vhFomTrP6D3l7_lX6Dyn1gH6zWpjC_dRqOSRv3AqS3buZiC-vNwCatLhu6WE74cykBAE2veIr8Gp_ebiITXJKiHBNaTlPk2WEfcJ1NL3g7nafy6l-V4h2-Vj3tapJQiLfpgReIXYIswFYH7En7qy94fL0eOUbZzQI9fOuiXvAN-owR3GYcbwz9Hll23VACWsekMJdDBEgUSdek9JOmRHGxko6FE79-_ClYvF1dGUgZB2mDwY_xF2TOG2q3XDi9Aw`
