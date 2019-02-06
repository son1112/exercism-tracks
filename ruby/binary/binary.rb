# Binary converts a binary string to its decimal integer
module Binary
  def self.to_decimal(bin_str)
    validate(bin_str)

    bin_str.chars.each_with_index.sum do |char, index|
      power = (bin_str.length - 1) - index
      binary_to_denary(char.to_i, power)
    end
  end

  def self.binary_to_denary(digit, power)
    digit * (2**power)
  end

  def self.validate(binary_string)
    raise ArgumentError unless binary_string[/[^0-1]/].nil?
  end
end
