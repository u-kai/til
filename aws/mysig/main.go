package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	sig "mysig/pkg"
	"net/http"
	"os"
	"time"

	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/config"
)

// q:aws のv4の署名を作成するにはどうすればいいのか？
// a: https://docs.aws.amazon.com/ja_jp/general/latest/gr/sigv4-signed-request-examples.html
func main() {

	ctx := context.Background()

	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		panic(err.Error())
	}

	credentials, err := cfg.Credentials.Retrieve(ctx)
	if err != nil {
		panic(err.Error())
	}

	type Test struct {
		Id   string `json:"id"`
		Name string `json:"name"`
	}
	data := Test{Id: "1", Name: "test"}
	b, err := json.Marshal(data)
	src := string(b)
	hash := sig.V4Sign(src)
	req, err := http.NewRequest(
		http.MethodPost,
		os.Getenv("API_URL"),
		//hash,
		bytes.NewReader([]byte(src)),
	)
	if err != nil {
		panic(err.Error())
	}

	signer := v4.NewSigner()
	signer.SignHTTP(ctx, credentials, req, hash, "execute-api", "ap-northeast-1", time.Now())

	if err != nil {
		panic(err.Error())
	}

	httpClient := new(http.Client)

	response, err := httpClient.Do(req)
	if err != nil {
		panic(err.Error())
	}

	defer response.Body.Close()

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err.Error())
	}

	fmt.Print(string(responseBody))
}
