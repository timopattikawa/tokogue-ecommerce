package com.cart.servicecartspring.v1.service;

import com.cart.servicecartspring.v1.dto.AddToCartRequest;
import com.cart.servicecartspring.v1.dto.AddToCartResponse;
import com.cart.servicecartspring.v1.repository.CartDetailRepository;
import com.cart.servicecartspring.v1.repository.CartRepository;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.mockito.Mock;
import org.mockito.junit.jupiter.MockitoExtension;


@ExtendWith(MockitoExtension.class)
class CartServiceTest {

    @Mock
    private CartRepository cartRepository;

    @Mock
    private CartDetailRepository cartDetailRepository;

    @Test
    public void TestAddToCart_SuccessWithNewCart() {
        AddToCartRequest request = new AddToCartRequest(1, 5);
        AddToCartResponse response = new AddToCartResponse("SUCCESS");
        M

    }


}