-- 用户表
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(32) UNIQUE NOT NULL,
    phone VARCHAR(16) UNIQUE,
    email VARCHAR(64) UNIQUE,
    password_hash VARCHAR(128) NOT NULL,
    avatar_url VARCHAR(256),
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now()
);

-- 好友关系
CREATE TABLE friendships (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    friend_id INT NOT NULL,
    status INT NOT NULL,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now(),
    UNIQUE(user_id, friend_id)
);

-- 聊天消息
CREATE TABLE messages (
    id BIGSERIAL PRIMARY KEY,
    session_id VARCHAR(64) NOT NULL,
    sender_id INT NOT NULL,
    receiver_id INT NOT NULL,
    msg_type INT NOT NULL,
    content TEXT,
    status INT NOT NULL,
    created_at TIMESTAMP DEFAULT now()
);