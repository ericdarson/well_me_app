create table T_NASABAH(
BCA_ID 	varchar2(20),
nama	varchar2(40),
email	varchar2(320),
no_rekening 	varchar2(10),
password	varchar2(32),
token	varchar2(32),
SID	varchar2(16),
tanggal_join	date,
date_updated	date,
status	number(1),
wrong_attempt	number(1),
date_locked 	date,
bobot_resiko	number(2,1),
no_hp		varchar2(13),
Creation_date		date,
Last_update_date	date,
Last_update_by	varchar2(30),
constraint t_nasabah_pk primary key(bca_id)
);

create table T_planner(
ID_PLAN		Number constraint t_planner_pk Primary key,
Nama_plan	varchar2(20),
Goal_Amount	Number,
Current_amount	number,
flag_checker	number(1),
periodic		date,
due_date		date,
puzzle_randomize	varchar2(9),
Creation_date		date,
Last_update_date	date,
Last_update_by	varchar2(30),
BCA_ID		varchar2(20),
constraint T_Planner_fk1 foreign key (BCA_ID) REFERENCES t_nasabah(BCA_ID)
);

create table t_promo(
kode_promo	varchar2(10),
tanggal_mulai		date,
tanggal_selesai		date,
deskripsi			varchar2(4000),
cashback			number,
gambar_promosi		blob,
target_akumulasi	number,
minimum_transaction	number,
Creation_date		date,
Last_update_date	date,
Last_update_by	varchar2(30),
constraint t_promo_pk primary key (kode_promo)
);

create table T_nasabah_promo(
BCA_ID		varchar2(20),
kode_promo varchar2(10),
usage_flag number(1),
akumulasi number,
target_akumulasi number,
Creation_date		date,
Last_update_date	date,
Last_update_by	varchar2(30),
constraint t_nasabah_promo_fk1 foreign key (kode_promo) references t_promo (kode_promo),
constraint t_nasabah_promo_fk2 foreign key (bca_id) references t_nasabah (bca_id),
constraint t_nasabah_promo_pk primary key (BCA_ID,kode_promo)
);

create table T_Reksadana_nasabah(
ID_reksadana_nasabah	number,
BCA_ID					varchar2(20),
ID_PRODUK				number,
ID_PLAN					number,
jumlah_unit				number,
nab_rerata				number,
nab_sekarang			number,
Creation_date		date,
Last_update_date	date,
Last_update_by	varchar2(30),
constraint t_reksadana_nasabah_pk primary key (id_reksadana_nasabah),
constraint t_reksadana_nasabah_fk1 foreign key (bca_id) references t_nasabah (bca_id),
constraint t_reksadana_nasabah_fk2 foreign key (ID_PLAN) references t_planner (ID_PLAN)
);

create table M_JENIS_REKSADANA(
id_jenis_reksadana	number,
nama_jenis			varchar2(30),
Creation_date		date,
Last_update_date	date,
Last_update_by	varchar2(30),
constraint M_JENIS_REKSADANA_pk primary key (id_jenis_reksadana)
);

create table T_Produk_reksadana(
id_produk	number,
nama_produk	varchar2(50),
minimum_unit	number,
alokasi_asset	number,
expense_ratio	number,
NAB		number,
TOTAL_AUM	NUMBER,
manager_investasi varchar2(50),
tingkat_resiko	varchar2(10),
bank_kustodian	varchar2(50),
bank_penampung	varchar2(50),
year_on_year_nab	number,
three_months_nab	number,
monthly_nab		number,
weekly_nab	number,
id_jenis_reksadana	number CONSTRAINT t_produk_reksadana_fk1 references M_JENIS_REKSADANA(id_jenis_reksadana),
Creation_date		date,
Last_update_date	date,
Last_update_by	varchar2(30),
CONSTRAINT t_produk_reksadana_pk primary key (id_produk)
);


create table T_daily_nab(
id_daily number,
nab_daily number,
date_daily	date,
ID_produk number,
Creation_date		date,
Last_update_date	date,
Last_update_by	varchar2(30),
constraint t_daily_nab_pk primary key (id_daily),
constraint t_daily_nab_fk1 foreign key (id_produk) references t_produk_reksadana (id_produk)
);



