-- +migrate Up
insert into types(id, name) values
    ('59943b9b-3f6d-4622-acff-07f455b67e7b', 'Платья'),
    ('fdf84c82-078e-496b-a061-15a743c2329a', 'Головные уборы'),
    ('816e7f75-2acf-451d-967b-5b7800a24403', 'Аксессуары'),
    ('0a45a9f1-992a-43c8-9ffd-5c5a00083297', 'Нижнее бельё');

insert into subtypes(type_id, name) values
    ('59943b9b-3f6d-4622-acff-07f455b67e7b', 'Вечерние'),
    ('59943b9b-3f6d-4622-acff-07f455b67e7b', 'Летние'),
    ('59943b9b-3f6d-4622-acff-07f455b67e7b', 'Деловые'),
    ('fdf84c82-078e-496b-a061-15a743c2329a', 'Кепки'),
    ('fdf84c82-078e-496b-a061-15a743c2329a', 'Панамы'),
    ('fdf84c82-078e-496b-a061-15a743c2329a', 'Зимние шапки'),
    ('fdf84c82-078e-496b-a061-15a743c2329a', 'Осенние шапки'),

    ('816e7f75-2acf-451d-967b-5b7800a24403', 'Украшения'),
    ('816e7f75-2acf-451d-967b-5b7800a24403', 'Часы'),
    ('816e7f75-2acf-451d-967b-5b7800a24403', 'Шарфы'),
    ('816e7f75-2acf-451d-967b-5b7800a24403', 'Перчатки'),
    ('816e7f75-2acf-451d-967b-5b7800a24403', 'Ремни'),
    ('816e7f75-2acf-451d-967b-5b7800a24403', 'Галстуки'),
    ('816e7f75-2acf-451d-967b-5b7800a24403', 'Сумки'),

    ('0a45a9f1-992a-43c8-9ffd-5c5a00083297', 'Майки'),
    ('0a45a9f1-992a-43c8-9ffd-5c5a00083297', 'Бюстгалтеры'),
    ('0a45a9f1-992a-43c8-9ffd-5c5a00083297', 'Трусы'),
    ('0a45a9f1-992a-43c8-9ffd-5c5a00083297', 'Носки'),
    ('0a45a9f1-992a-43c8-9ffd-5c5a00083297', 'Колготки'),
    ('0a45a9f1-992a-43c8-9ffd-5c5a00083297', 'Кальсоны'),
    ('0a45a9f1-992a-43c8-9ffd-5c5a00083297', 'Термобельё'),

    ('a4981358-9e59-45db-8ff4-ea8c11dfee66', 'Футболки'),
    ('a4981358-9e59-45db-8ff4-ea8c11dfee66', 'Поло'),
    ('a4981358-9e59-45db-8ff4-ea8c11dfee66', 'Майки'),
    ('a4981358-9e59-45db-8ff4-ea8c11dfee66', 'Лонгслив'),
    ('a4981358-9e59-45db-8ff4-ea8c11dfee66', 'Свитеры'),
    ('a4981358-9e59-45db-8ff4-ea8c11dfee66', 'Худи'),
    ('a4981358-9e59-45db-8ff4-ea8c11dfee66', 'Свитшоты'),
    ('a4981358-9e59-45db-8ff4-ea8c11dfee66', 'Топы'),

    ('99bfce00-b014-4458-9e26-a4675f72e352', 'Джинсы'),
    ('99bfce00-b014-4458-9e26-a4675f72e352', 'Брюки'),
    ('99bfce00-b014-4458-9e26-a4675f72e352', 'Шорты'),
    ('99bfce00-b014-4458-9e26-a4675f72e352', 'Юбки'),
    ('99bfce00-b014-4458-9e26-a4675f72e352', 'Леггинсы'),
    ('99bfce00-b014-4458-9e26-a4675f72e352', 'Бриджи'),

    ('b2e705a6-5e35-4957-93d4-fd801beac077', 'Туфли'),
    ('b2e705a6-5e35-4957-93d4-fd801beac077', 'Сапоги'),
    ('b2e705a6-5e35-4957-93d4-fd801beac077', 'Кеды'),
    ('b2e705a6-5e35-4957-93d4-fd801beac077', 'Босоножки'),
    ('b2e705a6-5e35-4957-93d4-fd801beac077', 'Ботинки'),
    ('b2e705a6-5e35-4957-93d4-fd801beac077', 'Полуботинки');
