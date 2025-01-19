-- +goose Up
-- +goose StatementBegin

CREATE SCHEMA users;

SET
SEARCH_PATH TO users, PUBLIC;

CREATE TABLE users (
  id           TEXT         NOT NULL,
  account_id   TEXT         NOT NULL UNIQUE,
  name         VARCHAR(50),
  age          INT CHECK (age >= 0 AND age <= 120),
  gender       VARCHAR(10),
  location     GEOGRAPHY(POINT, 4326),
  status       VARCHAR(10)  NOT NULL,
  created_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  deleted_at   TIMESTAMP,
  PRIMARY KEY (id)
);

CREATE INDEX idx_users_deleted_at ON users (deleted_at);
CREATE INDEX idx_users_location ON users USING GIST (location);

CREATE TABLE preferences (
  id           SERIAL,
  user_id      TEXT         NOT NULL,
  min_age      INT CHECK (min_age >= 0 AND min_age <= 120),
  max_age      INT CHECK (max_age >= 0 AND max_age <= 120),
  gender       VARCHAR(10),
  distance     INT CHECK (distance >= 0),
  created_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  deleted_at   TIMESTAMP,
  PRIMARY KEY (id),
  CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

--- Interest table
CREATE TABLE interests (
  id           SERIAL,
  name         VARCHAR(50)  NOT NULL UNIQUE,
  created_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  deleted_at   TIMESTAMP,
  PRIMARY KEY (id)
);
-- 為 deleted_at 添加索引
CREATE INDEX idx_interests_deleted_at ON interests (deleted_at);

CREATE TABLE user_interests (
  user_id      TEXT         NOT NULL,
  interest_id  INT          NOT NULL,
  CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
  CONSTRAINT fk_interest FOREIGN KEY (interest_id) REFERENCES interests(id) ON DELETE CASCADE,
  PRIMARY KEY (user_id, interest_id)
);

-- 為 user_id 添加索引
CREATE INDEX idx_user_interests_user_id ON user_interests (user_id);
-- 為 interest_id 添加索引
CREATE INDEX idx_user_interests_interest_id ON user_interests (interest_id);

-- Photos table
CREATE TABLE photos (
  id           TEXT          NOT NULL,
  url          VARCHAR(255)  NOT NULL UNIQUE,
  created_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  deleted_at   TIMESTAMP,
  PRIMARY KEY (id)
);

-- 為 deleted_at 添加索引
CREATE INDEX idx_photos_deleted_at ON photos (deleted_at);
-- 為 url 添加索引
CREATE INDEX idx_photos_url ON photos (url);

-- User photos table
CREATE TABLE user_photos (
  user_id      TEXT         NOT NULL,
  photo_id     TEXT         NOT NULL, -- 與 photos 表的主鍵一致
  CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
  CONSTRAINT fk_photo FOREIGN KEY (photo_id) REFERENCES photos(id) ON DELETE CASCADE,
  PRIMARY KEY (user_id, photo_id)
);
-- 為 user_id 添加索引
CREATE INDEX idx_user_photos_user_id ON user_interests (user_id);
-- 為 photo_id 添加索引
CREATE INDEX idx_user_photos_photo_id ON user_photos (photo_id);

-- Matches table
CREATE TABLE matches (
  id           SERIAL,
  user_a_id    TEXT         NOT NULL,
  user_b_id    TEXT         NOT NULL,
  status       VARCHAR(10)  NOT NULL,
  created_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  deleted_at   TIMESTAMP,
  PRIMARY KEY (id),
  CONSTRAINT unique_matches UNIQUE (user_a_id, user_b_id)
);

-- 為 deleted_at 添加索引
CREATE INDEX idx_matches_deleted_at ON matches (deleted_at);

CREATE TABLE chats (
  id           TEXT      NOT NULL,
  match_id     INT       NOT NULL,
  created_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  deleted_at   TIMESTAMP,
  PRIMARY KEY (id),
  CONSTRAINT fk_match FOREIGN KEY (match_id) REFERENCES matches(id) ON DELETE CASCADE
);

CREATE TABLE messages (
  id            SERIAL,
  chat_id       TEXT         NOT NULL,
  sender_id     TEXT         NOT NULL,
  content       TEXT         NOT NULL,
  created_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  deleted_at    TIMESTAMP,
  PRIMARY KEY (id),
  CONSTRAINT fk_chat_room FOREIGN KEY (chat_id) REFERENCES chats (id) ON DELETE CASCADE,
  CONSTRAINT fk_user FOREIGN KEY (sender_id) REFERENCES users (id) ON DELETE CASCADE
);

-- -- 設定隨機用戶資料生成
-- WITH random_users AS (
--   SELECT
--     -- 隨機生成名字
--     CASE
--       WHEN random() < 0.5 THEN 'John' ELSE 'Jane'
--     END || ' ' || substring(md5(random()::text), 1, 5) AS name,

--     -- 隨機生成年齡
--     (floor(random() * 80) + 18) AS age,

--     -- 隨機生成性別
--     CASE
--       WHEN random() < 0.5 THEN 'male' ELSE 'female'
--     END AS gender,

--     -- 隨機生成位置 (經度、緯度)
--     ST_SetSRID(ST_MakePoint(
--       random() * 360 - 180,  -- 隨機經度 (-180 到 180)
--       random() * 180 - 90   -- 隨機緯度 (-90 到 90)
--     ), 4326) AS location
--   FROM generate_series(1, 100)  -- 生成 100 條數據
-- )

-- -- 插入隨機生成的用戶資料到 users 表
-- INSERT INTO users (name, age, gender, location)
-- SELECT name, age, gender, location
-- FROM random_users;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP SCHEMA IF EXISTS users CASCADE;
-- +goose StatementEnd
