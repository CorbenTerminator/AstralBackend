SET NAMES utf8;
create schema if not exists shop DEFAULT CHARACTER SET utf8;

use shop;


create table if not exists products (
    product_id  serial primary key,
    name    varchar(200),
    price   decimal(10,2),
    category_id int
) DEFAULT CHARSET=utf8;

create table if not exists categories (
    category_id  int primary key,
    name    varchar(150)
) DEFAULT CHARSET=utf8;


create table if not exists orders (
    order_id    serial primary key,
    user_id bigint unsigned,
    createdAt   datetime,
    status_id   int unsigned
) DEFAULT CHARSET=utf8;

create table if not exists status (
    status_id   int UNSIGNED NOT NULL AUTO_INCREMENT UNIQUE primary key,
    name    varchar(150)
) DEFAULT CHARSET=utf8;

create table if not exists order_products(
    order_id    bigint unsigned,
    product_id bigint unsigned
) DEFAULT CHARSET=utf8;

create table if not exists users(
    user_id serial primary key,
    login   varchar(150),
    password    varchar(150)
) DEFAULT CHARSET=utf8;

alter table products 
    add constraint category_fk
    foreign key (category_id)
    references categories(category_id) 
    on delete cascade;

alter table orders
    add constraint customer_fk
    foreign key (user_id)
    references users(user_id) 
    on delete cascade,

    add constraint status_fk
    foreign key (status_id)
    references status(status_id)
    on delete cascade;

alter table order_products
    add constraint order_fk
    foreign key (order_id)
    references orders(order_id)
    on delete cascade,

    add constraint product_fk
    foreign key (product_id) 
    references products(product_id)
    on delete cascade;

insert into users(user_id, login, password) values
(1, 'admin', '$2a$14$rk6syVXtGQPGOz4cg6kFq.tHkqy4z9ZIyutCCVfiEMFyPT6o0AR4a'),
(2, 'user', '$2a$14$QnP0Nd8lR4PrR0Lvc3gPxuTbkNSXlk42rhhVoP33HO.X/JeZdtCwm'),
(3, 'user1', '$2a$14$DfbfTOHDxszUnZnzn2Z5BeInHgoa2.cYoBOEggwzU0S154iAFgErG');

insert into categories(category_id, name) values 
(1, 'Электроника'),
(2, 'Бытовая химия'),
(3, 'Бакалея');


insert into products(product_id, name, price, category_id) values
(1, 'Компьютер', 23654.42, 1),
(2, 'Видеокарта', 100000.99, 1),
(3, 'Перчатки', 23.56, 2),
(4, 'Мыло', 57.78, 2),
(5, 'Хлебцы', 87.56, 3),
(6, 'Листья для лазаньи', 125.00, 3);

insert into status(name) values
('создан'), ('оплачен'), ('в работе'), ('доставлен');


insert into orders(user_id, createdAt, status_id) values 
(1, NOW(), 1),
(2, NOW(), 2),
(3, NOW(), 3);

insert into order_products(order_id, product_id) values 
(1, 1),
(1, 4),
(2, 4),
(2, 5),
(3, 6),
(3, 3),
(3, 2);