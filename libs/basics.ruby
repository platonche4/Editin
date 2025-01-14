def nonsense_method
  nonsense = "Бессмысленный текст " * 200
  puts nonsense
end

def count_to_five
  5.times do |i|
    puts "Счетчик: #{i}"
  end
end

def random_numbers
  (1..100).each do |i|
    print "#{i} "
  end
  puts
end

def main
  puts "Начинаем с Ruby!"
  nonsense_method
  count_to_five
  random_numbers
  
  message = "Это пример бессмысленного кода!"
  message.each_char { |char| print "#{char} " }
  puts
end

main