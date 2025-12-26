-- Modify "children" table
ALTER TABLE "public"."children" ALTER COLUMN "id" SET DEFAULT gen_random_uuid();
