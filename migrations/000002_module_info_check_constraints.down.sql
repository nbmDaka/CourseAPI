-- Drop existing constraints for module_info table
ALTER TABLE module_info
    DROP CONSTRAINT updated_at_check,
    DROP CONSTRAINT module_duration_check;
