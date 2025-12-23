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
-- Create "foods" table
CREATE TABLE "public"."foods" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
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
