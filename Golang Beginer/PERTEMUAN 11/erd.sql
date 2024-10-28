-- Active: 1730083286169@@127.0.0.1@5432
CREATE DATABASE postgres;


-- Table Users
CREATE TABLE Users (
    user_id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    role VARCHAR(10) CHECK (role IN ('admin', 'student', 'mentor')) NOT NULL
);


CREATE TABLE ADMIN (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    user_id INT REFERENCES Users (user_id),
    class_id INT REFERENCES Users (user_id)
);

CREATE TABLE Mentor (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    degree VARCHAR(100) NOT NULL,
    experience INT,
    user_id INT REFERENCES Users(user_id) ON DELETE SET NULL
);



CREATE TABLE Student (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES Users (user_id),
    name VARCHAR(50) NOT NULL,
    phone_number VARCHAR(13) NOT NULL,
    address VARCHAR(255) NOT NULL
);



-- Table Materials
CREATE TABLE Materials (
    material_id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    video_url VARCHAR(255),
    mentor_id INT REFERENCES Users(user_id) ON DELETE SET NULL
);

-- Table Classes
CREATE TABLE Classes (
    class_id SERIAL PRIMARY KEY,
    class_name VARCHAR(100) NOT NULL,
    mentor_id INT REFERENCES Users(user_id) ON DELETE SET NULL
);


-- Table Class_Schedule
CREATE TABLE Class_Schedule (
    schedule_id SERIAL PRIMARY KEY,
    class_id INT REFERENCES Classes(class_id) ON DELETE CASCADE,
    date DATE NOT NULL,
    start_time TIME NOT NULL,
    end_time TIME NOT NULL
);

-- Table Attendance
CREATE TABLE Attendance (
    attendance_id SERIAL PRIMARY KEY,
    user_id INT REFERENCES Users(user_id) ON DELETE CASCADE,
    schedule_id INT REFERENCES Class_Schedule(schedule_id) ON DELETE CASCADE,
    status VARCHAR(10) CHECK (status IN ('present', 'absent')) NOT NULL
);

-- Table Assignments
CREATE TABLE Assignments (
    assignment_id SERIAL PRIMARY KEY,
    class_id INT REFERENCES Classes(class_id) ON DELETE CASCADE,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    deadline DATE NOT NULL
);

-- Table Grades
CREATE TABLE Grades (
    grade_id SERIAL PRIMARY KEY,
    user_id INT REFERENCES Users(user_id) ON DELETE CASCADE,
    assignment_id INT REFERENCES Assignments(assignment_id) ON DELETE CASCADE,
    grade INT CHECK (grade >= 0) NOT NULL
);

-- Table Announcements
CREATE TABLE Announcements (
    announcement_id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    date DATE DEFAULT CURRENT_DATE
);

-- Table Leaderboard
CREATE TABLE Leaderboard (
    leaderboard_id SERIAL PRIMARY KEY,
    user_id INT REFERENCES Users(user_id) ON DELETE CASCADE,
    score INT CHECK (score >= 0) NOT NULL,
    ranking INT CHECK (ranking > 0)
);
