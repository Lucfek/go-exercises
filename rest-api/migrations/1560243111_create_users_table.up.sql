
--
-- Name: users Type: TABLE; Schema: public; Owner: testuser
--
CREATE TABLE public.users (
    id bigserial,
    email text  NOT NULL UNIQUE,
    password bytea,
    created_at timestamp without time zone DEFAULT timezone('utc'::text, now()) NOT NULL
);

ALTER TABLE public.users OWNER TO testuser;

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);

