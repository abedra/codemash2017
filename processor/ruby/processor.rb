require 'redis'

class Processor
  attr_accessor :file, :threshold, :entries

  def initialize(file, threshold)
    @file = file
    @threshold = threshold
    @entries = Hash.new(0)
  end

  def run
    lines = File.readlines(file)
    lines.each do |line|
      parts = line.split(" ")
      address = parts[0]
      method = parts[5].gsub("\"", "")
      uri = parts[6]
      response_code = parts[8]

      if method == "POST" && response_code == "200"
        entries[address] += 1
      end
    end

    if entries.count > 0
      redis = Redis.new
      redis.multi do
        entries.each do |k,v|
          puts "Blacklisting #{k}. Threshold #{@threshold}, Actual: #{v}"
          redis.set "#{k}:repsheet:ip:blacklisted", "Failed login processor"
        end
      end
    end
  end
end

processor = Processor.new(ARGV[0], ARGV[1])
processor.run
