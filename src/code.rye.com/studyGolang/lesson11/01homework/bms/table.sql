-- ---------------
-- 这是注释
-- BMS
-- ---------------

CREATE TABLE book(
    id bigint(20) AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(20) NOT NULL,
    price DOUBLE(16,2) NOT NULL,
    publisher_id bigint(20) DEFAULT 0
)engine=InnoDB DEFAULT charset=utf8mb4;

-- 出版社表
CREATE TABLE publisher(
    id bigint(20) AUTO_INCREMENT PRIMARY KEY,
    province varchar(20) NOT NULL,
    city varchar(20) NOT NULL
)engine=InnoDB DEFAULT charset=utf8mb4;