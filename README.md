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

## 後で読む

[DockerのVolume機能について実験してみたことをまとめます](https://qiita.com/namutaka/items/f6a574f75f0997a1bb1d)

[makefileの参考記事](https://qiita.com/yoskeoka/items/317a3afab370155b3ae8)