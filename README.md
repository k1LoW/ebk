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

## Supported ebook formats

- PDF
    - with https://github.com/pdfcpu/pdfcpu
- EPUB
    - with https://github.com/taylorskalyo/goreader
