
CREATE TABLE public.publishing (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) NOT NULL,
    deleted boolean DEFAULT false
);

CREATE TABLE public.book (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) NOT NULL,
    year smallint NOT NULL,
    publishing UUID NOT NULL,
    deleted boolean DEFAULT false,
    CONSTRAINT publishing_fk FOREIGN KEY (publishing) REFERENCES public.publishing(id)
);

CREATE TABLE public.author (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    firstname VARCHAR(100) NOT NULL,
    lastname VARCHAR(100) NOT NULL,
    photo VARCHAR(50) NOT NULL,
    birth_place VARCHAR(100) NOT NULL,
    age smallint NOT NULL,
    date_of_birth VARCHAR(30) NOT NULL, 
    date_of_death VARCHAR(30),
    gender VARCHAR(15) NOT NULL,
    deleted boolean DEFAULT false
);

CREATE TABLE public.book_authors (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	book_id UUID NOT NULL,
	author_id UUID NOT NULL,
	CONSTRAINT book_fk FOREIGN KEY (book_id) REFERENCES public.book(id), 
	CONSTRAINT author_fk FOREIGN KEY (author_id) REFERENCES public.author(id), 
	CONSTRAINT book_author_unique UNIQUE (book_id, author_id)
);

CREATE TABLE public.genre (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(50) NOT NULL,
    deleted boolean DEFAULT false
);

CREATE TABLE public.book_genres (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    book_id UUID NOT NULL,
    genre_id UUID NOT NULL,
    CONSTRAINT book_fk FOREIGN KEY (book_id) REFERENCES public.book(id),
    CONSTRAINT genre_fk FOREIGN KEY (genre_id) REFERENCES public.genre(id),
    CONSTRAINT book_genre_unique UNIQUE (book_id, genre_id)
);

CREATE TABLE public.award (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100),
    deleted boolean DEFAULT false
)

CREATE TABLE public.book_awards (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    book_id UUID NOT NULL,
    award_id UUID NOT NULL,
    year smallint NOT NULL,
    CONSTRAINT book_fk FOREIGN KEY (book_id) REFERENCES public.book(id),
    CONSTRAINT award_fk FOREIGN KEY (award_id) REFERENCES public.award(id),
    CONSTRAINT book_genre_unique UNIQUE (book_id, award_id)
)


для массива авторов

SELECT a.id, a.firstname, a.lastname, a.photo, a.birth_place, a.age, a.date_of_birth, a.date_of_death, a.gender FROM public.book b
JOIN public.book_authors ba ON ba.book_id = '8b506f21-3fa1-459f-961a-9a6af8ffccf5'
JOIN public.author a ON a.id = ba.author_id 
WHERE b.id = '8b506f21-3fa1-459f-961a-9a6af8ffccf5';

для списка жанров книги

SELECT g.name FROM public.book_genres bg
JOIN public.genre g ON g.id = bg.genre_id
WHERE book_id = '8bf192c1-d396-4b39-bbfb-77872dcba7fa';

для получения списка награду у книги

SELECT a.id, a.name, ba.year FROM public.book_awards ba
JOIN public.award a ON a.id = ba.award_id
WHERE ba.book_id = $1 AND a.deleted = false;

для данных книги к списку авторов

SELECT b.id, b.name, b.year, p.id, p.name  FROM public.book b
JOIN public.publishing p ON p.id = b.publishing
WHERE b.id = '8b506f21-3fa1-459f-961a-9a6af8ffccf5' AND b.deleted = false;


INSERT

INSERT INTO public.publishing (id, name) VALUES 
('8b506f21-3fa1-459f-961a-9a6af8ffccf5', 'PUB1');

INSERT INTO public.book (id, name, year, publishing) VALUES
('8bf192c1-d396-4b39-bbfb-77872dcba7fa', 'Book1', 2001, '8b506f21-3fa1-459f-961a-9a6af8ffccf5'),
('b5790365-4659-40ec-adad-5eca43c80db9', 'Book2', 2002, '8b506f21-3fa1-459f-961a-9a6af8ffccf5'),
('dac83be8-eede-4661-9f1d-e78113b1688d', 'Book3', 2003, '8b506f21-3fa1-459f-961a-9a6af8ffccf5');

INSERT INTO public.author (id, firstname, lastname, photo, birth_place, age, date_of_birth, date_of_death, gender) VALUES
('0b0c5a9a-9a06-4019-a395-943e0935ee48', 'author1', '111', 'photo1', 'bornplace1', 20, 'date111', 'date111', 'male'),
('29c9c8b3-1ceb-433d-92d7-a5a31b0ef278', 'author2', '222', 'photo2', 'bornplace2', 20, 'date222', 'date222', 'male'),
('ba341666-6cae-4e0d-bc84-1d01e43c1095', 'author3', '333', 'photo3', 'bornplace3', 20, 'date333', 'date333', 'female');

INSERT INTO book_authors (id, book_id, author_id) VALUES
('1586acbd-dbd3-4b0e-8756-4708507fdc9c', '8bf192c1-d396-4b39-bbfb-77872dcba7fa', '0b0c5a9a-9a06-4019-a395-943e0935ee48'),
('37a754d7-87b8-4463-bddf-0976234ed3b9', '8bf192c1-d396-4b39-bbfb-77872dcba7fa', 'ba341666-6cae-4e0d-bc84-1d01e43c1095'),
('7c8baa4e-4788-4fb1-881e-35fcc674084a', 'b5790365-4659-40ec-adad-5eca43c80db9', '29c9c8b3-1ceb-433d-92d7-a5a31b0ef278'),
('837cfb44-7dbb-4470-88aa-fbb61b14426c', 'dac83be8-eede-4661-9f1d-e78113b1688d', '0b0c5a9a-9a06-4019-a395-943e0935ee48');

INSERT INTO public.genre (name) VALUES
('6db0c689-4669-44f5-8753-25473276e9d3', 'Роман'),
('ab6f948e-439f-4d58-9620-ca52abe6b039', 'Поэма'),
('947e5040-1ea8-4cca-8f82-53a2fa6f01e4', 'Фантистика');

INSERT INTO public.book_genres (book_id, genre_id) VALUES
('8bf192c1-d396-4b39-bbfb-77872dcba7fa', '6db0c689-4669-44f5-8753-25473276e9d3'),
('8bf192c1-d396-4b39-bbfb-77872dcba7fa', '947e5040-1ea8-4cca-8f82-53a2fa6f01e4');
