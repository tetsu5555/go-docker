## 使用技術

[echo](https://echo.labstack.com/)
[CircleCI](https://circleci.com/)

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