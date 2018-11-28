module Complement
  COMPLEMENTS = {
    'C': 'G',
    'G': 'C',
    'T': 'A',
    'A': 'U',
     '': ''
  }

  def self.of_dna(string)
    string_complement = []

    string.each_char { |char|
      string_complement << COMPLEMENTS[char.to_sym]
    }

    string_complement.join
  end
end
