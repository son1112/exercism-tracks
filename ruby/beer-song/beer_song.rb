require 'pry'

# BeerSong recites the class beer song from number of beers down by count
module BeerSong
  def self.recite(number, count)
    (0...count).map do |dec|
      statement(number - dec)
    end.join("\n")
  end

  def self.statement(number)
    dec = number - 1
    current_bottles = x_bottles(number)
    remaining_bottles = x_bottles(dec)

    <<-TEXT.gsub(/^ */, '')
    #{current_bottles.capitalize} of beer on the wall, #{current_bottles} of beer.
    #{take(number)}, #{remaining_bottles} of beer on the wall.
    TEXT
  end

  def self.x_bottles(number)
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
