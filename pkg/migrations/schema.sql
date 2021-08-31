--
-- PostgreSQL database dump
--

-- Dumped from database version 13.3 (Ubuntu 13.3-1.pgdg20.04+1)
-- Dumped by pg_dump version 13.3 (Ubuntu 13.3-1.pgdg20.04+1)

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

SET default_table_access_method = heap;

--
-- Name: contents; Type: TABLE; Schema: public; Owner: rahma
--

CREATE TABLE public.contents (
    id integer NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    title character varying(100) NOT NULL,
    details text NOT NULL
);


ALTER TABLE public.contents OWNER TO rahma;

--
-- Name: contents_id_seq; Type: SEQUENCE; Schema: public; Owner: rahma
--

CREATE SEQUENCE public.contents_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.contents_id_seq OWNER TO rahma;

--
-- Name: contents_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: rahma
--

ALTER SEQUENCE public.contents_id_seq OWNED BY public.contents.id;


--
-- Name: contenttype; Type: TABLE; Schema: public; Owner: rahma
--

CREATE TABLE public.contenttype (
    id integer NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    name character varying(50) NOT NULL
);


ALTER TABLE public.contenttype OWNER TO rahma;

--
-- Name: contenttype_id_seq; Type: SEQUENCE; Schema: public; Owner: rahma
--

CREATE SEQUENCE public.contenttype_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.contenttype_id_seq OWNER TO rahma;

--
-- Name: contenttype_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: rahma
--

ALTER SEQUENCE public.contenttype_id_seq OWNED BY public.contenttype.id;


--
-- Name: schema_migration; Type: TABLE; Schema: public; Owner: rahma
--

CREATE TABLE public.schema_migration (
    version character varying(14) NOT NULL
);


ALTER TABLE public.schema_migration OWNER TO rahma;

--
-- Name: users; Type: TABLE; Schema: public; Owner: rahma
--

CREATE TABLE public.users (
    id integer NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    email text NOT NULL,
    password text NOT NULL
);


ALTER TABLE public.users OWNER TO rahma;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: rahma
--

CREATE SEQUENCE public.users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.users_id_seq OWNER TO rahma;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: rahma
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- Name: contents id; Type: DEFAULT; Schema: public; Owner: rahma
--

ALTER TABLE ONLY public.contents ALTER COLUMN id SET DEFAULT nextval('public.contents_id_seq'::regclass);


--
-- Name: contenttype id; Type: DEFAULT; Schema: public; Owner: rahma
--

ALTER TABLE ONLY public.contenttype ALTER COLUMN id SET DEFAULT nextval('public.contenttype_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: rahma
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- Name: contents contents_pkey; Type: CONSTRAINT; Schema: public; Owner: rahma
--

ALTER TABLE ONLY public.contents
    ADD CONSTRAINT contents_pkey PRIMARY KEY (id);


--
-- Name: contenttype contenttype_pkey; Type: CONSTRAINT; Schema: public; Owner: rahma
--

ALTER TABLE ONLY public.contenttype
    ADD CONSTRAINT contenttype_pkey PRIMARY KEY (id);


--
-- Name: users users_email_key; Type: CONSTRAINT; Schema: public; Owner: rahma
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_email_key UNIQUE (email);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: rahma
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: schema_migration_version_idx; Type: INDEX; Schema: public; Owner: rahma
--

CREATE UNIQUE INDEX schema_migration_version_idx ON public.schema_migration USING btree (version);


--
-- PostgreSQL database dump complete
--

