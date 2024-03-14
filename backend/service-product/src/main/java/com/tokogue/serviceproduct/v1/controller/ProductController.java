package com.tokogue.serviceproduct.v1.controller;

import com.tokogue.serviceproduct.base.BaseResponse;
import com.tokogue.serviceproduct.v1.dto.CartCheckRequest;
import com.tokogue.serviceproduct.v1.dto.CartCheckResponse;
import com.tokogue.serviceproduct.v1.dto.ProductDTO;
import com.tokogue.serviceproduct.v1.service.ProductService;
import lombok.RequiredArgsConstructor;
import org.springframework.data.repository.query.Param;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;

import java.util.List;

@Controller
@RequiredArgsConstructor
@RequestMapping("/v1/product")
public class ProductController {

    private final ProductService productService;

    @GetMapping
    public ResponseEntity<BaseResponse<List<ProductDTO>>> getAllProduct(@Param("page") int page) {
        List<ProductDTO> products = productService.getProducts(page);
        BaseResponse<List<ProductDTO>> baseResponse = BaseResponse
                .of(HttpStatus.OK.value(), null, products);
        return new ResponseEntity<>(baseResponse, HttpStatus.OK);
    }

    @GetMapping("/search/")
    public ResponseEntity<String> searchProduct(@Param("query") String query) {
        return new ResponseEntity<>("", HttpStatus.OK);
    }

    @GetMapping("/suggest/")
    public ResponseEntity<String> suggestProduct(@Param("query") String query) {
        return new ResponseEntity<>("", HttpStatus.OK);
    }

    @PostMapping("/cart-stock")
    public ResponseEntity<BaseResponse<CartCheckResponse>> cartRequestStock(@RequestBody CartCheckRequest request) {
        CartCheckResponse cartCheckResponse = productService.cartRequestProduct(request);
        return new ResponseEntity<>(
                BaseResponse.of( HttpStatus.OK.value(), null, cartCheckResponse),
                HttpStatus.OK
        );
    }

}
