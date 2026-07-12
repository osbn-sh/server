ALTER TABLE vote
        DROP CONSTRAINT vote_pkey;

ALTER TABLE vote
    ADD CONSTRAINT vote_unique_key UNIQUE (
                                           user_id,
                                           option_id,
                                           university_id,
                                           professor_id,
                                           major_id,
                                           lesson_id
        );


ALTER TABLE vote
    ADD COLUMN id serial PRIMARY KEY