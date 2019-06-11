## 使用技術

- [echo](https://echo.labstack.com/)
- [CircleCI](https://circleci.com/)

## Set up

```
//start service
make up

// stop service
make down

// run shell interactively in backend container
make backend

// run test
make test
```

## メモ

Dockerfileからイメージをビルド

```
docker build --tag=golang-study .
```
golang-studyイメージからコンテナを立ち上げる

```
docker run -it -v=$(pwd):/go  golang-study
```

`WORKDIR /go/src`を指定すると`docker-compose exec web bash`でコンテナにログインした時に`/go/src`ディレクトリから始まる

### ERROR: Named volume "data:/var/lib/mysql:rw" is used in service "xxxx" but no declaration was found in the volumes section.
http://grmn.hatenablog.com/entry/2018/09/06/113204

## 参考

[DockerのVolume機能について実験してみたことをまとめます](https://qiita.com/namutaka/items/f6a574f75f0997a1bb1d)

[makefileの参考記事](https://qiita.com/yoskeoka/items/317a3afab370155b3ae8)

[fresh導入時に参考にした記事](https://qiita.com/po3rin/items/9acd41ef428436335c97)

[はじめてのgo test](https://qiita.com/marnie_ms4/items/e51cc6d879cc9ad07af3)

[CircleCIの結果をSlackに通知する](https://qiita.com/su-kun1899/items/640f6fa8b48749396c16)

[DBの設定の時にこれがあって助かった](https://stackoverflow.com/questions/52999940/connection-refused-while-trying-to-connect-golang-to-mssql-server-using-docker)

[Testify](https://re-engines.com/2018/10/16/go-testify%E3%82%92%E4%BD%BF%E3%81%A3%E3%81%A6%E3%83%86%E3%82%B9%E3%83%88%E4%BD%9C%E6%88%90/)

[【GORM】Go言語でORM触ってみた](https://qiita.com/chan-p/items/cf3e007b82cc7fce2d81)