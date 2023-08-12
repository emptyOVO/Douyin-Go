/*
SQLyog Enterprise v12.09 (64 bit)
MySQL - 8.0.31 : Database - douyin
*********************************************************************
*/

/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`douyin` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;

/*Table structure for table `comment` */

CREATE TABLE `comment` (
  `comment_id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `video_id` int NOT NULL,
  `comment_text` varchar(100) COLLATE utf8mb4_general_ci NOT NULL,
  `create_time` varchar(20) COLLATE utf8mb4_general_ci NOT NULL,
  `timestamp` int NOT NULL,
  PRIMARY KEY (`comment_id`),
  KEY `user_id` (`user_id`),
  KEY `video_id` (`video_id`),
  CONSTRAINT `comment_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `user` (`user_id`),
  CONSTRAINT `comment_ibfk_2` FOREIGN KEY (`video_id`) REFERENCES `video` (`video_id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

/*Data for the table `comment` */

insert  into `comment`(`comment_id`,`user_id`,`video_id`,`comment_text`,`create_time`,`timestamp`) values (1,1,6,'好厉害','08-06',1691307847),(3,6,2,'好大的雨','08-06',1691315306),(8,6,4,'好可爱','08-08',1691509000),(9,6,6,'确实厉害','08-11',1691753164);

/*Table structure for table `follow` */

CREATE TABLE `follow` (
  `id` int NOT NULL AUTO_INCREMENT,
  `follow_id` int NOT NULL,
  `followed_id` int NOT NULL,
  PRIMARY KEY (`id`),
  KEY `follow_id` (`follow_id`),
  KEY `followed_id` (`followed_id`),
  CONSTRAINT `follow_ibfk_1` FOREIGN KEY (`follow_id`) REFERENCES `user` (`user_id`),
  CONSTRAINT `follow_ibfk_2` FOREIGN KEY (`followed_id`) REFERENCES `user` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

/*Data for the table `follow` */

insert  into `follow`(`id`,`follow_id`,`followed_id`) values (1,3,2),(2,3,1),(3,2,3),(4,2,1),(6,6,1),(7,1,6),(8,6,3);

/*Table structure for table `like` */

CREATE TABLE `like` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `video_id` int NOT NULL,
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`),
  KEY `video_id` (`video_id`),
  CONSTRAINT `like_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `user` (`user_id`),
  CONSTRAINT `like_ibfk_2` FOREIGN KEY (`video_id`) REFERENCES `video` (`video_id`)
) ENGINE=InnoDB AUTO_INCREMENT=17 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

/*Data for the table `like` */

insert  into `like`(`id`,`user_id`,`video_id`) values (1,3,4),(2,3,3),(3,3,2),(4,3,1),(5,2,6),(8,2,2),(9,2,1),(12,6,2),(13,3,14),(14,3,8);

/*Table structure for table `message` */

CREATE TABLE `message` (
  `msg_id` int NOT NULL AUTO_INCREMENT,
  `from_user_id` int NOT NULL,
  `to_user_id` int NOT NULL,
  `content` varchar(255) COLLATE utf8mb4_general_ci NOT NULL,
  `create_time` int NOT NULL,
  PRIMARY KEY (`msg_id`),
  KEY `from_user_id` (`from_user_id`),
  KEY `to_user_id` (`to_user_id`),
  CONSTRAINT `message_ibfk_1` FOREIGN KEY (`from_user_id`) REFERENCES `user` (`user_id`),
  CONSTRAINT `message_ibfk_2` FOREIGN KEY (`to_user_id`) REFERENCES `user` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

/*Data for the table `message` */

insert  into `message`(`msg_id`,`from_user_id`,`to_user_id`,`content`,`create_time`) values (1,1,6,'你好呀',1691315390),(2,6,1,'你也好呀，很高兴认识你',1691315413);

/*Table structure for table `user` */

CREATE TABLE `user` (
  `user_id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `follow_count` int NOT NULL,
  `follower_count` int NOT NULL,
  `password` varchar(255) COLLATE utf8mb4_general_ci NOT NULL,
  `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `background_image` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `signature` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `total_favorited` int DEFAULT NULL,
  `work_count` int DEFAULT NULL,
  `favorite_count` int DEFAULT NULL,
  `is_follow` tinyint(1) DEFAULT '0',
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

/*Data for the table `user` */

insert  into `user`(`user_id`,`name`,`follow_count`,`follower_count`,`password`,`avatar`,`background_image`,`signature`,`total_favorited`,`work_count`,`favorite_count`,`is_follow`) values (1,'1',1,3,'e10adc3949ba59abbe56e057f20f883e','http://dummyimage.com/400x400','http://dummyimage.com/400x400','这是一条个性签名',0,0,0,0),(2,'2',2,2,'e10adc3949ba59abbe56e057f20f883e','http://dummyimage.com/400x400','http://dummyimage.com/400x400','这是一条个性签名',1,0,0,0),(3,'3',2,2,'e10adc3949ba59abbe56e057f20f883e','http://dummyimage.com/400x400','http://dummyimage.com/400x400','这是一条个性签名',0,0,0,0),(5,'hahaha',0,0,'e10adc3949ba59abbe56e057f20f883e','http://dummyimage.com/400x400','http://dummyimage.com/400x400','这是一条个性签名',0,0,0,0),(6,'feige',2,1,'e10adc3949ba59abbe56e057f20f883e','http://dummyimage.com/400x400','http://dummyimage.com/400x400','这是一条个性签名',0,2,0,0);

/*Table structure for table `video` */

CREATE TABLE `video` (
  `video_id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `play_url` varchar(255) COLLATE utf8mb4_general_ci NOT NULL,
  `cover_url` varchar(255) COLLATE utf8mb4_general_ci NOT NULL,
  `favorite_count` int NOT NULL,
  `comment_count` int NOT NULL,
  `is_favorite` tinyint(1) NOT NULL,
  `title` varchar(255) COLLATE utf8mb4_general_ci NOT NULL,
  `timestamp` int NOT NULL,
  PRIMARY KEY (`video_id`),
  KEY `user_id` (`user_id`),
  CONSTRAINT `video_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `user` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

/*Data for the table `video` */

insert  into `video`(`video_id`,`user_id`,`play_url`,`cover_url`,`favorite_count`,`comment_count`,`is_favorite`,`title`,`timestamp`) values (1,1,'http://192.168.2.166:80/public/1_wx_camera_1690344007531.mp4','http://192.168.2.166:80/public/1_wx_camera_1690344007531.mp4.png',2,0,0,'1',1691235846),(2,1,'http://192.168.2.166:80/public/1_video_20230728_090622.mp4','http://192.168.2.166:80/public/1_video_20230728_090622.mp4.png',3,1,0,'2',1691235857),(3,2,'http://192.168.2.166:80/public/2_mmexport1690506409070.mp4','http://192.168.2.166:80/public/2_mmexport1690506409070.mp4.png',1,0,0,'3',1691235883),(4,2,'http://192.168.2.166:80/public/2_392df987627643be6a16caada48b6d59.mp4','http://192.168.2.166:80/public/2_392df987627643be6a16caada48b6d59.mp4.png',2,1,0,'4',1691235892),(6,3,'http://192.168.2.166:80/public/3_wx_camera_1690458000098.mp4','http://192.168.2.166:80/public/3_wx_camera_1690458000098.mp4.png',1,2,0,'6',1691235952),(8,6,'http://192.168.2.166:80/public/6_fad34bce3055fe10fc9625c54a2d03ae.mp4','http://192.168.2.166:80/public/6_fad34bce3055fe10fc9625c54a2d03ae.mp4.png',1,0,0,'嘻嘻',1691315340),(10,6,'http://192.168.2.166:80/public/6_mmexport1687010961497.mp4','http://192.168.2.166:80/public/6_mmexport1687010961497.mp4.png',0,0,0,'乐队',1691504805),(14,6,'http://192.168.2.166:80/public/6_share_ea985702feb33b2eaedc78fd3acc11d1.mp4','http://192.168.2.166:80/public/6_share_ea985702feb33b2eaedc78fd3acc11d1.mp4.png',1,0,0,'手指',1691507244),(15,3,'http://192.168.2.166:80/public/3_VIDEO_20230808_234120648.mp4','http://192.168.2.166:80/public/3_VIDEO_20230808_234120648.mp4.png',0,0,0,'新ipad',1691509288);

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
