=begin
文字列の長さとして数値Nと文字列Sが標準入力で与えられます。
文字列Sは、1つ以上の単語を (間に空白などを挟まずに) 連結したものです。
ここで、各単語は2文字以上であり、最初の文字と最後の文字のみが英大文字、それ以外の文字は全て英小文字である。
これらの単語を辞書順に並べ (大文字小文字の違いは無視する)、同様に連結して出力するプログラムを作成してください。 
例えば、S='FisHDoGCaTAAAaAAbCAC'とする。これは 'FisH', 'DoG', 'CaT', 'AA', 'AaA', 'AbC', 'AC' の7つの単語を連結したものである。
これらを辞書順に並べると 'AA', 'AaA', 'AbC', 'AC', 'CaT', 'DoG', 'FisH' となるため、'AAAaAAbCACCaTDoGFisH' と出力すればよい。 
### 制約 * Sは長さ2以上100,000以下の文字列である。 * Sの各文字は英大文字または英小文字である。 * Sは問題文で述べたような単語の連結である。 
### 入力 入力ごとに改行した値が与えられます。 ``` N S ```
### 入力例 ``` 20 FisHDoGCaTAAAaAAbCAC ```
=end

lines = readlines 
i = 0
while i < lines.length 
  lines[i] = lines[i].chomp 
  i += 1 
end 
words = lines[1] 
array = [] 
string = "" 
i = 0 

lines[0].to_i.times do 
  if /\A[A-Z]+\z/.match?(words[i]) 
    if string != "" 
      string << words[i] 
      array.push(string) 
      string = "" 
      i += 1 
      next 
    end 
    string << words[i] 
  else 
    string << words[i] 
  end 
  i += 1 
end 
print array.sort.join('')