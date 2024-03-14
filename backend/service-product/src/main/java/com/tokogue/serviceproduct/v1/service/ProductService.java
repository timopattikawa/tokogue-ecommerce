package com.tokogue.serviceproduct.v1.service;

import com.tokogue.serviceproduct.v1.domain.Product;
import com.tokogue.serviceproduct.v1.dto.CartCheckRequest;
import com.tokogue.serviceproduct.v1.dto.CartCheckResponse;
import com.tokogue.serviceproduct.v1.dto.ProductDTO;
import org.apache.solr.client.solrj.SolrServerException;

import java.io.IOException;
import java.util.List;

public interface ProductService {
    public List<ProductDTO> getProducts(int page);
    public CartCheckResponse cartRequestProduct(CartCheckRequest request);
}
