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
    performance_preparation
    performance
  end

  private

  def performance_preparation
    puts "#{time.strftime('%H:%M')} ＼＼＼\\顔合わせ//／／／"
    @time += 20.minutes
    puts "START  [[[   #{time.strftime('%H:%M')}   ]]]"
  end

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
