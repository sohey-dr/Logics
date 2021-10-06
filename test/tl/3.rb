
n = 0
# 20000はとりあえず確認するために切り上げる位置
for i in 1..20000 do
  n += 1 / i.to_f
  if n > 9
    p i
    break
  end
end
