def check(x)
  if x % 2 == 0
      puts "偶数です"
  else
      puts "奇数です"
  end
end

check(ARGV[0].to_i)