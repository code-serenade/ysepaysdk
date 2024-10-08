package api

import "github.com/codingeasygo/util/xmap"

var sdk *Config

func init() {

	Verbose = true
	IsDev = true

	testPrivate := `-----BEGIN PRIVATE KEY-----
MIICdgIBADANBgkqhkiG9w0BAQEFAASCAmAwggJcAgEAAoGBAJyEcfptTx5RMZmYnZkbxh9n2PFQjuDqy9/la4D0TabG3IfHgoO7SUGawBNwOLloXOpevt2e9yLr3zV2wE0oNzoPa+bPw/Id5k+H3SW6FO2VNi+SroiS4tOTkRJsuddWMy3qeOFosvSAznlGW5xX87+5Eym6j65znSQGMDwwayU9AgMBAAECgYBQoQYaTO9CHe+gQXeGZUFl8lhWz/oasbn/CC+eSbwq2yzGIagBdpyhWDf0i44dLT8YWWCXyIlliv78isU7uijgSLY2ISiCzDd7KcTicDiD1gT3BEnX48hS/2I1BOIPfD/95HIjvfmRclCagmi6/7euZ7WiyxJWZEGYz4ujn7314QJBAM5gFp3wTi7W4lVcI+B6Zhay3oAA0TWuVoDuZKwk9f1grDszF1aRbvL6bcnUaNIc6KbNxx6+zVJz9ohBF5VWSbkCQQDCJzvKYOQ43ke02bVPLnmk4/0KFqMNhjWKKQTDYvXg67nBaQRnXaoiSEQ2boQ1AonSw/eahqA/UUB4awkhdymlAkEAkFi3GQsutCZKsqe61FToOPItHFz199UT3iRgN+O2Rt4qOVH3/e3FK6Ar0GKIWoYr2BfQMVoFZKmFgzH8G9D5yQJAFrRf2ZrbFYto900+VrQopY/D15ouFzbKrA1/8Rlxh4wJfQS67xuU+c1ZrMPf0hMp6uvf2MJFmyH1WB+sQa2dwQJAXzOahoBbYJKQDY419MUzntQh6Jk/1EzS7P9jpPbXtird4kS2kxwxPE7QL6xEf7255D3hj+U1lWMYDC2b6SaSeA==
-----END PRIVATE KEY-----`

	testPublic := `-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDNYenkt3tGBf3yGLBgNKXpj76gI3KFhcE6OKO+4NvbQTx70lg29s4ofikP5fVTHi4i//GTTPESyQz1F8+b2WasOBZUG3CS847o1KiyRiWURxK4LgagnBqnve8tQJ+qHT/H8iL6yFdipjq1Zkyu29pmhMgsbK++zHXj+SkrWdXWkwIDAQAB
-----END PUBLIC KEY-----`

	sdk = NewConfig(xmap.M{
		"cert_id":     "826440345119153",
		"private_key": testPrivate,
		"public_key":  testPublic,
	})
}
