Нужно описать модель библиотеки. Есть 3 сущности: "Автор", "Книга", "Читатель".
Физически книга только одна и может быть только у одного читателя.
Нужно составить таблицы для библиотеки так, чтобы это учесть.
Первый запрос — выбрать названия всех книг, которые на руках.
Второй запрос — выбрать названия всех книг в библиотеке, у которых больше 3 авторов.
Третий запрос — выбрать имена топ-3 читаемых авторов на данный момент.   в ответе таблицы и запросы выведи в sql


-- Нужно описать модель библиотеки. Есть 3 сущности: "Автор", "Книга", "Читатель".
-- Физически книга только одна и может быть только у одного читателя.
-- Нужно составить таблицы для библиотеки так, чтобы это учесть.
CREATE TABLE IF NOT EXISTS author (
    author_id SERIAL,
    name text NOT NULL,

    CONSTRAINT pk_author PRIMARY KEY (author_id)
);

CREATE TABLE IF NOT EXISTS book (
    book_id SERIAL,
    title TEXT NOT NULL,

    CONSTRAINT pk_book PRIMARY KEY (book_id)
);

CREATE TABLE IF NOT EXISTS reader (
    reader_id SERIAL,
    name TEXT NOT NULL,
    book_id INT,

    CONSTRAINT pk_reader PRIMARY KEY (reader_id),
    CONSTRAINT fk_reader_book FOREIGN KEY (book_id) REFERENCES book (book_id),
    CONSTRAINT uq_reader_book_id UNIQUE (book_id)
);

CREATE TABLE IF NOT EXISTS author_book (
    author_id INT,
    book_id INT,

    CONSTRAINT fk_author_book_author FOREIGN KEY (author_id) REFERENCES author (author_id),
    CONSTRAINT fk_author_book_book FOREIGN KEY (book_id) REFERENCES book (book_id),
    CONSTRAINT pk_author_book PRIMARY KEY (author_id, book_id)
);


-- Первый запрос — выбрать названия всех книг, которые на руках.
SELECT b.title
FROM book b
JOIN reader r
ON r.book_id = b.book_id
WHERE r.book_id IS NOT NULL;

-- Второй запрос — выбрать названия всех книг в библиотеке, у которых больше 3 авторов.
SELECT b.title, SUM(ab.author_id) AS author_num
FROM book b
JOIN author_book ab
ON ab.book_id = b.book_id
LEFT JOIN reader r
ON r.book_id = b.book_id
WHERE r.book_id IS NOT NULL
GROUP BY b.title
HAVING author_num > 3;

-- Третий запрос — выбрать имена топ-3 читаемых авторов на данный момент.
SELECT a.name, COUNT(ab.author_id) AS reader_num
FROM author a
JOIN author_book ab
ON ab.author_id = a.author_id
LEFT JOIN reader r
ON r.book_id = ab.book_id
WHERE r.book_id IS NOT NULL
GROUP BY a.name
ORDER BY reader_num DESC LIMIT 3;