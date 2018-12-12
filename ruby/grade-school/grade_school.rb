# School provides a class for creating a school roster
class School
  def initialize
    @students = {}
    @grade_list = []
  end

  def add(name, grade)
    (@students[grade] ||= []) << name
  end

  def students(grade)
    return [] if @students[grade].nil?

    @students[grade].sort
  end

  def students_by_grade
    collection = @students.collect do |grade, students|
      { grade: grade, students: students.sort }
    end

    collection.sort_by { |hash| hash[:grade] }
  end
end
