--
-- PostgreSQL database dump
--

-- Dumped from database version 11.9 (Debian 11.9-1.pgdg90+1)
-- Dumped by pg_dump version 11.9 (Debian 11.9-1.pgdg90+1)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: movies; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.movies (
    id bigint NOT NULL,
    name character varying(250),
    published_at timestamp without time zone
);


ALTER TABLE public.movies OWNER TO postgres;

--
-- Name: movies_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.movies_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.movies_id_seq OWNER TO postgres;

--
-- Name: movies_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.movies_id_seq OWNED BY public.movies.id;


--
-- Name: user_interactions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.user_interactions (
    id bigint NOT NULL,
    user_id bigint,
    movie_id bigint,
    rating integer,
    comment character varying(250)
);


ALTER TABLE public.user_interactions OWNER TO postgres;

--
-- Name: user_interactions_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.user_interactions_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.user_interactions_id_seq OWNER TO postgres;

--
-- Name: user_interactions_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.user_interactions_id_seq OWNED BY public.user_interactions.id;


--
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id bigint NOT NULL,
    full_name character varying(250),
    email character varying(250),
    password character varying(250),
    user_type character varying(250)
);


ALTER TABLE public.users OWNER TO postgres;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.users_id_seq OWNER TO postgres;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- Name: movies id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.movies ALTER COLUMN id SET DEFAULT nextval('public.movies_id_seq'::regclass);


--
-- Name: user_interactions id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_interactions ALTER COLUMN id SET DEFAULT nextval('public.user_interactions_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- Data for Name: movies; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.movies (id, name, published_at) FROM stdin;
1	AndhaDhun	2016-06-22 19:10:25
2	Baaghi	2016-06-01 19:10:25
3	Dhadak	2016-06-10 19:10:25
4	Falooda	2016-06-15 19:10:25
5	Gold	2016-07-01 19:10:25
6	Parasite	2020-12-22 01:10:00
\.


--
-- Data for Name: user_interactions; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.user_interactions (id, user_id, movie_id, rating, comment) FROM stdin;
1	1	1	3	This movie is okay, would not recommend though.
4	1	3	2	Okay movie, would only watch once.
6	1	5	4	great story line, loved the character development!!
7	2	1	5	Greatest Movie!!
8	3	1	3	Liked this movie, a must watch.
3	1	2	1	Terrible movie, bad directions.
5	1	4	5	Greatest Movie Ever!!
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.users (id, full_name, email, password, user_type) FROM stdin;
1	Amogh Yermalkar	ayermalkar99@gmail.com	amogh@123	user
2	Test User 2	testuser2@gmail.com	test@user2	user
3	Test User 3	testuser3@gmail.com	test@user3	user
4	Admin User 1	adminuser@gmail.com	admin@123	admin
\.


--
-- Name: movies_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.movies_id_seq', 30, true);


--
-- Name: user_interactions_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.user_interactions_id_seq', 8, true);


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.users_id_seq', 4, true);


--
-- Name: movies movies_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.movies
    ADD CONSTRAINT movies_pkey PRIMARY KEY (id);


--
-- Name: user_interactions user_interactions_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_interactions
    ADD CONSTRAINT user_interactions_pkey PRIMARY KEY (id);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: user_interactions fk_movie_id; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_interactions
    ADD CONSTRAINT fk_movie_id FOREIGN KEY (movie_id) REFERENCES public.movies(id);


--
-- Name: user_interactions fk_user_id; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_interactions
    ADD CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- PostgreSQL database dump complete
--

