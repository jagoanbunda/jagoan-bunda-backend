-- Create "users" table
CREATE TABLE "public"."users" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "name" character varying(255) NOT NULL,
  "email" character varying(255) NOT NULL,
  "phone" character varying(20) NULL,
  "password_hash" character varying(255) NOT NULL,
  "role" character varying(20) NOT NULL DEFAULT 'nakes',
  "is_verified" boolean NULL DEFAULT false,
  "address" text NULL,
  "nik" character varying(255) NULL,
  "supervisor_id" uuid NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "uni_users_phone" UNIQUE ("phone"),
  CONSTRAINT "fk_users_assigned_parents" FOREIGN KEY ("supervisor_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create index "idx_users_deleted_at" to table: "users"
CREATE INDEX "idx_users_deleted_at" ON "public"."users" ("deleted_at");
-- Create index "idx_users_email" to table: "users"
CREATE UNIQUE INDEX "idx_users_email" ON "public"."users" ("email");
-- Create "children" table
CREATE TABLE "public"."children" (
  "id" uuid NOT NULL,
  "parent_id" uuid NULL,
  "name" character varying(255) NOT NULL,
  "birthday" date NOT NULL,
  "gender" character varying(10) NOT NULL,
  "nik" character varying(255) NOT NULL,
  "birth_weight" numeric(5,2) NOT NULL,
  "birth_height" numeric(5,2) NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_children_parent" FOREIGN KEY ("parent_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create index "idx_children_deleted_at" to table: "children"
CREATE INDEX "idx_children_deleted_at" ON "public"."children" ("deleted_at");
-- Create "anthropometries" table
CREATE TABLE "public"."anthropometries" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "child_id" uuid NULL,
  "user_id" uuid NULL,
  "weight" numeric(5,2) NOT NULL,
  "height" numeric(5,2) NOT NULL,
  "head_circumference" numeric NULL,
  "z_score_bbu" numeric(4,2) NULL,
  "z_score_tbu" numeric(4,2) NULL,
  "z_score_bbtb" numeric(4,2) NULL,
  "status_bbu" numeric(4,2) NULL,
  "status_tbu" numeric(4,2) NULL,
  "status_bbtb" numeric(4,2) NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "uni_anthropometries_child_id" UNIQUE ("child_id"),
  CONSTRAINT "fk_anthropometries_child" FOREIGN KEY ("child_id") REFERENCES "public"."children" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_anthropometries_user" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create index "idx_anthropometries_deleted_at" to table: "anthropometries"
CREATE INDEX "idx_anthropometries_deleted_at" ON "public"."anthropometries" ("deleted_at", "deleted_at");
-- Create "food_intakes" table
CREATE TABLE "public"."food_intakes" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "child_id" uuid NULL,
  "user_id" uuid NULL,
  "meal_time" character varying(20) NOT NULL DEFAULT 'breakfast',
  "foods" jsonb NOT NULL,
  "total_energy" numeric(8,2) NULL,
  "total_protein" numeric(8,2) NULL,
  "total_fat" numeric(8,2) NULL,
  "total_carbohydrate" numeric(8,2) NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_food_intakes_child" FOREIGN KEY ("child_id") REFERENCES "public"."children" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_food_intakes_user" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create index "idx_food_intakes_deleted_at" to table: "food_intakes"
CREATE INDEX "idx_food_intakes_deleted_at" ON "public"."food_intakes" ("deleted_at", "deleted_at");
-- Create "kpsp_screenings" table
CREATE TABLE "public"."kpsp_screenings" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "nakes_id" uuid NULL,
  "child_id" uuid NULL,
  "date" date NOT NULL,
  "answers" json NULL,
  "result" text NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_kpsp_screenings_child" FOREIGN KEY ("child_id") REFERENCES "public"."children" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_kpsp_screenings_nakes" FOREIGN KEY ("nakes_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create index "idx_kpsp_screenings_deleted_at" to table: "kpsp_screenings"
CREATE INDEX "idx_kpsp_screenings_deleted_at" ON "public"."kpsp_screenings" ("deleted_at");
-- Create "pmt_programs" table
CREATE TABLE "public"."pmt_programs" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "child_id" uuid NULL,
  "user_id" uuid NULL,
  "start_date" date NOT NULL,
  "end_date" date NOT NULL,
  "status" character varying(20) NULL DEFAULT 'active',
  "notes" text NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_pmt_programs_child" FOREIGN KEY ("child_id") REFERENCES "public"."children" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_pmt_programs_user" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create index "idx_pmt_programs_deleted_at" to table: "pmt_programs"
CREATE INDEX "idx_pmt_programs_deleted_at" ON "public"."pmt_programs" ("deleted_at", "deleted_at");
-- Create "pmt_daily_records" table
CREATE TABLE "public"."pmt_daily_records" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "user_id" uuid NULL,
  "program_id" bigint NULL,
  "date" date NOT NULL,
  "consumed" boolean NOT NULL,
  "notes" text NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_pmt_daily_records_program" FOREIGN KEY ("program_id") REFERENCES "public"."pmt_programs" ("id") ON UPDATE CASCADE ON DELETE SET NULL,
  CONSTRAINT "fk_pmt_daily_records_user" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create index "idx_pmt_daily_records_deleted_at" to table: "pmt_daily_records"
CREATE INDEX "idx_pmt_daily_records_deleted_at" ON "public"."pmt_daily_records" ("deleted_at", "deleted_at");
-- Create index "idx_pmt_daily_records_program_id" to table: "pmt_daily_records"
CREATE INDEX "idx_pmt_daily_records_program_id" ON "public"."pmt_daily_records" ("program_id");
