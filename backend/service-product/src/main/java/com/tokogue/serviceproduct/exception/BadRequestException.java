package com.tokogue.serviceproduct.exception;

import lombok.Getter;

@Getter
public class BadRequestException extends RuntimeException {

    int trxStatus;
    String errorDetail;

    public BadRequestException(String message) {
        super(message);
        this.trxStatus = 0;
        this.errorDetail = "";
    }

    public BadRequestException(String message, int trxStatus, String errorDetail) {
        super(message);
        this.trxStatus = trxStatus;
        this.errorDetail = errorDetail;
    }

}
