-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Jul 17, 2025 at 12:41 PM
-- Server version: 10.4.32-MariaDB
-- PHP Version: 8.2.12

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `rental_kendaraan`
--

-- --------------------------------------------------------

--
-- Table structure for table `kendaraan`
--

CREATE TABLE `kendaraan` (
  `id` int(11) NOT NULL,
  `tipe` enum('mobil','motor') NOT NULL,
  `merk` varchar(100) NOT NULL,
  `nomor_polisi` varchar(20) NOT NULL,
  `harga_sewa` decimal(10,2) NOT NULL,
  `status` enum('tersedia','disewa') NOT NULL DEFAULT 'tersedia'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `kendaraan`
--

INSERT INTO `kendaraan` (`id`, `tipe`, `merk`, `nomor_polisi`, `harga_sewa`, `status`) VALUES
(1, 'mobil', 'Toyota Avanza', 'B1001AAA', 300000.00, 'tersedia'),
(2, 'mobil', 'Honda Mobilio', 'B1002BBB', 320000.00, 'tersedia'),
(3, 'mobil', 'Daihatsu Xenia', 'B1003CCC', 310000.00, 'tersedia'),
(4, 'mobil', 'Suzuki Ertiga', 'B1004DDD', 315000.00, 'tersedia'),
(5, 'mobil', 'Mitsubishi Xpander', 'B1005EEE', 350000.00, 'tersedia'),
(6, 'mobil', 'Wuling Confero', 'B1006FFF', 290000.00, 'tersedia'),
(7, 'mobil', 'Nissan Livina', 'B1007GGG', 330000.00, 'tersedia'),
(8, 'motor', 'Honda Vario', 'B2001HHH', 90000.00, 'tersedia'),
(9, 'motor', 'Yamaha Mio', 'B2002III', 85000.00, 'tersedia'),
(10, 'motor', 'Suzuki Nex', 'B2003JJJ', 80000.00, 'tersedia'),
(11, 'motor', 'Yamaha Aerox', 'B2004KKK', 95000.00, 'tersedia'),
(12, 'motor', 'Honda Beat', 'B2005LLL', 88000.00, 'tersedia'),
(13, 'motor', 'Kawasaki KLX', 'B2006MMM', 120000.00, 'tersedia'),
(14, 'motor', 'Vespa Sprint', 'B2007NNN', 150000.00, 'tersedia'),
(15, 'motor', 'Honda PCX', 'B2008OOO', 130000.00, 'tersedia'),
(16, 'mobil', 'Toyota Avanza', 'B1001AAA', 300000.00, 'tersedia'),
(17, 'mobil', 'Honda Mobilio', 'B1002BBB', 320000.00, 'tersedia'),
(18, 'mobil', 'Daihatsu Xenia', 'B1003CCC', 310000.00, 'tersedia'),
(19, 'mobil', 'Suzuki Ertiga', 'B1004DDD', 315000.00, 'tersedia'),
(20, 'mobil', 'Mitsubishi Xpander', 'B1005EEE', 350000.00, 'tersedia'),
(21, 'mobil', 'Wuling Confero', 'B1006FFF', 290000.00, 'tersedia'),
(22, 'mobil', 'Nissan Livina', 'B1007GGG', 330000.00, 'tersedia'),
(23, 'motor', 'Honda Vario', 'B2001HHH', 90000.00, 'tersedia'),
(24, 'motor', 'Yamaha Mio', 'B2002III', 85000.00, 'tersedia'),
(25, 'motor', 'Suzuki Nex', 'B2003JJJ', 80000.00, 'tersedia'),
(26, 'motor', 'Yamaha Aerox', 'B2004KKK', 95000.00, 'tersedia'),
(27, 'motor', 'Honda Beat', 'B2005LLL', 88000.00, 'tersedia'),
(28, 'motor', 'Kawasaki KLX', 'B2006MMM', 120000.00, 'tersedia'),
(29, 'motor', 'Vespa Sprint', 'B2007NNN', 150000.00, 'tersedia'),
(30, 'motor', 'Honda PCX', 'B2008OOO', 130000.00, 'tersedia'),
(31, 'motor', 'vario', 'b6780gk', 85000.00, 'tersedia');

-- --------------------------------------------------------

--
-- Table structure for table `pelanggan`
--

CREATE TABLE `pelanggan` (
  `id` int(11) NOT NULL,
  `nama` varchar(100) NOT NULL,
  `alamat` varchar(255) DEFAULT NULL,
  `telepon` varchar(20) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `pelanggan`
--

INSERT INTO `pelanggan` (`id`, `nama`, `alamat`, `telepon`) VALUES
(1, 'Andi Nugroho', 'Jl. Mawar No.1', '081234567890'),
(2, 'Budi Santoso', 'Jl. Melati No.2', '082234567891'),
(3, 'Citra Larasati', 'Jl. Anggrek No.3', '083334567892'),
(4, 'Dewi Lestari', 'Jl. Kamboja No.4', '084434567893'),
(5, 'Eko Prasetyo', 'Jl. Kenanga No.5', '085534567894'),
(6, 'Fajar Ramadhan', 'Jl. Dahlia No.6', '086634567895'),
(7, 'Gita Permata', 'Jl. Teratai No.7', '087734567896'),
(8, 'Hendra Wijaya', 'Jl. Flamboyan No.8', '088834567897'),
(9, 'Indah Wulandari', 'Jl. Cempaka No.9', '089934567898'),
(10, 'Joko Purnomo', 'Jl. Bougenville No.10', '080134567899'),
(11, 'Andi Nugroho', 'Jl. Mawar No.1', '081234567890'),
(12, 'Budi Santoso', 'Jl. Melati No.2', '082234567891'),
(13, 'Citra Larasati', 'Jl. Anggrek No.3', '083334567892'),
(14, 'Dewi Lestari', 'Jl. Kamboja No.4', '084434567893'),
(15, 'Eko Prasetyo', 'Jl. Kenanga No.5', '085534567894'),
(16, 'Fajar Ramadhan', 'Jl. Dahlia No.6', '086634567895'),
(17, 'Gita Permata', 'Jl. Teratai No.7', '087734567896'),
(18, 'Hendra Wijaya', 'Jl. Flamboyan No.8', '088834567897'),
(19, 'Indah Wulandari', 'Jl. Cempaka No.9', '089934567898'),
(20, 'Joko Purnomo', 'Jl. Bougenville No.10', '080134567899'),
(21, 'Aksara Angkasa', 'Jl. Merbabu 26', '085879000543');

-- --------------------------------------------------------

--
-- Table structure for table `transaksi`
--

CREATE TABLE `transaksi` (
  `id` int(11) NOT NULL,
  `pelanggan_id` int(11) DEFAULT NULL,
  `kendaraan_id` int(11) DEFAULT NULL,
  `tanggal_pinjam` date DEFAULT NULL,
  `tanggal_kembali` date DEFAULT NULL,
  `total` decimal(12,2) DEFAULT NULL,
  `status` enum('dipinjam','selesai') NOT NULL DEFAULT 'dipinjam'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `transaksi`
--

INSERT INTO `transaksi` (`id`, `pelanggan_id`, `kendaraan_id`, `tanggal_pinjam`, `tanggal_kembali`, `total`, `status`) VALUES
(1, 1, 1, '2025-07-01', '2025-07-03', 600000.00, 'selesai'),
(2, 2, 2, '2025-07-02', '2025-07-05', 960000.00, 'selesai'),
(3, 3, 8, '2025-07-03', '2025-07-04', 90000.00, 'selesai'),
(4, 4, 9, '2025-07-04', '2025-07-06', 170000.00, 'selesai'),
(5, 5, 5, '2025-07-06', '2025-07-09', 1050000.00, 'selesai'),
(6, 6, 6, '2025-07-07', '2025-07-08', 290000.00, 'selesai'),
(7, 7, 10, '2025-07-08', '2025-07-09', 80000.00, 'selesai'),
(8, 8, 3, '2025-07-09', '2025-07-10', 310000.00, 'dipinjam'),
(9, 9, 11, '2025-07-10', '2025-07-12', 190000.00, 'dipinjam'),
(10, 10, 4, '2025-07-11', '2025-07-13', 630000.00, 'selesai'),
(11, 1, 1, '2025-07-01', '2025-07-03', 600000.00, 'selesai'),
(12, 2, 2, '2025-07-02', '2025-07-05', 960000.00, 'selesai'),
(13, 3, 8, '2025-07-03', '2025-07-04', 90000.00, 'selesai'),
(14, 4, 9, '2025-07-04', '2025-07-06', 170000.00, 'selesai'),
(15, 5, 5, '2025-07-06', '2025-07-09', 1050000.00, 'selesai'),
(16, 6, 6, '2025-07-07', '2025-07-08', 290000.00, 'selesai'),
(17, 7, 10, '2025-07-08', '2025-07-09', 80000.00, 'selesai'),
(18, 8, 3, '2025-07-09', '2025-07-10', 310000.00, 'dipinjam'),
(19, 9, 11, '2025-07-10', '2025-07-12', 190000.00, 'dipinjam'),
(20, 10, 4, '2025-07-11', '2025-07-13', 630000.00, 'dipinjam');

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` int(11) NOT NULL,
  `username` varchar(50) NOT NULL,
  `password` varchar(100) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `username`, `password`) VALUES
(1, 'admin', 'admin123');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `kendaraan`
--
ALTER TABLE `kendaraan`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `pelanggan`
--
ALTER TABLE `pelanggan`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `transaksi`
--
ALTER TABLE `transaksi`
  ADD PRIMARY KEY (`id`),
  ADD KEY `pelanggan_id` (`pelanggan_id`),
  ADD KEY `kendaraan_id` (`kendaraan_id`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `username` (`username`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `kendaraan`
--
ALTER TABLE `kendaraan`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=37;

--
-- AUTO_INCREMENT for table `pelanggan`
--
ALTER TABLE `pelanggan`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=22;

--
-- AUTO_INCREMENT for table `transaksi`
--
ALTER TABLE `transaksi`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=21;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `transaksi`
--
ALTER TABLE `transaksi`
  ADD CONSTRAINT `transaksi_ibfk_1` FOREIGN KEY (`pelanggan_id`) REFERENCES `pelanggan` (`id`),
  ADD CONSTRAINT `transaksi_ibfk_2` FOREIGN KEY (`kendaraan_id`) REFERENCES `kendaraan` (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
