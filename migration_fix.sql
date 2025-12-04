-- Migration to add missing columns
-- Run this in Cloud SQL Query Editor

-- Add image column to articles table
ALTER TABLE articles ADD COLUMN IF NOT EXISTS image TEXT;

-- Add status column to donations table (if not exists)
DO $$ 
BEGIN
    IF NOT EXISTS (SELECT 1 FROM information_schema.columns 
                   WHERE table_name='donations' AND column_name='status') THEN
        ALTER TABLE donations ADD COLUMN status donation_status NOT NULL DEFAULT 'pending';
    END IF;
END $$;

-- Verify changes
SELECT table_name, column_name, data_type 
FROM information_schema.columns 
WHERE table_name IN ('articles', 'donations')
ORDER BY table_name, ordinal_position;
