package com.cart.servicecartspring.v1.dto;

import com.fasterxml.jackson.annotation.JsonProperty;
import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@AllArgsConstructor
@NoArgsConstructor
public class AddToCartRequest {

    @JsonProperty("product_id")
    int productId;

    @JsonProperty("qty")
    int qty;

}
