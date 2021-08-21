require 'active_support/time'

class TimeTable
  attr_reader :time, :bands

  def initialize(band_number)
    @time = Time.local(2021, 8, 26, 13, 00)

    # バンドの配列作成 ex) ["バンド1", "バンド2", "バンド3", "バンド4", "バンド5"]
    @bands = Array.new(band_number).map.with_index(1){ |_, index| "バンド#{index}"}
  end

  def output
    rehearsal
    performance
  end

  private

  def performance
    bands.each.with_index(1) { |band, index|
      play_performance(band)
      if index == bands.size
        break
      elsif index % 3 == 0
        ventilation
      else
        convert
      end
    }
  end

  def rehearsal
    bands.reverse.each.with_index(1) { |band, index|
      play_rehearsal(band)
      if index == bands.size
        break
      elsif index % 3 == 0
        ventilation
      else
        convert
      end
    }
  end

  def play_rehearsal(band)
    after_time = time + 15.minutes
    puts "#{time.strftime('%H:%M')}〜#{after_time.strftime('%H:%M')} #{band}"
    @time = after_time
  end

  def ventilation
    after_time = time + 5.minutes
    puts "#{time.strftime('%H:%M')}〜#{after_time.strftime('%H:%M')} <換気>"
    @time = after_time
  end

  def play_performance(band)
    after_time = time + 20.minutes
    puts "#{time.strftime('%H:%M')}〜#{after_time.strftime('%H:%M')} #{band}"
    @time = after_time
  end

  def convert
    @time += 5.minutes
  end
end

puts "バンド数を入力してください"
TimeTable.new(gets.to_i).output


# class TimeTable
#   def run(num)
#     bands = []
#     kanki = 1
#     (1..num).each do |n|
#       bands.push("バンド#{n}")
#     end

#     start = Time.local(2021,8,26,13,00)
#     puts "#{start.strftime('%H:%M')}〜#{rehear_band(start).strftime('%H:%M')} 音出しバンド"
#     start += 15.minutes

#     bands.reverse.each.with_index(1) { |band, index|
#       puts "#{start.strftime('%H:%M')}〜#{rehear_band(start).strftime('%H:%M')} #{band}"
#       start += 20.minutes
#       if index == bands.size
#         break
#       elsif kanki == 2
#         kanki = 0
#         puts "#{start.strftime('%H:%M')}〜#{rehear_kanki(start).strftime('%H:%M')} <換気>"
#         start += 5.minutes 
#       end
#       kanki += 1
#     }

#     puts "#{start.strftime('%H:%M')} ＼＼＼\\顔合わせ//／／／"
#     start += 20.minutes
#     puts "START  [[[   #{start.strftime('%H:%M')}   ]]]"
#     kanki = 1

#     bands.each.with_index(1) { |band, index|
#       puts "#{start.strftime('%H:%M')}〜#{honban_band(start).strftime('%H:%M')} #{band}"
#       start += 20.minutes
      
#       # if lunch == false && start > Time.local(2021,03,20,12,10)
#       #   lunch = true
#       #   puts "☆☆☆45分間のお昼休憩です☆☆☆"
#       #   start += 35.minutes
#       #   kanki = 0
#       # end

#       if index == bands.size
#         break
#       elsif kanki == 2
#         kanki = 0
#         puts "#{start.strftime('%H:%M')}〜#{honban_kanki(start).strftime('%H:%M')} <換気>"
#         start += 10.minutes 
#       else
#         start += 10.minutes
#       end
#       kanki += 1
#     }

#     start += 35.minutes
#     puts "完パケ#{start.strftime('%H:%M')}"
#     puts "※換気・転換10分 本番20分"
#   end

#   def rehear_band(start)
#     start + 15.minutes
#   end

#   def rehear_kanki(start)
#     start + 10.minutes
#   end

#   def honban_band(start)
#     start + 20.minutes
#   end

#   def honban_kanki(start)
#     start + 10.minutes
#   end
# end

# puts "バンド数を入力してください"
# TimeTable.new.run(gets.to_i)


