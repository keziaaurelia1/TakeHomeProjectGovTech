CREATE TABLE `image` (
  `product_id` bigint NOT NULL, 
  `image_id` bigint NOT NULL UNIQUE,
  `path` varchar(1000) NOT NULL,
  `description` varchar(10) NOT NULL,
  `created_at` datetime(6) NOT NULL,
  PRIMARY KEY (`image_id`),
  FOREIGN KEY (`product_id`) REFERENCES product(`product_id`)
);