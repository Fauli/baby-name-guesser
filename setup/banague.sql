--
-- PostgreSQL database dump
--


--
-- TOC entry 3647 (class 1262 OID 16420)
-- Name: banague; Type: DATABASE; Schema: -; Owner: banague
--

CREATE DATABASE banague WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'C';


ALTER DATABASE banague OWNER TO banague;

\connect banague

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

--
-- TOC entry 2 (class 3079 OID 16442)
-- Name: pgcrypto; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS pgcrypto WITH SCHEMA public;


--
-- TOC entry 3648 (class 0 OID 0)
-- Dependencies: 2
-- Name: EXTENSION pgcrypto; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION pgcrypto IS 'cryptographic functions';


SET default_tablespace = 'public';

--
-- TOC entry 216 (class 1259 OID 16428)
-- Name: names; Type: TABLE; Schema: public; Owner: banague
--

CREATE TABLE public.names (
    name character varying NOT NULL
);


ALTER TABLE public.names OWNER TO banague;

--
-- TOC entry 217 (class 1259 OID 16435)
-- Name: voters; Type: TABLE; Schema: public; Owner: banague
--

CREATE TABLE public.voters (
    name character varying NOT NULL,
    last_name character varying NOT NULL,
    email character varying NOT NULL,
    password character varying NOT NULL
);


ALTER TABLE public.voters OWNER TO banague;

--
-- TOC entry 218 (class 1259 OID 16485)
-- Name: votes; Type: TABLE; Schema: public; Owner: banague
--

CREATE TABLE public.votes (
    names_fk character varying NOT NULL,
    voter_fk character varying NOT NULL,
    is_paid boolean DEFAULT false NOT NULL
);


ALTER TABLE public.votes OWNER TO banague;


--
-- TOC entry 3489 (class 2606 OID 16434)
-- Name: names names_pkey; Type: CONSTRAINT; Schema: public; Owner: banague
--

ALTER TABLE ONLY public.names
    ADD CONSTRAINT names_pkey PRIMARY KEY (name);


--
-- TOC entry 3493 (class 2606 OID 16492)
-- Name: votes unique_vote; Type: CONSTRAINT; Schema: public; Owner: banague
--

ALTER TABLE ONLY public.votes
    ADD CONSTRAINT unique_vote PRIMARY KEY (names_fk, voter_fk);


--
-- TOC entry 3491 (class 2606 OID 16441)
-- Name: voters voters_pkey; Type: CONSTRAINT; Schema: public; Owner: banague
--

ALTER TABLE ONLY public.voters
    ADD CONSTRAINT voters_pkey PRIMARY KEY (email);


--
-- TOC entry 3494 (class 2606 OID 16493)
-- Name: votes votes_names_fk_fkey; Type: FK CONSTRAINT; Schema: public; Owner: banague
--

ALTER TABLE ONLY public.votes
    ADD CONSTRAINT votes_names_fk_fkey FOREIGN KEY (names_fk) REFERENCES public.names(name);


--
-- TOC entry 3495 (class 2606 OID 16498)
-- Name: votes votes_voter_fk_fkey; Type: FK CONSTRAINT; Schema: public; Owner: banague
--

ALTER TABLE ONLY public.votes
    ADD CONSTRAINT votes_voter_fk_fkey FOREIGN KEY (voter_fk) REFERENCES public.voters(email);


-- Completed on 2024-06-12 22:02:05 CEST

--
-- PostgreSQL database dump complete
--

