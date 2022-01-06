BEGIN TRANSACTION;
CREATE TABLE [activities] (
id INTEGER NOT NULL PRIMARY KEY,
time TEXT,
description TEXT
);
INSERT INTO "activities" VALUES(1,'2021-12-09T16:34:04Z','christmas eve bike class');
INSERT INTO "activities" VALUES(2,'2021-12-09T16:56:12Z','cross country skiing is horrible and cold');
INSERT INTO "activities" VALUES(3,'2021-12-09T16:56:23Z','sledding with nephew');
COMMIT;
