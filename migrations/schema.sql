--
-- PostgreSQL database dump
--

-- Dumped from database version 12.9 (Ubuntu 12.9-0ubuntu0.20.04.1)
-- Dumped by pg_dump version 12.9 (Ubuntu 12.9-0ubuntu0.20.04.1)

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
-- Name: authority_permissions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.authority_permissions (
    id bigint NOT NULL,
    name text
);


ALTER TABLE public.authority_permissions OWNER TO postgres;

--
-- Name: authority_permissions_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.authority_permissions_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.authority_permissions_id_seq OWNER TO postgres;

--
-- Name: authority_permissions_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.authority_permissions_id_seq OWNED BY public.authority_permissions.id;


--
-- Name: authority_role_permissions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.authority_role_permissions (
    id bigint NOT NULL,
    role_id bigint,
    permission_id bigint
);


ALTER TABLE public.authority_role_permissions OWNER TO postgres;

--
-- Name: authority_role_permissions_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.authority_role_permissions_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.authority_role_permissions_id_seq OWNER TO postgres;

--
-- Name: authority_role_permissions_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.authority_role_permissions_id_seq OWNED BY public.authority_role_permissions.id;


--
-- Name: authority_roles; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.authority_roles (
    id bigint NOT NULL,
    name text
);


ALTER TABLE public.authority_roles OWNER TO postgres;

--
-- Name: authority_roles_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.authority_roles_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.authority_roles_id_seq OWNER TO postgres;

--
-- Name: authority_roles_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.authority_roles_id_seq OWNED BY public.authority_roles.id;


--
-- Name: authority_user_roles; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.authority_user_roles (
    id bigint NOT NULL,
    user_id bigint,
    role_id bigint
);


ALTER TABLE public.authority_user_roles OWNER TO postgres;

--
-- Name: authority_user_roles_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.authority_user_roles_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.authority_user_roles_id_seq OWNER TO postgres;

--
-- Name: authority_user_roles_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.authority_user_roles_id_seq OWNED BY public.authority_user_roles.id;


--
-- Name: contenttype; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.contenttype (
    id integer NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    name character varying(50) NOT NULL
);


ALTER TABLE public.contenttype OWNER TO postgres;

--
-- Name: contenttype_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.contenttype_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.contenttype_id_seq OWNER TO postgres;

--
-- Name: contenttype_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.contenttype_id_seq OWNED BY public.contenttype.id;


--
-- Name: schema_migration; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.schema_migration (
    version character varying(14) NOT NULL
);


ALTER TABLE public.schema_migration OWNER TO postgres;

--
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id integer NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    email text NOT NULL,
    password text NOT NULL,
    otp text,
    email_verified boolean DEFAULT false
);


ALTER TABLE public.users OWNER TO postgres;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.users_id_seq
    AS integer
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
-- Name: x; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.x (
    id integer NOT NULL,
    title character varying(100) NOT NULL
);


ALTER TABLE public.x OWNER TO postgres;

--
-- Name: x_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.x_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.x_id_seq OWNER TO postgres;

--
-- Name: x_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.x_id_seq OWNED BY public.x.id;


--
-- Name: xx; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.xx (
    id integer NOT NULL,
    title character varying(100) NOT NULL,
    description character varying(100) NOT NULL
);


ALTER TABLE public.xx OWNER TO postgres;

--
-- Name: xx_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.xx_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.xx_id_seq OWNER TO postgres;

--
-- Name: xx_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.xx_id_seq OWNED BY public.xx.id;


--
-- Name: xxx; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.xxx (
    title character varying(100) NOT NULL,
    description character varying(100) NOT NULL,
    id integer NOT NULL
);


ALTER TABLE public.xxx OWNER TO postgres;

--
-- Name: xxx_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.xxx_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.xxx_id_seq OWNER TO postgres;

--
-- Name: xxx_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.xxx_id_seq OWNED BY public.xxx.id;


--
-- Name: authority_permissions id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.authority_permissions ALTER COLUMN id SET DEFAULT nextval('public.authority_permissions_id_seq'::regclass);


--
-- Name: authority_role_permissions id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.authority_role_permissions ALTER COLUMN id SET DEFAULT nextval('public.authority_role_permissions_id_seq'::regclass);


--
-- Name: authority_roles id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.authority_roles ALTER COLUMN id SET DEFAULT nextval('public.authority_roles_id_seq'::regclass);


--
-- Name: authority_user_roles id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.authority_user_roles ALTER COLUMN id SET DEFAULT nextval('public.authority_user_roles_id_seq'::regclass);


--
-- Name: contenttype id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.contenttype ALTER COLUMN id SET DEFAULT nextval('public.contenttype_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- Name: x id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.x ALTER COLUMN id SET DEFAULT nextval('public.x_id_seq'::regclass);


--
-- Name: xx id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.xx ALTER COLUMN id SET DEFAULT nextval('public.xx_id_seq'::regclass);


--
-- Name: xxx id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.xxx ALTER COLUMN id SET DEFAULT nextval('public.xxx_id_seq'::regclass);


--
-- Name: authority_permissions authority_permissions_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.authority_permissions
    ADD CONSTRAINT authority_permissions_pkey PRIMARY KEY (id);


--
-- Name: authority_role_permissions authority_role_permissions_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.authority_role_permissions
    ADD CONSTRAINT authority_role_permissions_pkey PRIMARY KEY (id);


--
-- Name: authority_roles authority_roles_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.authority_roles
    ADD CONSTRAINT authority_roles_pkey PRIMARY KEY (id);


--
-- Name: authority_user_roles authority_user_roles_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.authority_user_roles
    ADD CONSTRAINT authority_user_roles_pkey PRIMARY KEY (id);


--
-- Name: contenttype contenttype_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.contenttype
    ADD CONSTRAINT contenttype_pkey PRIMARY KEY (id);


--
-- Name: contenttype unique_name; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.contenttype
    ADD CONSTRAINT unique_name UNIQUE (name);


--
-- Name: users users_email_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_email_key UNIQUE (email);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: x x_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.x
    ADD CONSTRAINT x_pkey PRIMARY KEY (id);


--
-- Name: xx xx_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.xx
    ADD CONSTRAINT xx_pkey PRIMARY KEY (id);


--
-- Name: xxx xxx_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.xxx
    ADD CONSTRAINT xxx_pkey PRIMARY KEY (id);


--
-- Name: schema_migration_version_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX schema_migration_version_idx ON public.schema_migration USING btree (version);


--
-- PostgreSQL database dump complete
--

