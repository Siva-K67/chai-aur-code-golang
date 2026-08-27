-- Courses table schema for the courses API project
-- Run this manually in psql/pgAdmin, or it's auto-created by connectDB() in db.go

CREATE TABLE IF NOT EXISTS courses (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    price INT NOT NULL
);

-- Sample data (matches local dev database)
INSERT INTO courses (id, name, price) VALUES
    (1, 'Boxing Guide by Ippo Makunochi', 499),
    (2, 'One Piece by Takeshi Obada', 599),
    (4, 'How to Bully your friends by Takamura Mamoru', 888),
    (5, 'Naruto by Kishimoto', 300),
    (6, 'Dragon Ball by Akira Toriyama', 399),
    (7, 'Death Note by Tsugumi Ohba', 699),
    (8, 'The Way of the Shinobi by Jiraya', 999);

-- Keeps the auto-increment counter in sync with the manually specified IDs above,
-- so the next INSERT (without an explicit id) continues from 9, not 1
SELECT setval('courses_id_seq', (SELECT MAX(id) FROM courses));