-- Add constraint checks for updated_at column
ALTER TABLE module_info
    ADD CONSTRAINT updated_at_check
        CHECK (updated_at >= created_at);

-- Add constraint checks for module_duration column
ALTER TABLE module_info
    ADD CONSTRAINT module_duration_check
        CHECK (module_duration > 5 AND module_duration <= 15);