CREATE TABLE todos (
	id INT AUTO_INCREMENT PRIMARY KEY,
	user_id int not null,
	title VARCHAR(255) NOT NULL,
	description TEXT,
	status ENUM('pending', 'done') DEFAULT 'pending',
	created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
	updated_at DATETIME ON UPDATE CURRENT_TIMESTAMP
);