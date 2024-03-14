package com.cart.servicecartspring.v1.domain;


import jakarta.persistence.Column;
import jakarta.persistence.Entity;
import jakarta.persistence.Id;
import jakarta.persistence.Table;
import lombok.Data;

import java.sql.Date;

@Entity
@Data
@Table(name = "tb_cart")
public class Cart {

    @Id
    private Integer id;

    @Column(name = "user_id")
    private Integer userId;

    @Column(name = "total_amount")
    private Integer total_amount;

    @Column(name = "status_cart")
    private String statusCart;

    @Column(name = "created_at")
    private Date createdAt;

}
