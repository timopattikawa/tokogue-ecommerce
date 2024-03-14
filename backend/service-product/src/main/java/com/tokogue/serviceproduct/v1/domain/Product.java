package com.tokogue.serviceproduct.v1.domain;

import lombok.Data;
import org.springframework.data.annotation.Id;
import org.springframework.data.mongodb.core.mapping.Document;
import org.springframework.data.mongodb.core.mapping.Field;

@Data
@Document(collection = "tokogue_product")
public class Product {

    @Id
    String id;
    @Field(name = "name_product")
    String productName;
    @Field(name = "stock")
    int stock;
    @Field(name = "price")
    int price;

}
