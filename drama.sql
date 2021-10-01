-- Table: drama

-- DROP TABLE drama;

CREATE TABLE drama
(
 id serial NOT NULL PRIMARY KEY,
 title character varying NOT NULL,
 director character varying,
 release_date date
);
