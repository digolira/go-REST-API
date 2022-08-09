USE teams_test;

CREATE TABLE Teams (
    id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(70),
    city VARCHAR(40),
    state VARCHAR(5),
    stadium VARCHAR(70),
    LeagueDivision VARCHAR(3),
    PRIMARY KEY (id)
);
