CREATE TABLE pastes (
  shortlink char(7) NOT NULL,
  expiration_length_in_minutes int NOT NULL,
  created_at datetime NOT NULL,
  paste_path varchar(255) NOT NULL,
  PRIMARY KEY(shortlink)
);