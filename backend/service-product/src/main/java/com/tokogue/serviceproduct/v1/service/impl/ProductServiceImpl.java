package com.tokogue.serviceproduct.v1.service.impl;

import com.tokogue.serviceproduct.exception.BadRequestException;
import com.tokogue.serviceproduct.exception.InternalServerException;
import com.tokogue.serviceproduct.exception.NotFoundException;
import com.tokogue.serviceproduct.v1.domain.Product;
import com.tokogue.serviceproduct.v1.dto.CartCheckRequest;
import com.tokogue.serviceproduct.v1.dto.CartCheckResponse;
import com.tokogue.serviceproduct.v1.dto.ProductDTO;
import com.tokogue.serviceproduct.v1.repository.ProductRepository;
import com.tokogue.serviceproduct.v1.service.ProductService;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.apache.solr.client.solrj.SolrClient;
import org.apache.solr.client.solrj.response.QueryResponse;
import org.apache.solr.common.SolrDocument;
import org.apache.solr.common.params.MapSolrParams;
import org.springframework.stereotype.Service;

import java.util.*;

@Service
@Slf4j
@RequiredArgsConstructor
public class ProductServiceImpl implements ProductService {

    private final ProductRepository productRepository;
    private final SolrClient solrClient;

    @Override
    public List<ProductDTO> getProducts(int page){
        try {
            log.info("Try to query products in solr");
            int start = page * 10;
            final Map<String, String> queryParamMap = new HashMap<>();
            queryParamMap.put("q", "*:*");
            queryParamMap.put("start", Integer.toString(start));
            queryParamMap.put("rows", "10");

            log.info("Query to solr : {}", queryParamMap);
            MapSolrParams mapSolrParams = new MapSolrParams(queryParamMap);
            QueryResponse tokogueProduct = solrClient.query("TOKOGUE_PRODUCT", mapSolrParams);

            List<ProductDTO> productDTOList = new ArrayList<>();
            for (SolrDocument result : tokogueProduct.getResults()) {
                String id = (String) result.get("_id");
                String productName = (String) result.get("product_name");
                List<String> genre = (List<String>) result.get("genre");
                Integer price = (Integer) result.get("price");
                ProductDTO productDTO = new ProductDTO(id, productName, genre, price);
                productDTOList.add(productDTO);
            }
            log.info("Successful for query from solr");
            return productDTOList;
        } catch (Exception e) {
            log.info("Error in query from solr");
            log.error(e.getMessage());
            throw new InternalServerException("Fail to query from solr");
        }
    }

    @Override
    public CartCheckResponse cartRequestProduct(CartCheckRequest request) {

        Product product = productRepository.findById(request.getProductId()).orElseThrow(
                () -> new NotFoundException("Not Found Product")
        );

        if(product.getStock() < request.getQty()) {
            throw new BadRequestException("Stock not available", 9, "Stock avail is " + product.getStock());
        }

        product.setStock(product.getStock() - request.getQty());

        Product save = productRepository.save(product);

        return new CartCheckResponse(save.getStock(), product.getPrice(), true);
    }


}
