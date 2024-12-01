-- Create enum type "auth_providers"
CREATE TYPE "public"."auth_providers" AS ENUM ('local');
-- Create enum type "permissions"
CREATE TYPE "public"."permissions" AS ENUM ('user', 'admin');
-- Create "users" table
CREATE TABLE "public"."users" ("id" uuid NOT NULL DEFAULT gen_random_uuid(), "username" text NOT NULL DEFAULT 'anonymous', "email" text NOT NULL, "password" text NULL, "provider" text NOT NULL DEFAULT 'local', "permissions" text NOT NULL DEFAULT 'user', PRIMARY KEY ("id"), CONSTRAINT "password_length" CHECK (length(password) > 8), CONSTRAINT "username_length" CHECK (length(username) > 3));
-- Create index "unique_email" to table: "users"
CREATE UNIQUE INDEX "unique_email" ON "public"."users" ("email");
-- Create index "unique_username" to table: "users"
CREATE UNIQUE INDEX "unique_username" ON "public"."users" ("username");
