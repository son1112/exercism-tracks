require 'pry'

# Trinary converts trinary number to decimal equivalent
class Trinary
  def initialize(trinary)
    @trinary = trinary
  end

  def to_decimal
    return 0 unless valid_trinary?

    denary = []

    @trinary.each_char.with_index do |char, index|
      power = (@trinary.length - 1) - index
      denary << trinary_to_denary(char.to_i, power)
    end

    denary.sum
  end

  def trinary_to_denary(digit, power)
    digit * (3**power)
  end

  def valid_trinary?
    @trinary.count('^[0-3]').zero?
  end
end
