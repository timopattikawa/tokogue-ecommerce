package com.tokogue.serviceproduct.v1.dto;

import com.fasterxml.jackson.annotation.JsonProperty;
import lombok.AllArgsConstructor;
import lombok.Data;

@Data
@AllArgsConstructor
public class CartCheckResponse {

    @JsonProperty("price")
    int price;
    @JsonProperty("is_stock_reduced")
    boolean isStockReduced;

}
