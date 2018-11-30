module Complement
  COMPLEMENTS = {
    'C' => 'G',
    'G' => 'C',
    'T' => 'A',
    'A' => 'U'
  }

  def self.of_dna(string)
    string.each_char.sum('') { |char|  COMPLEMENTS[char] || "" }
  end
end
