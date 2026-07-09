-- + migrate up

CREATE TABLE plans (
    id SERIAL PRIMARY KEY,
    user_id INT NULL,
    name VARCHAR(100) NOT NULL,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_plans_user
        FOREIGN KEY (user_id)
        REFERENCES users(id)
        ON DELETE CASCADE
);

-- + migrate down

DROP TABLE plans;