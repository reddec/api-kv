-- +migrate Up

CREATE TABLE value
(
    ID           INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    EXPIRE_AT    INTEGER, -- DT in MS
    CONTENT_TYPE TEXT    NOT NULL,
    NAMESPACE    TEXT    NOT NULL,
    KEY          TEXT    NOT NULL,
    VALUE        BLOB    NOT NULL
);

CREATE VIEW value_key AS
SELECT id, key, namespace, expire_at
FROM value;

CREATE UNIQUE INDEX value_namespace_key ON value (NAMESPACE, KEY);
CREATE INDEX value_expire_at ON value (EXPIRE_AT);
CREATE INDEX value_namespace ON value (NAMESPACE);

-- +migrate Down
DROP INDEX value_namespace;
DROP INDEX value_expire_at;
DROP INDEX value_namespace_key;
DROP VIEW value_key;
DROP TABLE value;