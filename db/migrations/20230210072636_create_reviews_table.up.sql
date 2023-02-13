CREATE TABLE `review` (
  `product_id` bigint NOT NULL, 
  `review_id` bigint NOT NULL UNIQUE,
  `rating` float NOT NULL,
  `review_comment` varchar(1000) NOT NULL,
  `date_time_review` varchar(100) NOT NULL,
  `created_at` datetime(6) NOT NULL,
  PRIMARY KEY (`review_id`),
  FOREIGN KEY (`product_id`) REFERENCES product(`product_id`)
);