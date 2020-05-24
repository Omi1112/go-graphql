# GolangでGraphQL

GolangでGraphQLを実践する。
ライブラリは`gqlgen`を使用する。

公式サイト
https://gqlgen.com/getting-started/


# スケルトン作成
```
$ mkdir [プロジェクト名]
$ cd [プロジェクト名]
$ go mod init github.com/[username]/[プロジェクト名]
$ go get github.com/99designs/gqlgen

$ go run github.com/99designs/gqlgen init
```

# 開発の流れ

1. スキーマを定義する。

`/graph/schema.grahqls`にスキーマを定義する。
このスキーマは、`gqlgen`専用のものではなく、
`GraphQL`共通のファイル。
言語や、ライブラリに依存するものでは無い。

2. スキーマから、メソッドを生成する。

### スキーマからメソッドを自動生成する。

```
go run github.com/99designs/gqlgen generate
```

### 仕上げ
仕上げとして、以下を実施する。
詳細が何を行われているかは不明。

`/graph/resolver.go`のpackageとimport間に、以下コメントを記載。

```
//go:generate go run github.com/99designs/gqlgen
```

コマンド実施で仕上げが行われる。

```
go generate ./...
```

3. メソッド内の詳細処理を記述する。

スキーマで定義したメソッドが`/graph/schema.resolvers.go`に、生成されている。
初期値は、`panic()`が定義されているため、ここを好きな様に定義していく。

4. サーバーの起動

メソッドを書き換えたら、サーバーを起動する。
```
go run server.go
```
