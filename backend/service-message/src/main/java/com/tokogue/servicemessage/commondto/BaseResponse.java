package com.tokogue.servicemessage.commondto;

import com.fasterxml.jackson.annotation.JsonProperty;
import lombok.AllArgsConstructor;
import lombok.Data;

@Data
@AllArgsConstructor
public class BaseResponse<T> {

    @JsonProperty("status_code")
    Integer statusCode;
    BaseErrorResponse errorResponse;
    T data;

}
