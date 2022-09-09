CREATE TABLE user
(
  id INT NOT NULL,
  first_name VARCHAR(20) NOT NULL,
  last_name VARCHAR(20) NOT NULL,
  gender VARCHAR(10) NOT NULL,
  DOB DATE NOT NULL,
  address VARCHAR(50) NOT NULL,
  email VARCHAR(30) NOT NULL,
  password VARCHAR(20) NOT NULL,
  mob_no INT NOT NULL,
  role VARCHAR(10) NOT NULL,
  PRIMARY KEY (id),
  UNIQUE (email)
);