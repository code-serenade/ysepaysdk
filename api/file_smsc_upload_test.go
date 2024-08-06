package api

import (
	"fmt"
	"testing"
	"time"

	"github.com/codingeasygo/util/xmap"
)

var sdk *Config

func init() {

	testPrivate := `-----BEGIN PRIVATE KEY-----
MIICdgIBADANBgkqhkiG9w0BAQEFAASCAmAwggJcAgEAAoGBAJyEcfptTx5RMZmY
nZkbxh9n2PFQjuDqy9/la4D0TabG3IfHgoO7SUGawBNwOLloXOpevt2e9yLr3zV2
wE0oNzoPa+bPw/Id5k+H3SW6FO2VNi+SroiS4tOTkRJsuddWMy3qeOFosvSAznlG
W5xX87+5Eym6j65znSQGMDwwayU9AgMBAAECgYBQoQYaTO9CHe+gQXeGZUFl8lhW
z/oasbn/CC+eSbwq2yzGIagBdpyhWDf0i44dLT8YWWCXyIlliv78isU7uijgSLY2
ISiCzDd7KcTicDiD1gT3BEnX48hS/2I1BOIPfD/95HIjvfmRclCagmi6/7euZ7Wi
yxJWZEGYz4ujn7314QJBAM5gFp3wTi7W4lVcI+B6Zhay3oAA0TWuVoDuZKwk9f1g
rDszF1aRbvL6bcnUaNIc6KbNxx6+zVJz9ohBF5VWSbkCQQDCJzvKYOQ43ke02bVP
Lnmk4/0KFqMNhjWKKQTDYvXg67nBaQRnXaoiSEQ2boQ1AonSw/eahqA/UUB4awkh
dymlAkEAkFi3GQsutCZKsqe61FToOPItHFz199UT3iRgN+O2Rt4qOVH3/e3FK6Ar
0GKIWoYr2BfQMVoFZKmFgzH8G9D5yQJAFrRf2ZrbFYto900+VrQopY/D15ouFzbK
rA1/8Rlxh4wJfQS67xuU+c1ZrMPf0hMp6uvf2MJFmyH1WB+sQa2dwQJAXzOahoBb
YJKQDY419MUzntQh6Jk/1EzS7P9jpPbXtird4kS2kxwxPE7QL6xEf7255D3hj+U1
lWMYDC2b6SaSeA==
-----END PRIVATE KEY-----`

	testPublic := `-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDNYenkt3tGBf3yGLBgNKXpj76g
I3KFhcE6OKO+4NvbQTx70lg29s4ofikP5fVTHi4i//GTTPESyQz1F8+b2WasOBZU
G3CS847o1KiyRiWURxK4LgagnBqnve8tQJ+qHT/H8iL6yFdipjq1Zkyu29pmhMgs
bK++zHXj+SkrWdXWkwIDAQAB
-----END PUBLIC KEY-----`

	sdk = NewConfig(xmap.M{
		"cert_id":     "826440345119153",
		"private_key": testPrivate,
		"public_key":  testPublic,
	})
	Verbose = true
	IsDev = true
}

func TestRequstFileSmscUpload_ValidInputs(t *testing.T) {
	filePath := "../testdata/test.jpeg"
	picType := "A001"
	sysFlowID := fmt.Sprintf("%v", time.Now().UnixNano())
	data, err := sdk.RequstFileSmscUpload(filePath, picType, sysFlowID)
	if err != nil {
		t.Errorf("TestRequstFileSmscUpload_ValidInputs failed: %v", err)
		return
	}
	t.Logf("TestRequstFileSmscUpload_ValidInputs success: %v", data)

}
