=begin
  楽天ふるさと納税では自治体情報の表示位置が定まっていない
  しかし、商品詳細URLに自治体コードが入っている
  そのため、文字列のインデックスを指定して抽出する必要がある
=end

url = "https://www.rakuten.co.jp/f032063-fuchu/info.html"

num = url[27..32]
puts num[0] == "0" ? num[1..5] : num