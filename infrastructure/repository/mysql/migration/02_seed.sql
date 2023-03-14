-- migrate:up
INSERT INTO userbirthday.users (
  id, name, email, phone, is_verified, 
  birthdate, created_at, updated_at
) 
VALUES 
  (
    'USER_1', 'NOT_BIRTHDAY_NOT_VERIFIED', 
    'user.1@email.com', '08111', 0, CURRENT_TIMESTAMP + INTERVAL 7 DAY, 
    CURRENT_TIMESTAMP, CURRENT_TIMESTAMP
  ), 
  (
    'USER_2', 'OK_BIRTHDAY_NOT_VERIFIED', 
    'user.2@email.com', '08222', 0, CURRENT_TIMESTAMP, 
    CURRENT_TIMESTAMP, CURRENT_TIMESTAMP
  ), 
  (
    'USER_3', 'NOT_BIRTHDAY_OK_VERIFIED', 
    'user.3@email.com', '08333', 1, CURRENT_TIMESTAMP + INTERVAL 7 DAY, 
    CURRENT_TIMESTAMP, CURRENT_TIMESTAMP
  ), 
  (
    'USER_4', 'OK_BIRTHDAY_OK_VERIFIED', 
    'user.4@email.com', '08444', 1, CURRENT_TIMESTAMP, 
    CURRENT_TIMESTAMP, CURRENT_TIMESTAMP
  ), 
  (
    'USER_5', 'OK_BIRTHDAY_OK_VERIFIED_BUT_ALREADY_HAVE_PROMO', 
    'user.5@email.com', '08555', 1, CURRENT_TIMESTAMP, 
    CURRENT_TIMESTAMP, CURRENT_TIMESTAMP
  );

INSERT INTO promotions (
  id, code, `type`, amount, use_limit, 
  valid_from, valid_to, created_at, 
  updated_at
) 
VALUES 
  (
    'PROMO_1', 'HBDUSER52023', 'birthday', 
    10000, 1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 
    CURRENT_TIMESTAMP, CURRENT_TIMESTAMP
  );

INSERT INTO users_promotions (
  id, user_id, promotion_id, promotion_use_count, 
  created_at, updated_at
) 
VALUES 
  (
    'USER_PROMO_1', 'USER_5', 'PROMO_1', 
    0, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP
  );

-- migrate:down
TRUNCATE TABLE users;
TRUNCATE TABLE promotions;
TRUNCATE TABLE users_promotions;