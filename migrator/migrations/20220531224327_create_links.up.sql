CREATE TABLE links(
    id bigserial not null primary key,
    link varchar null unique,
    token varchar null unique,
    views bigint DEFAULT 0 null, 
    created_at timestamp null
);
CREATE INDEX idx_links_token
ON links(token);