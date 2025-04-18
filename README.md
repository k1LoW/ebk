# ebk

`ebk` is a tiny tool for ebook.

## Usage

### `ebk ls`

list directory ebook contents.

``` console
$ ebk ls ~/path/to/books/ | grep oreilly
NFS & NIS 第2版:oreilly-4-87311-078-5e.pdf
ウェブオペレーション:oreilly-978-4-87311-493-4e.pdf
SQLアンチパターン:oreilly-978-4-87311-589-4e.epub
[...]
Go言語による並行処理:oreilly-978-4-87311-846-8e.epub
入門　監視:oreilly-978-4-87311-864-2e.epub
入門 Prometheus:oreilly-978-4-87311-877-2e.epub
みんなでアジャイル:oreilly-978-4-87311-909-0e.epub
ユニコーン企業のひみつ:oreilly-978-4-87311-946-5e.epub
```

#### Select and Open ebook using ebk and [peco](https://github.com/peco/peco) (macOS)

``` console
$ open $(ebk ls /path/to/books/ --with-path -d '\t' | peco | awk -F '\t' '{print $2}')
```

``` console
QUERY> SQL                                            IgnoreCase [5 (1/1)]
15時間でわかるMySQL集中講座 /path/to/books/15時間てわかる-MySQL集中講座_00.epub
SQLパフォーマンス詳解 - 開発者のためのSQLパフォーマンスのすべて /path/to/books/SQLパフォ
SQLアンチパターン   /path/to/books/oreilly-978-4-87311-589-4e.epub
PostgreSQL Q&A実録  /path/to/books/postgresql-qanda-jitsuroku.pdf
??? /path/to/books/達人に学ふSQL徹底指南書 第2版.pdf
```

## Supported ebook formats

- PDF
    - implemented using https://github.com/pdfcpu/pdfcpu
- EPUB
    - implemented using https://github.com/taylorskalyo/goreader

## Install

**deb:**

``` console
$ export EBK_VERSION=X.X.X
$ curl -o ebk.deb -L https://github.com/k1LoW/ebk/releases/download/v$EBK_VERSION/ebk_$EBK_VERSION-1_amd64.deb
$ dpkg -i ebk.deb
```

**RPM:**

``` console
$ export EBK_VERSION=X.X.X
$ yum install https://github.com/k1LoW/ebk/releases/download/v$EBK_VERSION/ebk_$EBK_VERSION-1_amd64.rpm
```

**apk:**

``` console
$ export EBK_VERSION=X.X.X
$ curl -o ebk.apk -L https://github.com/k1LoW/ebk/releases/download/v$EBK_VERSION/ebk_$EBK_VERSION-1_amd64.apk
$ apk add ebk.apk
```

**homebrew tap:**

```console
$ brew install k1LoW/tap/ebk
```

**manually:**

Download binary from [releases page](https://github.com/k1LoW/ebk/releases)

**go get:**

```console
$ go get github.com/k1LoW/ebk
```
