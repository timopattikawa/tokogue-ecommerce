package com.cart.servicecartspring.v1.domain;


import jakarta.persistence.Column;
import jakarta.persistence.Entity;
import jakarta.persistence.Id;
import jakarta.persistence.Table;

@Entity
@Table(name = "tb_cart_detail")
public class CartDetail {

    @Id
    private Integer id;

    @Column(name = "product_id")
    private Integer productId;

    @Column(name = "qty")
    private Integer qty;

    @Column(name = "cart_id")
    private Integer cartId;

}
