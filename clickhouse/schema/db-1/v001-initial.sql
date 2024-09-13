DROP TABLE IF EXISTS ba_metrics.metrics;
CREATE TABLE IF NOT EXISTS ba_metrics.metrics
(   
    created_at DateTime,
    link_id UUID,
    ip String,
    city String,
    country String,
    timezone String,
    referer String,
    browser String,
    localization String,
    model String,
    platform String,
    os String
)
ENGINE = MergeTree
PRIMARY KEY (created_at)