-- ---------------
-- 这是注释
-- BMS
-- ---------------

CREATE TABLE book(
    id bigint(20) AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(20) NOT NULL,
    price DOUBLE(16,2) NOT NULL
)engine=InnoDB DEFAULT charset=utf8mb4;