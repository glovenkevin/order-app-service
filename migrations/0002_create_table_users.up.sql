CREATE TABLE IF NOT EXISTS users (
  id uuid NOT NULL,
  role_id uuid NOT NULL,
  name character varying NOT NULL,
  email character varying NOT NULL unique,
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
('a30ff022-3cb6-482d-a8b5-285c09492b2d', 'b9e5009a-326b-4d18-a9c1-825dc3234336', 'Administrator', 'admin@admin.com', '', '08561234', '', false, now(), now());