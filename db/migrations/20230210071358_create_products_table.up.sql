CREATE TABLE `product` (
  `product_id` bigint NOT NULL UNIQUE, 
  `sku` varchar(8) NOT NULL UNIQUE,
  `title` varchar(100) NOT NULL,
  `description` varchar(1000) NOT NULL,
  `category` varchar(50) NOT NULL,
  `etalase` varchar(50) NOT NULL,
  `weight` double NOT NULL,
  `price` bigint NOT NULL,
  `updated_at` date NOT NULL,
  `created_at` datetime(6) NOT NULL,
  PRIMARY KEY (`product_id`)
);