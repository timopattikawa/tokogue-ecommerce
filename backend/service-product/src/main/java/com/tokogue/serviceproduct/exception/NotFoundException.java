package com.tokogue.serviceproduct.exception;

import lombok.Getter;

@Getter
public class NotFoundException extends RuntimeException {
    private final int trxStatus;
    private final String errorDetail;

    public NotFoundException(String message) {
        super(message);
        this.trxStatus = 0;
        this.errorDetail = "";
    }

    public NotFoundException(String message, int trxStatus, String errorDetail) {
        super(message);
        this.trxStatus = trxStatus;
        this.errorDetail = errorDetail;
    }

}
