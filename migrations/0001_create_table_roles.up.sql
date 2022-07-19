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
('b9e5009a-326b-4d18-a9c1-825dc3234336', 'ADM', 'Administrator', now(), now()),
('3e93750c-6040-4a19-b382-9cc48215522f', 'USR', 'Users', now(), now());