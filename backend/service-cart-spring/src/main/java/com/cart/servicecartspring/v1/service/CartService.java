package com.cart.servicecartspring.v1.service;

import com.cart.servicecartspring.v1.dto.AddToCartRequest;
import com.cart.servicecartspring.v1.dto.AddToCartResponse;

public interface CartService {

    public AddToCartResponse addToCart(AddToCartRequest request);
}
