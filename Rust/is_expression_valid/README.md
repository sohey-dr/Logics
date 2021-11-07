# Is Expression Valid
不適切な表現があった場合にはアラートを表示を出す。Webassembly。

## Usage
以下コマンドを実行して、生成されるpkgディレクトリを用いる。
```bash
$ wasm-pack build --target web
```

アクセスする場合は以下のHTMLを追加し、コマンドを実行をして読み込んでみる。
```html
<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8" />
  </head>
  <body>
    <script type="module">
      import init, { is_expression_valid } from "./pkg/is_expression_valid.js";
      init().then(() => {
        is_expression_valid("WebAssembly");
      });
    </script>
  </body>
</html>
```

```bash
$ python -m SimpleHTTPServer 808
````