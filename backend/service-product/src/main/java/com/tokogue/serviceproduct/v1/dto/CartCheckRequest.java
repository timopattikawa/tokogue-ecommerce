package com.tokogue.serviceproduct.v1.dto;

import com.fasterxml.jackson.annotation.JsonProperty;
import lombok.Data;

@Data
public class CartCheckRequest {

    @JsonProperty("product_id")
    String productId;

    @JsonProperty("qty")
    int qty;

}
