module Complement

  def self.of_dna(string)
    string_complement = []

    string.split('').each { |char|
      string_complement << complements[char.to_sym]
    }

    string_complement.join
  end

  def self.complements
    {
      'C': 'G',
      'G': 'C',
      'T': 'A',
      'A': 'U',
      '': ''
    }
  end
end
