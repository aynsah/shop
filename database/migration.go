package database

import "github.com/GuiaBolso/darwin"

var (
	migrations = []darwin.Migration{
		{
			Version:     1,
			Description: "Creating table items",
			Script: `CREATE TABLE IF NOT EXISTS items (
						id INT 		auto_increment, 
						name 		VARCHAR(100) NOT NULL,
						deleted_at 	timestamp NULL DEFAULT NULL,
						PRIMARY KEY (id)
					 ) ENGINE=InnoDB CHARACTER SET=utf8mb4;`,
		},
		{
			Version:     2,
			Description: "Insert table items",
			Script: `INSERT INTO items (id, name, deleted_at) VALUES
						(1, 'item 1', NULL),
						(2, 'item 2', NULL),
						(14, 'Item 5', '2021-10-26 01:39:26'),
						(15, 'item test', NULL);`,
		},
		{
			Version:     3,
			Description: "Creating table taxes",
			Script: `CREATE TABLE IF NOT EXISTS taxes (
						id INT 		auto_increment, 
						name 		VARCHAR(100) NOT NULL,
						rate 		FLOAT NOT NULL,
						deleted_at 	timestamp NULL DEFAULT NULL,
						PRIMARY KEY (id)
					 ) ENGINE=InnoDB CHARACTER SET=utf8mb4;`,
		},
		{
			Version:     4,
			Description: "Insert table taxes",
			Script: `INSERT INTO taxes (id, name, rate, deleted_at) VALUES
						(1, 'tax 1', 5, NULL),
						(2, 'tax 2', 10, NULL),
						(3, 'tax 3', 2.5, NULL),
						(4, 'tax 4', 20, NULL),
						(5, 'tax test', 5, '2021-10-26 01:39:27'),
						(6, 'taxes test', 10, NULL),
						(7, 'taxes test', 10, NULL);`,
		},
		{
			Version:     5,
			Description: "Creating table pivot items_taxes",
			Script: `CREATE TABLE IF NOT EXISTS items_taxes (
						id INT 		auto_increment,
						item_id		INT NOT NULL,
						tax_id		INT NOT NULL,
						PRIMARY KEY (id),
						KEY item_id (item_id),
						KEY tax_id (tax_id),
						CONSTRAINT items_taxes_ibfk_1 FOREIGN KEY (item_id) REFERENCES items (id),
						CONSTRAINT items_taxes_ibfk_2 FOREIGN KEY (tax_id) REFERENCES taxes (id)
					 ) ENGINE=InnoDB CHARACTER SET=utf8mb4;`,
		},
		{
			Version:     6,
			Description: "Insert table items_taxes",
			Script: `INSERT INTO items_taxes (id, item_id, tax_id) VALUES
						(1, 1, 1),
						(2, 1, 3),
						(3, 2, 1),
						(4, 2, 2),
						(21, 15, 3),
						(22, 15, 4),
						(23, 14, 4),
						(24, 14, 1);`,
		},
	}
)
