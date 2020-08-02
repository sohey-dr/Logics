array = []
[{:volume=>350, :unit=>"ml"}, {:volume=>24, :unit=>"本"}].each do |hash|
  array.push("#{hash[:volume]}#{hash[:unit]}")
end
p array