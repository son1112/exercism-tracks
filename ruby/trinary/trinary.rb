# Trinary converts trinary number to decimal equivalent
class Trinary
  def initialize(trinary)
    @trinary = trinary
  end

  def to_decimal
    return 0 unless valid_trinary?

    @trinary.each_char.with_index.collect do |char, index|
      to_denary(char.to_i, power(index))
    end.sum
  end

  def power(num)
    (@trinary.length - 1) - num
  end

  def to_denary(digit, power)
    digit * (3**power)
  end

  def valid_trinary?
    @trinary.count('^[0-3]').zero?
  end
end
