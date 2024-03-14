package com.cart.servicecartspring.v1.repository;

import com.cart.servicecartspring.v1.domain.Cart;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

@Repository
public interface CartRepository extends JpaRepository<Cart, Integer> {
}
