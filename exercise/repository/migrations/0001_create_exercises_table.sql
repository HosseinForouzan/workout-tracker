-- + migrate up

CREATE TABLE exercises(
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    muscle_group TEXT NOT NULL
    CHECK (muscle_group IN ('chest','back','legs','shoulders', 'arms', 'core')),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);


-- + migrate down

DROP TABLE exersices;