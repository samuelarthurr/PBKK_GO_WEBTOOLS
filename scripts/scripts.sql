USE gowtdb;

-- Create Categories table
CREATE TABLE IF NOT EXISTS categories (
    id int(11) NOT NULL AUTO_INCREMENT,
    name varchar(80) COLLATE utf8_unicode_ci NOT NULL,
    description text COLLATE utf8_unicode_ci,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- Insert default categories based on existing unique categories
INSERT INTO categories (name, description)
SELECT DISTINCT category, 'Imported from original tools table'
FROM tools;

-- Create temporary table to store existing tools
CREATE TEMPORARY TABLE temp_tools SELECT * FROM tools;

-- Modify tools table to add category_id
ALTER TABLE tools
ADD COLUMN category_id int(11) AFTER name,
ADD FOREIGN KEY (category_id) REFERENCES categories(id);

-- Update tools with correct category IDs
UPDATE tools t
JOIN categories c ON t.category = c.name
SET t.category_id = c.id;

-- Now we can safely drop the old category column
ALTER TABLE tools DROP COLUMN category;