n = 100_000
m = 200_000

puts "#{n} #{m}"

m.times do
  puts "#{rand(n) + 1} #{rand(n) + 1}"
end
