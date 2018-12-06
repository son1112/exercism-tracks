# Binary converts a binary string to its decimal integer
module Binary
  def self.to_decimal(binary_string)
    validate(binary_string)
    denary = []

    binary_string.each_char.with_index do |char, index|
      power = (binary_string.length - 1) - index
      denary << binary_to_denary(char.to_i, power)
    end

    denary.sum
  end

  def self.binary_to_denary(digit, power)
    digit * (2**power)
  end

  def self.validate(binary_string)
    raise ArgumentError unless binary_string.count('^0 - 1').zero?
  end
end
