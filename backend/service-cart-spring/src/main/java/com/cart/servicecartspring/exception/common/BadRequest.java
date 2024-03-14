package com.cart.servicecartspring.exception.common;

public class BadRequest extends RuntimeException{
    public BadRequest(String message) {
        super(message);
    }
}
