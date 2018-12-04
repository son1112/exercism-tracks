module BeerSong
  def self.recite(number, count)
    statements = []

    while count > 0
      statements << statement(number, number - 1)
      number -= 1
      count -= 1
    end

    return statements.join("\n")
  end

  def self.statement(number, dec)
    first = "#{plurality(number, true)} of beer on the wall, #{plurality(number)} of beer."
    second = number > 0 ?
               "Take #{take(number)} down and pass it around, #{plurality(dec)} of beer on the wall."
               : "Go to the store and buy some more, 99 bottles of beer on the wall."

    <<-TEXT.gsub(/^ */, '')
    #{first}
    #{second}
    TEXT
  end

  def self.plurality(number, cap = nil)
    front = cap ? "No" : "no"
    return "#{front} more bottles" if number == 0
    number > 1 ? "#{number} bottles" : "#{number} bottle"
  end

  def self.take(number)
    number > 1 ? "one" : "it"
  end

  def self.remaining(number)
    number > 0 ? number : "no more"
  end
end
