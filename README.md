# go-lambda-api-sample
Go言語でAWS Lambdaを使用したAPIのサンプル

## Lambdaへのアップロード方法
1. ビルド
`GOOS=linux GOARCH=amd64 go build -o hello`

1. ZIPに圧縮
`zip handler.zip ./hello`

1. Lambdaの設定画面からアップロード

## 参考サイト
[AWS Lambda で Go が使えるようになったので試してみた](https://dev.classmethod.jp/cloud/aws/aws-lambda-supports-go/)  
[AWS公式のサンプルコード](https://docs.aws.amazon.com/ja_jp/lambda/latest/dg/with-on-demand-https-create-package.html#with-apigateway-example-deployment-pkg-go)  
[AWS公式のチュートリアル](https://docs.aws.amazon.com/ja_jp/lambda/latest/dg/with-on-demand-https-example.html)  