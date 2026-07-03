-- lesson
ALTER TABLE lesson DROP CONSTRAINT IF EXISTS unique_name;
ALTER TABLE lesson DROP CONSTRAINT IF EXISTS unique_name_english;

ALTER TABLE lesson
    ADD COLUMN IF NOT EXISTS href TEXT,
    ADD CONSTRAINT unique_lesson_href UNIQUE (href);

-- professor
ALTER TABLE professor
    ADD COLUMN IF NOT EXISTS href TEXT,
    ADD CONSTRAINT unique_professor_href UNIQUE (href);

-- university
ALTER TABLE university
    ADD COLUMN IF NOT EXISTS href TEXT,
    ADD CONSTRAINT unique_university_href UNIQUE (href);

-- major
ALTER TABLE major
    ADD COLUMN IF NOT EXISTS href TEXT,
    ADD CONSTRAINT unique_major_href UNIQUE (href);
