# Trinary converts trinary number to decimal
class Trinary
  def initialize(trinary)
    @trinary = trinary
  end

  def to_decimal
    return 0 unless valid_trinary?

    digits.each_with_index.sum do |digit, index|
      digit * (3**index)
    end
  end

  private

  def digits
    @trinary.to_i.digits
  end

  def valid_trinary?
    return false if contains_newlines?

    @trinary.match?(/^[0-3]*$/)
  end

  def contains_newlines?
    @trinary.match?(/\n/)
  end
end