create table M_Resiko(
ID_RESIKO			Number,
bobot_resiko		number(2,1),
persentase			number,
ID_JENIS_REKSADANA	NUMBER,
Creation_date		date,
Last_update_date	date,
Last_update_by	varchar2(30),
constraint m_resiko_pk primary key (id_resiko),
constraint m_resiko_fk foreign key (id_jenis_reksadana) references M_jenis_reksadana (id_jenis_reksadana)
);
--INSERT INTO M_RESIKO VALUES(1,1.5,);
create table T_Transaksi_jual(
ID_TRANS_JUAL	varchar2(11),
BCA_ID		varchar2(20),
ID_PRODUK	number,
ID_PLAN		number,
NAB		NUMBER,
JUMLAH_UNIT	number,
transaction_date	date,
status_order_jual	number,
tanggal_order_jual	date,
status_verifikasi_bank	number,
tanggal_verifikasi_bank	date,
status_penjualan	number,
tanggal_penjualan	date,
Creation_date		date,
Last_update_date	date,
Last_update_by	varchar2(30),
constraint t_transaksi_jual_pk primary key (id_trans_jual),
constraint t_transaksi_jual_fk1 foreign key (id_produk) references t_produk_reksadana (id_produk),
constraint t_transaksi_jual_fk2 foreign key (id_plan) references t_planner (id_plan)
);

create table T_Transaksi_beli(
ID_TRANS_beli	varchar2(11),
BCA_ID		varchar2(20),
ID_PRODUK	number,
ID_PLAN		number,
NAB		NUMBER,
JUMLAH_UNIT	number,
transaction_date	date,
status_order_beli	number,
tanggal_order_beli	date,
status_verifikasi_bank	number,
tanggal_verifikasi_bank	date,
status_pembelian	number,
tanggal_pembelian	date,
Creation_date		date,
Last_update_date	date,T
Last_update_by	varchar2(30),
constraint t_transaksi_beli_pk primary key (id_trans_beli),
constraint t_transaksi_beli_fk1 foreign key (id_produk) references t_produk_reksadana (id_produk),
constraint t_transaksi_beli_fk2 foreign key (id_plan) references t_planner (id_plan)
);

alter table t_reksadana_nasabah add constraint t_reksadana_nasabah3 foreign key (id_produk) references t_produk_reksadana (id_produk);

--tambahan james
ALTER TABLE T_NASABAH ADD (TOKEN_EXPIRED DATE);
UPDATE T_PLANNER SET PERIODIC=NULL;
ALTER TABLE T_PLANNER MODIFY (PERIODIC VARCHAR2(20));
ALTER TABLE T_PLANNER ADD (NEXT_TARGET_AMOUNT NUMBER );

ALTER TABLE T_PLANNER 
ADD (IS_DONE NUMBER DEFAULT 0 );

ALTER TABLE T_PLANNER 
ADD (IS_DELETED NUMBER DEFAULT 0 );

--tambahan natan
alter table T_TRANSAKSI_BELI drop column transaction_date;
alter table T_TRANSAKSI_JUAL drop column transaction_date;

ALTER TABLE T_PRODUK ADD (URL_VENDOR VARCHAR2(2048));
ALTER TABLE T_PRODUK ADD (PASSWORD_VENDOR VARCHAR2(32));

ALTER TABLE T_TRANSAKSI_BELI ADD (TOTAL_BELI NUMBER);

ALTER TABLE T_TRANSAKSI_BELI 
MODIFY ID_TRANS_BELI VARCHAR2(12);

ALTER TABLE T_TRANSAKSI_JUAL 
MODIFY ID_TRANS_JUAL VARCHAR2(12);

ALTER TABLE T_REKSADANA_NASABAH 
ADD (STATUS_DAILY_UPDATE NUMBER DEFAULT 0);

ALTER TABLE M_RESIKO ADD (LEVEL_RESIKO NUMBER);

ALTER TABLE T_Produk_reksadana ADD (LEVEL_RESIKO NUMBER);


ALTER TABLE T_PLANNER 
ADD (KATEGORI VARCHAR2(20) );

ALTER TABLE T_PLANNER  MODIFY (PUZZLE_RANDOMIZE VARCHAR2(10 BYTE) );

ALTER TABLE M_RESIKO ADD (TINGKAT_RESIKO varchar2(10));

ALTER TABLE T_PROMO ADD (TITLE VARCHAR2(50));
ALTER TABLE T_PROMO ADD (SUBTITLE VARCHAR2(75));

ALTER TABLE M_JENIS_REKSADANA ADD (AKTIF NUMBER(1) DEFAULT 1);
ALTER TABLE M_JENIS_REKSADANA DROP COLUMN AKTIF;
--tambahan marcell
ALTER TABLE t_produk_reksadana DROP COLUMN alokasi_asset;


--1 JAN 2015
ALTER TABLE T_NASABAH ADD (PIN VARCHAR2(32 BYTE));
ALTER TABLE T_PRODUK_REKSADANA ADD (BIAYA_PENJUALAN NUMBER DEFAULT 0);
ALTER TABLE T_PRODUK_REKSADANA ADD (MINIMUM_SISA_UNIT NUMBER DEFAULT 0);