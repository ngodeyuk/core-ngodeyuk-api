CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Table for Course
CREATE TABLE courses (
    course_id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    img VARCHAR(255)
);

-- Table for User
CREATE TABLE users (
    user_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    course_id INTEGER REFERENCES courses(course_id) ON DELETE SET NULL,
    name VARCHAR(255) NOT NULL,
    img_url VARCHAR(255),
    username VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    heart INTEGER DEFAULT 5,
    points INTEGER DEFAULT 0
);

-- Table for Unit
CREATE TABLE units (
    unit_id SERIAL PRIMARY KEY,
    course_id INTEGER REFERENCES courses(course_id) ON DELETE CASCADE,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    sequence INTEGER
);

-- Table for Lesson
CREATE TABLE lessons (
    lesson_id SERIAL PRIMARY KEY,
    unit_id INTEGER REFERENCES units(unit_id) ON DELETE CASCADE,
    title VARCHAR(255) NOT NULL,
    sequence INTEGER
);

-- Table for Challenge
CREATE TABLE challenges (
    challenge_id SERIAL PRIMARY KEY,
    lesson_id INTEGER REFERENCES lessons(lesson_id) ON DELETE CASCADE,
    question TEXT NOT NULL,
    sequence INTEGER
);

-- Table for ChallengeProgress
CREATE TABLE challenge_progresses (
    challenge_progress_id SERIAL PRIMARY KEY,
    user_id UUID REFERENCES users(user_id) ON DELETE CASCADE,
    challenge_id INTEGER REFERENCES challenges(challenge_id) ON DELETE CASCADE,
    completed BOOLEAN DEFAULT FALSE
);

-- Table for ChallengeOption
CREATE TABLE challenge_options (
    challenge_option_id SERIAL PRIMARY KEY,
    challenge_id INTEGER REFERENCES challenges(challenge_id) ON DELETE CASCADE,
    text TEXT NOT NULL,
    correct BOOLEAN DEFAULT FALSE,
    img_url VARCHAR(255)
);
