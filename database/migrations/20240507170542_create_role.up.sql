CREATE TYPE Role as ENUM ('USER', 'ADMIN');
ALTER TABLE public."user"
ADD COLUMN role Role;