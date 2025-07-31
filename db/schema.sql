-- Create hospitals table
CREATE TABLE IF NOT EXISTS hospitals (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ,
    name TEXT UNIQUE NOT NULL
);

-- Create staffs table
CREATE TABLE IF NOT EXISTS staffs (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ,
    username TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL,
    hospital_id BIGINT NOT NULL,
    CONSTRAINT fk_hospitals_staffs FOREIGN KEY (hospital_id) REFERENCES hospitals(id)
);

-- Create patients table
CREATE TABLE IF NOT EXISTS patients (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ,
    first_name_th TEXT,
    middle_name_th TEXT,
    last_name_th TEXT,
    first_name_en TEXT,
    middle_name_en TEXT,
    last_name_en TEXT,
    date_of_birth TEXT,
    patient_hn TEXT,
    national_id TEXT,
    passport_id TEXT,
    phone_number TEXT,
    email TEXT,
    gender TEXT,
    hospital_id BIGINT,
    CONSTRAINT fk_hospitals_patients FOREIGN KEY (hospital_id) REFERENCES hospitals(id)
);

-- Insert initial hospital data
INSERT INTO hospitals (name) VALUES
('โรงพยาบาลศิริราช'),
('โรงพยาบาลรามาธิบดี')
ON CONFLICT (name) DO NOTHING; 