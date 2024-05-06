CREATE DATABASE IF NOT EXISTS dev;
USE dev;

create user if not exists 'dev_user'@'localhost' IDENTIFIED by '1234';
create user if not exists 'dev_user'@'%' IDENTIFIED by '1234';
grant all PRIVILEGES ON dev.* TO 'dev_user'@'localhost';
grant all PRIVILEGES ON dev.* TO 'dev_user'@'%';

create table IF NOT EXISTS users
(
  id         varchar(191) not null
    primary key,
  username   varchar(191) not null,
  password   varchar(191) not null,
  created_at datetime(3)  not null,
  updated_at datetime(3)  null
);

create table IF NOT EXISTS accounts
(
  id         varchar(191)                not null
    primary key,
  user_id    varchar(191)                not null,
  balance    bigint unsigned default '0' not null,
  created_at datetime(3)                 not null,
  updated_at datetime(3)                 null,
  constraint fk_users_account
    foreign key (user_id) references users (id)
);

create table IF NOT EXISTS login_logs
(
  id         varchar(191) not null
    primary key,
  user_id    varchar(191) not null,
  login_type varchar(191) not null,
  created_at datetime(3)  not null,
  constraint fk_users_login_logs
    foreign key (user_id) references users (id)
);

create table IF NOT EXISTS portfolios
(
  id         varchar(191) not null
    primary key,
  user_id    varchar(191) not null,
  risk_type  varchar(191) not null,
  created_at datetime(3)  not null,
  updated_at datetime(3)  null,
  constraint fk_users_portfolios
    foreign key (user_id) references users (id)
);

create table IF NOT EXISTS portfolio_items
(
  id           varchar(191)    not null
    primary key,
  portfolio_id varchar(191)    not null,
  stock_code   varchar(191)    not null,
  quantity     bigint unsigned not null,
  created_at   datetime(3)     not null,
  updated_at   datetime(3)     null,
  constraint fk_portfolios_portfolio_items
    foreign key (portfolio_id) references portfolios (id),
  constraint portfolio_items_stocks_id_fk
    foreign key (stock_code) references stocks (id)
);

create table IF NOT EXISTS stocks
(
  id         varchar(191)    not null
    primary key,
  stock_name varchar(191)    not null,
  price      bigint unsigned not null,
  created_at datetime(3)     not null,
  updated_at datetime(3)     null
);

create table IF NOT EXISTS stock_transactions
(
  id         varchar(191)    not null
    primary key,
  user_id    varchar(191)    not null,
  stock_id   varchar(191)    not null,
  quantity   bigint unsigned not null,
  price      bigint unsigned not null,
  created_at datetime(3)     not null,
  constraint stock_transactions_stocks_id_fk
    foreign key (stock_id) references stocks (id)
      on update cascade,
  constraint stock_transactions_users_id_fk
    foreign key (user_id) references users (id)
      on update cascade
);

create table IF NOT EXISTS transactions
(
  id            varchar(191)    not null
    primary key,
  account_id    varchar(191)    not null,
  amount        bigint unsigned not null,
  type          varchar(191)    not null,
  after_balance bigint unsigned not null,
  created_at    datetime(3)     not null,
  constraint fk_accounts_transactions
    foreign key (account_id) references accounts (id)
);


-- 현재 날짜 및 시간 정보 추출
SET @now = NOW();

-- 테이블에 10개의 임의 주식 데이터 삽입
INSERT INTO stocks
VALUES
  (UUID(), '삼성전자', 720000, @now, @now),
  (UUID(), '현대자동차', 682000, @now, @now),
  (UUID(), '카카오', 938000, @now, @now),
  (UUID(), '네이버', 715000, @now, @now),
  (UUID(), 'LG 전자', 153000, @now, @now),
  (UUID(), 'SK 하이닉스', 126000, @now, @now),
  (UUID(), '아모레퍼시픽', 304000, @now, @now),
  (UUID(), '카카오뱅크', 464000, @now, @now),
  (UUID(), '쿠팡', 353000, @now, @now),
  (UUID(), 'Amazon Inc.', 252000, @now, @now);











