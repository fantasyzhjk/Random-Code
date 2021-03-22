require 'parallel'

Parallel.each(1..50000000) do |i|
    puts "process #{Process.pid} receive #{i}"
end