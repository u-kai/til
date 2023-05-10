package sig

type HttpMethods int

const (
	Get HttpMethods = iota
	Post
	Delete
	Patch
	Connect
)

func (m HttpMethods) String() string {
	return []string{"GET", "POST", "DELETE", "PATCH", "CONNECT"}[m]
}

type CanonicalUri string

func (c CanonicalUri) String() string {
	return string(c)
}
func NewCanonicalUriFromUrl(url string) CanonicalUri {
	return CanonicalUri("/")
}

type CanonicalQueryString string

func (c CanonicalQueryString) String() string {
	return string(c)
}

func NewCanonicalQueryStringFromUrl(url string) CanonicalQueryString {
	return CanonicalQueryString("")
}

//	AWSの署名を計算する
//
// https://docs.aws.amazon.com/ja_jp/general/latest/gr/sigv4_signing.html
// https://docs.aws.amazon.com/ja_jp/general/latest/gr/sigv4-create-canonical-request.html
// https://docs.aws.amazon.com/ja_jp/general/latest/gr/sigv4-create-string-to-sign.html
// https://docs.aws.amazon.com/ja_jp/general/latest/gr/sigv4-calculate-signature.html
// https://docs.aws.amazon.com/ja_jp/general/latest/gr/sigv4-add-signature-to-request.html

type AwsSignatureBuilder struct {
	accessKeyId     string
	secretAccessKey string
	sessionToken    string
	timeStr         string
	apiUrl          string
	service         string
	method          HttpMethods
	region          string
}

func (b *AwsSignatureBuilder) AccessKeyId(accessKeyId string) *AwsSignatureBuilder {
	b.accessKeyId = accessKeyId
	return b
}
func (b *AwsSignatureBuilder) SecretAccessKey(secretAccessKey string) *AwsSignatureBuilder {
	b.secretAccessKey = secretAccessKey
	return b
}
func (b *AwsSignatureBuilder) SessionToken(sessionToken string) *AwsSignatureBuilder {
	b.sessionToken = sessionToken
	return b
}
func (b *AwsSignatureBuilder) TimeStr(timeStr string) *AwsSignatureBuilder {
	b.timeStr = timeStr
	return b
}
func (b *AwsSignatureBuilder) ApiUrl(apiUrl string) *AwsSignatureBuilder {
	b.apiUrl = apiUrl
	return b
}
func (b *AwsSignatureBuilder) Service(service string) *AwsSignatureBuilder {
	b.service = service
	return b
}
func (b *AwsSignatureBuilder) Method(method HttpMethods) *AwsSignatureBuilder {
	b.method = method
	return b
}
func (b *AwsSignatureBuilder) Region(region string) *AwsSignatureBuilder {
	b.region = region
	return b
}
func (b *AwsSignatureBuilder) BuildSignature() AwsSignature {
	return AwsSignature{}
}

type AwsSignature struct {
	Signature         string
	XAmzContentSha256 *string
}
