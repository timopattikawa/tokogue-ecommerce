package com.tokogue.serviceproduct.base;

import com.fasterxml.jackson.annotation.JsonProperty;
import lombok.Data;

@Data
public class BaseErrorResponse {
    @JsonProperty("error_code")
    int errorCode;
    @JsonProperty("error_message")
    String errorMessage;
    @JsonProperty("error_detail")
    GlobalErrorDetail errorDetail;
}
