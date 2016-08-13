# goFitbit
Get activity information by fitbit api
https://dev.fitbit.com/jp
Fitbit Developer APIを使ってOAuth 2.0 Authorization Code Grantで認証してsleep（睡眠情報)
を取得する。

## generateOauthURL.go
認証ページのURLを作成するだけのもの

## fitbit.go
認証後のページを表示する。
authorization codeを元にaccess tokenを作成して、sleepの情報を取得するところまで行っている。
webサーバをlocalhost:8080で立ち上げている。

#.env
.envに環境変数を書いていて、それをLoadしている。
go get github.com/joho/godotenv
でライブラリをgetする必要あり。
