-- +goose Up
-- +goose StatementBegin
INSERT INTO accounts.accounts (id, email)
  VALUES ('019478c9-5e52-770c-9bde-a7d6ea0bb768', 'john_doe@example.com');

INSERT INTO accounts.accounts (id, phone_number)
  VALUES ('019478c9-5e52-73af-a486-ccb32a31239d', '0912345678');

INSERT INTO auths.otps (account_id, code,  is_verified, expiration)
  VALUES ('019478c9-5e52-73af-a486-ccb32a31239d', '123456', true, '2050-01-01 00:00:00');

INSERT INTO users.users (id, account_id, name, age, gender, location, status)
  VALUES ('01947d89-ea65-7f59-b0e4-66c60415b807', '019478c9-5e52-770c-9bde-a7d6ea0bb768', 'John Doe', 30, 'male', ST_SetSRID(ST_GeomFromText('POINT(121.597366 25.105497)'),4326), 'active');

INSERT INTO users.users (id, account_id, name, age, gender, location, status)
  VALUES ('01947d89-ea65-7d4d-99e1-fa68871e8803', '019478c9-5e52-73af-a486-ccb32a31239d', 'Jenny Cheng', 20, 'female', ST_SetSRID(ST_GeomFromText('POINT(120.266670 22.633333)'),4326), 'active');

INSERT INTO users.interests (name)
  VALUES ('travel'), ('music'), ('food'), ('movie'), ('sport'), ('game');

INSERT INTO users.matches (id, user_a_id, user_b_id, status)
  VALUES (1, '01947d89-ea65-7f59-b0e4-66c60415b807', '01947d89-ea65-7d4d-99e1-fa68871e8803', 'accepted');

INSERT INTO users.chats (id, match_id)
  VALUES ('01948385-805d-7400-bcde-8f4b84efef4d', 1);

INSERT INTO users.messages (chat_id, sender_id, content)
  VALUES ('01948385-805d-7400-bcde-8f4b84efef4d', '01947d89-ea65-7f59-b0e4-66c60415b807', 'Hello, how are you?'),
        ('01948385-805d-7400-bcde-8f4b84efef4d', '01947d89-ea65-7d4d-99e1-fa68871e8803', 'Hi, I am fine. How about you?');

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

-- +goose StatementEnd
