-- Create "users" table
CREATE TABLE "public"."users" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "name" character varying(255) NOT NULL,
  "profile_picture" character varying(500) NULL,
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
-- Create "asq_recommendations" table
CREATE TABLE "public"."asq_recommendations" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "domain" character varying(20) NOT NULL,
  "age_months_min" integer NOT NULL,
  "age_months_max" integer NOT NULL,
  "activity_title" character varying(255) NOT NULL,
  "activity_description" text NOT NULL,
  "video_url" character varying(500) NULL,
  "sort_order" integer NULL DEFAULT 0,
  PRIMARY KEY ("id")
);
-- Create "education_articles" table
CREATE TABLE "public"."education_articles" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "title" character varying(255) NOT NULL,
  "category" character varying(50) NOT NULL,
  "content" text NOT NULL,
  "thumbnail_url" character varying(500) NULL,
  "published_at" timestamp NULL,
  "view_count" integer NULL DEFAULT 0,
  "is_active" boolean NULL DEFAULT true,
  "created_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  PRIMARY KEY ("id")
);
-- Create index "idx_education_articles_deleted_at" to table: "education_articles"
CREATE INDEX "idx_education_articles_deleted_at" ON "public"."education_articles" ("deleted_at");
-- Create "children" table
CREATE TABLE "public"."children" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
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
  CONSTRAINT "fk_anthropometries_child" FOREIGN KEY ("child_id") REFERENCES "public"."children" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_anthropometries_user" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create index "idx_anthropometries_deleted_at" to table: "anthropometries"
CREATE INDEX "idx_anthropometries_deleted_at" ON "public"."anthropometries" ("deleted_at", "deleted_at");
-- Create "asq_questionnaires" table
CREATE TABLE "public"."asq_questionnaires" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "age_months" integer NOT NULL,
  "age_range_min" integer NOT NULL,
  "age_range_max" integer NOT NULL,
  "version" character varying(20) NULL DEFAULT '1.0',
  "is_active" boolean NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "uni_asq_questionnaires_age_months" UNIQUE ("age_months"),
  CONSTRAINT "uni_asq_questionnaires_age_range_max" UNIQUE ("age_range_max"),
  CONSTRAINT "uni_asq_questionnaires_age_range_min" UNIQUE ("age_range_min")
);
-- Create index "idx_asq_questionnaires_deleted_at" to table: "asq_questionnaires"
CREATE INDEX "idx_asq_questionnaires_deleted_at" ON "public"."asq_questionnaires" ("deleted_at");
-- Create "asq_questions" table
CREATE TABLE "public"."asq_questions" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "questionnaire_id" bigint NULL,
  "domain" character varying(20) NOT NULL,
  "question_number" integer NOT NULL,
  "question_text" text NOT NULL,
  "question_text_id" text NULL,
  "illustration_url" character varying(500) NULL,
  "how_to_check" text NULL,
  "sort_order" integer NULL DEFAULT 0,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_asq_questions_questionnaire" FOREIGN KEY ("questionnaire_id") REFERENCES "public"."asq_questionnaires" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create "asq_screenings" table
CREATE TABLE "public"."asq_screenings" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "child_id" uuid NULL,
  "questionnaire_id" bigint NULL,
  "screening_date" date NOT NULL,
  "age_at_screening" integer NOT NULL,
  "completed_by_id" uuid NULL,
  "reviewed_by_id" uuid NULL,
  "status" character varying(20) NULL DEFAULT 'completed',
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_asq_screenings_child" FOREIGN KEY ("child_id") REFERENCES "public"."children" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_asq_screenings_completed_by" FOREIGN KEY ("completed_by_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_asq_screenings_questionnaire" FOREIGN KEY ("questionnaire_id") REFERENCES "public"."asq_questionnaires" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_asq_screenings_reviewed_by" FOREIGN KEY ("reviewed_by_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create index "idx_asq_screenings_deleted_at" to table: "asq_screenings"
CREATE INDEX "idx_asq_screenings_deleted_at" ON "public"."asq_screenings" ("deleted_at");
-- Create "asq_answers" table
CREATE TABLE "public"."asq_answers" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "screening_id" uuid NULL,
  "question_id" uuid NULL,
  "answer" character varying(20) NOT NULL,
  "score" integer NOT NULL,
  "notes" text NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_asq_answers_question" FOREIGN KEY ("question_id") REFERENCES "public"."asq_questions" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_asq_answers_screening" FOREIGN KEY ("screening_id") REFERENCES "public"."asq_screenings" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create "asq_concerns" table
CREATE TABLE "public"."asq_concerns" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "screening_id" uuid NULL,
  "has_vision_concern" boolean NULL DEFAULT false,
  "has_hearing_concern" boolean NULL DEFAULT false,
  "has_behavior_concern" boolean NULL DEFAULT false,
  "has_other_concern" boolean NULL DEFAULT false,
  "concern_details" text NULL,
  "created_at" timestamptz NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_asq_concerns_screening" FOREIGN KEY ("screening_id") REFERENCES "public"."asq_screenings" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create "asq_cutoffs" table
CREATE TABLE "public"."asq_cutoffs" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "questionnaire_id" bigint NULL,
  "domain" character varying(20) NOT NULL,
  "cutoff_score" numeric(4,1) NOT NULL,
  "monitoring_zone" numeric(4,1) NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_asq_cutoffs_questionnaire" FOREIGN KEY ("questionnaire_id") REFERENCES "public"."asq_questionnaires" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create "asq_domain_results" table
CREATE TABLE "public"."asq_domain_results" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "screening_id" uuid NULL,
  "domain" character varying(20) NOT NULL,
  "total_score" numeric(4,1) NOT NULL,
  "cutoff_score" numeric(4,1) NOT NULL,
  "monitoring_zone" numeric(4,1) NOT NULL,
  "result" character varying(20) NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_asq_domain_results_screening" FOREIGN KEY ("screening_id") REFERENCES "public"."asq_screenings" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create "foods" table
CREATE TABLE "public"."foods" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "name" character varying(255) NOT NULL,
  "category" character varying(100) NULL,
  "portion_name" character varying(100) NOT NULL,
  "portion_gram" numeric(8,2) NOT NULL,
  "energy_kcal" numeric(8,2) NOT NULL,
  "protein_g" numeric(8,2) NOT NULL,
  "fat_g" numeric(8,2) NOT NULL,
  "carbohydrate_g" numeric(8,2) NOT NULL,
  "is_active" boolean NULL DEFAULT true,
  PRIMARY KEY ("id")
);
-- Create index "idx_foods_deleted_at" to table: "foods"
CREATE INDEX "idx_foods_deleted_at" ON "public"."foods" ("deleted_at");
-- Create "food_intakes" table
CREATE TABLE "public"."food_intakes" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "child_id" uuid NULL,
  "user_id" uuid NULL,
  "meal_time" character varying(20) NOT NULL DEFAULT 'breakfast',
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
-- Create "food_intake_foods" table
CREATE TABLE "public"."food_intake_foods" (
  "food_intake_id" bigint NOT NULL,
  "food_id" bigint NOT NULL,
  PRIMARY KEY ("food_intake_id", "food_id"),
  CONSTRAINT "fk_food_intake_foods_food" FOREIGN KEY ("food_id") REFERENCES "public"."foods" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_food_intake_foods_food_intake" FOREIGN KEY ("food_intake_id") REFERENCES "public"."food_intakes" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
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
