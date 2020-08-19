GO言語版

Standard Go Project Layoutに沿ったディレクトリ構成を自動作成するシェル

すでに同じ名前のディレクトリが存在する場合は作成されない

Standard Go Project Layoutをそのまま適用すると多すぎるので、現在個人的に利用するディレクトリだけ作成。

ディレクトリとファイルはGOPATH/src内に展開

```
{project_name}
  ├── cmd 
  │    └──{project_name}  //起動スクリプト置き場
  │
  ├── web                 //テンプレート、HTML、CSSとかをおく
  │
  ├── pkg                 //ライブラリ
  │
  └── assets              //画像などのアセット置き場

```

## 参考

本家

https://github.com/golang-standards/project-layout


その他Tips等

https://qiita.com/sueken/items/87093e5941bfbc09bea8

https://techdo.mediado.jp/entry/2019/01/18/112503

https://future-architect.github.io/articles/20200528/
