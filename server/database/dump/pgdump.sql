--
-- PostgreSQL database dump
--

-- Dumped from database version 9.6.4
-- Dumped by pg_dump version 9.6.4

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


SET search_path = public, pg_catalog;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: user_results; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE user_results (
    id bigint NOT NULL,
    user_id integer NOT NULL,
    result integer,
    datetime timestamp with time zone
);


ALTER TABLE user_results OWNER TO admin;

--
-- Name: TABLE user_results; Type: COMMENT; Schema: public; Owner: admin
--

COMMENT ON TABLE user_results IS 'Results of the measurements';


--
-- Name: user_results_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE user_results_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE user_results_id_seq OWNER TO admin;

--
-- Name: user_results_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE user_results_id_seq OWNED BY user_results.id;


--
-- Name: user_sessions; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE user_sessions (
    session_id bigint NOT NULL,
    user_id integer NOT NULL,
    expiry_time timestamp with time zone,
    starting_time timestamp with time zone DEFAULT now(),
    auth_token text NOT NULL
);


ALTER TABLE user_sessions OWNER TO admin;

--
-- Name: TABLE user_sessions; Type: COMMENT; Schema: public; Owner: admin
--

COMMENT ON TABLE user_sessions IS 'Information about user authentications';


--
-- Name: user_sessions_session_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE user_sessions_session_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE user_sessions_session_id_seq OWNER TO admin;

--
-- Name: user_sessions_session_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE user_sessions_session_id_seq OWNED BY user_sessions.session_id;


--
-- Name: users; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE users (
    id integer NOT NULL,
    username character varying(64) NOT NULL,
    password text NOT NULL,
    registration_date timestamp with time zone NOT NULL
);


ALTER TABLE users OWNER TO admin;

--
-- Name: TABLE users; Type: COMMENT; Schema: public; Owner: admin
--

COMMENT ON TABLE users IS 'List of registered users';


--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE users_id_seq OWNER TO admin;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE users_id_seq OWNED BY users.id;


--
-- Name: user_results id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY user_results ALTER COLUMN id SET DEFAULT nextval('user_results_id_seq'::regclass);


--
-- Name: user_sessions session_id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY user_sessions ALTER COLUMN session_id SET DEFAULT nextval('user_sessions_session_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY users ALTER COLUMN id SET DEFAULT nextval('users_id_seq'::regclass);


--
-- Data for Name: user_results; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY user_results (id, user_id, result, datetime) FROM stdin;
\.


--
-- Name: user_results_id_seq; Type: SEQUENCE SET; Schema: public; Owner: admin
--

SELECT pg_catalog.setval('user_results_id_seq', 1, false);


--
-- Data for Name: user_sessions; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY user_sessions (session_id, user_id, expiry_time, starting_time, auth_token) FROM stdin;
\.


--
-- Name: user_sessions_session_id_seq; Type: SEQUENCE SET; Schema: public; Owner: admin
--

SELECT pg_catalog.setval('user_sessions_session_id_seq', 1, false);


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY users (id, username, password, registration_date) FROM stdin;
\.


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: admin
--

SELECT pg_catalog.setval('users_id_seq', 1, false);


--
-- Name: user_results user_results_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY user_results
    ADD CONSTRAINT user_results_pkey PRIMARY KEY (id);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: user_sessions_auth_token_uindex; Type: INDEX; Schema: public; Owner: admin
--

CREATE UNIQUE INDEX user_sessions_auth_token_uindex ON user_sessions USING btree (auth_token);


--
-- Name: user_results user_results_users_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY user_results
    ADD CONSTRAINT user_results_users_id_fk FOREIGN KEY (user_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: user_sessions user_sessions_users_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY user_sessions
    ADD CONSTRAINT user_sessions_users_id_fk FOREIGN KEY (user_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

