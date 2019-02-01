# Binary converts a binary string to its decimal integer
module Binary
  def self.to_decimal(binary_string)
    validate(binary_string)

    binary_string.each_char.with_index.collect do |char, index|
      power = (binary_string.length - 1) - index
      binary_to_denary(char.to_i, power)
    end.sum
  end

  def self.binary_to_denary(digit, power)
    digit * (2**power)
  end

  def self.validate(binary_string)
    raise ArgumentError unless binary_string[/[^0-1]/].nil?
  end
end
