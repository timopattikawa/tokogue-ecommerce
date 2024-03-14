package com.tokogue.serviceproduct.v1.dto;

import com.fasterxml.jackson.annotation.JsonProperty;
import lombok.AllArgsConstructor;
import lombok.Data;

import java.util.List;

@Data
@AllArgsConstructor
public class ProductDTO {

    @JsonProperty("id")
    String id;
    @JsonProperty("name_product")
    String productName;
    @JsonProperty("genre")
    List<String> genre;
    @JsonProperty("price")
    Integer price;

}
