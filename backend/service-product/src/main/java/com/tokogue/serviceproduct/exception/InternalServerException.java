package com.tokogue.serviceproduct.exception;

import lombok.Getter;

@Getter
public class InternalServerException extends RuntimeException{

    private int trxStatus;
    private String errorDetail;

    public InternalServerException(String message) {
        super(message);
        this.trxStatus = 0;
        this.errorDetail = "";
    }

    public InternalServerException(String message, int trxStatus, String errorDetail) {
        super(message);
        this.trxStatus = trxStatus;
        this.errorDetail = errorDetail;
    }

}
