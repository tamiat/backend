CREATE TABLE contents (
   id bigserial PRIMARY KEY,
   title varchar(150) NOT NULL,
   details TEXT NOT NULL,
   created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
   updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);