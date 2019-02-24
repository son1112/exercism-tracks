require 'pry'

# Trinary converts trinary number to denary
class Trinary
  def initialize(trinary)
    @trinary = trinary
  end

  def to_decimal
    return 0 unless valid_trinary?

    @trinary.chars.each_with_index.sum do |char, index|
      to_denary(char.to_i, power(index))
    end
  end

  def power(num)
    (@trinary.length - 1) - num
  end

  def to_denary(digit, power)
    digit * (3**power)
  end

  def valid_trinary?

    binding.pry if @trinary[/Invalid/]

    @trinary.match?(/^[0-3]*$/)
  end
end
