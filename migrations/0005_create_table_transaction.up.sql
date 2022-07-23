create table if not exists order_status (
    id uuid not null,
    code character varying not null,
    "description" character varying,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    primary key (id)
);

create table if not exists order_types (
    id uuid not null,
    code character varying not null,
    "description" character varying,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    primary key (id)
);

create table if not exists saved_locations (
    id uuid not null,
    user_id uuid not null,
    "name" character varying not null,
    "address" character varying not null,
    notes character varying,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    primary key (id),
    constraint fk_location_user_id FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE SET NULL 
);

create table if not exists transaction_order (
    id uuid not null,
    transaction_order_number character varying not null,
    user_id uuid not null,
    status_id uuid not null,
    order_type_id uuid not null,
    location_id uuid not null,
    is_cancel uuid not null,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    primary key (id),
    constraint fk_location_id FOREIGN key (location_id) references saved_locations(id) on DELETE set null,
    constraint fk_order_type foreign key (order_type_id) references order_types(id) on DELETE set null,
    constraint fk_order_status foreign key (status_id) references order_status(id) on DELETE set null
);

insert into order_status (id, code, "description", created_at, updated_at) 
values
('49e8b0cf-d138-4f65-97f5-6e6338efba23', 'waiting', 'waiting admin to accept', now(), now()),
('72e2d104-6abf-49f6-8861-c56ff266295b', 'progress', 'your order are on cook', now(), now()),
('f23b5af3-0b51-44a4-9b51-7bf38155a143', 'accepted', 'your order has been accepted by admin', now(), now()),
('b01f69e1-2449-4d9b-988d-9f8654d6055c', 'done', 'order finished', now(), now());

insert into order_types (id, code, "description", created_at, updated_at) 
values
('e5fdd208-e0ff-45b5-8fed-9813344dfd82', 'cod', 'Buy on site', now(), now()),
('dba0639f-d3a6-4bb7-b52b-ac796e105870', 'po', 'Open Purchase Order', now(), now());
