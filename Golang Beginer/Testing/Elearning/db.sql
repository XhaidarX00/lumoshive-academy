-- Active: 1730083286169@@127.0.0.1@5432@elearning@public

-- Table Users
CREATE TABLE Users (
    user_id VARCHAR(10) PRIMARY KEY,
    username VARCHAR(100) NOT NULL UNIQUE,
    email VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    role VARCHAR(10) CHECK (role IN ('admin', 'student', 'mentor')) NOT NULL
);

SELECT * FROM grades;


CREATE TABLE ADMIN (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    user_id VARCHAR(10) REFERENCES Users (user_id),
    class_id VARCHAR(10) REFERENCES Users (user_id)
);

CREATE TABLE Mentor (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    degree VARCHAR(100) NOT NULL,
    experience INT,
    user_id VARCHAR(10) REFERENCES Users(user_id) ON DELETE SET NULL
);

SELECT * FROM Mentor;


CREATE TABLE Student (
    id SERIAL PRIMARY KEY,
    user_id VARCHAR(10) REFERENCES Users (user_id),
    name VARCHAR(50) NOT NULL,
    phone_number VARCHAR(13) NOT NULL,
    address VARCHAR(255) NOT NULL
);

SELECT * FROM student;

-- Table Materials
CREATE TABLE Materials (
    material_id VARCHAR(10) PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    video_url VARCHAR(255),
    mentor_id VARCHAR(10) REFERENCES Users(user_id) ON DELETE SET NULL
);

-- Table Classes
CREATE TABLE Classes (
    class_id VARCHAR(10) PRIMARY KEY,
    class_name VARCHAR(100) NOT NULL,
    mentor_id VARCHAR(10) REFERENCES Users(user_id) ON DELETE SET NULL
);


-- Table Class_Schedule
CREATE TABLE Class_Schedule (
    schedule_id VARCHAR(10) PRIMARY KEY,
    class_id VARCHAR(10) REFERENCES Classes(class_id) ON DELETE CASCADE,
    date DATE NOT NULL,
    start_time TIME NOT NULL,
    end_time TIME NOT NULL
);

-- Table Attendance
CREATE TABLE Attendance (
    attendance_id SERIAL PRIMARY KEY,
    user_id VARCHAR(10) REFERENCES Users(user_id) ON DELETE CASCADE,
    schedule_id VARCHAR(10) REFERENCES Class_Schedule(schedule_id) ON DELETE CASCADE,
    status VARCHAR(10) CHECK (status IN ('present', 'absent')) NOT NULL
);

-- Table Assignments
CREATE TABLE Assignments (
    assignment_id VARCHAR(10) PRIMARY KEY,
    class_id VARCHAR(10) REFERENCES Classes(class_id) ON DELETE CASCADE,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    deadline DATE NOT NULL
);

-- Table Grades
CREATE TABLE Grades (
    grade_id SERIAL PRIMARY KEY,
    user_id VARCHAR(10) REFERENCES Users(user_id) ON DELETE CASCADE,
    assignment_id VARCHAR(10) REFERENCES Assignments(assignment_id) ON DELETE CASCADE,
    grade INT CHECK (grade >= 0) NOT NULL
);

-- Table Announcements
CREATE TABLE Announcements (
    announcement_id VARCHAR(10) PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    date DATE DEFAULT CURRENT_DATE
);

-- Table Leaderboard
CREATE TABLE Leaderboard (
    leaderboard_id SERIAL PRIMARY KEY,
    user_id VARCHAR(10) REFERENCES Users(user_id) ON DELETE CASCADE,
    score INT CHECK (score >= 0) NOT NULL,
    ranking INT CHECK (ranking > 0)
);



-- Dummy Data untuk tabel Users
INSERT INTO Users (user_id, username, email, password, role) VALUES
('adm1', 'admin1', 'admin1@example.com', 'password', 'admin'),
('std1', 'student1', 'student1@example.com', 'password', 'student'),
('mnt1', 'mentor1', 'mentor1@example.com', 'password', 'mentor');


SELECT * FROM users;

-- Data untuk Admin, Mentor, dan Student
INSERT INTO Admin (user_id, name) VALUES ('adm1', 'Haidar');
INSERT INTO Mentor (user_id, name, degree, experience) VALUES ('mnt1', 'Agustina', 'PhD Computer Science', 10);
INSERT INTO Student (user_id, name, phone_number, address) VALUES ('std1', 'Nadila', '081234567890', '123 Elm Street');

-- Dummy Data untuk Classes
INSERT INTO Classes (class_id, class_name, mentor_id) VALUES ('cls1', 'Programming 101', 'mnt1');

TRUNCATE Classes;
-- Dummy Data untuk Materials
INSERT INTO Materials (material_id, title, description, video_url, mentor_id) VALUES
('mat1', 'Intro to Programming', 'Basics of programming', 'http://example.com/video1', 'mnt1');

TRUNCATE Class_Schedule;

SELECT * FROM Class_Schedule;

-- Dummy Data untuk Class_Schedule
INSERT INTO Class_Schedule (schedule_id, class_id, date, start_time, end_time) VALUES
('sch1', 'std1', '2024-12-01', '10:00', '12:00');

-- Dummy Data untuk Attendance
INSERT INTO Attendance (user_id, schedule_id, status) VALUES
('std1', 'sch1', 'present');

INSERT INTO Attendance (user_id, schedule_id, status) VALUES
('mnt1', 'sch1', 'present');

SELECT * FROM attendance;

-- Dummy Data untuk Assignments
INSERT INTO Assignments (assignment_id, class_id, title, description, deadline) VALUES
('asg1', 'cls1', 'Assignment 1', 'Solve basic programming problems', '2024-12-10');

TRUNCATE Assignments;

SELECT * FROM Assignments;

-- Dummy Data untuk Grades
INSERT INTO Grades (user_id, assignment_id, grade) VALUES
('std1', 'asg1', 85);

TRUNCATE grades;

-- Dummy Data untuk Announcements
INSERT INTO Announcements (announcement_id, title, content) VALUES
('anc1', 'Welcome!', 'Welcome to the new semester!');

SELECT * FROM Announcements;

-- Dummy Data untuk Leaderboard
INSERT INTO Leaderboard (user_id, score, ranking) VALUES
('std1', 100, 1);



UPDATE student SET name='Daman', phone_number='098340989324', address='jakarta' WHERE id='2';

SELECT * FROM student;

SELECT * FROM users;

TRUNCATE mentor;

SELECT * FROM student;
INSERT INTO student (user_id, name, phone_number, address) VALUES ('std2', 'testing', '08097897', 'bogor');
