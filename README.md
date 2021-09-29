# URLのコンテンツ抽出実験

## py/test1.py

readability-xml (<https://pypi.org/project/readability-lxml/>) を使ったシンプルな実装

使い方はこう

```console
$ python test1.py {URLs...}

$ python test1.py < url_list.txt
```

サンプルはこう。実際は `...(snip)` 以降にも本文が続く

```console
$ python test1.py https://nazology.net/archives/95980
OK      https://nazology.net/archives/95980     刺し身好きにとって...(snip)
```

成功時の出力は次のフォーマット

    OK\t{URL}\t{CONTENT_IN_PLAIN_TEXT}

失敗時のフォーマット

    NG\t{URL}\t{FAILURE_REASON}

### 評価

URL | 成否 | 分類 | 備考
----|:----:|------|------
https://news.yahoo.co.jp/articles/acba39a56e3a15a628b2c6e7c288f103222ce86c | BAD | Yahooニュースの1記事 | 最初は大丈夫だったが、のちにダメになった
https://news.yahoo.co.jp/pickup/6404492 | BAD | 上記のYahooニュースのピックアップ | アクセスランキングのタイトルが採用される。ピックアップは本文が短いから?
https://www.asahi.com/articles/ASP9G5DG1P9GUTIL011.html | GOOD | 朝日新聞の記事 | (n/a)
https://www.asahi.com/articles/ASP9G66HWP9GUCLV00D.html | GOOD | 朝日新聞の記事 | (n/a)
https://www.asahi.com/ads/springvalley202109/ | BAD | 朝日新聞のPR記事 | CP932のエンコードエラー。Windowsを使ってるせいかも
https://www.yomiuri.co.jp/economy/20210915-OYT1T50049/ | GOOD | 読売新聞の記事 | (n/a)
https://www.yomiuri.co.jp/culture/20210914-OYT1T50344/ | GOOD | 読売新聞の記事 | (n/a)
https://yab.yomiuri.co.jp/adv/chuo/opinion/20210311.php | GOOD | 読売新聞のPR記事 | (n/a)
https://www.sankei.com/article/20210915-HGJUDSPJKNBCXMXGCO4I5VTBGY/ | GOOD | 産経新聞の記事 | 画像キャプションが先頭に含まれてる
https://www.sankei.com/article/20210911-E523ZF5SKJI7BBYWAZG6WYIQXM/ | GOOD | 産経新聞の記事 | (n/a)
https://mainichi.jp/articles/20210915/k00/00m/040/074000c | GOOD | 毎日新聞の記事 | 画像キャプションが先頭に含まれてる
https://mainichi.jp/articles/20210914/dde/012/040/017000c | GOOD | 毎日新聞の記事 | (n/a)
https://mainichi.jp/sp/sekai/ | BAD | 毎日新聞のPRページ | ワンピースタブロイドの広告
https://japanese.engadget.com/jp-pr-oppo-reno-5-a-020005713.html | GOOD | engadget日本版 | (n/a)
https://nazology.net/archives/95980 | GOOD | ナゾロジーの記事 | (n/a)

## go/test2.go

Mozillaによる[Readability.js](https://github.com/mozilla/readability)のGo移植版を用いたやつ。

使い方

```console
$ cd go
$ go buid ./test2.go
$ ./test2 < ../list.txt
```

出力形式は test1 と同じ。

### 評価

[結果全体](./go/test2_out.txt)

URL | 成否 | 分類 | 備考
----|:----:|------|------
https://news.yahoo.co.jp/articles/acba39a56e3a15a628b2c6e7c288f103222ce86c | GOOD | Yahooニュースの1記事 | (n/a)
https://news.yahoo.co.jp/pickup/6404492 | BAD | 上記のYahooニュースのピックアップ | サイドバーの内容が採用されている
https://www.asahi.com/articles/ASP9G5DG1P9GUTIL011.html | GOOD | 朝日新聞の記事 | (n/a)
https://www.asahi.com/articles/ASP9G66HWP9GUCLV00D.html | GOOD | 朝日新聞の記事 | (n/a)
https://www.asahi.com/ads/springvalley202109/ | GOOD | 朝日新聞のPR記事 | (n/a)
https://www.yomiuri.co.jp/economy/20210915-OYT1T50049/ | GOOD | 読売新聞の記事 | (n/a)
https://www.yomiuri.co.jp/culture/20210914-OYT1T50344/ | GOOD | 読売新聞の記事 | (n/a)
https://yab.yomiuri.co.jp/adv/chuo/opinion/20210311.php | GOOD | 読売新聞のPR記事 | (n/a)
https://www.sankei.com/article/20210915-HGJUDSPJKNBCXMXGCO4I5VTBGY/ | GOOD | 産経新聞の記事 | 画像キャプションが先頭に含まれてる
https://www.sankei.com/article/20210911-E523ZF5SKJI7BBYWAZG6WYIQXM/ | GOOD | 産経新聞の記事 | (n/a)
https://mainichi.jp/articles/20210915/k00/00m/040/074000c | GOOD | 毎日新聞の記事 | 画像キャプションが先頭に含まれてる
https://mainichi.jp/articles/20210914/dde/012/040/017000c | GOOD | 毎日新聞の記事 | (n/a)
https://mainichi.jp/sp/sekai/ | GOOD | 毎日新聞のPRページ | なんかいい感じに取れてるが、ブラウザで見た際の印象とは乖離がある
https://japanese.engadget.com/jp-pr-oppo-reno-5-a-020005713.html | GOOD | engadget日本版 | (n/a)
https://nazology.net/archives/95980 | GOOD | ナゾロジーの記事 | (n/a)

## java/test3

Javaの boilerpipe を使ったやつ。

使い方

```console
$ ./gradlew run

# URL一覧はソースコードに埋め込んである
```

### 評価

[結果全体](./java/test3_out.txt)

URL | 成否 | 分類 | 備考
----|:----:|------|------
https://news.yahoo.co.jp/articles/acba39a56e3a15a628b2c6e7c288f103222ce86c | GOOD | Yahooニュースの1記事 | (n/a)
https://news.yahoo.co.jp/pickup/6404492 | BAD | 上記のYahooニュースのピックアップ | 既に無効になっている
https://www.asahi.com/articles/ASP9G5DG1P9GUTIL011.html | GOOD | 朝日新聞の記事 | タイトルは含まれないが、画像キャプションも含まれない。末尾に「関連ニュース」のヘッダーが含まれる。
https://www.asahi.com/articles/ASP9G66HWP9GUCLV00D.html | GOOD | 朝日新聞の記事 | キャプションが含まれる
https://www.asahi.com/ads/springvalley202109/ | GOOD | 朝日新聞のPR記事 | (n/a)
https://www.yomiuri.co.jp/economy/20210915-OYT1T50049/ | GOOD | 読売新聞の記事 | (n/a)
https://www.yomiuri.co.jp/culture/20210914-OYT1T50344/ | GOOD | 読売新聞の記事 | (n/a)
https://yab.yomiuri.co.jp/adv/chuo/opinion/20210311.php | GOOD | 読売新聞のPR記事 | (n/a)
https://www.sankei.com/article/20210915-HGJUDSPJKNBCXMXGCO4I5VTBGY/ | GOOD | 産経新聞の記事 | (n/a)
https://www.sankei.com/article/20210911-E523ZF5SKJI7BBYWAZG6WYIQXM/ | GOOD | 産経新聞の記事 | (n/a)
https://mainichi.jp/articles/20210915/k00/00m/040/074000c | GOOD | 毎日新聞の記事 | (n/a)
https://mainichi.jp/articles/20210914/dde/012/040/017000c | GOOD | 毎日新聞の記事 | (n/a)
https://mainichi.jp/sp/sekai/ | GOOD | 毎日新聞のPRページ | (n/a)
https://japanese.engadget.com/jp-pr-oppo-reno-5-a-020005713.html | GOOD | (n/a)
https://nazology.net/archives/95980 | GOOD | ナゾロジーの記事 | (n/a)

* 全体的に末尾に余計なものが付きがち
* タイトルやキャプションは入ったり入らなかったりする

### 解析

<https://github.com/pvdlg/boilerpipe/blob/master/src/main/java/de/l3s/boilerpipe/extractors/ArticleExtractor.java>

複数のフィルタを組み合わせる形になっている。
