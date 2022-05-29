CREATE TABLE IF NOT EXISTS roles (
  id uuid NOT NULL,
  code character(4) NOT NULL,
  name character varying NOT NULL,
  created_at timestamp without time zone NOT NULL,
  updated_at timestamp without time zone NOT NULL,
  PRIMARY KEY (id)
);

INSERT INTO roles (id, code, name, created_at, updated_at) 
VALUES 
(uuid_generate_v4(), 'ADM', 'Administrator', now(), now()),
(uuid_generate_v4(), 'USR', 'Users', now(), now());