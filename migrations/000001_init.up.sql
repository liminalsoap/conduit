CREATE TABLE "users" (
                         "id" bigserial PRIMARY KEY,
                         "email" varchar(255) UNIQUE NOT NULL,
                         "username" varchar(255) UNIQUE NOT NULL,
                         "password" varchar(255) NOT NULL,
                         "bio" varchar,
                         "image" varchar(255)
);

CREATE TABLE "follows" (
                           "id" bigserial PRIMARY KEY,
                           "following_user_id" bigint NOT NULL,
                           "followed_user_id" bigint NOT NULL,
                           "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "tags" (
                        "id" bigserial PRIMARY KEY,
                        "title" varchar(255) UNIQUE NOT NULL
);

CREATE TABLE "articles" (
                            "id" bigserial PRIMARY KEY,
                            "slug" varchar(255) UNIQUE NOT NULL,
                            "title" varchar(255) NOT NULL,
                            "description" varchar NOT NULL,
                            "body" text NOT NULL,
                            "created_at" timestamp NOT NULL DEFAULT (now()),
                            "updated_at" timestamp NOT NULL DEFAULT (now()),
                            "user_id" bigint NOT NULL
);

CREATE TABLE "likes" (
                         "id" bigserial PRIMARY KEY,
                         "article_id" bigint NOT NULL,
                         "user_id" bigint NOT NULL
);

CREATE TABLE "articles_tags" (
                                 "id" bigserial PRIMARY KEY,
                                 "article_id" bigint NOT NULL,
                                 "tag_id" bigint NOT NULL
);

CREATE TABLE "comments" (
                            "id" bigserial PRIMARY KEY,
                            "created_at" timestamp NOT NULL DEFAULT (now()),
                            "updated_at" timestamp NOT NULL DEFAULT (now()),
                            "body" text NOT NULL,
                            "user_id" bigint NOT NULL
);

CREATE INDEX ON "users" ("username");

CREATE INDEX ON "tags" ("title");

CREATE INDEX ON "articles" ("slug");

CREATE INDEX ON "articles" ("title");

COMMENT ON COLUMN "articles"."user_id" IS 'author id';

COMMENT ON COLUMN "comments"."user_id" IS 'author id';

ALTER TABLE "follows" ADD FOREIGN KEY ("followed_user_id") REFERENCES "users" ("id") ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE "follows" ADD FOREIGN KEY ("following_user_id") REFERENCES "users" ("id") ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE "articles" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE NO ACTION;

ALTER TABLE "likes" ADD FOREIGN KEY ("article_id") REFERENCES "articles" ("id") ON DELETE CASCADE;

ALTER TABLE "likes" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE;

ALTER TABLE "articles_tags" ADD FOREIGN KEY ("article_id") REFERENCES "articles" ("id") ON DELETE CASCADE;

ALTER TABLE "articles_tags" ADD FOREIGN KEY ("tag_id") REFERENCES "tags" ("id") ON DELETE CASCADE;

ALTER TABLE "comments" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION;