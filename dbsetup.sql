CREATE TABLE players (
	id integer NOT NULL PRIMARY KEY,
	name varying(255) NOT NULL,
	rank float NOT NULL,
	aga_id integer NOT NULL
);

CREATE TABLE results (
	id integer NOT NULL PRIMARY KEY,
	white_player_id integer REFERENCES players (id),
	black_player_id integer REFERENCES players (id),
	white_won boolean,
	time timestamp
);

/*
CREATE TABLE ladder_cache (

);

CREATE TABLE event_types (

);

CREATE TABLE events (

);
*/
