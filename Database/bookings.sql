-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: localhost
-- Generation Time: Jul 18, 2023 at 07:32 PM
-- Server version: 10.4.28-MariaDB
-- PHP Version: 8.0.28

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `bookings`
--

-- --------------------------------------------------------

--
-- Table structure for table `reservations`
--

CREATE TABLE `reservations` (
  `id` int(11) UNSIGNED NOT NULL,
  `first_name` varchar(255) DEFAULT NULL,
  `last_name` varchar(255) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `phone` varchar(255) DEFAULT NULL,
  `start_date` date DEFAULT NULL,
  `end_date` date DEFAULT NULL,
  `room_id` int(11) UNSIGNED DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `processed` int(11) DEFAULT 0
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `reservations`
--

INSERT INTO `reservations` (`id`, `first_name`, `last_name`, `email`, `phone`, `start_date`, `end_date`, `room_id`, `created_at`, `updated_at`, `processed`) VALUES
(1, 'Muhammad', 'Selim', 'salimmia3745@gmail.com', '01521435652', '2020-11-30', '2020-12-02', 1, '2023-07-13 15:40:09', '2023-07-13 09:40:09', 1),
(5, 'Muhammad', 'Halim', 'halim@gmail.com', '555-222-666', '2021-02-01', '2021-02-04', 1, '2023-06-16 07:15:10', '2023-06-16 07:15:10', 0),
(6, 'Muhammad', 'Tomal', 'tomal@gmail.com', '555-666-567', '2020-11-26', '2020-11-27', 2, '2023-07-13 15:39:24', '2023-07-13 09:34:37', 1),
(7, 'Delowar', 'Hossain Shakil', 'shakil@gmail.com', '555-444-333', '2021-01-04', '2021-01-05', 1, '2023-07-14 10:54:32', '2023-07-14 04:54:32', 0),
(9, 'MH Shadhin', 'Prodhan', 'shadhin@gmail.com', '111-222-333', '2021-03-09', '2021-03-10', 2, '2023-06-16 12:24:37', '2023-06-16 12:24:37', 0),
(10, 'Shiddharth', 'Joy', 'joy@gmail.com', '222-111-333', '2023-06-07', '2023-06-08', 1, '2023-07-14 10:53:52', '2023-07-14 04:53:52', 1),
(18, 'Ashik', 'asd', 'd@g.xo', '333', '2023-06-17', '2023-06-17', 1, '2023-07-14 10:24:43', '2023-06-21 08:23:13', 1),
(21, 'selim', 'vai', 'a@gmail.com', '111-222-333', '2023-06-19', '2023-06-19', 1, '2023-06-17 10:43:27', '2023-06-17 10:43:27', 0),
(22, 'testing', 'email', 'emai@gmail.com', '111-222-333', '2023-06-19', '2023-06-22', 1, '2023-07-15 13:26:26', '2023-07-15 07:26:26', 1),
(34, 'Rrid', 'Rishan', 'rrid@r.com', '', '2023-07-20', '2023-07-28', 1, '2023-07-18 11:05:02', '2023-07-18 11:05:02', 0),
(35, 'Rrid', 'Rishan', 'rrid@r.com', '', '2023-07-20', '2023-07-28', 1, '2023-07-18 11:05:55', '2023-07-18 11:05:55', 0),
(36, 'Rrid', 'Rishan', 'rrid@r.com', '555555555', '2023-07-20', '2023-07-28', 1, '2023-07-18 11:06:26', '2023-07-18 11:06:26', 0),
(37, 'kkk', 'iii', 'rrid@r.com', '55555555', '2023-07-20', '2023-07-28', 2, '2023-07-18 11:07:01', '2023-07-18 11:07:01', 0),
(38, 'selim', 'mia', 'a@a.com', '888888888', '2023-10-26', '2023-10-28', 1, '2023-07-18 11:12:20', '2023-07-18 11:12:20', 0),
(39, 'salim', 'vai', 'a@b.com', '444444444', '2024-03-13', '2024-03-14', 1, '2023-07-18 11:17:16', '2023-07-18 11:17:16', 0),
(40, 'ashik', 'khan', 'b@a.com', '3333333333', '2024-03-12', '2024-03-14', 2, '2023-07-18 11:18:22', '2023-07-18 11:18:22', 0);

-- --------------------------------------------------------

--
-- Table structure for table `restrictions`
--

CREATE TABLE `restrictions` (
  `id` int(11) UNSIGNED NOT NULL,
  `restriction_name` varchar(255) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `restrictions`
--

INSERT INTO `restrictions` (`id`, `restriction_name`, `created_at`, `updated_at`) VALUES
(1, 'Reservation', '2020-11-17 18:00:00', '2020-11-17 18:00:00'),
(2, 'Owner Block', '2020-11-18 18:00:00', '2020-11-18 18:00:00');

-- --------------------------------------------------------

--
-- Table structure for table `rooms`
--

CREATE TABLE `rooms` (
  `id` int(11) UNSIGNED NOT NULL,
  `room_name` varchar(255) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `rooms`
--

INSERT INTO `rooms` (`id`, `room_name`, `created_at`, `updated_at`) VALUES
(1, 'General\'s Quarter', '2020-11-20 12:06:30', '2020-12-02 12:06:30'),
(2, 'Major\'s Suite', '2020-06-11 14:03:54', '2020-06-18 14:03:54');

-- --------------------------------------------------------

--
-- Table structure for table `room_restrictions`
--

CREATE TABLE `room_restrictions` (
  `id` int(11) UNSIGNED NOT NULL,
  `start_date` date NOT NULL,
  `end_date` date NOT NULL,
  `room_id` int(11) UNSIGNED NOT NULL,
  `reservation_id` int(11) UNSIGNED DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  `restriction_id` int(11) UNSIGNED NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `room_restrictions`
--

INSERT INTO `room_restrictions` (`id`, `start_date`, `end_date`, `room_id`, `reservation_id`, `created_at`, `updated_at`, `restriction_id`) VALUES
(1, '2021-02-01', '2021-02-04', 1, 5, '2023-06-16 07:15:10', '2023-06-16 07:15:10', 1),
(2, '2021-02-01', '2021-02-04', 2, NULL, '2020-11-17 18:00:00', '2020-11-17 18:00:00', 2),
(3, '2020-11-26', '2020-11-27', 2, 6, '2023-06-16 11:31:40', '2023-06-16 11:31:40', 1),
(4, '2021-01-04', '2021-01-05', 1, 7, '2023-06-16 11:40:24', '2023-06-16 11:40:24', 1),
(6, '2021-03-09', '2021-03-10', 2, 9, '2023-06-16 12:24:37', '2023-06-16 12:24:37', 1),
(7, '2023-06-07', '2023-06-08', 1, 10, '2023-06-17 00:44:24', '2023-06-17 00:44:24', 1),
(15, '2023-06-17', '2023-06-17', 1, 18, '2023-06-17 06:08:21', '2023-06-17 06:08:21', 1),
(18, '2023-06-19', '2023-06-19', 1, 21, '2023-06-17 10:43:27', '2023-06-17 10:43:27', 1),
(19, '2023-06-19', '2023-06-22', 1, 22, '2023-06-17 11:32:01', '2023-06-17 11:32:01', 1),
(30, '2023-07-14', '2023-07-17', 1, NULL, '2023-07-14 06:56:34', '0000-00-00 00:00:00', 2),
(31, '2023-06-13', '2023-06-14', 1, NULL, '2023-07-14 04:23:31', '2023-07-14 04:23:31', 2),
(32, '2023-06-14', '2023-06-15', 1, NULL, '2023-07-14 04:23:31', '2023-07-14 04:23:31', 2),
(33, '2023-06-14', '2023-06-15', 2, NULL, '2023-07-14 04:23:31', '2023-07-14 04:23:31', 2),
(34, '2023-06-10', '2023-06-11', 1, NULL, '2023-07-14 04:45:20', '2023-07-14 04:45:20', 2),
(35, '2023-07-13', '2023-07-14', 2, NULL, '2023-07-14 04:45:34', '2023-07-14 04:45:34', 2),
(36, '2023-07-28', '2023-07-29', 2, NULL, '2023-07-14 04:45:34', '2023-07-14 04:45:34', 2),
(37, '2023-07-11', '2023-07-12', 1, NULL, '2023-07-14 04:45:34', '2023-07-14 04:45:34', 2),
(38, '2023-07-20', '2023-07-28', 1, 34, '2023-07-18 11:05:02', '2023-07-18 11:05:02', 1),
(39, '2023-07-20', '2023-07-28', 1, 35, '2023-07-18 11:05:55', '2023-07-18 11:05:55', 1),
(40, '2023-07-20', '2023-07-28', 1, 36, '2023-07-18 11:06:26', '2023-07-18 11:06:26', 1),
(41, '2023-07-20', '2023-07-28', 2, 37, '2023-07-18 11:07:01', '2023-07-18 11:07:01', 1),
(42, '2023-10-26', '2023-10-28', 1, 38, '2023-07-18 11:12:20', '2023-07-18 11:12:20', 1),
(43, '2024-03-13', '2024-03-14', 1, 39, '2023-07-18 11:17:16', '2023-07-18 11:17:16', 1),
(44, '2024-03-12', '2024-03-14', 2, 40, '2023-07-18 11:18:22', '2023-07-18 11:18:22', 1);

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` int(11) UNSIGNED NOT NULL,
  `first_name` varchar(255) DEFAULT NULL,
  `last_name` varchar(255) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  `phone` varchar(255) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  `access_level` int(11) UNSIGNED DEFAULT 0
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `first_name`, `last_name`, `email`, `password`, `phone`, `created_at`, `updated_at`, `access_level`) VALUES
(1, 'Muhammad', 'Selim', 'admin@admin.com', '$2a$12$4omoC3s2QHI2o3REbG/bE.v38qBlL56DFW//aN3jZOE72OvM5XOB2', '+8801521435652', '2023-07-15 13:32:56', '2023-06-18 12:04:35', 1),
(16, 'Selim', 'Mia', 'admin@aa.com', '$2a$10$ggVlsrg3ojlrMvBVNQjd8eNWSI05qiv00bY.m/oumpUwVplQO9H56', '111', '2023-07-17 12:35:28', '2023-07-17 12:35:28', 0),
(20, 'asdfasdfasdfasd', 'asd', 'admin@a.com', '$2a$10$U7voN5odetsap1/6wk5O6.fLwlSslt3mrjsRJ2rIsjH30JyifrhRm', 'asdfasd', '2023-07-18 00:58:35', '2023-07-18 00:58:35', 0),
(21, 'aaaa', 'aa', 'a@a.com', '$2a$10$iMIoe/0cX6IUrtoPRcnBbOFpFJf7lUH2rc048X8CmWKWufTybtDxK', 'aaa', '2023-07-18 15:31:19', '2023-07-18 09:29:33', 0),
(22, 'Rrid ', 'Rishan', 'rrid@r.com', '$2a$10$ESAISDSUwv3jXiVawedqn.88/r9DiWvB2XgEOe0MRwVNUp29TqiWC', '017411111111', '2023-07-18 11:03:58', '2023-07-18 11:03:58', 0);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `reservations`
--
ALTER TABLE `reservations`
  ADD PRIMARY KEY (`id`),
  ADD KEY `room_id` (`room_id`);

--
-- Indexes for table `restrictions`
--
ALTER TABLE `restrictions`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `rooms`
--
ALTER TABLE `rooms`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `room_restrictions`
--
ALTER TABLE `room_restrictions`
  ADD PRIMARY KEY (`id`),
  ADD KEY `room_id` (`room_id`),
  ADD KEY `reservation_id` (`reservation_id`),
  ADD KEY `restriction_id` (`restriction_id`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `reservations`
--
ALTER TABLE `reservations`
  MODIFY `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=41;

--
-- AUTO_INCREMENT for table `restrictions`
--
ALTER TABLE `restrictions`
  MODIFY `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT for table `rooms`
--
ALTER TABLE `rooms`
  MODIFY `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT for table `room_restrictions`
--
ALTER TABLE `room_restrictions`
  MODIFY `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=45;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=23;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `reservations`
--
ALTER TABLE `reservations`
  ADD CONSTRAINT `reservations_ibfk_1` FOREIGN KEY (`room_id`) REFERENCES `rooms` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `reservations_ibfk_2` FOREIGN KEY (`room_id`) REFERENCES `rooms` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Constraints for table `room_restrictions`
--
ALTER TABLE `room_restrictions`
  ADD CONSTRAINT `room_restrictions_ibfk_1` FOREIGN KEY (`room_id`) REFERENCES `rooms` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `room_restrictions_ibfk_2` FOREIGN KEY (`reservation_id`) REFERENCES `reservations` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `room_restrictions_ibfk_3` FOREIGN KEY (`restriction_id`) REFERENCES `restrictions` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `room_restrictions_ibfk_4` FOREIGN KEY (`room_id`) REFERENCES `rooms` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `room_restrictions_ibfk_5` FOREIGN KEY (`reservation_id`) REFERENCES `reservations` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `room_restrictions_ibfk_6` FOREIGN KEY (`restriction_id`) REFERENCES `restrictions` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
