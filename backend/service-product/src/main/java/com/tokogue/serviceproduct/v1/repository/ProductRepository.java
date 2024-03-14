package com.tokogue.serviceproduct.v1.repository;

import com.tokogue.serviceproduct.v1.domain.Product;
import org.springframework.data.mongodb.repository.MongoRepository;
import org.springframework.stereotype.Repository;

@Repository
public interface ProductRepository extends MongoRepository<Product, String> {
}
