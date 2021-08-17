CREATE TABLE contents(
                         id SERIAL PRIMARY KEY,
                         created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                         updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                         title VARCHAR(100) NOT NULL,
                         details TEXT NOT NULL
)