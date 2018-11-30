class School
  def initialize
    @students = {}
    @grade_list = []
  end

  def add(name, grade)
    student_list = @students[grade] || []
    @students[grade] = student_list << name
  end

  def students(grade)
    @students[grade] ? @students[grade].sort : []
  end

  def students_by_grade
    @students.each { |grade, students|
      @grade_list.push(
        { grade: grade, students: students.sort }
      )
    }
    @grade_list.sort_by { |hash| hash[:grade] }
  end
end
