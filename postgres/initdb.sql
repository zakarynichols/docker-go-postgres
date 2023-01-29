CREATE TYPE school_type AS ENUM ('public', 'private');
CREATE TYPE faculty_role AS ENUM ('instructor');

CREATE TABLE schools (
    school_id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    location VARCHAR(255) NOT NULL,
    type school_type NOT NULL
);

CREATE TABLE faculty_members (
    faculty_member_id SERIAL PRIMARY KEY,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    role faculty_role NOT NULL,
    school_id INT NOT NULL,
    FOREIGN KEY (school_id) REFERENCES schools(school_id)
);

CREATE TABLE courses (
    course_id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    school_id INT NOT NULL,
    faculty_member_id INT NOT NULL,
    credits INT NOT NULL,
    FOREIGN KEY (school_id) REFERENCES schools(school_id),
    FOREIGN KEY (faculty_member_id) REFERENCES faculty_members(faculty_member_id)
);

CREATE TABLE students (
    student_id SERIAL PRIMARY KEY,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    school_id INT NOT NULL,
    FOREIGN KEY (school_id) REFERENCES schools(school_id)
);

CREATE TABLE course_grades (
    grade_id SERIAL PRIMARY KEY,
    student_id INT NOT NULL,
    course_id INT NOT NULL,
    grade FLOAT NOT NULL,
    FOREIGN KEY (student_id) REFERENCES students(student_id),
    FOREIGN KEY (course_id) REFERENCES courses(course_id)
);

CREATE TABLE assignments (
    assignment_id SERIAL PRIMARY KEY,
    course_id INT NOT NULL,
    name VARCHAR(255) NOT NULL,
    due_date DATE NOT NULL,
    FOREIGN KEY (course_id) REFERENCES courses(course_id)
);

CREATE TABLE submissions (
    submission_id SERIAL PRIMARY KEY,
    student_id INT NOT NULL,
    assignment_id INT NOT NULL,
    submission_date DATE NOT NULL,
    FOREIGN KEY (student_id) REFERENCES students(student_id),
    FOREIGN KEY (assignment_id) REFERENCES assignments(assignment_id)
);

CREATE TABLE submission_grades (
    grade_id SERIAL PRIMARY KEY,
    submission_id INT NOT NULL,
    faculty_member_id INT NOT NULL,
    grade FLOAT NOT NULL,
    FOREIGN KEY (submission_id) REFERENCES submissions(submission_id),
    FOREIGN KEY (faculty_member_id) REFERENCES faculty_members(faculty_member_id)
);

CREATE TABLE attendance (
    attendance_id SERIAL PRIMARY KEY,
    student_id INT NOT NULL,
    course_id INT NOT NULL,
    date DATE NOT NULL,
    present BOOLEAN NOT NULL,
    FOREIGN KEY (student_id) REFERENCES students(student_id),
    FOREIGN KEY (course_id) REFERENCES courses(course_id)
);

INSERT INTO schools (school_id, name, location, type) VALUES
    (1, 'University of XYZ', 'City A', 'public'),
    (2, 'College of ABC', 'City B', 'private');

INSERT INTO faculty_members (faculty_member_id, first_name, last_name, email, role, school_id) VALUES
    (1, 'John', 'Doe', 'johndoe@email.com', 'instructor', 1),
    (2, 'Jane', 'Smith', 'janesmith@email.com', 'instructor', 1),
    (3, 'Bob', 'Johnson', 'bobjohnson@email.com', 'instructor', 2);

INSERT INTO courses (course_id, name, school_id, faculty_member_id, credits) VALUES
    (1, 'Introduction to Computer Science', 1, 1, 3),
    (2, 'Calculus', 1, 2, 4),
    (3, 'Physics', 2, 3, 5);

INSERT INTO students (student_id, first_name, last_name, email, school_id) VALUES 
    (1, 'John', 'Doe', 'johndoe@email.com', 1),
    (2, 'Jane', 'Smith', 'janesmith@email.com', 1),
    (3, 'Bob', 'Johnson', 'bobjohnson@email.com', 2);

INSERT INTO course_grades (grade_id, student_id, course_id, grade) VALUES
    (1, 1, 1, 90),
    (2, 1, 2, 85),
    (3, 2, 1, 95),
    (4, 3, 3, 80);

INSERT INTO assignments (assignment_id, course_id, name, due_date) VALUES
    (1, 1, 'Introduction to Programming', '2022-02-28'),
    (2, 1, 'Data Structures', '2022-03-15'),
    (3, 2, 'Calculus I', '2022-03-12'),
    (4, 2, 'Calculus II', '2022-04-02'),
    (5, 3, 'Introduction to Database Systems', '2022-04-20'),
    (6, 3, 'SQL Queries', '2022-05-04');

INSERT INTO submissions (submission_id, student_id, assignment_id, submission_date) VALUES
    (1, 1, 1, '2022-02-28'),
    (2, 2, 1, '2022-03-01'),
    (3, 3, 2, '2022-03-11'),
    (4, 3, 1, '2022-03-13'),
    (5, 2, 3, '2022-04-01'),
    (6, 1, 4, '2022-04-03'),
    (7, 3, 5, '2022-04-19'),
    (8, 2, 6, '2022-05-03');

INSERT INTO attendance (attendance_id, student_id, course_id, date, present) VALUES
    (1, 1, 1, '12-21-2022', true),
    (2, 2, 2, '12-21-2022', false),
    (3, 3, 1, '12-21-2022', false),
    (4, 3, 1, '12-21-2022', false),
    (5, 3, 1, '12-21-2022', false),
    (6, 3, 1, '12-21-2022', true);

INSERT INTO submission_grades (grade_id, submission_id, faculty_member_id, grade) VALUES
    (1, 1, 1, 88.0),
    (2, 2, 1, 85.0),
    (3, 3, 2, 92.5),
    (4, 4, 2, 90.0);

SELECT * FROM schools;
SELECT * FROM courses;
SELECT * FROM students;
SELECT * FROM faculty_members;
SELECT * FROM course_grades;
SELECT * FROM assignments;
SELECT * FROM submissions;
SELECT * FROM submission_grades;
SELECT * FROM attendance;