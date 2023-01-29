-- Get all of a students submitted assignments that have been graded.
SELECT 
    *
FROM 
    submissions
JOIN 
    submission_grades ON submissions.submission_id = submission_grades.submission_id
JOIN 
    students ON submissions.student_id = students.student_id
WHERE 
    submissions.student_id = <student_id>;

-- Get the entire grade for a student's course. Could probably just use the average of all submissions.
SELECT 
    *
FROM 
    course_grades
JOIN 
    students ON course_grades.student_id = students.student_id
JOIN 
    courses ON course_grades.course_id = courses.course_id
WHERE 
    course_grades.student_id = <student_id>;

-- Get all assignment's for a student.
SELECT assignments.*
FROM submissions
JOIN assignments ON submissions.assignment_id = assignments.assignment_id
WHERE submissions.student_id = <student_id>;

-- Get all student's for a course.
SELECT students.*
FROM students
INNER JOIN courses
ON students.school_id = courses.school_id
WHERE courses.course_id = <course_id>;