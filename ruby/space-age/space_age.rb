# SpaceAge calculates how old someone would be on different planets from seconds
class SpaceAge
  EARTH_ORBITAL = 315_576_00.0 # seconds

  PLANETS = {
    earth: 1.0,
    mercury: 0.2408467,
    venus: 0.61519726,
    mars: 1.8808158,
    jupiter: 11.862615,
    saturn: 29.447498,
    uranus: 84.016846,
    neptune: 164.79132
  }.freeze

  PLANETS.each do |planet, earth_years|
    method_name = "on_#{planet}"
    define_method(method_name) do
      calculate_age(earth_years)
    end
  end

  attr_reader :seconds

  def initialize(seconds)
    @seconds = seconds
  end

  def calculate_age(earth_years)
    seconds / (EARTH_ORBITAL * earth_years)
  end
end
