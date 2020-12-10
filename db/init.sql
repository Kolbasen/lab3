-- Create tables.
DROP TABLE IF EXISTS "dishes";
CREATE TABLE "dishes"
(
    "id"   SERIAL PRIMARY KEY,
    "name" VARCHAR(50) NOT NULL UNIQUE,
    "price" INTEGER NOT NULL
);

-- Insert demo data.
INSERT INTO "dishes" (name, price) VALUES ('Borsch', 100);
INSERT INTO "dishes" (name, price) VALUES ('Pasta', 130);
INSERT INTO "dishes" (name, price) VALUES ('Salo', 35);
INSERT INTO "dishes" (name, price) VALUES ('Salad', 80);
INSERT INTO "dishes" (name, price) VALUES ('Pizza', 150);
INSERT INTO "dishes" (name, price) VALUES ('Fish', 140);
