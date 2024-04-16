-- Add constraint checks for updated_at column
ALTER TABLE module_info
    ADD CONSTRAINT updated_at_check
        CHECK (updatedAt >= createdAt);

-- Add constraint checks for module_duration column
ALTER TABLE module_info
    ADD CONSTRAINT module_duration_check
        CHECK (moduleDuration > 5 AND moduleDuration <= 15);