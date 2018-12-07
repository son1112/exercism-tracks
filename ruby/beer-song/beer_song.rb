require 'pry'

# BeerSong recites the class beer song from number of beers down by count
module BeerSong
  def self.recite(number, count)
    statements = []

    (0...count).to_a.map do |dec|
      statements << statement(number - dec)
    end

    statements.join("\n")
  end

  def self.statement(number)
    dec = number - 1
    current_bottles = plurality(number)
    remaining_bottles = plurality(dec)

    <<-TEXT.gsub(/^ */, '')
    #{current_bottles.capitalize} of beer on the wall, #{current_bottles} of beer.
    #{take(number)}, #{remaining_bottles} of beer on the wall.
    TEXT
  end

  def self.plurality(number_of_bottles)
    number = number_of_bottles

    if number > 1
      "#{number} bottles"
    elsif number == 1
      "#{number} bottle"
    elsif number.zero?
      'no more bottles'
    elsif number < 0
      '99 bottles'
    end
  end

  def self.take(number)
    if number > 1
      'Take one down and pass it around'
    elsif number == 1
      'Take it down and pass it around'
    elsif number.zero?
      'Go to the store and buy some more'
    end
  end
end
