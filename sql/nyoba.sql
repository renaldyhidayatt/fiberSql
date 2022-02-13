CREATE TABLE sumber(
    id serial PRIMARY KEY,
    nama VARCHAR(50) NOT NULL
);

CREATE TABLE uang(
    id serial PRIMARY KEY,
    tgl_uang VARCHAR(50) NOT NULL,
    id_pengeluaran INT NOT NULL CHECK (id_pengeluaran >= 0),
    id_pendapatan INT NOT NULL CHECK (id_pendapatan >= 0),
    jumlah INT NOT NULL CHECK (jumlah >= 0)
);

CREATE TABLE pengeluaran (
    id serial PRIMARY KEY,
    tgl_pengeluaran VARCHAR(50) NOT NULL,
    jumlah INT NOT NULL,
    id_sumber INT NOT NULL
);

CREATE TABLE pemasukan (
    id serial PRIMARY KEY,
    tgl_pemasukan DATE NOT NULL,
    jumlah INT NOT NULL,
    id_sumber INT NOT NULL
);


CREATE TABLE karyawan(
    id serial PRIMARY KEY,
    nama VARCHAR(50) NOT NULL,
    posisi VARCHAR(50) NOT NULL,
    alamat VARCHAR(50) NOT NULL,
    umur INT NOT NULL,
    kontak VARCHAR(40) NOT NULL
);


CREATE TABLE users(
    id serial PRIMARY KEY,
    email VARCHAR(50) NOT NULL,
    name VARCHAR(50) NOT NULL,
    password VARCHAR(255) NOT NULL
);