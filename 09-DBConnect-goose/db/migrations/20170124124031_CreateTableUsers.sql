
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
create table Users (
  key            char(008)     primary key,
  data1          int8,
  data2          int8,
  data3          int8
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE Users;
