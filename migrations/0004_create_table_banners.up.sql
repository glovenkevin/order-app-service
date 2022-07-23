create table if not EXISTS banners (
    id uuid not null,
    "name" character varying not null,
    "description" character varying, 
    image_url character varying not null,
    is_deleted boolean not null default false,
    is_show boolean not null default true,
    seq integer not null,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone,
    primary key (id)
);