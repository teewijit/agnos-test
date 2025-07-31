INSERT INTO hospitals (id, name) VALUES
(1, 'โรงพยาบาลศิริราช'),
(2, 'โรงพยาบาลรามาธิบดี');

INSERT INTO patients (
    id,
    first_name_th, middle_name_th, last_name_th,
    first_name_en, middle_name_en, last_name_en,
    date_of_birth, patient_hn,
    national_id, passport_id,
    phone_number, email, gender, hospital_id
) VALUES
-- Hospital 1 - 5 คน
(1, 'สมชาย', '', 'ใจดี', 'Somchai', '', 'Jaidee', '1980-07-15', 'HN6800001', '1234567890123', NULL, '0812345678', 'somchai.jaidee@example.com', 'M', 1),
(2, 'ชัชชัย', '', 'ทองดี', 'Chatchai', '', 'Thongdee', '1975-10-01', 'HN6800003', '1234567890124', NULL, '0811111111', 'chatchai@example.com', 'M', 1),
(3, 'ปาริชาติ', '', 'วัฒนศิลป์', 'Parichat', '', 'Wattanasin', '1988-05-23', 'HN6800004', '1234567890125', NULL, '0822222222', 'parichat@example.com', 'F', 1),
(4, 'อดิศร', '', 'สมบุญ', 'Adisorn', '', 'Somboon', '1990-11-11', 'HN6800005', NULL, 'B98765432', '0833333333', 'adisorn@example.com', 'M', 1),
(5, 'ณัฐพร', 'ศิริ', 'เพชรรัตน์', 'Nattaporn', 'Siri', 'Phetcharat', '1995-09-09', 'HN6800006', '1234567890126', NULL, '0844444444', 'nattaporn@example.com', 'F', 1),

-- Hospital 2 - 3 คน
(6, 'สุภาวดี', 'พิมพ์ใจ', 'ทองคำ', 'Supawadee', 'Pimjai', 'Thongkham', '1992-03-22', 'HN6800002', NULL, 'A12345678', '0898765432', 'supawadee.thongkham@example.com', 'F', 2),
(7, 'ปวริศ', '', 'รัตนโกสินทร์', 'Pawarit', '', 'Rattanakosin', '1983-04-18', 'HN6800007', '1234567890127', NULL, '0855555555', 'pawarit@example.com', 'M', 2),
(8, 'นงนภัส', '', 'สุนทรารักษ์', 'Nongnapat', '', 'Suntararak', '1999-12-31', 'HN6800008', '1234567890128', NULL, '0866666666', 'nongnapat@example.com', 'F', 2);