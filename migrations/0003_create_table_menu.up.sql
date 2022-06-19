CREATE TABLE menu (
    id uuid NOT NULL,
    name character varying NOT NULL,
    description character varying,
    price numeric NOT NULL,
    stock integer NOT NULL,
    image_url character varying,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    PRIMARY KEY (id)
);

INSERT INTO menu (id, name, description, price, stock, image_url, created_at, updated_at)
VALUES 
    ('f1740163-1bf3-478c-95ab-f5cc2277d8d7', 'Ote-Ote Porong', 'Isi Ayam dan Rumput Laut', 8000, 10, 'https://via.placeholder.com/150', NOW(), NOW()),
    ('52f0645e-6af7-42eb-a643-5da3a29e296c', 'Ote-Ote Porong', 'Isi Babi dan Rumput Laut', 10000, 10, 'https://via.placeholder.com/150', NOW(), NOW());