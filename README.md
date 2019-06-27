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

https://scrapbox.io/forme/go-docker%E5%8F%82%E8%80%83%E8%A8%98%E4%BA%8B