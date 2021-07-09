require "csv"

=begin
CSVの内容を多重配列として取り出して
その中の一要素だけ置き換える処理
今回は多重配列の中の配列のインデックスが7個目のものを特定の数字に置き換えている(1に)
=end


rows = CSV.read("slack/tech_bowl/users (1) copy.csv")

rows.each do |row|
  row[7] = 1
end


CSV.open("slack/tech_bowl/users (1)-(2)-(1).csv", "w") do |csv|
  rows.each do |row|
    csv << row
  end
end
