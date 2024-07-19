--
-- PostgreSQL database dump
--

-- Dumped from database version 16.3
-- Dumped by pg_dump version 16.3

-- Started on 2024-07-19 10:30:59

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
-- TOC entry 216 (class 1259 OID 16407)
-- Name: encryption_history; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.encryption_history (
    id integer NOT NULL,
    key text NOT NULL,
    value text NOT NULL
);


ALTER TABLE public.encryption_history OWNER TO postgres;

--
-- TOC entry 215 (class 1259 OID 16406)
-- Name: encryprion_history_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.encryprion_history_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.encryprion_history_id_seq OWNER TO postgres;

--
-- TOC entry 4839 (class 0 OID 0)
-- Dependencies: 215
-- Name: encryprion_history_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.encryprion_history_id_seq OWNED BY public.encryption_history.id;


--
-- TOC entry 4688 (class 2604 OID 16410)
-- Name: encryption_history id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.encryption_history ALTER COLUMN id SET DEFAULT nextval('public.encryprion_history_id_seq'::regclass);


--
-- TOC entry 4690 (class 2606 OID 16414)
-- Name: encryption_history encryprion_history_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.encryption_history
    ADD CONSTRAINT encryprion_history_pkey PRIMARY KEY (id);


-- Completed on 2024-07-19 10:30:59

--
-- PostgreSQL database dump complete
--

