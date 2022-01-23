CREATE TABLE IF NOT EXISTS fav_items (
  bid TEXT PRIMARY KEY,
  title TEXT NOT NULL,
  description TEXT NOT NULL,
  add_fav_time INTEGER,
  link TEXT NOT NULL,
  author TEXT NOT NULL
);
