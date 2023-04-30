create database ttgoksh;
use ttgoksh;

select * from members;
select * from products;
select * from review_products;
select * from like_reviews;

insert into products (name_product,price) values ("shinsui",25000);
insert into products (name_product,price) values ('nivea',35000);
insert into products (name_product,price) values ('makarea',67000);

insert into review_products (id_product,id_member,desc_review) values (1,1,"product bagus cocok buat saya");
insert into review_products (id_product,id_member,desc_review) values (1,2,"product mantul");
insert into review_products (id_product,id_member,desc_review) values (2,2,"product bagus banget");
insert into review_products (id_product,id_member,desc_review) values (3,3,"Kurang cocok sama productnya,tapi ok lah");

-- USERNAME, GENDER, SKINTYPE,SKINCOLOR, DESC_REVIEW, JUMLAH_LIKE_REVIEW;
select username,gender,skin_type,skin_color,desc_review,count(like_reviews.id_member) as JUMLAH_LIKE_REVIEW
from review_products
left join members on members.id_member =  review_products.id_member
left join like_reviews on like_reviews.id_review = review_products.id_review
where id_product=1
group by review_products.id_review;