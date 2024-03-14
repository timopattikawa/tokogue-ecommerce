package com.tokogue.serviceproduct.base;

import lombok.Data;
import lombok.Value;

@Data
@Value(staticConstructor = "of")
public class BaseResponse<T> {

    int status;
    BaseErrorResponse error;
    T data;

}
