package com.tokogue.serviceproduct.base;

import com.fasterxml.jackson.annotation.JsonProperty;
import lombok.AllArgsConstructor;
import lombok.Data;

@Data
@AllArgsConstructor
public class GlobalErrorDetail {

    @JsonProperty("trx_status")
    int trxStatus;
    @JsonProperty("error_detail")
    String errorDetail;

}
