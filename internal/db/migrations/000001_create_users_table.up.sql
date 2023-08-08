CREATE TABLE "users" (
"id" bigserial PRIMARY KEY,
"username" varchar NOT NULL,
"first_name" varchar NOT NULL,
"second_name" varchar NOT NULL,
"email" varchar NOT NULL,
"hashed_password" varchar NOT NULL,
"created_at" timestamptz NOT NULL DEFAULT (now())
);