CREATE TABLE IF NOT EXISTS users (
  id uuid NOT NULL,
  role_id uuid NOT NULL,
  name character varying NOT NULL,
  email character varying NOT NULL,
  password character varying NOT NULL,
  phone_number character varying,
  fcm_token character varying,
  is_blocked boolean NOT NULL DEFAULT false,
  created_at timestamp without time zone NOT NULL,
  updated_at timestamp without time zone NOT NULL,
  PRIMARY KEY (id),
  CONSTRAINT fk_roles FOREIGN KEY (role_id) REFERENCES roles(id) ON DELETE SET NULL
);

INSERT INTO users (id, role_id, name, email, password, phone_number, fcm_token, is_blocked, created_at, updated_at)
VALUES 
(uuid_generate_v4(), (SELECT id FROM roles WHERE code = 'ADM'), 'Administrator', 'admin@admin.com', '', '08561234', '', false, now(), now());