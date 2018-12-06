BASE = 315_576_00.0

PLANETS = {
  'earth' => 1.0,
  'mercury' => 0.2408467,
  'venus' => 0.61519726,
  'mars' => 1.8808158,
  'jupiter' => 11.862615,
  'saturn' => 29.447498,
  'uranus' => 84.016846,
  'neptune' => 164.79132
}.freeze

# SpaceAge calculates how old someone would be on different planets from seconds
class SpaceAge
  def initialize(seconds)
    @seconds = seconds
    define_planet_methods
  end

  def define_planet_methods
    PLANETS.each do |planet, earth_years|
      method_name = "on_#{planet}".to_sym
      self.class.define_method(method_name) do
        @seconds / (BASE * earth_years)
      end
    end
  end
end
