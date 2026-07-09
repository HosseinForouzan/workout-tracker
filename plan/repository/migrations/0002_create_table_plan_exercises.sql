-- + migrate up

CREATE TABLE plan_exercises (
    id SERIAL PRIMARY KEY,

    plan_id INT NOT NULL,
    exercise_id INT NOT NULL,

    target_sets SMALLINT NOT NULL,
    target_reps SMALLINT NOT NULL,
    target_weight NUMERIC(5,2),

    sort_order SMALLINT NOT NULL,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_plan_exercises_plan
        FOREIGN KEY (plan_id)
        REFERENCES plans(id)
        ON DELETE CASCADE,

    CONSTRAINT fk_plan_exercises_exercise
        FOREIGN KEY (exercise_id)
        REFERENCES exercises(id)

);

-- + migrate down

DROP TABLE plans;