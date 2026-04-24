DROP TABLE IF EXISTS album;
CREATE TABLE album (
  id         INT AUTO_INCREMENT NOT NULL,
  title      VARCHAR(128) NOT NULL,
  artist     VARCHAR(255) NOT NULL,
  released   YEAR NOT NULL,
  PRIMARY KEY (`id`)
);

INSERT INTO album
  (title, artist, released)
VALUES
  ('Going Under', 'Evanescence', 2003),
  ('Now You re Gone', 'Basshunter', 2007),
  ('GO BABY', 'Justin Bieber', 2025),
  ('Gone with the Sin', 'HIM', 1999);
