package sig_test

import (
	sig "mysig/pkg"
	"testing"
)

type testWithBody struct {
	*testWithoutBody
	XAmzContentSha256 string
	TimeStr           string
	Body              string
}

func newTestBody() testWithBody {
	without := newTestWithoutBody()
	body := `{"key":"value"}`
	without.ExpectedSignature = "2ddd7741c55669cffafb4ea9dc60346d0bd026d972edca79e45f129e1d31401e"
	return testWithBody{
		XAmzContentSha256: "beaead3198f7da1e70d03ab969765e0821b24fc913697e929e726aeaebf0eba3",
		testWithoutBody:   newTestWithoutBody(),
		TimeStr:           "20230503T071829Z",
		Body:              body,
	}
}

type testWithoutBody struct {
	TimeStr           string
	AccessKeyId       string
	SecretAccessKey   string
	SessionToken      string
	ApiUrl            string
	Service           string
	Region            string
	ExpectedSignature string
}

func newTestWithoutBody() *testWithoutBody {
	return &testWithoutBody{
		AccessKeyId:       "AKIDEXAMPLE",
		SecretAccessKey:   "TESTSECRET",
		SessionToken:      "SESSIONTOKEN",
		ApiUrl:            "https://test.execute-api.ap-northeast-1.amazonaws.com/prod",
		Service:           "apigateway",
		Region:            "ap-northeast-1",
		ExpectedSignature: "f3ce4787656af82573fd8d93b2829aafdc7315f30d87cc3f5f60456b34f08300",
		TimeStr:           "20230503T072244Z",
	}
}

func TestCalculateAwsSignatureWithoutBody(t *testing.T) {
	test := newTestWithoutBody()
	sigBuilder := sig.AwsSignatureBuilder{}
	signature := sigBuilder.AccessKeyId(test.AccessKeyId).SecretAccessKey(test.SecretAccessKey).SessionToken(test.SessionToken).Region(test.Region).TimeStr(test.TimeStr).ApiUrl(test.ApiUrl).Service(test.Service).BuildSignature()
	if signature.Signature != test.ExpectedSignature {
		t.Errorf("Expected signature %s, got %s", test.ExpectedSignature, signature.Signature)
	}
	if signature.XAmzContentSha256 != nil {
		t.Errorf("Expected nil, got %s", *signature.XAmzContentSha256)
	}
}
