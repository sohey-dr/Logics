require 'active_support/time'


# class TimeTable
#   attr_reader :time, :band_num
#   # attr_writer :time

#   # def initialize(start,band_num)
#   #   @time = Time.parse(start)
#   #   @band_num = band_num
#   # end

#   @@time = Time.parse("202103201000")

#   def bands
#     @bands ||= bands_arry
#   end

#   def bands_arry
#    arry = []
#    for num in 1..band_num do
#      arry.push("バンド#{num}")
#    end
#    arry
#   end

#   def add_time
#     @@time += 5.minutes
#   end
# end

# p TimeTable.new.add_time
# p TimeTable.new("202103201000",20).bands

class TimeTable
  def run(num)
    bands = []
    kanki = 1
    (1..num).each do |n|
      bands.push("バンド#{n}")
    end

    start = Time.local(2021,8,26,10,00)
    puts "#{start.strftime('%H:%M')}〜#{rehear_band(start).strftime('%H:%M')} 音出しバンド"
    start += 15.minutes

    bands.reverse.each.with_index(1) { |band, index|
      puts "#{start.strftime('%H:%M')}〜#{rehear_band(start).strftime('%H:%M')} #{band}"
      start += 15.minutes
      # if index == bands.size
      #   break
      # elsif kanki == 3
      #   kanki = 0
      #   puts "#{start.strftime('%H:%M')}〜#{rehear_kanki(start).strftime('%H:%M')} <換気>"
      #   start += 5.minutes 
      # end
      kanki += 1
    }

    puts "#{start.strftime('%H:%M')} ＼＼＼\\顔合わせ//／／／"
    start += 30.minutes
    puts "START  [[[   #{start.strftime('%H:%M')}   ]]]"
    kanki = 1

    bands.each.with_index(1) { |band, index|
      puts "#{start.strftime('%H:%M')}〜#{honban_band(start).strftime('%H:%M')} #{band}"
      start += 20.minutes
      
      # if lunch == false && start > Time.local(2021,03,20,12,10)
      #   lunch = true
      #   puts "☆☆☆45分間のお昼休憩です☆☆☆"
      #   start += 35.minutes
      #   kanki = 0
      # end

      # if index == bands.size
      #   break
      # elsif kanki == 3
      #   kanki = 0
      #   puts "#{start.strftime('%H:%M')}〜#{honban_kanki(start).strftime('%H:%M')} <換気>"
      #   start += 10.minutes 
      # else
      #   start += 10.minutes
      # end
      start += 10.minutes
      kanki += 1
    }

    start += 35.minutes
    puts "完パケ#{start.strftime('%H:%M')}"
    puts "※換気・転換10分 本番15分"
  end

  def rehear_band(start)
    start + 15.minutes
  end

  def rehear_kanki(start)
    start + 10.minutes
  end

  def honban_band(start)
    start + 20.minutes
  end

  def honban_kanki(start)
    start + 10.minutes
  end
end

puts "バンド数を入力してください"
TimeTable.new.run(gets.to_i)
