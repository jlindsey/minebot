puts 'Starting'

resp_map = {
  'list' => <<-LIST,
Line 1
Line 2
Line 3
    LIST
}

loop do
  line = gets.strip
  break if line == 'exit'
  puts resp_map[line] || "Unknown command: #{line}"
end

puts 'Finished'
