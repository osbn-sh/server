ALTER TABLE option
    ADD COLUMN owner TEXT NOT NULL
        CHECK (owner IN ('university', 'professor', 'major', 'lesson'));