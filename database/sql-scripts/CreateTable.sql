CREATE TABLE CardItem (
    ID NOT NULL PRIMARY KEY,
    NAME VARCHAR(100),
    CardMarket_id INTEGER,
    Multiverse_id INTEGER,
    TcgPlayer_id INTEGER,
    MTGO_id INTEGER,
    Reserved NOT NULL VARCHAR(10),
    PriceNormal VARCHAR(256),
    PriceFoil VARCHAR(256),
    MTGO_Tix VARCHAR(256)
);